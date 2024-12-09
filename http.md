
## 1. Introduction
Welcome to the comprehensive learning course on HTTP and the `net/http` package in Go. In this course, you will learn how to create web servers, handle HTTP requests and responses, work with routing and middleware, and make HTTP requests from your Go applications.

## 2. Understanding HTTP

### 2.1. What is HTTP?
HTTP (Hypertext Transfer Protocol) is an application protocol used for transmitting hypertext via the Internet. It is the foundation of any data exchange on the Web and a protocol used for fetching resources, such as HTML documents.

### 2.2. HTTP Methods
HTTP defines several methods to indicate the desired action to be performed on the identified resource:
- **GET**: Retrieve data from a server.
- **POST**: Send data to a server.
- **PUT**: Update a resource.
- **DELETE**: Remove a resource.
- **PATCH**: Apply partial modifications to a resource.

### 2.3. HTTP Status Codes
HTTP status codes indicate the outcome of a request. Here are some common codes:
- **200 OK**: The request has succeeded.
- **201 Created**: The request has been fulfilled and a new resource has been created.
- **400 Bad Request**: The server could not understand the request due to invalid syntax.
- **404 Not Found**: The server could not find the requested resource.
- **500 Internal Server Error**: The server encountered an unexpected condition.

## 3. Setting Up Your Go Environment
1. Install Go from the [official site](https://golang.org/dl/).
2. Set up your workspace and create a new directory for your project.
3. Initialize a new Go module:
   ```bash
   go mod init your-module-name
   ```

## 4. Getting Started with net/http

### 4.1. Creating a Basic HTTP Server
You can create a basic HTTP server using the `net/http` package:

```go
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, World!")
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```

### 4.2. Handling Requests and Responses
You can handle different types of requests and return responses using the `http.ResponseWriter` and `*http.Request` parameters in your handler function.

## 5. Routing in Go

### 5.1. Using the `http.ServeMux`
The `http.ServeMux` is a request multiplexer that matches incoming requests to their respective handler.

```go
mux := http.NewServeMux()
mux.HandleFunc("/", handler)
mux.HandleFunc("/about", aboutHandler)
http.ListenAndServe(":8080", mux)
```

### 5.2. Third-Party Routers
Consider using third-party routers for more complex routing requirements. Popular ones include:
- [gorilla/mux](https://github.com/gorilla/mux)
- [chi](https://github.com/go-chi/chi)

## 6. Middleware
Middleware is a function that wraps around your handler to perform actions before or after the handler executes. For example:

```go
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Printf("Received request: %s\n", r.URL.Path)
        next.ServeHTTP(w, r)
    })
}

// Using middleware
http.Handle("/", loggingMiddleware(http.HandlerFunc(handler)))
```

## 7. Working with JSON
Go provides the `encoding/json` package for working with JSON data. You can easily encode and decode JSON in your HTTP handlers.

```go
type Response struct {
    Message string `json:"message"`
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    response := Response{Message: "Hello, JSON!"}
    json.NewEncoder(w).Encode(response)
}
```

## 8. Making HTTP Requests

### 8.1. GET Requests
You can make GET requests to external APIs or services using `http.Get`.

```go
response, err := http.Get("https://api.example.com/data")
if err != nil {
    log.Fatal(err)
}
defer response.Body.Close()
// Process the response...
```

### 8.2. POST Requests
To make a POST request, use `http.Post` or create a custom request.

```go
data := url.Values{}
data.Set("key", "value")
response, err := http.PostForm("https://api.example.com/submit", data)
if err != nil {
    log.Fatal(err)
}
defer response.Body.Close()
// Process the response...
```

## 9. Error Handling
Always handle errors appropriately in your HTTP server and client code.

```go
if err != nil {
    http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    return
}
```

## 10. Best Practices
- Use proper status codes for responses.
- Validate incoming data.
- Structure your code logically, separating handlers, middleware, and routing.
- Use context for cancellation and timeout handling.

## 11. Advanced Topics

### 11.1. Context
Use the `context` package to manage deadlines, cancellation signals, and other request-scoped values.

```go
func handlerWithContext(ctx context.Context, w http.ResponseWriter, r *http.Request) {
    select {
    case <-ctx.Done():
        http.Error(w, "Request canceled", http.StatusRequestTimeout)
        return
    default:
        // Continue processing...
    }
}
```

### 11.2. Testing HTTP Handlers
You can test your HTTP handlers using the `httptest` package.

```go
func TestHandler(t *testing.T) {
    req := httptest.NewRequest("GET", "/", nil)
    w := httptest.NewRecorder()
    handler(w, req)

    res := w.Result()
    if res.StatusCode != http.StatusOK {
        t.Errorf("Expected status OK, got %v", res.Status)
    }
}
```

### 11.3. HTTP/2 in Go
Go has built-in support for HTTP/2. You can enable it easily when starting your server:

```go
srv := &http.Server{
    Addr: ":8080",
    TLSConfig: &tls.Config{
        MinVersion: tls.VersionTLS13,
    },
}

srv.ListenAndServeTLS("cert.pem", "key.pem")
```
