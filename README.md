# ✅ Task Manager

A lightweight **task management service** written in **Go**, supporting both **CLI mode** and a **REST API**.  
Built to demonstrate clean architecture, proper separation of layers, and simple task automation tools.

---

## 🧠 Overview

Task Manager helps organize personal or development tasks with basic operations such as **create**, **list**, **update**, and **delete**.

The project demonstrates:

- Practical CLI application development  
- REST API basics  
- Layered architecture (handler → service → repository)  
- File-based storage (CSV / JSON)  
- Error handling & validation  

---

## 🛠️ Tech Stack

- **Go 1.22+**
- **net/http**
- **gorilla/mux**
- **CSV / JSON storage**
- **Makefile**
- **Docker (optional)**

---

## ✨ Features

### ✅ Implemented
- CLI and REST API modes  
- CRUD operations for tasks  
- CSV/JSON file storage  
- Logging system  
- Task priority (low / medium / high)  
- Task status (todo / in_progress / done)  

### 🚧 In Progress
- REST API search & filtering  
- PostgreSQL persistent storage  

### 🔮 Planned
- Authentication system  
- Task reminders & deadlines  
- TUI (terminal UI)  
- JSON import/export  
- Swagger documentation  

---

# 📦 Setup & Installation

## 1. Clone repository
```bash
git clone https://github.com/adal4ik/task-manager.git
cd task-manager
```
## 2. Run
```bash
make up
```
## 3. Down
```bash
make down
```
---
# 🧪 Usage Examples (REST API)

## Create Task
```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{
        "title": "Write documentation",
        "priority": "high"
      }'
```
## List Tasks
```bash
curl http://localhost:8080/tasks
```
## Update Task
```bash
curl -X PUT http://localhost:8080/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{ "status": "in_progress" }'
```
## Delete Task
```bash
curl -X DELETE http://localhost:8080/tasks/1
```
---
# Architecture
```bash
.
├── cmd
│   └── task-manager
│       └── main.go
├── docker-compose.yml
├── Dockerfile
├── ERD.png
├── go.mod
├── go.sum
├── init.sql
├── internal
│   ├── adapters
│   │   ├── driven
│   │   │   └── database
│   │   │       ├── conn.go
│   │   │       ├── conn_test.go
│   │   │       └── repository
│   │   │           ├── register_repo.go
│   │   │           └── repository.go
│   │   └── driver
│   │       └── http
│   │           ├── handlers
│   │           │   ├── handlers.go
│   │           │   └── register_handler.go
│   │           ├── middleware
│   │           │   └── middleware.go
│   │           └── router.go
│   ├── config
│   │   └── config.go
│   ├── core
│   │   ├── domain
│   │   │   ├── tasks.go
│   │   │   ├── users.go
│   │   │   ├── workspace.go
│   │   │   └── workspace_member.go
│   │   ├── interfaces
│   │   │   ├── driven
│   │   │   │   ├── driveninterface.go
│   │   │   │   └── register_interface.go
│   │   │   └── driver
│   │   │       ├── driverinterface.go
│   │   │       └── register_interface.go
│   │   └── service
│   │       ├── register_service.go
│   │       └── service.go
│   └── utils
│       ├── logger.go
│       ├── response.go
│       └── utils.go
├── LICENSE
├── Makefile
├── README.md
└── web
    └── templates
        ├── login.html
        ├── register.html
        ├── reset-password.html
        └── tasks.html
```
---
# Author
Adilet Rabaev(adal4ik)
