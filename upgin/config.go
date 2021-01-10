/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  config
 * @Version: 1.0.0
 * @Date: 2021/1/10 15:27
 */
package upgin

import (
	"apiDemo/upgin/config"
	"apiDemo/upgin/utils"
	"os"
	"path/filepath"
)

const (
	// DEV is for develop
	DEV = "dev"
	// PROD is for production
	PROD = "prod"
)

var (
	UpConfig  *Config
	AppConfig *upGinAppConfig
	AppPath   string
	WorkPath  string
)

type Config struct {
	AppName  string // Application name
	RunMode  string // Running Mode: dev | prod
	ConfPath string // Config app.ini path
	ConfName string // Config name

	WebConfig WebConfig
}

type WebConfig struct {
	TemplateLeft  string
	TemplateRight string
}

type upGinAppConfig struct {
	// 继承Configer配置文件
	innerConfig config.Configer // 配置
}

func init() {
	UpConfig = &Config{
		AppName:  "kcgin",
		RunMode:  "dev",
		ConfPath: "conf",
		ConfName: "app.conf",
		WebConfig: WebConfig{
			TemplateLeft:  "{{",
			TemplateRight: "}}",
		},
	}
	//
	if os.Getenv("KCGIN_APPNAME") != "" {
		UpConfig.AppName = os.Getenv("KCGIN_APPNAME")
	}
	if os.Getenv("KCGIN_RUNMODE") != "" {
		UpConfig.RunMode = os.Getenv("KCGIN_RUNMODE")
	}
	if os.Getenv("KCGIN_CONFPATH") != "" {
		UpConfig.ConfPath = os.Getenv("KCGIN_CONFPATH")
	}
	if os.Getenv("KCGIN_CONFNAME") != "" {
		UpConfig.ConfName = os.Getenv("KCGIN_CONFNAME")
	}
	var err error
	WorkPath, err = os.Getwd()
	if err != nil {
		panic(err)
	}

	// 绝对路径 absolute
	AppPath, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	// 路径拼接
	appConfigPath := filepath.Join(AppPath, UpConfig.ConfPath, UpConfig.AppName)
	if !utils.FileExists(appConfigPath) {
		appConfigPath = filepath.Join(WorkPath, UpConfig.ConfPath, UpConfig.ConfName)
		if !utils.FileExists(appConfigPath) {
			// 在读取不到就去配置文件中读取
			AppConfig = &upGinAppConfig{innerConfig: config.NewFakeConfig()}
			return
		}
	}
	//
	conf, err := newAppConfig("ini", appConfigPath)
	if err != nil {
		return
	}
	if appname := conf.String("appname"); appname != "" {
		UpConfig.AppName = appname
	}
	if runmode := conf.String("runmode"); runmode != "" {
		UpConfig.RunMode = runmode
	}
	AppConfig = conf

}

func newAppConfig(adapterName, appConfigPath string) (*upGinAppConfig, error) {
	// NewConfig adapterName is ini/json/xml/yaml.
	//filename is the config file path.
	ac, err := config.NewConfig(adapterName, appConfigPath)
	if err != nil {
		return nil, err
	}
	return &upGinAppConfig{innerConfig: ac}, nil

}

//type Configer interface 对Configer接口实现
func (u *upGinAppConfig) Set(key, val string) error {
	err := u.innerConfig.Set(UpConfig.RunMode+"::"+key, val)
	if err != nil {
		return err
	}
	return u.innerConfig.Set(key, val)
}
func (u *upGinAppConfig) String(key string) string {
	val := u.innerConfig.String(UpConfig.RunMode + "::" + key)
	if val != "" {
		return val
	}
	return u.innerConfig.String(key)
}

func (u *upGinAppConfig) Strings(key string) []string {
	if v := u.innerConfig.Strings(UpConfig.RunMode + "::" + key); len(v) > 0 {
		return v
	}
	return u.innerConfig.Strings(key)
}

func (u *upGinAppConfig) Int(key string) (int, error) {
	if v, err := u.innerConfig.Int(UpConfig.RunMode + "::" + key); err == nil {
		return v, nil
	}
	return u.innerConfig.Int(key)
}

func (u *upGinAppConfig) Int64(key string) (int64, error) {
	if v, err := u.innerConfig.Int64(UpConfig.RunMode + "::" + key); err == nil {
		return v, nil
	}
	return u.innerConfig.Int64(key)
}

func (u *upGinAppConfig) Bool(key string) (bool, error) {
	if v, err := u.innerConfig.Bool(UpConfig.RunMode + "::" + key); err == nil {
		return v, nil
	}
	return u.innerConfig.Bool(key)
}

func (u *upGinAppConfig) Float(key string) (float64, error) {
	if v, err := u.innerConfig.Float(UpConfig.RunMode + "::" + key); err == nil {
		return v, nil
	}
	return u.innerConfig.Float(key)
}

func (u *upGinAppConfig) DefaultString(key string, defaultVal string) string {
	if v := u.String(key); v != "" {
		return v
	}
	return defaultVal
}

func (u *upGinAppConfig) DefaultStrings(key string, defaultVal []string) []string {
	if v := u.Strings(key); len(v) != 0 {
		return v
	}
	return defaultVal
}

func (u *upGinAppConfig) DefaultInt(key string, defaultVal int) int {
	if v, err := u.Int(key); err == nil {
		return v
	}
	return defaultVal
}

func (u *upGinAppConfig) DefaultInt64(key string, defaultVal int64) int64 {
	if v, err := u.Int64(key); err == nil {
		return v
	}
	return defaultVal
}

func (u *upGinAppConfig) DefaultBool(key string, defaultVal bool) bool {
	if v, err := u.Bool(key); err == nil {
		return v
	}
	return defaultVal
}

func (u *upGinAppConfig) DefaultFloat(key string, defaultVal float64) float64 {
	if v, err := u.Float(key); err == nil {
		return v
	}
	return defaultVal
}

func (u *upGinAppConfig) DIY(key string) (interface{}, error) {
	return u.innerConfig.DIY(key)
}

func (u *upGinAppConfig) GetSection(section string) (map[string]string, error) {
	return u.innerConfig.GetSection(section)
}

func (u *upGinAppConfig) SaveConfigFile(filename string) error {
	return u.innerConfig.SaveConfigFile(filename)
}
