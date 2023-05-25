package redigo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestSet(t *testing.T) {
	fmt.Println("----------------------------------------------")
	fmt.Println(Set("name", 123))
}

func init() {
	data, err := ioutil.ReadFile("./env.json")
	if err != nil {
		panic("配置文件读取失败: " + err.Error())
	}
	config := &EnvConfig{}
	if err = json.Unmarshal(data, &config); err != nil {
		panic("配置文件解析失败: " + err.Error())
	}
	NewRedis(WithHost(config.Host), WithPort(config.Port), WithDatabase(config.Database), WithPassword(config.Password))
}
