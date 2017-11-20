/**********************************************
** @Des: This file ...
** @Author: haodaquan
** @Date:   2017-09-08 00:18:02
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-09-16 17:26:48
***********************************************/

package models

import (
	"net/url"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func Init() {
	dbhost := beego.AppConfig.String("db.host")
	dbport := beego.AppConfig.String("db.port")
	dbuser := beego.AppConfig.String("db.user")
	dbpassword := beego.AppConfig.String("db.password")
	dbname := beego.AppConfig.String("db.name")
	timezone := beego.AppConfig.String("db.timezone")
	if dbport == "" {
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	// fmt.Println(dsn)

	if timezone != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(timezone)
	}
	// 参数1:数据库的别名，用来在 ORM 中切换数据库使用
	// 参数2:driverName
	// 参数3:对应的链接字符串
	// 参数4(可选)  设置最大空闲连接
	// 参数5(可选)  设置最大数据库连接 (go >= 1.2)
	orm.RegisterDataBase("default", "mysql", dsn)
	orm.RegisterModel(new(Auth), new(Role), new(RoleAuth), new(Admin),
		new(Group), new(Env), new(Code), new(Api), new(ApiDetail), new(ApiParam))

	//后一个使用true会带上很多打印信息，数据库操作和建表操作的；第二个为true代表强制创建表
	//orm.RunSyncdb("default", false, false)

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}
//自定义数据库前缀和表名
func TableName(name string) string {
	return beego.AppConfig.String("db.prefix") + name
}
