# What Should I Cook

**What Should I Cook** — RESTful API для генерации рациона на основе загруженных пользователем рецептов. Сервис написан на Go с использованием стандартной библиотеки net/http.

## Технический стек
- **Язык и библиотеки:** Go 1.26 (net/http, chi)
- **База данных:** PostgreSQL 15 & golang-migrate CLI
- **Аутентификация:** JWT (access token)
- **Запуск программы:** есть Makefile и docker-compose для тестовой базы данных.
- **Тесты:** testing

## Возможности
- Регистрация и авторизация
- Рандомная генерация меню по заданным параметрам с автоматическим формированием списка покупок
- Создание, редактирование и удаление рецептов (CRUD с валидацией входных данных)
- Возможность публиковать рецепты (настройка приватности) и забирать рецепты других пользователей к себе в коллекцию
- Возможность задавать рецептам категории в виде определенной кастомизируемой кухни (например, "японской кухня", "осеннее меню" и т.д.)
- Возможность присваивать рецептам тип блюда ("основное", "напиток", "салат" и т.д.)

---

## Установка и запуск

#### Необходимое ПО:
- **Go** (версии 1.26 или выше)
- **Docker** & **Docker Compose**
- **Golang-Migrate CLI** (для управления миграциями)
- **Make** (опционально)

### 1. Клонирование репозитория

```bash
git clone https://github.com/Anguilla-anguilla/WhatShouldICook.git
cd WhatShouldICook
```

### 2. Настройка переменных окружения
Создайте  `.env`  файл и задайте настройки окружения. Например:
```env
SERVER_PORT=8080
SERVER_HOST=localhost
DB_URL=postgres://testuser:testpass@localhost:5433/testdb?sslmode=disable
JWT_SECRET=your_secret_key_here
JWT_TTL=24
```

### 3. Запуск PostgreSQL через Docker
Прежде, чем запускать сам контейнер, убедитесь, что Docker (или Docker Desktop) активен.

**Поднять базу данных:**
```bash
docker-compose up -d
```

**Остановить базу данных:**
```bash
docker-compose down          # Остановить контейнеры ИЛИ
docker-compose down -v       # Остановить и удалить том с данными
```

### 4. Применение миграций
**Через Makefile:**
```bash
make migrate-up
```
**Вручную:**
```bash
migrate -path internal/migrations -database postgres://testuser:testpass@localhost:5433/testdb?sslmode=disable up
```
### 5. Запуск
**Через Makefile:**
```bash
make run
```

**Вручную:**
```bash
go run cmd/api/main.go
```

---

## Эндпоинты

| Метод | Эндпоинт | Доступ | Описание | Пример тела запроса |
|-------|----------|--------|----------|---------------------|
| `POST` | `/api/v1/auth/register` | публичный | Регистрация нового пользователя | `{"username": "User", "email": "email@email.com", "password": "password"}` |
| `POST` | `/api/v1/auth/login` | публичный | Вход в систему | `{"username": "User", "password": "password"}` |
| `GET` | `/api/v1/users/profile` | приватный | Получение данных текущего пользователя | — |
| `PUT` | `/api/v1/users/profile` | приватный | Обновление username и/или email | `{"username": "Updated", "email": "new@mail.com"}` |
| `PUT` | `/api/v1/users/password` | приватный | Смена пароля | `{"old_password": "password", "new_password": "wordpass"}` |
| `DELETE` | `/api/v1/users/profile` | приватный | Удаление пользователя | — |
| `POST` | `/api/v1/cuisines/` | приватный | Создание категории кухни | `{"name": "Cuisine", "description": ""}` |
| `GET` | `/api/v1/cuisines/{id}` | приватный | Получение данных о конкретной кухне | — |
| `GET` | `/api/v1/cuisines/` | приватный | Получение всех кухонь пользователя | — |
| `PUT` | `/api/v1/cuisines/{id}` | приватный | Обновление данных кухни | `{"name": "Cuisine Updated", "description": ""}` |
| `DELETE` | `/api/v1/cuisines/{id}` | приватный | Удаление кухни | — |
| `GET` | `/api/v1/categories/{id}` | приватный | Получение категории блюда | — |
| `GET` | `/api/v1/categories/` | приватный | Получение всех категорий блюд | — |
| `POST` | `/api/v1/ingredients/` | приватный | Создание ингредиента | `{"name": "ice"}` |
| `GET` | `/api/v1/ingredients/{id}` | приватный | Получение ингредиента | — |
| `GET` | `/api/v1/ingredients/` | приватный | Получение списка всех ингредиентов | — |
| `PUT` | `/api/v1/ingredients/{id}` | приватный | Обновление названия ингредиента | `{"name": "new_name"}` |
| `DELETE` | `/api/v1/ingredients/{id}` | приватный | Удаление ингредиента | — |
| `POST` | `/api/v1/recipes/` | приватный | Создание рецепта | См. [пример создания рецепта](#пример-создания-рецепта) |
| `POST` | `/api/v1/recipes/{id}/copy` | приватный | Копирование публичного рецепта другого пользователя | `{"cuisine_id": 2, "owner_id": 2}` |
| `GET` | `/api/v1/recipes/{id}` | приватный | Получение рецепта | — |
| `GET` | `/api/v1/recipes/` | приватный | Получение списка рецептов. **Параметры (опционально):** `?favorite` (bool), `?public` (bool), `?cuisine_id` (int), `?category_id` (int) | — |
| `PUT` | `/api/v1/recipes/{id}` | приватный | Обновление рецепта | См. [пример обновления рецепта](#пример-обновления-рецепта) |
| `DELETE` | `/api/v1/recipes/{id}` | приватный | Удаление рецепта | — |
| `POST` | `/api/v1/menu/generate` | приватный | Генерация рациона с автоматическим списком покупок | См. [пример генерации меню](#пример-генерации-меню) |
| `GET` | `/api/v1/ration/{id}` | приватный | Получение списка блюд в рационе | — |

---

## Примеры запросов

### Пример создания рецепта
```json
{
  "name": "Какая-то еда",
  "description": "Взять еду, помыть, порезать и приготовить",
  "cooking_time": 20,
  "price": 404,
  "expires_after": 24,
  "store_in_freezer": false,
  "favorite": false,
  "fridgeless_store": 1,
  "public": true,
  "category_id": 1,
  "cuisine_id": 1,
  "ingredients": [
    {"name": "Сосиски", "quantity": 2},
    {"name": "Макарошки", "quantity": 1},
    {"name": "Авокадо", "quantity": 1},
    {"name": "Огурец", "quantity": 2},
    {"name": "Кетчуп", "quantity": 1}
  ]
}
```

### Пример обновления рецепта

``` json
{
  "name": "Какая-то еда upd",
  "description": "Если у вас в холодильнике ничего не осталось...",
  "cooking_time": 5,
  "price": 200,
  "expires_after": 24,
  "store_in_freezer": false,
  "favorite": false,
  "fridgeless_store": 1,
  "public": true,
  "category_id": 1,
  "cuisine_id": 1,
  "ingredients": [
    {"name": "Сосиски", "quantity": 2},
    {"name": "Авокадо", "quantity": 1}
  ]
}
```

### Пример генерации меню
```json
{
  "cuisine_id": 2,
  "category_count": {
    "1": 4,
    "3": 2
  },
  "duration": 7
}
```