# 🧠 Go Backend – Feedback Management System

This is the backend API for a feedback management system built in **Go**. It supports authentication, role-based access, and feedback CRUD operations. Works seamlessly with a Flutter frontend.

---

## 🚀 Getting Started

### 📦 Install Dependencies

```bash
go mod tidy
```
## ▶️ Run the Server

```bash
go run main.go
```
## 🔐 Authentication

Users and Admins can **register** and **login**.

A **JWT token** is returned on login.

Protected routes require the following header:

```bash
Authorization: Bearer <your_token_here>
```
Middleware enforces role-based access.
## 🌐 API Endpoints

### 🔑 Auth Routes

| Method | Endpoint     | Description         |
|--------|--------------|---------------------|
| POST   | `/register`  | Register a new user |
| POST   | `/login`     | Login and get token |

### 💬 Feedback Routes

| Method | Endpoint         | Role        | Description                 |
|--------|------------------|-------------|-----------------------------|
| POST   | `/feedback`      | user        | Submit feedback             |
| GET    | `/feedbacks`     | user/admin  | Get feedbacks (role-based) |
| DELETE | `/feedback/{id}` | user/admin  | Delete a feedback by ID     |

> 🔐 All feedback routes are protected via **JWT**.
## 🗃 Database

- **Default**: SQLite  
- Can switch to **PostgreSQL** or **MySQL** by modifying the `database.Connect()` method

### 🧾 Feedback Model (GORM)

```go
type Feedback struct {
  gorm.Model
  Name    string
  Email   string
  Rating  int
  Comment string
  UserID  uint
}
```
## 🧪 Testing

### 📬 Using Postman

1. Register or login via:
   POST /register
   POST /login

2. Copy the **JWT token** from the login response.

3. For authenticated routes, add the following header:
   Authorization: Bearer <your_token_here>

## 🛠 Technologies Used

- **Go** – Backend programming language  
- **GORM** – ORM for database models  
- **JWT** – Token-based authentication  
- **Gorilla Mux** – HTTP routing  
- **SQLite** – Default database (can switch to PostgreSQL/MySQL)



