package main

import (
	"PPGo_ApiAdmin/models"
	_ "PPGo_ApiAdmin/routers"

	"github.com/astaxie/beego"
)

func main() {
	//fmt.Println(libs.Md5([]byte("admin"+"kmcB")))
	models.Init()
	beego.Run()
}
