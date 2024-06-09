package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	ExcelPath string `json:"excel_path"`
}

func ReadJsonConfig() string {
	// 读取配置文件
	filePath := "./config.json"
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("读取文件错误：", err)
		return ""
	}

	// 解析 JSON 数据
	config := Config{}
	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("解析JSON错误：", err)
		return ""
	}
	// 打印配置信息
	return config.ExcelPath
}
