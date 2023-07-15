package main

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/procfs"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"io"
	"k8s.io/klog/v2"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	initGorm()

	router := gin.Default()

	go func() {
		time.Sleep(2 * time.Second)
		WatchSysInfo()
	}()

	// Define a GET endpoint
	router.GET("/", func(c *gin.Context) {

		data := GetSysInfo()

		c.JSON(http.StatusOK, gin.H{"data": data})
	})

	// Run the server on port 8080
	router.Run(":8000")
}

var gormOnce = sync.Once{}
var db *gorm.DB

var NodeName string

func initGorm() *gorm.DB {
	gormOnce.Do(func() {
		var err error
		db, err = gorm.Open(sqlite.Open("/data/agent.db"), &gorm.Config{})
		if err != nil {
			klog.Errorf("connect sqlite error:%v", err)
			return
		}
	})

	NodeName = "service2"

	if db == nil {
		klog.Errorf("open sqlite error")
		return nil
	}

	err := db.AutoMigrate(&RegisCone{}, &SysInfo{})
	if err != nil {
		klog.Errorf("auto migrate error:%v", err)
		return nil
	}
	return db
}

type RegisCone struct {
	TaskId    string `json:"task-id"`
	SvcName   string `json:"svc-name"`
	Port      string `json:"port"`
	Namespace string `json:"namespace"`
}

// AddCone 添加模型连接
func AddCone(taskId string, svcName string, port string, ns string) {
	var regiscone RegisCone
	db.Where("task_id = ?", taskId).First(&regiscone)
	if regiscone != (RegisCone{}) {
		db.Model(&regiscone).Where("task_id = ?", taskId).Updates(map[string]interface{}{"svc_name": svcName, "port": port, "namespace": ns})
	} else {
		db.Create(&RegisCone{TaskId: taskId, SvcName: svcName, Port: port, Namespace: ns})
	}
}

func GetConeByTaskId(taskId string) *RegisCone {
	var cone RegisCone
	db.Where("task_id = ?", taskId).First(&cone)
	return &cone
}

func DelConeByTaskId(taskId string) {
	db.Delete(&RegisCone{}, "task_id = ?", taskId)
}

func GetAllCone() []*RegisCone {
	var cone []*RegisCone
	db.Find(&cone)
	return cone
}

type SysInfo struct {
	Cpu      float64 `json:"cpu"`
	Mem      float64 `json:"mem"`
	HostName string  `json:"hostname"`
}

// WatchSysInfo 监控系统信息并存储
func WatchSysInfo() {
	//cpu := sys.GetCpu()
	//mem := sys.GetMem()

	cpu, mem := NewGetSysInfo()

	var sysinfo SysInfo
	if db == nil {
		klog.Errorf("db is nil")
		return
	}

	db.Where("host_name = ?", NodeName).First(&sysinfo)
	if sysinfo == (SysInfo{}) {
		//klog.Info("创建系统信息:", NodeName)
		db.Create(&SysInfo{Cpu: cpu, Mem: mem, HostName: NodeName})
	} else {
		//klog.Info("修改系统信息:", NodeName)
		db.Model(&sysinfo).Where("host_name = ?", NodeName).Updates(map[string]interface{}{"cpu": cpu, "mem": mem})
	}
}

func GetSysInfo() []*SysInfo {
	var sysInfo []*SysInfo
	db.Find(&sysInfo)
	return sysInfo
}

var OldCpuAll float64
var uses float64

func NewGetSysInfo() (float64, float64) {

	OldCpuAll, uses = GetStats()

	time.Sleep(time.Second * 1)
	newCpuAll, newUses := GetStats()

	cpu := (newUses - uses) / (newCpuAll - OldCpuAll) * 100
	OldCpuAll, uses = newCpuAll, newUses

	meminfo, err := getMemInfo()
	if err != nil {
		klog.Errorf("getMemInfo error:%v", err)
		return 0, 0
	}
	MemUsed := meminfo["MemTotal_bytes"] - meminfo["MemFree_bytes"] - meminfo["Buffers_bytes"] - meminfo["Cached"]
	MemUsed = MemUsed / meminfo["MemTotal_bytes"] * 100

	//klog.Info("内存占用率为:", MemUsed)
	return cpu, MemUsed
}

func GetStats() (float64, float64) {
	fs, err := procfs.NewFS("/proc")
	if err != nil {
		klog.Errorf(err.Error())
		return 0, 0
	}
	stats, err := fs.Stat()
	if err != nil {
		klog.Errorf(err.Error())
		return 0, 0
	}

	cpuAll := stats.CPUTotal.System + stats.CPUTotal.User + stats.CPUTotal.Nice + stats.CPUTotal.Idle + stats.CPUTotal.Iowait + stats.CPUTotal.IRQ + stats.CPUTotal.SoftIRQ + stats.CPUTotal.Steal + stats.CPUTotal.Guest + stats.CPUTotal.GuestNice
	htcpu := cpuAll - stats.CPUTotal.Idle - stats.CPUTotal.Iowait

	return cpuAll, htcpu
}

var (
	reParens = regexp.MustCompile(`\((.*)\)`)
)

func getMemInfo() (map[string]float64, error) {
	file, err := os.Open(procFilePath("meminfo"))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return parseMemInfo(file)
}

func procFilePath(name string) string {
	return filepath.Join("/proc", name)
}

func parseMemInfo(r io.Reader) (map[string]float64, error) {
	var (
		memInfo = map[string]float64{}
		scanner = bufio.NewScanner(r)
	)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		// Workaround for empty lines occasionally occur in CentOS 6.2 kernel 3.10.90.
		if len(parts) == 0 {
			continue
		}
		fv, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value in meminfo: %w", err)
		}
		key := parts[0][:len(parts[0])-1] // remove trailing : from key
		// Active(anon) -> Active_anon
		key = reParens.ReplaceAllString(key, "_${1}")
		switch len(parts) {
		case 2: // no unit
		case 3: // has unit, we presume kB
			fv *= 1024
			key = key + "_bytes"
		default:
			return nil, fmt.Errorf("invalid line in meminfo: %s", line)
		}
		memInfo[key] = fv
	}

	return memInfo, scanner.Err()
}
