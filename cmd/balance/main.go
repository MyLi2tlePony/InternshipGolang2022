package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/MyLi2tlePony/AvitoInternshipGolang2022/internal/balance"
	databaseConfig "github.com/MyLi2tlePony/AvitoInternshipGolang2022/internal/config/database"
	loggerConfig "github.com/MyLi2tlePony/AvitoInternshipGolang2022/internal/config/logger"
	serverConfig "github.com/MyLi2tlePony/AvitoInternshipGolang2022/internal/config/server"
	"github.com/MyLi2tlePony/AvitoInternshipGolang2022/internal/logger"
	internalhttp "github.com/MyLi2tlePony/AvitoInternshipGolang2022/internal/server/http"
	storage "github.com/MyLi2tlePony/AvitoInternshipGolang2022/internal/storage/postgres"
	"net/http"
	"os"
	"os/signal"
	"path"
	"syscall"
)

var configPath string

func init() {
	defaultConfigPath := path.Join("configs", "balance", "config.toml")
	flag.StringVar(&configPath, "config", defaultConfigPath, "Path to configuration file")
}

func main() {
	flag.Parse()

	dbConfig, err := databaseConfig.New(configPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	dbConnectionString := dbConfig.GetConnectionString()
	sqlStorage := balance.Storage(storage.New(dbConnectionString))

	err = sqlStorage.Connect(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	logConfig, err := loggerConfig.New(configPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	srvConfig, err := serverConfig.New(configPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	srv := internalhttp.NewServer(logger.New(logConfig), balance.New(sqlStorage), srvConfig.GetHostPort())

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()

		if err := srv.Stop(); err != nil {
			fmt.Println("failed to stop serv: " + err.Error())
		}
	}()

	fmt.Println("app is running...")

	err = srv.Start()
	if !errors.Is(err, http.ErrServerClosed) && err != nil {
		fmt.Println("failed to start serv: " + err.Error())
		cancel()
		os.Exit(1) //nolint:gocritic
	}

	fmt.Println("serv closed")
}
