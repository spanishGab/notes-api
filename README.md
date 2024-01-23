# tasks-api

## Requisitos

A API deverá permitir:

- **Cadastrar uma nova conta de usuário:**

  - route: /api/v1/signup
  - method: POST

  - **Request:**

    - headers:
      - Content-Type: application/json
    - body:
      ```json
      {
        "firstName": "firstName",
        "lastName": "lastName",
        "nickname": "nickname",
        "birthday": "16/04/2000",
        "email": "email@example.com",
        "password": "password"
      }
      ```

  - **Response:**
    - https status code: **201 Created**
      - Location: http://127.0.0.1:5000/api/v1/users/:id
      - Content-Type: application/json
      - body:
        ```json
        {
          "user": {
            "userId": 123,
            "firstName": "firstName",
            "lastName": "lastName",
            "birthday": "16/04/2000",
            "nickname": "nickname",
            "email": "firstName.lastName@example.com"
          }
        }
        ```

- **Entrar em uma conta na API:**

  - route: /api/v1/login
  - method: POST

  - **Request:**

    - headers:
      - Content-Type: application/json
    - body:
      ```json
      {
        "email": "firstName.lastName@example.com",
        "password": "password"
      }
      ```

  - **Response:**
  - https status code: **200 OK**
    - Content-Type: application/json
    - body:
      ```json
      {
        "token": "token",
        "userId": 123
      }
      ```

- **Sair de uma conta na API:**

  - route: /api/v1/logout
  - method: POST

  - **Request:**

    - headers:
      - Content-Type: application/json
      - Authorization: (JWT)
    - body:
      ```json
      {
        "userId": 123
      }
      ```

  - **Response:**

    - headers:
      - Content-Type: application/json
    - https status code: **204 No Content**

- **Criar uma nova tarefa:**

  - route: /api/v1/tasks
  - method: POST

  - **Request:**

    - headers:
      - Content-Type: application/json
      - Authorization: (JWT)
    - body:
      ```json
      {
        "title": "any title",
        "content": "any content"
      }
      ```

  - **Response:**

    - headers:
      - Content-Type: application/json
      - Location: http://127.0.0.1:5000/api/v1/tasks
    - body:
      ```json
      {
        "id": 123,
        "title": "any title",
        "content": "any content",
        "status": "PENDING"
      }
      ```
    - https status code: **201 Created**

* **Buscar tarefa pertencente ao usuário:**

  - route: /api/v1/tasks/:id
  - method: GET

  - **Request:**

    - headers:
      - Content-Type: application/json
      - Authorization: (JWT)

  - **Response:**

    - headers:
      - Content-Type: application/json
      - Authorization: (JWT)
    - body:
      ```json
      {
        "id": 123,
        "title": "any title",
        "content": "any content"
      }
      ```
    - https status code: **200 OK**

* **Listar as tarefas do usuário:**

  - route: /api/v1/tasks
  - method: GET

  - **Request:**

    - headers:
      - Content-Type: application/json
      - Authorization: (JWT)

  - **Response:**

    - headers:
      - Content-Type: application/json
      - Authorization: (JWT)
    - body:
      ```json
      [
        {
          "id": 123,
          "title": "any title",
          "content": "any content",
        },
        {
          "id": 123,
          "title": "any title",
          "content": "any content",
        },
      ];
      ```
    - https status code: **200 OK**

* **Editar o titulo de uma tarefa:**

  - route: /api/v1/tasks/:id/title
  - method: /PATCH

  - **Request:**

    - headers:
      - Content-Type: application/json
      - Authorization: (JWT)
    - body:
      ```json
      {
        "title": "any title"
      }
      ```

  - **Response:**

    - headers:
      - Content-Type: application/json
      - Authorization: (JWT)
    - body:
      ```json
      {
        "id": 123,
        "title": "any title",
        "content": "any content",
        "status": "FINISHED"
      }
      ```
    - https status code: **200 OK**

* **Editar o texto de uma tarefa:**

  - route: /api/v1/tasks/:id/text
  - method: /PATCH

  - **Request:**

    - headers:
      - Content-Type: application/json
      - Authorization: (JWT)
    - body:
      ```json
      {
        "content": "any content"
      }
      ```

  - **Response:**

    - headers:
      - Content-Type: application/json
      - Authorization: (JWT)
    - body:
      ```json
      {
        "id": 123,
        "title": "any title",
        "content": "any content",
        "status": "FINISHED"
      }
      ```
    - https status code: **200 OK**

* **Editar o status de uma tarefa**

  - route: /api/v1/tasks/:id/status
  - method: PATCH

  - **Request:**

    - headers:
      - Content-Type: application/json
      - Authorization: (JWT)
    - body:
      ```json
      {
        "status": "FINISHED"
      }
      ```

  - **Response:**

    - headers:
      - Content-Type: application/json
      - Authorization: (JWT)
    - body:
      ```json
      {
        "id": 123,
        "title": "any title",
        "content": "any content",
        "status": "FINISHED"
      }
      ```
    - https status code: **200 OK**

* **Apagar uma tarefa do sistema:**

  - route: /api/v1/tasks/:id
  - method: DELETE

  - **Request:**

    - headers:
      - Content-Type: application/json
      - Authorization: (JWT)
    - body:
      ```json
      {
        "id": 123
      }
      ```

  - **Response:**

    - headers:
      - Content-Type: application/json
      - Authorization: (JWT)
    - https status code: **200 OK**
