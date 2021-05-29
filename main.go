package main

import (
	"blog-service/global"
	"blog-service/internal/model"
	"blog-service/internal/routers"
	"blog-service/pkg/setting"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

func main() {
	engine := routers.NewRouter()

	server := &http.Server{
		Addr:              ":" + global.ServerSetting.HttpPort,
		Handler:           engine,
		ReadHeaderTimeout: global.ServerSetting.ReadTimeout,
		WriteTimeout:      global.ServerSetting.WriteTimeout,
		MaxHeaderBytes:    1 << 20,
	}

	_ = server.ListenAndServe()
}

func init() {

	setUpLog()

	c, err := settingUp()

	err = setUpDbEngine()

	SetUpNacos(c)

	if err != nil {
		log.Fatal("init config error")
	}
	log.Info("----------------------------------Init OK----------------------------------")
}

func settingUp() (chan interface{}, error) {
	setting, err := setting.NewSetting()
	if err != nil {
		return nil, err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return nil, err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return nil, err
	}
	err = setting.ReadSection("Database", &global.DataBaseSetting)
	if err != nil {
		return nil, err
	}

	err = setting.ReadSection("Nacos", &global.NacosSetting)
	if err != nil {
		return nil, err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return setting.C, nil
}

func setUpDbEngine() error {
	db, err := model.NewEngine(global.DataBaseSetting)
	if err != nil {
		panic(err)
	}
	global.DBEngine = db

	return nil
}

func setUpLog() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func SetUpNacos(c chan interface{}) {
	_ = setting.NewNacos(c, global.NacosSetting.NamespaceId,
		global.NacosSetting.Group,
		global.NacosSetting.DataId,
		global.NacosSetting.IpAddr,
		global.NacosSetting.Port)
	//global.Nacos = nacos
}
