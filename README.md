# 🛒 Distributed E-Commerce Backend (Microservices in Go)

A production-style distributed e-commerce backend built using **Go**, following a **microservices architecture** with an **API Gateway**, **JWT authentication**, and **Docker Compose orchestration**.

This project demonstrates service isolation, inter-service communication, and scalable backend design patterns commonly used in real-world systems.

---

## 🚀 Features

* 🔐 JWT-based authentication
* 🌐 API Gateway for centralized routing
* 🧩 Independent microservices (User, Product, Order)
* 🗄️ Separate PostgreSQL database per service
* 🔁 Inter-service HTTP communication
* 🐳 Fully containerized with Docker Compose
* ⚙️ Environment-based configuration using `.env`
* 🧱 Modular and maintainable project structure

---

## 🏗️ Architecture

```
Client
   ↓
API Gateway (JWT validation)
   ↓
-----------------------------------------
| User Service    (Auth + Users)        |
| Product Service (Products CRUD)       |
| Order Service   (Order + Validation)  |
-----------------------------------------
Each service → Own PostgreSQL database
```

### Service Responsibilities

* **API Gateway**

  * Central entry point
  * JWT validation middleware
  * Request routing to services

* **User Service**

  * User registration & login
  * Password hashing with bcrypt
  * JWT token generation

* **Product Service**

  * Product CRUD operations
  * Product catalog management

* **Order Service**

  * Order creation
  * Validates user and product via service calls
  * Persists order data

---

## 🧰 Tech Stack

* **Language:** Go (Golang)
* **Framework:** Gin
* **Database:** PostgreSQL
* **ORM:** GORM
* **Auth:** JWT + bcrypt
* **Containerization:** Docker & Docker Compose
* **Architecture:** Microservices + API Gateway

---

## 📁 Project Structure

```
ecommerce-microservices/
│
├── api-gateway/
├── user-service/
├── product-service/
├── order-service/
└── docker-compose.yml
```

Each service is independently buildable and deployable.

---

## ⚙️ Environment Variables

Each service uses a `.env` file for configuration.

Example:

```
PORT=8001
DB_HOST=user-db
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=userdb
DB_PORT=5432
JWT_SECRET=supersecret
```

> ⚠️ `.env` files are excluded from version control via `.gitignore`.

---

## 🐳 Running the Project

### 1️⃣ Clone the repository

```bash
git clone https://github.com/PIXELL07/ecommerce-microservices.git
cd ecommerce-microservices
```

### 2️⃣ Start all services

```bash
docker-compose up --build
```

### 3️⃣ API Gateway

The system will be available at:

```
http://localhost:8000
```

---

## 🔌 Main API Endpoints

### Auth (Public)

* `POST /register`
* `POST /login`

### Products (Protected)

* `GET /products`
* `POST /products`

### Orders (Protected)

* `POST /orders`

> Protected routes require `Authorization: Bearer <token>` header.

---

## 🎯 Learning Outcomes

This project demonstrates:

* Designing microservices in Go
* Implementing API Gateway pattern
* JWT authentication flow
* Service-to-service communication
* Container orchestration with Docker Compose
* Environment-driven configuration
* Clean modular backend structure

---

## 🚧 Future Improvements

* gRPC between services
* Redis caching layer
* Message queue (Kafka/RabbitMQ)
* Distributed tracing
* Kubernetes deployment
* CI/CD pipeline

---

## 👨‍💻 Author

Built as part of backend engineering practice and microservices learning journey.
