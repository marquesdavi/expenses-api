
# Expenses API Documentation

## Introduction

This is a simple API for managing expenses. The API allows you to create, read, update, and delete expenses. The application was built using Go, Gin, and SQLite, and runs inside a Docker container.

## Setup and Execution

### Prerequisites

- Docker
- Docker Compose

### Setup and Execution Steps

1. Clone the project repository:
   ```sh
   git clone https://github.com/marquesdavi/expenses-api
   cd expenses-api
   ```

2. Build and start the application using Docker Compose:
   ```sh
   docker-compose up --build
   ```

3. The API will be available at `http://localhost:8080`.

## API Endpoints

### 1. Create an Expense

- **URL:** `/expenses`
- **Method:** `POST`
- **Headers:**
  - `Content-Type: application/json`
- **Request Body:**
  ```json
  {
    "description": "Grocery shopping",
    "amount": 150.75
  }
  ```
- **Success Response:**
  - **Code:** `201 Created`
  - **Body:**
    ```json
    {
      "ID": 1,
      "CreatedAt": "2024-06-09T23:00:28.000Z",
      "UpdatedAt": "2024-06-09T23:00:28.000Z",
      "DeletedAt": null,
      "description": "Grocery shopping",
      "amount": 150.75
    }
    ```

### 2. List All Expenses

- **URL:** `/expenses`
- **Method:** `GET`
- **Headers:** `None`
- **Request Body:** `None`
- **Success Response:**
  - **Code:** `200 OK`
  - **Body:**
    ```json
    [
      {
        "ID": 1,
        "CreatedAt": "2024-06-09T23:00:28.000Z",
        "UpdatedAt": "2024-06-09T23:00:28.000Z",
        "DeletedAt": null,
        "description": "Grocery shopping",
        "amount": 150.75
      }
    ]
    ```

### 3. Get an Expense by ID

- **URL:** `/expenses/{id}`
- **Method:** `GET`
- **Headers:** `None`
- **Request Body:** `None`
- **Success Response:**
  - **Code:** `200 OK`
  - **Body:**
    ```json
    {
      "ID": 1,
      "CreatedAt": "2024-06-09T23:00:28.000Z",
      "UpdatedAt": "2024-06-09T23:00:28.000Z",
      "DeletedAt": null,
      "description": "Grocery shopping",
      "amount": 150.75
    }
    ```

### 4. Update an Expense

- **URL:** `/expenses/{id}`
- **Method:** `PUT`
- **Headers:**
  - `Content-Type: application/json`
- **Request Body:**
  ```json
  {
    "description": "Updated grocery shopping",
    "amount": 200.00
  }
  ```
- **Success Response:**
  - **Code:** `200 OK`
  - **Body:**
    ```json
    {
      "ID": 1,
      "CreatedAt": "2024-06-09T23:00:28.000Z",
      "UpdatedAt": "2024-06-09T23:10:28.000Z",
      "DeletedAt": null,
      "description": "Updated grocery shopping",
      "amount": 200.00
    }
    ```

### 5. Delete an Expense

- **URL:** `/expenses/{id}`
- **Method:** `DELETE`
- **Headers:** `None`
- **Request Body:** `None`
- **Success Response:**
  - **Code:** `204 No Content`
  - **Body:** `None`

## cURL Request Examples

### Create an Expense

```sh
curl -X POST http://localhost:8080/expenses \
-H "Content-Type: application/json" \
-d '{
  "description": "Grocery shopping",
  "amount": 150.75
}'
```

### List All Expenses

```sh
curl -X GET http://localhost:8080/expenses
```

### Get an Expense by ID

```sh
curl -X GET http://localhost:8080/expenses/1
```

### Update an Expense

```sh
curl -X PUT http://localhost:8080/expenses/1 \
-H "Content-Type: application/json" \
-d '{
  "description": "Updated grocery shopping",
  "amount": 200.00
}'
```

### Delete an Expense

```sh
curl -X DELETE http://localhost:8080/expenses/1
```

