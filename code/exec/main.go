package main

import (
	"debug/elf"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/prometheus/procfs"
)

type InstrumentableType int

const (
	InstrumentableGolang InstrumentableType = iota + 1
	InstrumentableJava
	InstrumentableDotnet
	InstrumentablePython
	InstrumentableRuby
	InstrumentableNodejs
	InstrumentableRust
	InstrumentableGeneric
	InstrumentablePHP
)

func main() {
	pid := os.Getenv("targetPID")
	if pid == "" {
		fmt.Println("未设置目标服务的PID")
		return
	}

	nPid, err := strconv.ParseInt(pid, 10, 32)
	if err != nil {
		fmt.Println("PID 格式不正确")
	}

	proc := FindProcLanguage(int32(nPid), nil, "")
	switch proc {
	case InstrumentableGolang:
		fmt.Println("目标应用为Go语言编写")
	default:
		fmt.Println("目标应用为其他编程语言编写")
	}
}

var rubyModule = regexp.MustCompile(`^(.*/)?ruby[\d.]*$`)
var pythonModule = regexp.MustCompile(`^(.*/)?python[\d.]*$`)

func instrumentableFromModuleMap(moduleName string) InstrumentableType {
	if strings.Contains(moduleName, "libcoreclr.so") {
		return InstrumentableDotnet
	}
	if strings.Contains(moduleName, "libjvm.so") {
		return InstrumentableJava
	}
	if strings.HasSuffix(moduleName, "/node") || moduleName == "node" {
		return InstrumentableNodejs
	}
	// 通过正则表达
	if rubyModule.MatchString(moduleName) {
		return InstrumentableRuby
	}
	if pythonModule.MatchString(moduleName) {
		return InstrumentablePython
	}

	return InstrumentableGeneric
}

func FindProcLanguage(pid int32, elfF *elf.File, path string) InstrumentableType {
	// 查找Lib信息
	maps, err := FindLibMaps(pid)

	if err != nil {
		return InstrumentableGeneric
	}

	// 根据Lib推断是哪种编程语言
	for _, m := range maps {
		t := instrumentableFromModuleMap(m.Pathname)
		if t != InstrumentableGeneric {
			return t
		}
	}

	// 如果文件格式为空，那么通过PID的路径获取ELF文件信息
	if elfF == nil {
		pidPath := fmt.Sprintf("/proc/%d/exe", pid)
		elfF, err = elf.Open(pidPath)

		if err != nil || elfF == nil {
			return InstrumentableGeneric
		}
	}

	// 通过ELF文件来推断是哪种编程语言
	t := findLanguageFromElf(elfF)
	if t != InstrumentableGeneric {
		return t
	}

	// 判断是否是PHP
	t = instrumentableFromPath(path)
	if t != InstrumentableGeneric {
		return t
	}

	bytes, err := os.ReadFile(fmt.Sprintf("/proc/%d/environ", pid))
	if err != nil {
		return InstrumentableGeneric
	}
	// 最后判断是否是.Net
	return instrumentableFromEnviron(string(bytes))
}

func findLanguageFromElf(elfF *elf.File) InstrumentableType {
	// go语言的特征
	gosyms := elfF.Section(".gosymtab")

	if gosyms != nil {
		return InstrumentableGolang
	}

	return matchExeSymbols(elfF)
}

func FindLibMaps(pid int32) ([]*procfs.ProcMap, error) {
	proc, err := procfs.NewProc(int(pid))

	if err != nil {
		return nil, err
	}

	return proc.ProcMaps()
}

func matchExeSymbols(f *elf.File) InstrumentableType {
	syms, err := f.Symbols()
	if err != nil && !errors.Is(err, elf.ErrNoSymbols) {
		return InstrumentableGeneric
	}

	t := matchSymbols(syms)
	if t != InstrumentableGeneric {
		return t
	}

	dynsyms, err := f.DynamicSymbols()
	if err != nil && !errors.Is(err, elf.ErrNoSymbols) {
		return InstrumentableGeneric
	}

	return matchSymbols(dynsyms)
}

func matchSymbols(syms []elf.Symbol) InstrumentableType {
	for _, s := range syms {
		if elf.ST_TYPE(s.Info) != elf.STT_FUNC {
			// Symbol not associated with a function or other executable code.
			continue
		}
		t := instrumentableFromSymbolName(s.Name)
		if t != InstrumentableGeneric {
			return t
		}
	}

	return InstrumentableGeneric
}

func instrumentableFromSymbolName(symbol string) InstrumentableType {
	if strings.Contains(symbol, "rust_panic") {
		return InstrumentableRust
	}
	// 如果包含JVM 代表是JAVA
	if strings.HasPrefix(symbol, "JVM_") || strings.HasPrefix(symbol, "graal_") {
		return InstrumentableJava
	}

	return InstrumentableGeneric
}

func instrumentableFromPath(path string) InstrumentableType {
	if strings.Contains(path, "php") {
		return InstrumentablePHP
	}
	return InstrumentableGeneric
}

// 判断是否是.Net
func instrumentableFromEnviron(environ string) InstrumentableType {
	if strings.Contains(environ, "ASPNET") || strings.Contains(environ, "DOTNET") {
		return InstrumentableDotnet
	}
	return InstrumentableGeneric
}
