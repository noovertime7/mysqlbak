package modles

import (
	"github.com/noovertime7/mysqlbak/pkg/log"
	"github.com/spf13/viper"
	_ "gopkg.in/yaml.v2"
	"os"
)

func init() {
	path, err := os.Getwd()
	if err != nil {
		log.Logger.Error(err)
		return
	}
	config := viper.New()
	config.SetConfigType("yaml")           //设置文件的类型
	config.AddConfigPath(path + "/config") //设置读取的文件路径
	config.SetConfigName("config")         //设置读取的文件名
	if err := config.ReadInConfig(); err != nil {
		log.Logger.Error("读取配置文件失败", err)
		return
	}
	if err = config.Unmarshal(SystemConf); err != nil {
		log.Logger.Error("配置文件解析失败", err)
		return
	}
}

var SystemConf = new(System)

type Database struct {
	Host        string `yaml:"host"`
	Password    string `yaml:"password"`
	User        string `yaml:"user"`
	Port        string `yaml:"port"`
	DBName      string
	BackupCycle string
	KeepNumber  int
	isAllDBBak  bool `yaml:"isAllDBBak"`
	DingConf    DingConf
	OssConf     OssConf
}

type System struct {
	BakConfig
}

type BakConfig struct {
	Database []Database `yaml:"Database"`
}
type DingConf struct {
	IsDingSend bool   `yaml:"isDingSend"`
	Webhook    string `yaml:"webhook"`
}

type OssConf struct {
	Type       string `yaml:"type"`
	IsOssSave  bool   `yaml:"isOssSave"`
	Endpoint   string `yaml:"endpoint"`
	Accesskey  string `yaml:"accesskey"`
	Secretkey  string `yaml:"secretkey"`
	BucketName string `yaml:"bucketname"`
	Directory  string `yaml:"directory"`
}
