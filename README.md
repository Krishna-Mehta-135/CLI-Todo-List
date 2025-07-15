# 📝 CLI Todo List App (Written in Go)

A simple, minimalistic command-line Todo list manager written in Go.  
Easily add, list, edit, delete, and toggle your tasks — all from your terminal.

## 🚀 Features

- ✅ Add todos from the command line
- 🗒️ List todos with `✅` / `❌` status markers
- ✏️ Edit any todo by index
- 🔁 Toggle completion status
- ❌ Delete tasks
- 💾 Persistent storage using local JSON files

## 📦 Usage

### ⏳ Build & Run

```bash
go run .
```

### 🔧 Available Flags

```bash
-a string      Add a new todo
-e string      Edit a todo by index and specify a new title. Format: index:new_title
-d int         Delete a todo by index (default -1)
-t int         Toggle a todo by index (default -1)
-l             List all todos
-h             Show help message
```

---

### 💡 Examples

```bash
# Add todos
go run . -a "Buy milk"
go run . -a "Read Clean Code"

# List all todos
go run . -l

# Edit the second todo (index 1)
go run . -e 1:"Read The Go Programming Language"

# Toggle completion status of first todo
go run . -t 0

# Delete third todo (index 2)
go run . -d 2
```

---

## 🧠 Tech Stack

- **Go (Golang)** — CLI logic, JSON storage
- Standard library only (`flag`, `os`, `encoding/json`)

---

## 📁 File Structure

```
.
├── main.go         # Entry point — parses flags & executes commands
├── todo.go         # Todo struct & methods (add, edit, delete, toggle, validate)
├── storage.go      # Generic JSON file persistence using generics
├── README.md       # You’re here!
```

---

## 📌 Todo (Improvements)

- [ ] Add search/filter support
- [ ] Add due dates / priority
- [ ] Export to markdown or CSV
- [ ] Use Cobra or urfave/cli for better UX

---

## 🙌 Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you’d like to change.

---

### Made with 💻 and ☕ by Krishna Mehta