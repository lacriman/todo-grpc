# 🛡️ Secure gRPC ToDo CLI App (Go + OAuth 2.0 / OIDC)

This project is a secure, full-stack command-line ToDo application written in **Go**, using **gRPC** and **OAuth 2.0 / OpenID Connect (OIDC)** for authentication. It demonstrates how to build a robust backend API, protect it with JWT tokens, and interact with it via a CLI client.

---

## 📦 Features

- 🧠 gRPC backend with full CRUD support for ToDo items
- 🔐 JWT-based authentication via OAuth2 / OIDC (Auth0, Keycloak, etc.)
- 💬 CLI client using `cobra` for user-friendly commands
- 🗂️ Local token storage in `~/.todo-cli/token`
- ✅ Production-ready patterns (interceptors, token validation, modular structure)
- 🐳 Docker support & automation-ready (planned)

---

## 🔐 Authentication Flow

1. User logs in via browser (OIDC provider) and gets an `access_token`
2. CLI client attaches `Authorization: Bearer <token>` to each gRPC call
3. Backend verifies token using provider's JWKS
4. If valid, user claims are extracted and request is authorized

---

## 🚀 Quick Start

### 1. Requirements

- Go 1.22+
- `protoc` compiler with Go plugins
- OAuth2/OIDC provider (e.g. Auth0, Google, Keycloak)

### 2. Clone & Setup

```bash
git clone https://github.com/your-username/todo-grpc-cli.git
cd todo-grpc-cli

# Install Go dependencies
go mod tidy

# Generate gRPC bindings
protoc --go_out=. --go-grpc_out=. proto/todo.proto
```

### 3. Run gRPC Server

```bash
go run cmd/server/main.go
```

### 4. Use CLI Client

```bash
go run cmd/client/main.go create "Buy milk"
go run cmd/client/main.go list
go run cmd/client/main.go delete 2
```
