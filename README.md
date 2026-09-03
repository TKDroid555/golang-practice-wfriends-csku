# Note App (Next.js + Golang)

This is a simple Note-taking application built with a Next.js frontend and a Golang backend. The primary purpose of this project is to demonstrate the implementation of various Gang of Four (GoF) Design Patterns in a Golang REST API.

## 🚀 How to Run

### 1. Run the Backend (Golang)
Open a terminal in the root directory and start the Go server:
```bash
cd backend
go run main.go
```
The backend API will start on `http://localhost:8080`.

### 2. Run the Frontend (Next.js)
Open another terminal in the root directory and start the Next.js development server:
```bash
cd frontend
npm install  # (If you haven't installed dependencies yet)
npm run dev
```
The frontend UI will be available at `http://localhost:3000`.

---

## 🏗️ Design Patterns Implementation (Golang Backend)

The Golang backend purposefully utilizes 5 Gang of Four Design Patterns to structure the application. Here is a breakdown of where they are used:

### 1. Singleton Pattern
- **Location:** `backend/pkg/storage/memory.go`
- **Description:** Uses `sync.Once` to ensure that only one instance of the `MemoryStorage` struct is created. This ensures we have a single, global in-memory database instance throughout the application lifecycle.

### 2. Strategy Pattern
- **Location:** `backend/pkg/storage/storage.go`
- **Description:** The `NoteStorage` interface acts as the strategy. `MemoryStorage` is the concrete strategy that implements it. The HTTP handlers only depend on the interface, meaning we can easily swap out `MemoryStorage` with a database (like Postgres) in the future without changing the handler logic.

### 3. Factory Method Pattern
- **Location:** `backend/pkg/patterns/factory.go`
- **Description:** The `NoteFactory` interface and its implementation encapsulate the creation of a `Note` object. Instead of initializing the struct manually in the handler, the handler calls `factory.CreateNote(title, content)`. The factory handles generating a unique ID and setting the `CreatedAt` timestamp automatically.

### 4. Builder Pattern
- **Location:** `backend/pkg/patterns/builder.go`
- **Description:** The `ResponseBuilder` struct simplifies sending JSON HTTP responses. Instead of manually creating structs and writing headers every time, we can chain methods sequentially:
  `patterns.NewResponseBuilder(w).Status(http.StatusCreated).Message("Success").Data(note).Send()`

### 5. Decorator Pattern
- **Location:** `backend/pkg/patterns/decorator.go`
- **Description:** We implemented HTTP Middleware as decorators. The `ApplyDecorators` function takes an `http.HandlerFunc` and wraps it with `LoggingDecorator` and `CORSDecorator`. This adds logging and CORS header functionality dynamically to the routes without modifying the core routing or handler logic in `main.go`.

---

## 🎨 Tech Stack
- **Frontend**: Next.js (App Router), React, Tailwind CSS, TypeScript
- **Backend**: Golang (Standard `net/http` library)
