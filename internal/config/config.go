package config

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Address string `yaml:"address" env-required:"true"`
}

// If yaml files des not have cong variables read it from env
type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
	StoragePath string `yaml:"storage_path" env:"STORAGE_PATH" env-required:"true" env-default:"storage/storage.db" `
	HTTPServer  `yaml:"http_server"`
}

// Must load the config else the program terminates here
// in such functions like mustLoad which are necessary to load never return error just throw fatal error and
// stop the program
func MustLoad() *Config {

	var configPath string
	configPath = os.Getenv("CONFIG_PATH")
	fmt.Println(configPath, "123")
	if configPath == "" {
		flags := flag.String("config", "", "path to the config file")
		flag.Parse()
		configPath = *flags
		if configPath == "" {
			log.Fatal(color.RedString("Config Path not provided!"))
		}

	}
	_, err := os.Stat(configPath)

	if os.IsNotExist(err) {
		errMsg := color.RedString("Config Path not found %s", err.Error())
		log.Fatalf("%s", errMsg)
	}

	var cnfg Config
	e := cleanenv.ReadConfig(configPath, &cnfg)

	if e != nil {
		errMsg := color.RedString("Error reading config path %s", err.Error())
		log.Fatalf("%s", errMsg)
	}

	fmt.Println(color.GreenString("Config File Loaded Successfully by Project"))

	return &cnfg
}
