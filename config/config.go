package config

import (
	"fmt"
	"github.com/ethMatch/proxy/types"
	"gopkg.in/yaml.v2"
	"os"
)

var (
	ENV          *types.Config
	DEFAULT_PATH = defaultPath()
)

func defaultPath() (path string) {
	path = "/config/config.yml"
	home, _ := os.Getwd()
	if home != "" {
		path = home + path
	}
	return
}
func init() {
	ENV = &types.Config{}
}

func InitConfig(path string) (err error) {
	if path == "" {
		path = DEFAULT_PATH
	}
	fmt.Println(path)
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	d := yaml.NewDecoder(file)
	if err = d.Decode(&ENV); err != nil {
		return err
	}
	fmt.Println(err)
	return
}
