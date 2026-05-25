# SwiftChat Backend

[![Go 1.23](https://img.shields.io/badge/Go-1.23%2B-00ADD8.svg?logo=go&logoColor=white)](https://go.dev/)
[![Go Build](https://github.com/kwa0x2/SwiftChat-Backend/actions/workflows/ci.yml/badge.svg)](https://github.com/kwa0x2/SwiftChat-Backend/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/kwa0x2/swiftchat-backend?style=flat-square)](https://goreportcard.com/report/github.com/kwa0x2/swiftchat-backend)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Website](https://img.shields.io/badge/Live%20Demo-chat.alperkarakoyun.com-red.svg)](https://chat.alperkarakoyun.com/)

> A production-ready, real-time chat backend built with Go. Powers [SwiftChat](https://chat.alperkarakoyun.com/) — a modern messaging platform with friends, rooms, file sharing and live presence.

---

## Overview

SwiftChat Backend is a real-time messaging service written in Go using the [Gin](https://github.com/gin-gonic/gin) HTTP framework and [Socket.IO](https://github.com/zishang520/socket.io) for bi-directional communication. It exposes a clean REST API for resource management and a WebSocket layer for live events such as messaging, presence and friend updates.

The codebase follows a layered architecture (`controller → service → repository`) with dependency injection, making it easy to test, extend and reason about.

## Features

- **Real-time messaging** over Socket.IO with room-based broadcasting
- **Authentication** via Google OAuth 2.0 and email/password sign-up
- **Session management** backed by Redis with secure HTTP-only cookies
- **Friends system** — send / accept / reject requests, block & unblock users
- **Chat rooms** with persisted message history
- **File uploads** to Amazon S3 (profile photos and attachments)
- **Transactional emails** powered by Resend with HTML templates
- **Containerized deployment** with Docker Compose (API + Postgres + Redis)
- **CI pipeline** with GitHub Actions

## Tech Stack

| Layer            | Technology                                                                 |
|------------------|----------------------------------------------------------------------------|
| Language         | Go 1.23+                                                                   |
| HTTP Framework   | [Gin](https://github.com/gin-gonic/gin)                                    |
| Real-time        | [Socket.IO](https://github.com/zishang520/socket.io) (Go)                  |
| Database         | PostgreSQL 17 + [GORM](https://gorm.io/)                                   |
| Cache / Sessions | Redis                                                                      |
| Object Storage   | Amazon S3 (AWS SDK v2)                                                     |
| Auth             | Google OAuth 2.0, JWT, Sessions                                            |
| Email            | [Resend](https://resend.com/)                                              |
| Containerization | Docker & Docker Compose                                                    |

## Architecture

![Architecture diagram](./diagram.jpg)

## Getting Started

### Prerequisites

- [Docker](https://www.docker.com/) & Docker Compose
- (Optional, for local dev without Docker) Go 1.23+, PostgreSQL 17, Redis

### 1. Clone the repository

```bash
git clone https://github.com/kwa0x2/SwiftChat-Backend.git
cd SwiftChat-Backend
```

### 2. Configure environment variables

Create a `.env` file in the project root:

```env
# Server
PORT=9000
GIN_MODE=release

# PostgreSQL
POSTGRE_DB=swiftchat
POSTGRE_USER=postgres
POSTGRE_PASSWORD=your_strong_password
POSTGRE_HOST=swiftchat_postgres
POSTGRE_PORT=5432

# Redis
REDIS_HOST=swiftchat_redis:6379
REDIS_PASSWORD=your_redis_password

# Session & JWT
SESSION_SECRET=your_session_secret
JWT_SECRET=your_jwt_secret

# Google OAuth
GOOGLE_CLIENT_ID=your_google_client_id
GOOGLE_CLIENT_SECRET=your_google_client_secret
GOOGLE_REDIRECT_URL=http://localhost:9000/api/v1/auth/login/google/callback

# AWS S3
AWS_ACCESS_KEY_ID=your_aws_key
AWS_SECRET_ACCESS_KEY=your_aws_secret
AWS_REGION=eu-central-1
AWS_BUCKET_NAME=swiftchat-bucket

# Resend (email)
RESEND_API_KEY=your_resend_api_key
```

### 3. Run with Docker Compose

```bash
docker compose up -d --build
```

The API will be available at **http://localhost:9000**.

### 4. (Optional) Run locally without Docker

```bash
go mod download
make run
```

## API Reference

All endpoints are prefixed with `/api/v1`.

### Auth

| Method | Endpoint                       | Description                  |
|--------|--------------------------------|------------------------------|
| GET    | `/auth/login/google`           | Initiate Google OAuth flow   |
| GET    | `/auth/login/google/callback`  | Google OAuth callback        |
| POST   | `/auth/signup`                 | Sign up with email/password  |
| POST   | `/auth/logout`                 | Log out current session      |
| GET    | `/auth/status`                 | Check current auth state     |

### Users

| Method | Endpoint                  | Description                |
|--------|---------------------------|----------------------------|
| PATCH  | `/user/username`          | Update username            |
| PATCH  | `/user/profile-photo`     | Upload / update avatar     |

### Friends

| Method | Endpoint              | Description               |
|--------|-----------------------|---------------------------|
| GET    | `/friends`            | List all friends          |
| GET    | `/friends/blocked`    | List blocked users        |
| PATCH  | `/friends/block`      | Block a user              |
| DELETE | `/friends`            | Remove a friend           |

### Friend Requests

| Method | Endpoint              | Description                       |
|--------|-----------------------|-----------------------------------|
| POST   | `/friend-requests`    | Send a friend request             |
| GET    | `/friend-requests`    | List incoming requests            |
| PATCH  | `/friend-requests`    | Accept / reject a request         |

### Rooms

| Method | Endpoint                | Description                    |
|--------|-------------------------|--------------------------------|
| POST   | `/rooms/check`          | Get or create a chat room      |
| GET    | `/rooms/chat-list`      | Fetch the user's chat list     |

### Messages

| Method | Endpoint               | Description           |
|--------|------------------------|-----------------------|
| POST   | `/messages/history`    | Get chat history      |

### Files

| Method | Endpoint           | Description     |
|--------|--------------------|-----------------|
| POST   | `/files/upload`    | Upload a file   |

### Real-time (Socket.IO)

Connect to `ws://localhost:9000/socket.io/` with an authenticated session cookie. The server emits events for new messages, friend updates, presence changes and more.

## Project Structure

```
SwiftChat-Backend/
├── config/            # App configuration & environment loading
├── controller/        # HTTP handlers (Gin)
├── email_templates/   # HTML templates for transactional emails
├── internal/          # Internal packages (not exported)
├── middlewares/       # Auth, session and CORS middleware
├── models/            # GORM models (User, Room, Message, …)
├── repository/        # Database access layer
├── routes/            # Route registration
├── service/           # Business logic
├── socket/            # Socket.IO setup and event handlers
├── sql/               # Database init scripts
├── types/             # Shared types & DTOs
├── utils/             # Helpers (hashing, validation, …)
├── compose.yaml       # Docker Compose definition
├── Dockerfile         # API container image
└── main.go            # Application entrypoint
```

## Development

Common tasks are wrapped in the included `Makefile`:

```bash
make run        # run the API locally with hot reload (requires air)
make compile    # build a Windows binary into bin/
make compose    # start the full stack with Docker Compose
```

## Related Repositories

- [SwiftChat Frontend](https://github.com/kwa0x2/SwiftChat-Frontend) — Next.js client

## License

Distributed under the MIT License. See [`LICENSE`](LICENSE) for more information.
