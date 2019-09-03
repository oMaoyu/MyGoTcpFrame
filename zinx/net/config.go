package net

import (
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
)

// 定义config 对应的结构体
type server struct {
	Port    uint32
	IP      string
	Name    string
	Version string
}

var MyConfig server

// 获取config文件数据
// json格式
func LoadConfig() error {
	confInfo, err := ioutil.ReadFile("./conf/conf.json")
	if err != nil {
		return err
	}
	// 序列化数据
	err = json.Unmarshal(confInfo, &MyConfig)
	if err != nil {
		return err
	}
	return nil
}

//toml格式
func LoadTomlConf() error {
	_, err := toml.DecodeFile("./conf/conf.toml", &MyConfig)
	if err != nil {
		return err
	}
	//fmt.Println(MyConfig)
	return nil
}

// 初始化配置
func init() {
	//err := LoadConfig()
	//if err != nil {
	//	fmt.Println("配置初始化失败")
	//}
	err := LoadTomlConf()
	if err != nil {
		fmt.Println("toml配置失败")
		fmt.Println(err)
		//os.Exit(-1)
	}
}
