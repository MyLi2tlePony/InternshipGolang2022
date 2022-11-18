package migration

import "github.com/spf13/viper"

type config struct {
	Path string
}

func New(configPath string) (*config, error) {
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return &config{
		Path: viper.GetString("migration.path"),
	}, nil
}

func (l config) GetPath() string {
	return l.Path
}
