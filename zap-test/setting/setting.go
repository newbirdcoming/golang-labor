package setting

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(ZapConfig)

// ZapConfig
// @Date: 2023-01-09 16:37:05
// @Description: zap日志配置结构体
type ZapConfig struct {
	Prefix     string         `yaml:"prefix" mapstructure:"prefix"`
	TimeFormat string         `yaml:"timeFormat" mapstructure:"timeFormat"`
	Level      string         `yaml:"level" mapstructure:"level"`
	Caller     bool           `yaml:"caller" mapstructure:"caller"`
	StackTrace bool           `yaml:"stackTrace" mapstructure:"stackTrace"`
	Writer     string         `yaml:"writer" mapstructure:"writer"`
	Encode     string         `yaml:"encode" mapstructure:"encode"`
	LogFile    *LogFileConfig `yaml:"logFile" mapstructure:"logFile"`
}

// LogFileConfig
// @Date: 2023-01-09 16:38:45
// @Description: 日志文件配置结构体
type LogFileConfig struct {
	MaxSize  int      `yaml:"maxSize" mapstructure:"maxSize"`
	BackUps  int      `yaml:"backups" mapstructure:"backups"`
	Compress bool     `yaml:"compress" mapstructure:"compress"`
	Output   []string `yaml:"output" mapstructure:"output"`
	Errput   []string `yaml:"errput" mapstructure:"errput"`
}

// 业务逻辑函数
func ReadConfigFile() error {
	viper.SetConfigName("config.yml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
	}
	// 监听配置文件的变化
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		}
	})
	viper.WatchConfig()
	return nil
}
