🛒 Shopping Cart – Full Stack Project

Built with Go (Gin + GORM) + SQLite + React
By: Shyamjee Srivastav

This is a simple full-stack Shopping Cart application built as per the assignment PDF.
It includes user signup/login, items, cart system, order checkout, and token-based authentication.

📌 Features Overview
Backend (Go + Gin + GORM + SQLite)

User Signup & Login

Token-based Authentication (stored in DB per user)

Items API (POST + GET)

Cart API (Add to Cart + View Cart)

Order API (Checkout + View Orders)

SQLite Database with Auto-Migrations

Simple, readable student-friendly code

Ginkgo Test Suite for:

signup

login

add to cart

checkout

Frontend (React)

Login Page

Items Page

Add to Cart Button

Checkout Button

View Cart

View Orders

Session Token stored in localStorage

Calls backend via simple fetch()

Minimal styling (assignment requirement)

🗂️ Project Structure
Shopping-cart/
│
├── backend/
│   ├── main.go
│   ├── database.go
│   ├── go.mod
│   ├── models/
│   ├── controllers/
│   ├── routes/
│   ├── tests/
│   └── shopping.db (auto-created)
│
└── frontend/
    ├── package.json
    ├── src/
    │   ├── App.js
    │   ├── api.js
    │   ├── pages/
    │   │   ├── Login.js
    │   │   └── Items.js
    │   └── components/
    │       └── ItemList.js
    └── public/

🚀 How to Run the Project
1️⃣ Start Backend
cd backend
go run .


Backend runs on:
👉 http://localhost:8080

2️⃣ Start Frontend
cd frontend
npm install
npm start


Frontend runs on:
👉 http://localhost:3003

🔐 Authentication Flow

User logs in through /users/login

Backend validates credentials

Token is generated (simple random string)

Token stored in:

user table (CurrentToken)

browser localStorage

Protected APIs require:

Authorization: <token>

🧪 API Endpoints
Users
Method	Endpoint	Description
POST	/users	Create user
GET	/users	List users
POST	/users/login	Login + Get token
Items
Method	Endpoint	Description
POST	/items	Create item
GET	/items	List items
Carts
Method	Endpoint	Description
POST	/carts	Add item to cart
GET	/carts	View cart
Orders
Method	Endpoint	Description
POST	/orders	Checkout
GET	/orders	View orders
🧪 Running Tests (Ginkgo)

Inside backend folder:

go test ./...


Tests cover:

Signup

Login

Token validation

Add to cart

Checkout flow

📸 Screenshots (Add your own)
/screenshots/
  - login.png
  - items.png
  - cart.png
  - checkout.png
  - orders.png

📦 Postman Collection

A Postman collection is included for all API calls:
shopping-cart.postman_collection.json

🔧 Technologies Used
Backend

Go 1.21+

Gin Web Framework

GORM ORM

SQLite Database

Ginkgo Testing Framework

Frontend

React.js

Fetch API

LocalStorage

🎯 Key Learning Outcomes

Full-stack development workflow

REST API design

Token authentication

Database modeling

Frontend API integration

Testing backend logic

Clean folder structure

🙋‍♂️ Author

Shyamjee Srivastav
GitHub: https://github.com/shyamfitness
