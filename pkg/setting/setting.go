package setting

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
}

var Conf = new(TotalConfig)

type TotalConfig struct {
	*BotConfig `mapstructure:"bot"`
}

type BotConfig struct {
	*DingTalkConfig `mapstructure:"dingtalk"`
	*WxChatConfig   `mapstructure:"wxchat"`
}

type DingTalkConfig struct {
	Enable          bool     `mapstructure:"enable"`
	CallBackUrl     string   `mapstructure:"call_back_url"`
	Secrets         []string `mapstructure:"secrets"`
	Keywords        []string `mapstructure:"keywords"`
	Token           string   `mapstructure:"token"`
	TemplateFile    string   `mapstructure:"template_file"`
	MessageTemplate string   `mapstructure:"message_template"`
}

type WxChatConfig struct {
	Enable          bool     `mapstructure:"enable"`
	CallBackUrl     string   `mapstructure:"call_back_url"`
	Secrets         []string `mapstructure:"secrets"`
	Keywords        []string `mapstructure:"keywords"`
	Token           string   `mapstructure:"token"`
	TemplateFile    string   `mapstructure:"template_file"`
	MessageTemplate string   `mapstructure:"message_template"`
}

func init() {
	viper.SetConfigFile("conf/conf.yaml")

	// 实时读取配置文件
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	// 获取环境变量
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("viper.ReadInConfig failed, err:%v\n", err)
		return
	}

	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
	}

	// fmt.Println(Conf.BotConfig)
}
