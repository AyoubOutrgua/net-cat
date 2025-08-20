# TCP-Chat: NetCat-like Group Chat in Go

A simple TCP-based group chat server and client written in Go, inspired by NetCat. The server listens on a specified port and allows multiple clients to connect, send messages, and interact in real-time.

---

## 🛠️ Features

- ✅ TCP server-client architecture (1 server, max 10 clients)
- ✅ Clients must provide a **non-empty name** to join
- ✅ All messages include **timestamp** and **sender name**  
  Format: `[YYYY-MM-DD HH:MM:SS][username]: message`
- ✅ **Previous messages** are sent to newly connected clients
- ✅ Clients are notified when someone **joins** or **leaves**
- ✅ **Empty messages** are ignored
- ✅ Default port is **8989**
- ✅ Uses **goroutines** and **mutexes** (or channels) for concurrency
- ✅ Robust **error handling** on both server and client side

---

## 🚀 Usage

### 🖥️ Server
```sh
go run .              # uses default port 8989
go run . 2525         # specify custom port
```
### Client
```
nc \$IP $port
```

## 🗂️ Project Structure
```
/TCP-Chat
├── functions/
│   ├── atoi.go     
│   ├── handleClient.go   
│   ├── isPrintableRange.go
│   ├── isvalidusename.go
│   ├── listenning.go
│   └── sendMessage.go
├── go.mod 
├── main.go 
└── Readme.md
```


## 📚 Learning Objectives

- TCP socket programming in Go
- Client-server communication
- Managing concurrency using goroutines and mutexes
- Working with channels and synchronization
- Building real-time applications
- Clean code practices and modular project design

---

## 🧪 Requirements

- **Language:** Go
- **Concurrency:** goroutines, channels or sync.Mutex
- **Max connections:** 10 clients
- **Allowed packages:**  
  `io`, `log`, `os`, `fmt`, `net`, `sync`, `time`, `bufio`, `errors`, `strings`, `reflect`
- Code must handle runtime errors (connection drops, invalid input, etc.)

## 👨‍💻 Author
```
- MOHAMMED EL GHAMARI (melghama)
- AYYOUB OUTRGUA (aoutrgua)
- ABD-EL-KAFY BOURAZZA (abourazz)
```