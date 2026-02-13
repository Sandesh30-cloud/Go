# Go User API

A production-style REST API built with Go using clean architecture principles.

This project demonstrates:
- Layered architecture (handler → service → model)
- RESTful API design
- JSON handling
- Proper HTTP status codes
- Go modules
- Dependency management

---

## 🚀 Tech Stack

- Go (1.22+)
- Gorilla Mux (Router)
- Standard net/http
- JSON encoding/decoding

---

## 📂 Project Structure

go-user-api/
│
├── cmd/
│ └── main.go # Application entry point
│
├── internal/
│ ├── handler/ # HTTP layer
│ ├── service/ # Business logic
│ └── model/ # Data structures
│
├── go.mod
└── README.md



---

## ⚙️ Setup & Run

### 1️⃣ Clone Repository

```bash
git clone https://github.com/Sandesh30-cloud/Go/go-user-api.git
cd go-user-api
```
```
go mod tidy
```
```
go run cmd/main.go
```
```
http://localhost:8080
```
### Example
#### Create User
```
curl -X POST http://localhost:8080/users \
-H "Content-Type: application/json" \
-d '{"id":1,"name":"Sandesh","age":23}'
```

#### Response
```
{
  "id": 1,
  "name": "Sandesh",
  "age": 23
}
```

#### Get All Users
```
curl http://localhost:8080/users
```
#### Get User By ID
```
curl http://localhost:8080/users/1
```

#### Delete User
```
curl -X DELETE http://localhost:8080/users/1
```

