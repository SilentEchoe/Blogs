package main

import "fmt"

type Backup struct {
	images string
	Args   []string
}

func main() {
	backupInfo := newBackup()
	backupInfo.SetArgs("127.0.0.1:3306")
	backupInfo.SetArgsTwo("mongo://127.0.0.1")
	fmt.Println(backupInfo)
}

func newBackup() *Backup {
	return &Backup{
		images: "nginx:latest",
	}
}

func (b *Backup) SetArgs(mysqladdress string) {
	newArrg := []string{"InMysqlAddress=" + mysqladdress}
	b.Args = append(b.Args, newArrg...)
}

func (b *Backup) SetArgsTwo(mongoAddress string) {
	newArrg := []string{"InmongoAddress=" + mongoAddress}
	b.Args = append(b.Args, newArrg...)
}
