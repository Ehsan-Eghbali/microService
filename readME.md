
# Go Microservices Project

A simple example of a microservices architecture implemented in Go. This project demonstrates a basic structure for a RESTful API with layers for controllers, services, and repositories, as well as database integration and middleware support.

## 📖 Overview
This project is built with a layered architecture to ensure scalability, testability, and maintainability. It uses PostgreSQL as the database and includes configuration management via `.env` files.

### 📂 Project Structure
```
project/
├── main.go                   # Application entry point
├── config/
│   └── config.go             # Configuration loader
├── middleware/
│   └── logging.go            # HTTP logging middleware
├── controller/
│   └── user_controller.go    # Handles HTTP requests
├── service/
│   └── user_service.go       # Business logic
├── repository/
│   └── user_repository.go    # Data access logic
├── model/
│   └── user.go               # User data model
├── db/
│   └── connection.go         # Database connection logic
├── .env                      # Environment variables
├── go.mod                    # Module dependencies
├── go.sum                    # Dependency checksums
```

---

## 🛠️ Features
- **Layered Architecture**: Separation of concerns into Controller, Service, and Repository layers.
- **Database Integration**: PostgreSQL as the database.
- **Logging Middleware**: Logs HTTP requests and response times.
- **Configuration Management**: Environment variables for flexible configuration.

---

## 🚀 How to Run

### Prerequisites
- Go 1.20+
- PostgreSQL installed and running
- Environment variables configured in `.env` file.

### 1️⃣ Clone the repository
```bash
git clone https://github.com/yourusername/go-microservices.git
cd go-microservices
```

### 2️⃣ Install dependencies
```bash
go mod tidy
```

### 3️⃣ Set up the database
Run the following SQL command to create the `users` table:
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL
);
```

### 4️⃣ Configure `.env` file
Create a `.env` file in the project root:
```env
SERVER_PORT=8080
DB_URL=postgres://user:password@localhost:5432/mydb
```

### 5️⃣ Run the application
```bash
go run main.go
```

### 6️⃣ Test the API
Use Postman or `curl` to test the endpoints:
- **Get all users**: `GET /users`
- **Get a user by ID**: `GET /user?id=1`
- **Create a user**: `POST /users` (Body: JSON `{ "name": "John", "email": "john@example.com" }`)

---

## 🧩 Project Layers

### 1️⃣ Controller Layer
Handles HTTP requests and formats responses. Located in the `controller` directory.

### 2️⃣ Service Layer
Implements business logic and communicates with the repository layer. Located in the `service` directory.

### 3️⃣ Repository Layer
Handles data access logic (e.g., database queries). Located in the `repository` directory.

---

## 🌟 Highlights
- **Clean and maintainable architecture.**
- **Uses `gorilla/mux` for routing.**
- **Environment-based configuration for flexibility.**

---

## 📋 License
This project is licensed under the MIT License.

---

## ✉️ Feedback
Have questions or suggestions? Feel free to reach out or create an issue. Contributions are welcome! 😊
