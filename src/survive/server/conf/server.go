package conf

import (
	"encoding/json"
	"github.com/name5566/leaf/log"
	"io/ioutil"
)

//定义server.json对应的配置项对象
var Server struct {
	LogLevel   string //日志等级
	LogPath    string //日志路径
	WSAddr     string //websocket服务地址
	TCPAddr    string //tcp服务地址
	MaxConnNum int //最大连接数
}
var DB struct{
	DBUrl        string //数据库地址
	DBMaxConnNum int    //数据库最大连接数
}
func loadServerConfig(){
	data, err := ioutil.ReadFile("conf/server.json")
	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(data, &Server)
	if err != nil {
		log.Fatal("%v", err)
	}
}

func loadDBConfig(){
	data, err := ioutil.ReadFile("conf/db.json")
	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(data, &DB)
	if err != nil {
		log.Fatal("%v", err)
	}
}
func init() {
	loadServerConfig()
	loadDBConfig()
}
