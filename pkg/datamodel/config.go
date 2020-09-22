package datamodel

import "github.com/rs/zerolog"

//config
type Logger struct {
	Level      zerolog.Level `yaml:"Level"`
	Filepath   string        `yaml:"Filepath"`
	MaxSize    int           `yaml:"MaxSize"`
	MaxAge     int           `yaml:"MaxAge"`
	MaxBackups int           `yaml:"MaxBackups"`
}

type Config struct {
	Logger Logger `yaml:"Logger"`
	SnKey  SnKey  `yaml:"SnKey"`
	Data   Data   `yaml:"Data"`
}