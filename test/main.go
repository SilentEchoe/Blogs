package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	// 指定源DLL文件和目标目录

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Current directory:", currentDir)

	sourceDLL := currentDir + "/libmcfgthread-1.dll"
	destinationDir := "C:\\Windows\\System32"

	// 检查源DLL文件是否存在
	_, err = os.Stat(sourceDLL)
	if os.IsNotExist(err) {
		fmt.Printf("Source DLL file '%s' does not exist\n", sourceDLL)
		os.Exit(1)
	}

	// 打开源DLL文件
	source, err := os.Open(sourceDLL)
	if err != nil {
		fmt.Printf("Error opening source DLL file: %v\n", err)
		os.Exit(1)
	}
	defer source.Close()

	// 获取DLL文件名
	dllFileName := filepath.Base(sourceDLL)

	// 创建目标目录（如果不存在）
	err = os.MkdirAll(destinationDir, os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating destination directory: %v\n", err)
		os.Exit(1)
	}

	// 创建目标DLL文件
	destinationDLL := filepath.Join(destinationDir, dllFileName)
	destination, err := os.Create(destinationDLL)
	if err != nil {
		fmt.Printf("Error creating destination DLL file: %v\n", err)
		os.Exit(1)
	}
	defer destination.Close()

	// 拷贝DLL文件内容
	_, err = io.Copy(destination, source)
	if err != nil {
		fmt.Printf("Error copying DLL file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("DLL file '%s' copied successfully to '%s'!\n", sourceDLL, destinationDLL)
}
