# ✅ Task Manager

A lightweight **task management system** built in **Go**, supporting both **CLI** and **REST API** modes.  
It allows users to create, update, delete, and track tasks — designed with **clean architecture** principles for flexibility and maintainability.

---

## 🧠 Overview

Task Manager helps organize daily workflows through a simple interface that can run both from the **terminal** and via **HTTP API**.  
It demonstrates Go fundamentals such as file handling, data persistence, HTTP routing, and modular design.

---

## 🛠️ Tech Stack

- **Language:** Go 1.22  
- **Storage:** CSV / JSON file-based persistence  
- **Architecture:** Layered (Handler → Service → Repository)  
- **API Framework:** `net/http` + `gorilla/mux`  
- **Logging:** Custom logger  
- **Build:** `make`, Dockerfile (optional)

---

## ✨ Features

✅ **Implemented**
- CLI and REST API interfaces  
- CRUD operations for tasks (create, update, delete, list)  
- Task priority and status management  
- File-based storage (CSV/JSON)  
- Input validation and error handling  

🚧 **In Progress**
- REST API enhancements (search, filtering)  
- Persistent storage using PostgreSQL  
- Improved structured logging  

🔮 **Planned**
- User authentication and assignments  
- Task deadlines and reminders  
- TUI (Terminal UI) interface  
- JSON import/export between instances  

---

## 👨‍💻 Author

Adilet Rabaev

