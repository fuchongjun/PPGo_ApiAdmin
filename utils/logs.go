/**********************************************
** @Des: 
** @Author: June
** @Date: 2017-11-19-14:59
** @Last Modified by: June  
** @Last Modified time: 2017-11-19-14:59
***********************************************/
package utils
import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/session"
	"strings"
	"github.com/astaxie/beego"
	"os"
)

//GlobalSessions 全局session
var GlobalSessions *session.Manager

// consoleLogs开发模式下日志
var consoleLogs *logs.BeeLogger

// fileLogs 生产环境下日志
var fileLogs *logs.BeeLogger
var runmode string
type LogLevel int
const (
	EmergencyLog LogLevel = iota
	AlertLog
	CriticalLog
	ErrorLog
	WarningLog
	NoticeLog
	InformationalLog
	DebugLog
	WarnLog
	InfoLog
	TraceLog
)


func init() {
	consoleLogs = logs.NewLogger(1)
	consoleLogs.SetLogger(logs.AdapterConsole)
	fileLogs = logs.NewLogger(10000)
	currentDir, _ := os.Getwd() //当前的目录
	logDir:=beego.AppConfig.String("logdir")
	if logDir!="" {
		creatDir(currentDir+"/"+logDir)
	}else {
		creatDir(currentDir+"/logs")
		logDir="logs"
	}
	//主要的参数如下说明(除 separate 外,均与file相同)：
	//filename 保存的文件名
	//maxlines 每个文件保存的最大行数，默认值 1000000
	//maxsize 每个文件保存的最大尺寸，默认值是 1 << 28, //256 MB
	//	daily 是否按照每天 logrotate，默认是 true
	//maxdays 文件最多保存多少天，默认保存 7 天
	//rotate 是否开启 logrotate，默认是 true
	//level 日志保存的时候的级别，默认是 Trace 级别
	//perm 日志文件权限
	//separate 需要单独写入文件的日志级别,设置后命名类似 test.error.log
	fileLogs.SetLogger(logs.AdapterMultiFile, `{"filename":"`+logDir+`/adminSys.log","separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"],"level":7,"daily":true,"maxdays":20}`)
	runmode = strings.TrimSpace(strings.ToLower(beego.AppConfig.String("runmode")))
	if runmode == "" {
		runmode = "dev"
	}
}

//LogOut 输出日志
// @Title LogOut
// @Param	body		body 	models.AccountAccountTag	true		"body for AccountAccountTag content"
func LogOut(logLevel LogLevel,format string, v interface{}) {
	if runmode == "dev" {
		switch logLevel {
		case EmergencyLog:
			fileLogs.Emergency(format, v)
		case AlertLog:
			fileLogs.Alert(format, v)
		case CriticalLog:
			fileLogs.Critical(format, v)
		case ErrorLog:
			fileLogs.Error(format, v)
		case WarningLog:
			fileLogs.Warning(format, v)
		case NoticeLog:
			fileLogs.Notice(format, v)
		case InformationalLog:
			fileLogs.Informational(format, v)
		case DebugLog:
			fileLogs.Debug(format, v)
		case WarnLog:
			fileLogs.Warn(format, v)
		case InfoLog:
			fileLogs.Info(format, v)
		case TraceLog:
			fileLogs.Trace(format, v)
		default:
			fileLogs.Debug(format, v)
		}
	}
	switch logLevel {
	case EmergencyLog:
		consoleLogs.Emergency(format, v)
	case AlertLog:
		consoleLogs.Alert(format, v)
	case CriticalLog:
		consoleLogs.Critical(format, v)
	case ErrorLog:
		consoleLogs.Error(format, v)
	case WarningLog:
		consoleLogs.Warning(format, v)
	case NoticeLog:
		consoleLogs.Notice(format, v)
	case InformationalLog:
		fileLogs.Informational(format, v)
	case DebugLog:
		consoleLogs.Debug(format, v)
	case WarnLog:
		consoleLogs.Warn(format, v)
	case InfoLog:
		consoleLogs.Info(format, v)
	case TraceLog:
		consoleLogs.Trace(format, v)
	default:
		consoleLogs.Debug(format, v)
	}

}
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
func creatDir(dirname string) (string, error) {
	var path string
	if os.IsPathSeparator('\\') { //前边的判断是否是系统的分隔符
		path = "\\"
	} else {
		path = "/"
	}
	if ok,_:=pathExists(dirname);!ok{
		err := os.MkdirAll(dirname, os.ModePerm) //在当前目录下生成多级目录
		if err != nil {
			return dirname,err
		}
	}
	dirname += path
	return dirname,nil
}