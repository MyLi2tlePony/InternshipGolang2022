package migration

import (
	"context"
	"errors"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"

	"github.com/MyLi2tlePony/sql-migrator/pkg/migration"
)

type Logger interface {
	Error(string)
	Info(string)
}

type migrator struct {
	logger Logger
}

type localMigration struct {
	Version int
	Name    string

	Up   string
	Down string
}

var (
	ErrInvalidMigrationName = errors.New("invalid migration name")

	regGetVersion       = regexp.MustCompile(`^\d+`)
	regGetUpMigration   = regexp.MustCompile(`^.+_up\.sql$`)
	regGetDownMigration = regexp.MustCompile(`^.+_down\.sql$`)
)

func New(logger Logger) *migrator {
	return &migrator{
		logger: logger,
	}
}

func (m *migrator) Up(filePath, connString string) {
	migrator := migration.New(connString, m.logger)
	migrations, err := getMigrations(filePath)
	if err != nil {
		m.logger.Error(err.Error())
		return
	}

	for i := 1; ; i++ {
		if _, ok := migrations[i]; !ok {
			break
		}

		migrator.Create(migrations[i].Name, migrations[i].Up, migrations[i].Down)
	}

	ctx := context.Background()
	if err = migrator.Connect(ctx); err != nil {
		return
	}

	if err = migrator.Up(ctx); err != nil {
		return
	}

	if err = migrator.Close(ctx); err != nil {
		return
	}
}

func (m *migrator) Status(connString string) {
	migrator := migration.New(connString, m.logger)
	ctx := context.Background()
	var err error

	if err = migrator.Connect(ctx); err != nil {
		return
	}

	if err = migrator.Status(ctx); err != nil {
		return
	}

	if err = migrator.Close(ctx); err != nil {
		return
	}
}

func getMigrations(filePath string) (map[int]*localMigration, error) {
	files, err := os.ReadDir(filePath)
	if err != nil {
		return nil, err
	}

	migrations := make(map[int]*localMigration)

	for _, file := range files {
		strVersion := regGetVersion.FindString(file.Name())

		if strVersion != "" {
			version, err := strconv.Atoi(strVersion)
			if err != nil {
				return nil, err
			}

			parts := strings.Split(file.Name(), "_")
			if len(parts) != 3 {
				return nil, ErrInvalidMigrationName
			}

			sql, err := os.ReadFile(path.Join(filePath, file.Name()))
			if err != nil {
				return nil, err
			}

			if regGetUpMigration.MatchString(file.Name()) {
				if _, ok := migrations[version]; ok {
					migrations[version].Up = string(sql)
				} else {
					migrations[version] = &localMigration{
						Version: version,
						Name:    parts[1],
						Up:      string(sql),
					}
				}
			} else if regGetDownMigration.MatchString(file.Name()) {
				if _, ok := migrations[version]; ok {
					migrations[version].Down = string(sql)
				} else {
					migrations[version] = &localMigration{
						Version: version,
						Name:    parts[1],
						Down:    string(sql),
					}
				}
			} else {
				return nil, ErrInvalidMigrationName
			}
		}
	}

	return migrations, nil
}
