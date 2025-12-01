**Library Management System (Go + Gin + PostgreSQL)**

A production-style library management backend built using Golang, Gin, SQLX, PostgreSQL, and JWT Authentication.
This service manages books, members, borrowing, returning, categories, and includes role-based admin access.

**Features**-
Book Management (CRUD + Search by Category)
Member Management (CRUD)
Borrow & Return Workflow
Borrow History per Member
Category Management (Add/Delete/List)
JWT Login (Admin Protected Routes)
Clean Folder Structure (Repository Pattern)
SQLX for DB mapping

**Folder Structure**
.
├── go.mod
├── go.sum
├── main.go
└── service/
    ├── handler/          # Business logic layer
    │   ├── book_service.go
    │   ├── member_service.go
    │   ├── borrow_service.go
    │   ├── borrowdetails_service.go
    │   ├── category_service.go
    │   └── user_service.go
    │
    ├── libhttp/          # HTTP layer (Gin handlers + routes)
    │   ├── book_handler.go
    │   ├── member_handler.go
    │   ├── borrow_handler.go
    │   ├── borrowdetails_handler.go
    │   ├── category_handler.go
    │   ├── auth_handler.go
    │   └── routes.go
    │
    ├── models/           # Data models (DB + JSON)
    │   └── models.go
    │
    └── repository/
        ├── repository.go # Interfaces for repos
        └── db/
            ├── db.go     # PostgreSQL connection
            ├── queries.go
            ├── book_repo.go
            ├── member_repo.go
            ├── borrow_repo.go
            ├── borrowdetails_repo.go
            ├── category_repo.go
            └── user_repo.go


-Each layer has a single responsibility.
-Handlers NEVER interact with the database directly — they call services.
-Services use interfaces, not concrete DB implementations → makes system testable & maintainable.
-SQLX is used for simpler mapping between SQL rows and Go structs.
-JWT authentication protects admin routes.
-All major entities (Book, Member, Borrow, Category, User) have dedicated models, repository implementations, and service logic.

**Architecture Overview**
HTTP Layer (Gin Handlers)
        ↓
Service Layer (Business Logic)
        ↓
Repository Layer (Interfaces)
        ↓
DB Layer (SQLX + PostgreSQL)


**Database Setup**
1)Create the database:
CREATE DATABASE librarymanagement;
2)Create required tables
Tables used:
book
member
borrow
borrowdetails
category
users
Example: Users table (for login)
CREATE TABLE users (
  user_id SERIAL PRIMARY KEY,
  username VARCHAR(100),
  password VARCHAR(100),
  firstname VARCHAR(100),
  lastname VARCHAR(100),
  role VARCHAR(20)
);
3)Environment variables- Change the DATABASE_URL accroding to your database setup 

**HOW TO RUN**
1) Clone the repository-
   git clone https://github.com/vanitha-1111/LibraryManagement-GO.git
   cd LibraryManagement-GO
2)Install dependencies- go mod tidy
3)Configure database connection-
  -Inside **main.go** update: dsn := "postgres://postgres:<password>@localhost:5432/librarymanagement?sslmode=disable"
4)start the server- go run main.go
The API runs at http://localhost:8080

**Authentication (JWT)**
Login as admin to receive a token:
**POST /auth/login**
{
  "username": "admin",
  "password": "admin"
}
Use the token generated across protected endpoints:
Authorization: Bearer <token>

**KEY ENDPOINTS**

📚**Books**
POST /books
Creates a new book. Body must include book title, category ID, author, publication details, and status. Only admins can call this.

GET /books
Returns the list of all books in the library.

GET /books/:id
Returns details of a single book.
:id should be replaced with the book_id of the book you want to fetch.

GET /books/category/:name
Returns all books under a specific category.
:name is the category name, such as “science”, “math”, “english”.


👤 **Members**
POST /members
Creates a new member. Body must contain basic details like firstname, lastname, contact, and status. Only admins can call this.

GET /members
Fetches all members registered in the system.

GET /members/:id
Fetches details of a specific member.
:id is the member_id.

PUT /members/:id
Updates information for an existing member.
:id is the member_id you want to update.

DELETE /members/:id
Deletes the member with the given ID. Only admins can delete members.


📦 **Borrow & Borrow Details**

POST /borrow
Creates a borrow transaction for a member. The body must include member_id, borrow date, and due date.
This only creates the transaction — books are added separately.

POST /borrow/details
Adds a book to an existing borrow transaction. You must provide borrow_id and book_id.
Also validates member status and reduces book copies.

GET /borrow/:borrow_id/details
Fetches all books associated with a borrow transaction.
Replace :borrow_id with the ID from the borrow table.

GET /members/:id/history
Returns all borrow transactions for a specific member, along with the list of books in each transaction.
:id is the member_id.

🗂 **Categories**

POST /categories
Creates a new book category. Only admins can call this.

GET /categories
Returns all available categories.

DELETE /categories/:id
Deletes a category by its ID.
:id is the category_id from the category table.

