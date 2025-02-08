# 🔗 URL Shortener

A simple and fast URL shortening service built with Go.

## ✨ Features

- **URL Shortening**: Convert long URLs into short, memorable codes
- **Permanent Redirects**: Automatically redirect users to the original URL
- **JSON API**: Simple REST API interface
- **Input Validation**: Ensures valid URLs are provided
- **Logging**: Structured JSON logging for better observability

## 🚀 Getting Started

### Prerequisites

- Go 1.22 or higher

### Running the Server

```bash
go run main.go
```

The server will start on port 8080.

## 📡 API Endpoints

### Shorten a URL

**POST** `/api/shorten`  
**Content-Type:** `application/json`

Request:

```json
{
    "url": "https://your-long-url.com"
}
```

Response:

```json
{
    "data": "Xa4bK9mN"
}
```

### Access Shortened URL

**GET** `/{code}`

## ⚡ Technical Details

- Built with pure Go
- Uses Chi router for HTTP handling
- In-memory storage (map-based)
- Generates 8-character random codes
- Includes request logging and recovery middleware
- 10-second timeout for read/write operations

## 🔒 Security Features

- Request ID tracking
- Panic recovery middleware
- URL validation
- Content-Type enforcement

## 💻 Development

This is a learning project demonstrating Go best practices including:

- Clean code organization
- Error handling
- HTTP server configuration
- Middleware usage
- API design

---

Built with ❤️ using Go
