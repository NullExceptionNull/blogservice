package setting

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

type Setting struct {
	vp *viper.Viper
	C  chan interface{}
}

var sections = make(map[string]interface{})

//var c = make(chan interface{})

func NewSetting() (*Setting, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.SetConfigFile("configs/config.yaml")
	v.AddConfigPath("configs")
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	setting := Setting{vp: v, C: make(chan interface{})}

	go func() {
		setting.reload()
	}()

	return &setting, err
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	if _, ok := sections[k]; !ok {
		sections[k] = v
	}
	return nil
}

func (s *Setting) ReLoadSection() {
	for k, setting := range sections {
		_ = s.ReadSection(k, setting)
	}
}

func (s *Setting) reload() {
	for {
		select {
		case <-s.C:
			logrus.Info("Config is refreshing")
			s.ReLoadSection()
		case <-time.Tick(30 * time.Second):
			logrus.Info("The setting reload goroutine is running ")
		}
	}
}
