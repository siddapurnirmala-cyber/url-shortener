# 🔗 Go URL Shortener

A high-performance, production-ready URL shortener service built with Go, featuring multi-layer storage with MongoDB for persistence and Redis for blazing-fast caching.

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![MongoDB](https://img.shields.io/badge/MongoDB-4EA94B?style=for-the-badge&logo=mongodb&logoColor=white)
![Redis](https://img.shields.io/badge/redis-%23DD0031.svg?style=for-the-badge&logo=redis&logoColor=white)
![Gin](https://img.shields.io/badge/Gin-008ECF?style=for-the-badge&logo=gin&logoColor=white)

## 🚀 Features

- **Efficient Shortening**: Converts long URLs into short, unique codes using Hashids.
- **Lightning Fast Redirection**: Uses Redis caching to minimize database lookups for frequent redirects.
- **Robust Persistence**: Stores all URL mappings in MongoDB.
- **Optimized Database Queries**: Leverages MongoDB database indexing on short codes for blazingly fast     lookup times, ensuring maximum performance at scale.
- **Clean Architecture**: Modular code structure for easy maintenance and scalability.
- **Validation**: Ensures only valid URLs are shortened.
- **Rate Limiting**: Protects the API using a Redis-backed Token Bucket algorithm (per IP).

## 🛠️ Tech Stack

- **Language**: Go (Golang)
- **Web Framework**: [Gin Gonic](https://gin-gonic.com/)
- **Database**: [MongoDB](https://www.mongodb.com/)
- **Cache**: [Redis](https://redis.io/)
- **ID Generation**: [Hashids](https://hashids.org/go/)

## 📂 Project Structure

```text
.
├── cmd/
│   └── main.go           # Application entry point
├── config/
│   └── config.go         # Database & Cache connections
├── internal/
│   ├── handler/          # HTTP request handlers
│   ├── middleware/       # Custom middleware (Token Bucket Limiter)
│   ├── model/            # Data models
│   ├── repository/       # Data access layer (MongoDB)
│   ├── service/          # Business logic layer
│   └── utils/            # Utility functions (Hashids encoding)
├── go.mod                # Go module dependencies
└── README.md             # Project documentation
```

## 🏁 Getting Started

### 📋 Prerequisites

Before running the project, you need to install **Go (1.25.1+)**, **MongoDB**, and **Redis**. Here is how to set them up on your operating system:

#### 🪟 Windows
- **Go**: Download and run the installer from the [official Go website](https://go.dev/dl/).
- **MongoDB**: Download the [MongoDB Community Server](https://www.mongodb.com/try/download/community) `.msi` installer. Install it as a Windows Service (it will run automatically on `localhost:27017`).
- **Redis**: Redis is not officially supported natively on Windows. The best options are:
  - **WSL2 (Recommended):** Open Ubuntu in WSL and run `sudo apt install redis-server` followed by `sudo service redis-server start`.
  - **Docker:** Run `docker run -d -p 6379:6379 redis`.

#### 🍎 macOS
The easiest way to install the dependencies is using [Homebrew](https://brew.sh/):
```bash
# 1. Install Go
brew install go

# 2. Install & Start MongoDB
brew tap mongodb/brew
brew install mongodb-community
brew services start mongodb-community

# 3. Install & Start Redis
brew install redis
brew services start redis
```

#### 🐧 Linux (Ubuntu/Debian)
Use the `apt` package manager:
```bash
# 1. Install Go
sudo apt update
sudo apt install golang-go  # Note: You may want to download from golang.org to get the absolute latest version

# 2. Install & Start MongoDB (v7.0 Example)
curl -fsSL https://www.mongodb.org/static/pgp/server-7.0.asc | sudo gpg -o /usr/share/keyrings/mongodb-server-7.0.gpg --dearmor
echo "deb [ arch=amd64,arm64 signed-by=/usr/share/keyrings/mongodb-server-7.0.gpg ] https://repo.mongodb.org/apt/ubuntu jammy/mongodb-org/7.0 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-7.0.list
sudo apt update
sudo apt install -y mongodb-org
sudo systemctl enable --now mongod

# 3. Install & Start Redis
sudo apt install redis-server
sudo systemctl enable --now redis-server
```

### Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd url-shortner
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Run the application**

   Because this project is built in Go, it works cross-platform. The server will start on `http://localhost:9002` across all systems.

   **🪟 On Windows:**
   Run directly using Command Prompt or PowerShell:
   ```powershell
   go run cmd\main.go
   ```
   *To build a native executable:*
   ```powershell
   go build -o url-shortener.exe cmd\main.go
   .\url-shortener.exe
   ```

   **🍎 On macOS:**
   Run directly from your terminal:
   ```bash
   go run cmd/main.go
   ```
   *To build a native executable:*
   ```bash
   go build -o url-shortener cmd/main.go
   ./url-shortener
   ```

   **🐧 On Linux:**
   Run directly from your terminal:
   ```bash
   go run cmd/main.go
   ```
   *To build a native executable:*
   ```bash
   go build -o url-shortener cmd/main.go
   ./url-shortener
   ```

   **🌍 Cross-Compilation (Bonus)**
   Go makes it incredibly easy to build executables for other operating systems from your current machine. You can do this by setting `GOOS` and `GOARCH`:
   ```bash
   # Build for Linux (from Windows/Mac)
   GOOS=linux GOARCH=amd64 go build -o url-shortener-linux cmd/main.go
   
   # Build for Windows (from Mac/Linux)
   GOOS=windows GOARCH=amd64 go build -o url-shortener.exe cmd/main.go
   ```

## 📡 API Endpoints

### 1. Shorten a URL
**Endpoint:** `POST /shorten`

**Request Body:**
```json
{
  "url": "https://www.google.com"
}
```

**Example (cURL):**
```bash
curl -X POST http://localhost:9002/shorten \
     -H "Content-Type: application/json" \
     -d '{"url": "https://www.google.com"}'
```

**Success Response:**
```json
{
  "short_url": "http://localhost:9002/NjAzNzg"
}
```

---

### 2. Redirect to Original URL
**Endpoint:** `GET /:code`

**Example:**
Just open `http://localhost:9002/NjAzNzg` in your browser.

---

### 3. Rate Limiting
The API is protected by a custom Redis-backed **Token Bucket** rate limiter middleware.
- **Limit**: 10 tokens per IP.
- **Refill Rate**: 1 token per second.
- **Behavior**: Returns `429 Too Many Requests` if the limit is exceeded.

---

## ⚙️ Configuration

The application currently uses default local configurations. You can modify them in `config/config.go`:

- **MongoDB URI**: `mongodb://localhost:27017`
- **Redis Addr**: `localhost:6379`
- **Server Port**: `:9002` (in `cmd/main.go`)

## 📝 License

Distributed under the MIT License. See `LICENSE` for more information.
