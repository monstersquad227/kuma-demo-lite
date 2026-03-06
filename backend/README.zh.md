# backend

**Languages:** [English](README.md) | [中文](README.zh.md)

## 项目简介
`backend` 是一个基于 **Go** 与 **Gin** 框架编写的简单 HTTP 服务，主要用于演示 **Kuma Service Mesh** 环境下的后端服务功能。

它提供了两个接口：
+ 健康检查端口
+ 随机数生成接口

该项目可作为 **Kuma Mesh** 中的后端应用示例，方便进行服务发现、路由、mTLS、流量策略等功能的测试。

## 接口说明

### 1. 健康检查接口

#### 请求
```bash
GET /backend/actuator/health
```

#### 响应
```json
"backend OK"
```

### 2. 随机数生成接口

#### 请求
```bash
GET /backend/random
```

#### 响应
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

> 说明：
> + 一次请求会返回 10 条数据
> + id 从 1 到 10 递增
> + random 为 1~100 之间的随机整数
> + created_at / updated_at 为生成时间

## 运行方式
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

---

[English](README.md)
