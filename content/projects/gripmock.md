---
title: "GripMock - gRPC Mock Server"
date: "2025-11-21"
author: "Maxim Babichev"
excerpt: "Мощный gRPC Mock Server с поддержкой streaming, динамических стабов и Web UI"
tags: ["go", "grpc", "mock", "testing"]
links:
  github: "https://github.com/bavix/gripmock"
  docs: "https://gripmock.org/"
  homepage: "https://gripmock.org/"
---

# GripMock

GripMock — gRPC Mock Server для тестирования gRPC приложений.

## Описание

Мощный и гибкий mock-сервер для gRPC с поддержкой статических и динамических стабов, streaming, Web UI и REST API.

## Основные возможности

- Статические и динамические стабы с использованием шаблонов
- Поддержка всех типов streaming (server-side, client-side, bidirectional)
- Web интерфейс для управления стабами
- REST API для программного управления
- Приоритетная система стабов
- Header matching и pattern matching

## Быстрый старт

### Установка через Homebrew

```bash
brew tap gripmock/tap
brew install gripmock
```

### Использование без Docker

```bash
gripmock --stub=./stubs ./proto
```

### Использование Docker

```bash
docker run -p 4770:4770 -p 4771:4771 \
  -v $(pwd)/stubs:/stubs \
  -v $(pwd)/proto:/proto \
  bavix/gripmock --stub=/stubs /proto/service.proto
```

- **Порт 4770**: gRPC сервер
- **Порт 4771**: Web UI и REST API

## Ссылки

- **GitHub**: [bavix/gripmock](https://github.com/bavix/gripmock)
- **Документация**: [gripmock.org](https://gripmock.org/)

## Лицензия

MIT License
