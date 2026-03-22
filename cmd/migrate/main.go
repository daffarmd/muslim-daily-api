package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"api-go-test/app"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("usage: go run ./cmd/migrate [up|down|version|force <version>|create <name>]")
	}

	migrationsDir := "migrations"
	command := os.Args[1]
	if command == "create" {
		if len(os.Args) < 3 {
			log.Fatal("usage: go run ./cmd/migrate create <name>")
		}

		must(createMigrationFiles(migrationsDir, os.Args[2]))
		return
	}

	config := app.LoadConfig()
	if strings.TrimSpace(config.DBDSN) == "" {
		log.Fatal("DB_DSN is required")
	}

	migrationPath := "file://" + filepath.ToSlash(migrationsDir)
	migrator, err := migrate.New(migrationPath, config.DBDSN)
	if err != nil {
		log.Fatalf("failed to create migrator: %v", err)
	}
	defer func() {
		sourceErr, databaseErr := migrator.Close()
		if sourceErr != nil {
			log.Printf("failed to close migration source: %v", sourceErr)
		}
		if databaseErr != nil {
			log.Printf("failed to close migration database: %v", databaseErr)
		}
	}()

	switch command {
	case "up":
		err = migrator.Up()
		if errors.Is(err, migrate.ErrNoChange) {
			log.Println("no change")
			return
		}
		must(err)
		log.Println("migration up completed")
	case "down":
		err = migrator.Steps(-1)
		if errors.Is(err, migrate.ErrNoChange) {
			log.Println("no change")
			return
		}
		must(err)
		log.Println("migration down completed")
	case "version":
		version, dirty, err := migrator.Version()
		if errors.Is(err, migrate.ErrNilVersion) {
			log.Println("migration version: none")
			return
		}
		must(err)
		log.Printf("migration version=%d dirty=%t", version, dirty)
	case "force":
		if len(os.Args) < 3 {
			log.Fatal("usage: go run ./cmd/migrate force <version>")
		}

		var version int
		_, err := fmt.Sscanf(os.Args[2], "%d", &version)
		must(err)

		must(migrator.Force(version))
		log.Printf("migration forced to version=%d", version)
	default:
		log.Fatalf("unknown command %q", command)
	}
}

func createMigrationFiles(migrationsDir, name string) error {
	safeName := strings.TrimSpace(strings.ToLower(name))
	safeName = strings.ReplaceAll(safeName, " ", "_")
	safeName = strings.ReplaceAll(safeName, "-", "_")
	if safeName == "" {
		return errors.New("migration name cannot be empty")
	}

	version := time.Now().UTC().Format("20060102150405")
	upPath := filepath.Join(migrationsDir, version+"_"+safeName+".up.sql")
	downPath := filepath.Join(migrationsDir, version+"_"+safeName+".down.sql")

	if err := os.WriteFile(upPath, []byte("-- Write your UP migration here.\n"), 0o644); err != nil {
		return err
	}

	if err := os.WriteFile(downPath, []byte("-- Write your DOWN migration here.\n"), 0o644); err != nil {
		return err
	}

	log.Printf("created migration files:\n- %s\n- %s", upPath, downPath)
	return nil
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
