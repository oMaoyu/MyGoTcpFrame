package net

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// 定义config 对应的结构体

type Config struct {
	IP      string
	Port    uint32
	Name    string
	Version string
}
var MyConfig Config
// 获取config文件数据
func LoadConfig() error {
	confInfo,err := ioutil.ReadFile("./conf/conf.json")
	if err != nil {
		return err
	}
	// 序列化数据
	err = json.Unmarshal(confInfo,&MyConfig)
	if err != nil {
		return err
	}
	return nil
}

// 初始化配置
func init(){
	err := LoadConfig()
	if err != nil {
		fmt.Println("配置初始化失败")
		os.Exit(-1)
	}
}