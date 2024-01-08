package configuration

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

var Configuration = struct {
	GinMode                string        `yaml:"ginMode"`
	ChannelUpdateFrequency time.Duration `yaml:"channelUpdateFrequency"`
	ChannelTestFrequency   time.Duration `yaml:"channelTestFrequency"`
	BatchUpdateEnabled     bool          `yaml:"batchUpdateEnabled"`
	PORT                   int           `yaml:"port"`
	// mysql
	SqlDsn string `yaml:"sqlDsn"`

	SessionSecret string `yaml:"sessionSecret"`
	SqlitePath    string `yaml:"sqlitePath"`
	// redis
	RedisConnString string `yaml:"redisConnString"`
	SyncFrequency   string `yaml:"syncFrequency"`

	FrontendBaseUrl string `yaml:"frontendBaseUrl"`
}{}

func init() {
	yamlFile, err := os.ReadFile("./configuration/demo.yaml")
	if err != nil {
		fmt.Println(err)
	}
	if err = yaml.Unmarshal(yamlFile, &Configuration); err != nil {
		fmt.Println("unmarshal", err)
	}
}
