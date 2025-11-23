# Fullstack Shopping Cart (Go + Gin + GORM + React)

This is a very simple shopping cart project that feels like a student side project.  
The Go backend exposes a small REST API with token-based login, carts, and orders.  
The React frontend lets you log in, view items, add to cart, and checkout with a few buttons.

## Tech Stack
- Backend: Go, Gin, GORM, SQLite
- Frontend: React (CRA, plain JavaScript)
- Tests: Ginkgo + Gomega

## Folder Structure
```
backend/   # Go API (Gin + GORM)
frontend/  # React app created with create-react-app
README.md
shopping-cart.postman_collection.json
```

## Backend Setup
```bash
cd backend
go mod tidy          # install deps
go run .             # starts on http://localhost:8080
```

## Frontend Setup
```bash
cd frontend
npm install          # installs packages
npm start            # starts React on http://localhost:3000
```

## Token-Based Login
1. Create a user with `POST /users`.
2. Call `POST /users/login` with username/password.
3. The API returns a `token`. Only one token per user (overwrites the previous one).
4. Store the token in `localStorage` (frontend already does this).
5. Send the token in the `Authorization` header when calling protected endpoints (`POST /carts`, `POST /orders`).

## Cart → Order Flow
1. User logs in to get a token.
2. User clicks an item in the UI (sends `POST /carts` with `item_id` and `quantity`).
3. Cart is unique per user.
4. Checkout button calls `POST /orders`, which copies cart items into `Order` and empties the cart.
5. Orders button fetches the list to show the history.

## API Endpoints
```text
POST /users
{
  "username": "alice",
  "password": "pass"
}

GET /users

POST /users/login
{
  "username": "alice",
  "password": "pass"
}

POST /items
{
  "name": "Pen",
  "price": 2.5
}

GET /items

POST /carts        (needs Authorization header)
{
  "item_id": 1,
  "quantity": 1
}

GET /carts

POST /orders       (needs Authorization header)

GET /orders
```

## Screenshots
*(Add screenshots here if you have them. You can drop PNG/JPG into a `docs/` folder.)*

## Running Ginkgo Tests
```bash
cd backend
go test ./...
```
This runs the specs inside `backend/tests`, which use a temporary SQLite database.

## Troubleshooting
- **Go command not found**: Install Go 1.21+ and re-open your terminal so `$PATH` updates.
- **React API calls failing**: Make sure the Go backend is running on `http://localhost:8080` and CORS is allowed (Gin default is OK).
- **Token errors**: Ensure you are sending the `Authorization` header exactly as returned from `/users/login`.
- **SQLite file locked**: Stop any running backend before re-running tests so the DB file can be deleted safely.
- **Port already in use**: Change the React dev server port with `set PORT=3001 && npm start` or adjust the backend port in `main.go`.

