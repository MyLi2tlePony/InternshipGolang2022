package logger

import "github.com/spf13/viper"

type config struct {
	Level string
}

func New(configPath string) (*config, error) {
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return &config{
		Level: viper.GetString("logger.level"),
	}, nil
}

func (l config) GetLevel() string {
	return l.Level
}
