# What Should I Cook

[Читать на русском языке 🇷🇺](README.ru.md)

**What Should I Cook** is a RESTful API for automated meal planning and recipe management. The service generates custom menus based on user preferences and automatically compiles shopping lists. Built with Go 1.26 using only the standard `net/http` library and the `chi` router.

## Tech Stack
- **Backend:** Go 1.26 (net/http, chi router)
- **Database:** PostgreSQL 15 & golang-migrate CLI
- **Auth:** JWT (Access Tokens)
- **Environment:** Docker, Docker Compose, Makefile
- **Testing:** Standard `testing` package

## Core Features
- **JWT Authentication:** Secure user registration and login.
- **Automated Meal Planning:** Generates random menus tailored to user criteria and builds consolidated grocery lists.
- **Recipe CRUD:** Complete recipe management with robust input validation.
- **Social Features:** Public/private visibility toggles for recipes, allowing users to copy public community recipes into their personal collections.
- **Flexible Organization:** Support for custom cuisines (e.g., "Japanese", "Autumn Menu") and distinct dish types (e.g., "Main Course", "Drink", "Salad").

---

## Getting Started

### Prerequisites
Make sure you have the following installed:
- **Go** (v1.26+)
- **Docker** & **Docker Compose**
- **Golang-Migrate CLI**
- **Make** (optional, for shortcuts)

### 1. Clone the Repository
```bash
git clone https://github.com/Anguilla-Anguilla/whatshouldicook.git
cd whatshouldicook
```

### 2. Configure Environment Variables
Create a `.env` file in the root directory. Example configuration:
```env
SERVER_PORT=8080
SERVER_HOST=localhost
DB_URL=postgres://testuser:testpass@localhost:5433/testdb?sslmode=disable
JWT_SECRET=your_secret_key_here
JWT_TTL=24
```

### 3. Spin Up PostgreSQL via Docker
Ensure Docker Daemon (or Docker Desktop) is running.

**Start the database:**
```bash
docker-compose up -d
```

**Stop the database:**
```bash
docker-compose down          # Stops containers
docker-compose down -v       # Stops containers and wipes database volumes
```

### 4. Run Database Migrations
**Using Makefile:**
```bash
make migrate-up
```

**Manually:**
```bash
migrate -path internal/migrations -database postgres://testuser:testpass@localhost:5433/testdb?sslmode=disable up
```
### 5. Launch the Application
**Using Makefile:**
```bash
make run
```

**Manually:**
```bash
go run cmd/api/main.go
```
---

## API Reference

| Method | Endpoint | Access | Description | Request Body |
|-------|----------|--------|----------|---------------------|
| `POST` | `/api/v1/auth/register` | Public | Register a new user | `{"username": "User", "email": "email@email.com", "password": "password"}` |
| `POST` | `/api/v1/auth/login` | Public | Log in and receive a JWT | `{"username": "User", "password": "password"}` |
| `GET` | `/api/v1/users/profile` | Private | Fetch current user's profile | — |
| `PUT` | `/api/v1/users/profile` | Private | Update username and/or email | `{"username": "Updated", "email": "new@mail.com"}` |
| `PUT` | `/api/v1/users/password` | Private | Change password | `{"old_password": "password", "new_password": "wordpass"}` |
| `DELETE` | `/api/v1/users/profile` | Private | Delete user account | — |
| `POST` | `/api/v1/cuisines/` | Private | Create a custom cuisine category | `{"name": "Cuisine", "description": ""}` |
| `GET` | `/api/v1/cuisines/{id}` | Private | Get details for a specific cuisine | — |
| `GET` | `/api/v1/cuisines/` | Private | List all user cuisines | — |
| `PUT` | `/api/v1/cuisines/{id}` | Private | Update cuisine details | `{"name": "Cuisine Updated", "description": ""}` |
| `DELETE` | `/api/v1/cuisines/{id}` | Private | Delete a cuisine category | — |
| `GET` | `/api/v1/categories/{id}` | Private | Get a dish type | — |
| `GET` | `/api/v1/categories/` | Private | List all dish types | — |
| `POST` | `/api/v1/ingredients/` | Private | Create a new ingredient | `{"name": "ice"}` |
| `GET` | `/api/v1/ingredients/{id}` | Private | Fetch a specific ingredient | — |
| `GET` | `/api/v1/ingredients/` | Private | List all ingredients | — |
| `PUT` | `/api/v1/ingredients/{id}` | Private | Rename an ingredient | `{"name": "new_name"}` |
| `DELETE` | `/api/v1/ingredients/{id}` | Private | Delete an ingredient | — |
| `POST` | `/api/v1/recipes/` | Private | Create a new recipe | See [Create Recipe Example](#сreate-recipe-example) |
| `POST` | `/api/v1/recipes/{id}/copy` | Private | Copy another user's public recipe | `{"cuisine_id": 2, "owner_id": 2}` |
| `GET` | `/api/v1/recipes/{id}` | Private | Fetch a specific recipe | — |
| `GET` | `/api/v1/recipes/` | Private | List and filter recipes.  **Query (optional):** `favorite` (bool), `public` (bool), `cuisine_id` (int), `category_id` (int) (int), `?category_id` (int) | — |
| `PUT` | `/api/v1/recipes/{id}` | Private | Update a recipe | See [Update Recipe Example](#update-recipe-example) |
| `DELETE` | `/api/v1/recipes/{id}` | Private | Delete a recipe | — |
| `POST` | `/api/v1/menu/generate` | Private | Generate a meal plan and shopping list | See [Generate Menu Example](#generate-menu-example) |
| `GET` | `/api/v1/ration/{id}` | Private | Get dishes within a specific ration | — |

---

## Request Examples

### Create Recipe Example
```json
{
  "name": "Some Food",
  "description": "Take food, wash it, chop it, and cook it",
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
    {"name": "Sausages", "quantity": 2},
    {"name": "Pasta", "quantity": 1},
    {"name": "Mango", "quantity": 1},
    {"name": "Cucumber", "quantity": 2},
    {"name": "Ketchup", "quantity": 1}
  ]
}
```

### Update Recipe Example
``` json
{
  "name": "Some Food upd",
  "description": "I'll just leave it to your imagination",
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
    {"name": "Ketchup", "quantity": 2},
    {"name": "Mango", "quantity": 1}
  ]
}
```

### Generate Menu Example
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