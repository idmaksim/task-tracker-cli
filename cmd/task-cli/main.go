package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/idmaksim/task-tracker-cli/internal/delivery/cli/commands"
	"github.com/idmaksim/task-tracker-cli/internal/delivery/cli/handlers"
	"github.com/idmaksim/task-tracker-cli/internal/infrastructure/storage"
	"github.com/idmaksim/task-tracker-cli/internal/usecases"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	dataDir := filepath.Join(homeDir, ".task-tracker")
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		panic(err)
	}

	dataPath := filepath.Join(dataDir, "data.json")
	store := storage.NewJSONStorage(dataPath)
	if err := store.Init(); err != nil {
		panic(err)
	}

	service := usecases.NewTaskService(store)
	handler := handlers.NewTaskHandler(service)
	cmd := commands.NewCommands(handler)

	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
