package litelib

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"github.com/freelifer/litelib/module"
	"github.com/gin-gonic/gin"
)

type LiteLib struct {
	configPath string
	r          *gin.Engine
}

func NewLiteLib() *LiteLib {
	r := gin.Default()
	return &LiteLib{configPath: "config/conf.ini", r: r}
}

func (this *LiteLib) SetConfigPath(path string) {
	this.configPath = path
}

func (this *LiteLib) GetGinEngine() *gin.Engine {
	return this.r
}

func (this *LiteLib) Run() {
	// 配置加载
	c, err := goconfig.LoadConfigFile(this.configPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	name, _ := c.GetValue("", "name")
	version, _ := c.GetValue("", "version")
	port, _ := c.GetValue("", "port")

	fmt.Println("name: " + name)
	fmt.Println("version: " + version)
	fmt.Println("port: " + port)

	// 模块初始化
	list := c.GetSectionList()
	for _, v := range list {
		if v == "DEFAULT" {
			continue
		}

		adapter, err := module.NewModule(v)
		if err != nil {
			fmt.Println(err)
			return
		}
		config, err := c.GetSection(v)

		if err != nil {
			fmt.Println(err)
			return
		}
		adapter.Setup(config)
	}

	// gin运行
	this.r.Run(port)
}
