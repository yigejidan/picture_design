package common

import "github.com/spf13/viper"

type Setting struct {
	vp *viper.Viper
}

var sections = make(map[string]interface{})

// NewSetting 读取配置
func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("config")
	vp.SetConfigType("yaml")
	if err := vp.ReadInConfig(); err != nil {
		return nil, err
	}
	s := &Setting{vp}
	return s, nil
}

// ReadSection 读取指定的一段
func (s *Setting) ReadSection(k string, v interface{}) error {
	if err := s.vp.UnmarshalKey(k, v); err != nil {
		return err
	}
	if _, ok := sections[k]; !ok {
		sections[k] = v
	}
	return nil
}

// ReloadAllSection 重新加载
func (s *Setting) ReloadAllSection() error {
	for k, v := range sections {
		if err := s.ReadSection(k, v); err != nil {
			return err
		}
	}
	return nil
}
