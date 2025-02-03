# 📝 Task Tracker CLI

A powerful command-line task management tool built with Go, following the clean architecture principles. This project was implemented based on the requirements from [roadmap.sh Task Tracker project](https://roadmap.sh/projects/task-tracker).

## ✨ Features

- 📌 Add new tasks
- 📝 Update task descriptions
- 🗑️ Delete tasks
- 🔄 Change task status (todo/in-progress/done)
- 📋 List tasks with status filtering

## 🚀 Getting Started

### Prerequisites

- Go 1.21 or higher

### Installation

```bash
go install github.com/idmaksim/task-tracker-cli/cmd/task-cli@latest
```

## 💡 Usage

### Managing Tasks

```bash
# Add a new task
task-cli add "Buy groceries"

# Update task description
task-cli update 1 "Buy groceries and cook dinner"

# Delete a task
task-cli delete 1
```

### Task Status Management

```bash
# Mark task as in progress
task-cli mark-in-progress 1

# Mark task as done
task-cli mark-done 1
```

### Listing Tasks

```bash
# List all tasks
task-cli list

# List tasks by status
task-cli list todo
task-cli list in-progress
task-cli list done
```

## 🏗️ Project Structure

```
.
├── cmd/
│   └── task-cli/          # Application entry point
├── internal/
│   ├── domain/            # Business logic and entities
│   │   ├── models/        # Domain models
│   │   └── repositories/  # Repository interfaces
│   ├── usecases/         # Application use cases
│   ├── infrastructure/   # External implementations
│   │   └── storage/      # JSON storage implementation
│   └── delivery/         # CLI interface
│       └── cli/
│           ├── commands/ # CLI commands
│           └── handlers/ # Command handlers
└── pkg/                  # Public packages
    └── constants/       # Shared constants
```

## 🎯 Task Properties

Each task includes:

- 🔑 Unique ID
- 📄 Description
- 📊 Status (todo/in-progress/done)
- 📅 Creation date
- 🔄 Last update date

## 🛠️ Technical Details

- **Clean Architecture**: Separation of concerns with clear layer boundaries
- **Concurrent Safe**: JSON file storage with mutex protection
- **Error Handling**: Graceful error handling and user-friendly messages
- **CLI Framework**: Built with Cobra for robust command-line interface
- **Time Tracking**: Automatic creation and update time management

## ✨ Acknowledgments

- Project structure inspired by [roadmap.sh Task Tracker project](https://roadmap.sh/projects/task-tracker)
- Built with Go and ❤️
