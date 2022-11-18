package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/jackc/pgx/v4"
	"path"
	"time"

	databaseConfig "github.com/MyLi2tlePony/AvitoInternshipGolang2022/internal/config/database"
	loggerConfig "github.com/MyLi2tlePony/AvitoInternshipGolang2022/internal/config/logger"
	migrationConfig "github.com/MyLi2tlePony/AvitoInternshipGolang2022/internal/config/migration"
	"github.com/MyLi2tlePony/AvitoInternshipGolang2022/internal/logger"
	"github.com/MyLi2tlePony/AvitoInternshipGolang2022/internal/migration"
)

var configPath string

func init() {
	defaultConfigPath := path.Join("configs", "migration", "config.toml")
	flag.StringVar(&configPath, "config", defaultConfigPath, "Path to configuration file")
}

func main() {
	flag.Parse()

	logConfig, err := loggerConfig.New(configPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	l := logger.New(logConfig)

	dbConf, err := databaseConfig.New(configPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	connString := fmt.Sprintf("%s://%s:%s@%s:%s/%s",
		dbConf.Prefix, dbConf.UserName, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.DatabaseName)

	migratorConfig, err := migrationConfig.New(configPath)
	if err != nil {
		return
	}

	TryDBConnect(connString)

	migrator := migration.New(l)

	migrator.Up(migratorConfig.GetPath(), connString)
	migrator.Status(connString)
}

func TryDBConnect(configPath string) {
	ctx := context.Background()

	for i := 0; i < 10; i++ {
		_, err := pgx.Connect(ctx, configPath)
		if err != nil {
			time.Sleep(time.Second * 2)
		} else {
			break
		}
	}
}
