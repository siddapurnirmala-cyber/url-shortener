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

### Prerequisites

- Go 1.25.1+
- MongoDB (running on `localhost:27017`)
- Redis (running on `localhost:6379`)

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
   ```bash
   go run cmd/main.go
   ```
   The server will start on `http://localhost:9002`

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
