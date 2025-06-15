
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

# 📘 Database Configuration PostgreSQL

# 🛠 STEP 1: Create PostgreSQL Database (`postdb`)

This step explains how to create the PostgreSQL database required by the Go CRUD API project.

---

## ✅ Option 1: Using `psql` Command Line

If you have PostgreSQL and `psql` CLI installed, follow these steps:

1. Open a terminal and run:

```bash
psql -U postgres
```

> Replace `postgres` with your PostgreSQL username if it's different. You’ll be prompted to enter your password (default might be `postgres`).

2. Inside the PostgreSQL shell, run:

```sql
CREATE DATABASE postdb;
```

3. Exit the shell:

```sql
\q
```

---

## ✅ Option 2: Using pgAdmin (GUI)

If you prefer a graphical interface:

1. **Open pgAdmin 4** (installed along with PostgreSQL).
2. Login with your PostgreSQL credentials.
3. In the left panel:
   - Expand `Servers → PostgreSQL → Databases`
4. Right-click on **Databases** → Click **Create → Database**.
5. **Name** the database: `postdb`
6. Leave other settings as default and click **Save**.

---

Once the database is created, your application will be able to connect to it as long as your `.env` file is correctly configured.




## 📫 Author

**Tinku Gupta**

> If this project helped you, consider giving it a ⭐ on GitHub!
