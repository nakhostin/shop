package config

import (
	"io/ioutil"
	"log"
	"os"

	validator "github.com/asaskevich/govalidator"
	"gopkg.in/yaml.v2"
)

var (
	CFG Config = Config{
		Version:   "v1.0.0",
		App:       "event",
		MongoDB:   "",
		MongoHost: "",
		Debug:     true,
	}

	cfgPath = "configs/config.yaml"
)

// ConfigStruct struct config
type Config struct {
	Version string `yaml:"version" valid:"required"`
	App     string `yaml:"app" valid:"required"`

	MongoDB            string `yaml:"mongodb" valid:"required"`
	MongoHost          string `yaml:"mongo_host" valid:"required"`
	MongoUser          string `yaml:"mongo_user"`
	MongoPass          string `yaml:"mongo_pass"`
	MongoAuthMechanism string `yaml:"mongo_auth_mechanism"`

	Debug          bool     `yaml:"debug"`
	ListenerPort   string   `yaml:"listener_port" valid:"required"`
	AllowMethods   []string `yaml:"allow_methods"`
	StaticFilePath string   `yaml:"static_file_path"`

	Cmd []string `yaml:"cmd"`
}

func InitConfig() {
	CFG = *loadConfig(cfgPath)
}

// LoadConfig load config
func loadConfig(path string) *Config {
	config := new(Config)
	if _, err := os.Stat(path); err == nil {
		f, _ := ioutil.ReadFile(path)
		if err := yaml.Unmarshal(f, &config); err != nil {
			panic("config error : " + err.Error())
		}
	}
	if result, _ := validator.ValidateStruct(config); !result {
		log.Fatal("Loading config required failed")
	}
	return config
}
