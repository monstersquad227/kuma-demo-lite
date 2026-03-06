# backend

**Languages:** [English](README.md) | [中文](README.zh.md)

## Overview

`backend` is a simple HTTP service written in **Go** with the **Gin** framework. It serves as the backend in the **Kuma Service Mesh** demo.

It exposes two endpoints:

- Health check
- Random number API

Use it as a **Kuma Mesh** backend example for testing service discovery, routing, mTLS, and traffic policies.

## API

### 1. Health check

**Request**
```bash
GET /backend/actuator/health
```

**Response**
```json
"backend OK"
```

### 2. Random numbers

**Request**
```bash
GET /backend/random
```

**Response**
```json
[
  {
    "id": 1,
    "random": "42",
    "created_at": "2025-08-06 16:35:12",
    "updated_at": "2025-08-06 16:35:12"
  },
  {
    "id": 2,
    "random": "87",
    "created_at": "2025-08-06 16:35:12",
    "updated_at": "2025-08-06 16:35:12"
  }
]
```

> Notes:
> - Each request returns 10 items.
> - `id` goes from 1 to 10.
> - `random` is an integer between 1 and 100.
> - `created_at` / `updated_at` are generation timestamps.

## Run

```bash
backend % make help

Usage:
  make <command>

🎯 The commands are:
  help             Display this help.
  run              Running backend.
  build-linux      Build backend Linux binary.
  build-docker     Build docker image with the backend.
```
