# Secure Full-Stack gRPC App with OAuth 2.0 & OpenID Connect (OIDC)

This project is a modern full-stack web application built with **Go**, **gRPC**, and **OAuth 2.0 / OpenID Connect**. It serves as a learning sandbox and potential foundation for secure, scalable, and high-performance backend systems with user authentication and authorization.

---

## 🚀 Project Goals

- ✅ Implement a gRPC-based backend in Go
- ✅ Authenticate users using OAuth 2.0 + OIDC
- ✅ Protect gRPC endpoints with JWT-based access tokens
- ✅ Serve a simple frontend (HTML, CSS, JS) that interacts with the backend
- ✅ Understand token flows: Authorization Code, ID Token, Access Token
- ✅ Practice structuring a secure service-oriented architecture in Go

---

## 🧱 Tech Stack

| Layer    | Tech                                                   |
| -------- | ------------------------------------------------------ |
| Backend  | Go, gRPC, Protocol Buffers                             |
| Auth     | OAuth 2.0, OpenID Connect, JWT (via Auth0 or Keycloak) |
| Frontend | Vanilla JS, HTML, CSS                                  |
| Gateway  | grpc-gateway (optional REST interface)                 |
| Security | HTTPS (via local TLS or reverse proxy)                 |

---

## 🗂️ Project Structure

secure-app/
│
├── backend/ # Go-based gRPC server
│ ├── proto/ # .proto files
│ ├── auth/ # OIDC validation, token parsing
│ ├── handlers/ # gRPC service logic
│ ├── middleware/ # Token verification interceptors
│ └── main.go # Entry point
│
├── web/ # Static frontend
│ ├── index.html
│ ├── app.js # OIDC client logic
│ └── style.css
│
├── gateway/ # (optional) REST proxy using grpc-gateway
│ └── main.go
│
└── README.md

---

## 🔐 How Authentication Works

1. **Frontend Login Flow**

   - User clicks "Login"
   - Redirects to OIDC provider (Auth0 / Keycloak)
   - Provider authenticates user and redirects back with `code`
   - JS client exchanges `code` for `access_token` and `id_token`

2. **Token Usage**

   - `access_token`: sent with each request (`Authorization: Bearer <token>`)
   - `id_token`: contains identity claims (name, email, sub)

3. **Backend Protection**
   - Every gRPC call passes through an interceptor
   - Token is verified, validated, and claims are extracted
   - Request proceeds only if the token is valid and allowed

---

## 🛠️ Running the Project (WIP)

1. Install Go & `protoc`
2. Register an app on your OAuth/OIDC provider (Auth0, Google, or Keycloak)
3. Configure client ID, secret, and redirect URIs
4. Generate gRPC code from `proto/`
5. Run backend and frontend locally
6. Access the site, login, and interact with the gRPC-secured API

---

## 📚 Concepts Practiced

- gRPC service design with Protocol Buffers
- Secure API design with bearer tokens
- Identity & claims validation with OIDC
- Stateless authentication using JWT
- Full-stack Go app architecture

---

## 🧪 TODO / Roadmap

- [ ] Add grpc-gateway for REST fallback
- [ ] Token refresh logic on frontend
- [ ] Use HTTPS with self-signed cert or Caddy
- [ ] Add role-based access control via `roles` claim
- [ ] Persist user data in a database

---

## 📖 References

- [OAuth 2.0 RFC 6749](https://tools.ietf.org/html/rfc6749)
- [OpenID Connect Spec](https://openid.net/specs/openid-connect-core-1_0.html)
- [gRPC Docs](https://grpc.io/docs/)
- [CoreOS go-oidc](https://github.com/coreos/go-oidc)
