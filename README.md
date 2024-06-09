
# Documentação da API de Despesas

## Introdução

Esta é uma API simples para gerenciar despesas. A API permite criar, ler, atualizar e deletar despesas. A aplicação foi construída usando Go, Gin e SQLite, e é executada dentro de um contêiner Docker.

## Configuração e Execução

### Pré-requisitos

- Docker
- Docker Compose

### Passos para Configuração e Execução

1. Clone o repositório do projeto:
   ```sh
   git clone https://github.com/marquesdavi/expenses-api
   cd expenses-api
   ```

2. Construa e inicie a aplicação usando Docker Compose:
   ```sh
   docker-compose up --build
   ```

3. A API estará disponível em `http://localhost:8080`.

## Endpoints da API

### 1. Criar uma Despesa

- **URL:** `/expenses`
- **Método:** `POST`
- **Headers:**
  - `Content-Type: application/json`
- **Body de Requisição:**
  ```json
  {
    "description": "Compra de supermercado",
    "amount": 150.75
  }
  ```
- **Resposta de Sucesso:**
  - **Código:** `201 Created`
  - **Body:**
    ```json
    {
      "ID": 1,
      "CreatedAt": "2024-06-09T23:00:28.000Z",
      "UpdatedAt": "2024-06-09T23:00:28.000Z",
      "DeletedAt": null,
      "description": "Compra de supermercado",
      "amount": 150.75
    }
    ```

### 2. Listar Todas as Despesas

- **URL:** `/expenses`
- **Método:** `GET`
- **Headers:** `None`
- **Body de Requisição:** `None`
- **Resposta de Sucesso:**
  - **Código:** `200 OK`
  - **Body:**
    ```json
    [
      {
        "ID": 1,
        "CreatedAt": "2024-06-09T23:00:28.000Z",
        "UpdatedAt": "2024-06-09T23:00:28.000Z",
        "DeletedAt": null,
        "description": "Compra de supermercado",
        "amount": 150.75
      }
    ]
    ```

### 3. Obter uma Despesa pelo ID

- **URL:** `/expenses/{id}`
- **Método:** `GET`
- **Headers:** `None`
- **Body de Requisição:** `None`
- **Resposta de Sucesso:**
  - **Código:** `200 OK`
  - **Body:**
    ```json
    {
      "ID": 1,
      "CreatedAt": "2024-06-09T23:00:28.000Z",
      "UpdatedAt": "2024-06-09T23:00:28.000Z",
      "DeletedAt": null,
      "description": "Compra de supermercado",
      "amount": 150.75
    }
    ```

### 4. Atualizar uma Despesa

- **URL:** `/expenses/{id}`
- **Método:** `PUT`
- **Headers:**
  - `Content-Type: application/json`
- **Body de Requisição:**
  ```json
  {
    "description": "Compra de supermercado atualizada",
    "amount": 200.00
  }
  ```
- **Resposta de Sucesso:**
  - **Código:** `200 OK`
  - **Body:**
    ```json
    {
      "ID": 1,
      "CreatedAt": "2024-06-09T23:00:28.000Z",
      "UpdatedAt": "2024-06-09T23:10:28.000Z",
      "DeletedAt": null,
      "description": "Compra de supermercado atualizada",
      "amount": 200.00
    }
    ```

### 5. Deletar uma Despesa

- **URL:** `/expenses/{id}`
- **Método:** `DELETE`
- **Headers:** `None`
- **Body de Requisição:** `None`
- **Resposta de Sucesso:**
  - **Código:** `204 No Content`
  - **Body:** `None`

## Exemplos de Requisição usando cURL

### Criar uma Despesa

```sh
curl -X POST http://localhost:8080/expenses \
-H "Content-Type: application/json" \
-d '{
  "description": "Compra de supermercado",
  "amount": 150.75
}'
```

### Listar Todas as Despesas

```sh
curl -X GET http://localhost:8080/expenses
```

### Obter uma Despesa pelo ID

```sh
curl -X GET http://localhost:8080/expenses/1
```

### Atualizar uma Despesa

```sh
curl -X PUT http://localhost:8080/expenses/1 \
-H "Content-Type: application/json" \
-d '{
  "description": "Compra de supermercado atualizada",
  "amount": 200.00
}'
```

### Deletar uma Despesa

```sh
curl -X DELETE http://localhost:8080/expenses/1
```

## Conclusão

Esta documentação fornece uma visão geral de como configurar, executar e usar a API de Despesas. Com os exemplos fornecidos, você deve ser capaz de interagir com a API para gerenciar suas despesas facilmente. Se tiver alguma dúvida ou precisar de mais informações, sinta-se à vontade para entrar em contato.
