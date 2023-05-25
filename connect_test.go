package redigo

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

type EnvConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database int    `json:"database"`
	Password string `json:"password"`
}

func TestNewRedis(t *testing.T) {
	data, err := ioutil.ReadFile("./env.json")
	if err != nil {
		t.Fatal("配置文件读取失败", err)
	}
	config := &EnvConfig{}
	if err = json.Unmarshal(data, &config); err != nil {
		t.Fatal("配置文件解析失败", err)
	}
	NewRedis(WithHost(config.Host), WithPort(config.Port), WithDatabase(config.Database), WithPassword(config.Password))
}
