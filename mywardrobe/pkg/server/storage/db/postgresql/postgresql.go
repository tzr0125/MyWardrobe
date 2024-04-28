package db

import (
	"os"
	"path/filepath"
	"example.com/m/v2/pkg/logger"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"
	"fmt"
)

/*
	提供对数据库的连接
*/


/*postgresql:
  host: "localhost"
  port: 
  user:
  password:
  dbname:
*/
type postgresqlConfig struct {
	host string `yaml:"host"`
	port string `yaml:"port"`
	user string `yaml:"user"`
	password string `yaml:"password"`
	dbname string `yaml:"dbname"`
}

func getConfig() (*postgresqlConfig, error) {
	StdLogger := logger.StdLogger()
	dir, err := os.Getwd()
	if err != nil {
		StdLogger.Errorln("Could not get working directory")
		return nil, err
	}
	filePath := filepath.Join(dir, "settings/db.yaml")

	content, err := os.ReadFile(filePath)
	if err != nil {
		StdLogger.Errorln("Could not read settings file:", filePath)
		return nil, err
	}

	var config postgresqlConfig
	
	err = yaml.Unmarshal([]byte(content), &config)
	if err != nil {
		StdLogger.Errorln("Could not unmarshal settings file:", filePath)
		return nil, err
	}

	return &config, nil
}

func InitPostgresql() (*xorm.Engine, error){ 
	config, err := getConfig()
	if err != nil{
		return nil, err
	}
	configStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", 
	config.user, config.password, config.dbname, config.host, config.port)
	engine, err := xorm.NewEngine("postgresql", configStr)
	if err != nil {
		return nil, err
	} 
	return engine, err
}