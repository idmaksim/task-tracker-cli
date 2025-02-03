package main

import (
	"fmt"
	"os"

	"github.com/idmaksim/task-tracker-cli/internal/delivery/cli/commands"
	"github.com/idmaksim/task-tracker-cli/internal/delivery/cli/handlers"
	"github.com/idmaksim/task-tracker-cli/internal/infrastructure/storage"
	"github.com/idmaksim/task-tracker-cli/internal/usecases"
)

func main() {
	// Setup error handling
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	// Initialize storage
	store := storage.NewJSONStorage("data.json")
	if err := store.Init(); err != nil {
		panic(err)
	}

	// Initialize service and handler
	service := usecases.NewTaskService(store)
	handler := handlers.NewTaskHandler(service)
	cmd := commands.NewCommands(handler)

	// Setup and run CLI
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
