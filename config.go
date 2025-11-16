package main

import (
	"os"
	"github.com/spf13/viper"
)

type Config struct {
	Loglevel string `mapstructure:"loglevel"`

	AccessKey string `mapstructure:"access_key"`
	SecretKey string `mapstructure:"secret_key"`
	BucketName string `mapstructure:"bucket_name"`
	Region string `mapstructure:"region"`
	Endpoint string `mapstructure:"endpoint"`
	
	BindAddress string `mapstructure:"bind_address"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config_sample")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetDefault("loglevel", "INFO")

	if err := viper.ReadInConfig(); err != nil {
		Logoutput("unable to read config_sample", "error_force")
		return nil, err
	}
	viper.AutomaticEnv()

	viper.BindEnv("loglevel","loglevel")
	viper.BindEnv("access_key","access_key")
	viper.BindEnv("secret_key","secret_key")
	viper.BindEnv("bucket_name","bucket_name")
	viper.BindEnv("region","region")
	viper.BindEnv("endpoint","endpoint")
	viper.BindEnv("bind_address","bind_address")

	if _, err := os.Stat("conf/config.yaml"); err == nil {
		viper.SetConfigName("config")
		viper.AddConfigPath("conf")
		viper.MergeInConfig()
		Logoutput("Using config file: conf/config.yaml", "info_force")
	}else{
		Logoutput("Using Environment Variables", "info_force")
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		Logoutput("unable to unmarshal config", "error_force")
		return nil, err
	}
	return &config, nil
}
