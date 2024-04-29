package db

import (
	"fmt"
	"os"
	"path/filepath"

	"example.com/m/v2/pkg/logger"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"
)

/*
	提供对数据库的连接
*/

/*
postgresql:

	host: "localhost"
	port:
	user:
	password:
	dbname:
*/
type PostgresqlConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
}

func getConfig() (*PostgresqlConfig, error) {
	StdLogger := logger.StdLogger()
	dir, err := os.Getwd()
	if err != nil {
		StdLogger.Errorln("Could not get working directory")
		StdLogger.Errorln("detail:", err)
		return nil, err
	}
	filePath := filepath.Join(dir, "settings/postgresql.yaml")

	content, err := os.ReadFile(filePath)
	if err != nil {
		StdLogger.Errorln("Could not read settings file:", filePath)
		StdLogger.Errorln("detail:", err)
		return nil, err
	}

	var config PostgresqlConfig

	err = yaml.Unmarshal([]byte(content), &config)
	if err != nil {
		StdLogger.Errorln("Could not unmarshal settings file:", filePath)
		StdLogger.Errorln("detail:", err)
		return nil, err
	}

	return &config, nil
}

func InitPostgresql() (*xorm.Engine, error) {
	StdLogger := logger.StdLogger()
	config, err := getConfig()
	if err != nil {
		return nil, err
	}
	fmt.Println(config)
	configStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		config.User, config.Password, config.Dbname, config.Host, config.Port)
	engine, err := xorm.NewEngine("postgres", configStr)
	if err != nil {
		StdLogger.Errorln("Could not connect database: postgresql")
		StdLogger.Errorln("detail:", err)
		return nil, err
	}
	return engine, err
}
