
# 📘 Go CRUD API with PostgreSQL

This project is a RESTful API built using **Golang (Gin framework)** and **PostgreSQL**. It provides basic user authentication and allows users to create, view, update, and delete posts, as well as add comments to posts.

---

## ✨ Features

- ✅ **User Registration & Login** (with JWT authentication)
- 📝 **Create, Read, Update, Delete (CRUD)** operations on Posts
- 💬 **Comment System** (any user can comment on any post)
- 🔐 **Authorization**: 
  - Only the **owner** of a post can update or delete it
  - Comments are **immutable** (cannot be edited or deleted)
- 🔄 **Cascade delete**: When a post is deleted, all associated comments are also deleted
- 🧾 **Post Structure**: `title` and `description`

---

## 🛠️ Technologies Used

- [Golang](https://golang.org/)
- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [JWT Authentication](https://github.com/golang-jwt/jwt)
- [PostgreSQL](https://www.postgresql.org/)
- [GORM](https://gorm.io/) for ORM
- [dotenv](https://github.com/joho/godotenv) for environment variable management

---

## 📁 Folder Structure

```
.
├── config/         # Database connection and environment config
├── controllers/    # Business logic and API controllers
├── middleware/     # JWT authentication middleware
├── models/         # GORM models (User, Post, Comment)
├── routes/         # Route definitions
├── main.go         # Entry point
├── go.mod          # Module definition
├── .env            # Environment variables (DB credentials, JWT secret, etc.)
```

---

## 🚀 Getting Started

### 1. Clone the repo

```bash
git clone https://github.com/yourusername/go-crud-post-api.git
cd go-crud-post-api
```

### 2. Configure `.env`

Create a `.env` file with your PostgreSQL credentials:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=postdb
JWT_SECRET=your_jwt_secret
```

### 3. Create Database

Create a PostgreSQL database named `postdb` using pgAdmin or psql CLI.

### 4. Run the Server

```bash
go run main.go
```

Server will run at: `http://localhost:8080`

---

## 🔌 API Endpoints

| Method | Endpoint                 | Description                  |
|--------|--------------------------|------------------------------|
| POST   | `/register`              | Register a new user          |
| POST   | `/login`                 | Login and get JWT token      |
| POST   | `/posts`                 | Create a new post (auth)     |
| GET    | `/posts`                 | View all posts               |
| GET    | `/posts/:id`             | View single post             |
| PUT    | `/posts/:id`             | Update post (owner only)     |
| DELETE | `/posts/:id`             | Delete post (owner only)     |
| POST   | `/posts/:postId/comments`| Add comment to a post (auth) |

---

## 📫 Author

**Tinku Gupta**

> If this project helped you, consider giving it a ⭐ on GitHub!
