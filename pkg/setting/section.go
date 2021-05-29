package setting

import "time"

type ServerSettingS struct {
	AppName      string
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettingS struct {
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string

	UploadSavePath       string
	UploadServerUrl      string
	UploadImageMaxSize   int
	UploadImageAllowExts []string
}

type DataBaseSettingS struct {
	DbType       string
	UserName     string
	PassWord     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type NacosSettingS struct {
	IpAddr      string
	Port        uint64
	NamespaceId string
	Group       string
	DataId      string
}
