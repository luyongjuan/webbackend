package cfg

import (
	md "webbackend/pkg/datamodle"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"strings"
	"time"
)

var GConfig *md.Config

func Load(path string) {
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil {
		log.Panic().Msgf("Error reading config.yaml: %s", err)
	}
	GConfig = &md.Config{}
	if err := viper.Unmarshal(GConfig); err != nil {
		log.Panic().Msgf("Error unmashal config.yaml: %s", err)
	}

	logConfig := GConfig.Logger
	lumberjackOutput := &lumberjack.Logger{
		Filename:   logConfig.Filepath,
		MaxSize:    logConfig.MaxSize, // megabytes
		MaxBackups: logConfig.MaxBackups,
		MaxAge:     logConfig.MaxAge, //days
	}

	output := io.MultiWriter(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}, lumberjackOutput)
	log.Logger = zerolog.New(output).With().Timestamp().Logger().Level(logConfig.Level)

	log.Debug().Msgf("logger: %v", GConfig.Logger)
	log.Debug().Msgf("GateWay: %v", GConfig.SnKey)
}
