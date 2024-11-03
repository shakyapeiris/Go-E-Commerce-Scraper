package models

import "fmt"

type DBConfig struct {
	Host     string
	Port     int32
	Username string
	Password string
	Database string
}

func (config DBConfig) ConnectionString() string {
	return fmt.Sprintf("host=%s "+
		"port=%d "+
		"user=%s "+
		"password=%s "+
		"dbname=%s",
		config.Host, config.Port, config.Username, config.Password, config.Database)
}
