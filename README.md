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
        "description": "any description",
        "isLocked": false,
        "color": "#000000"
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
        "description": "any description",
        "status": "TODO",
        "isLocked": false,
        "color": "#000000"
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
    - body:
      ```json
      {
        "id": 123,
        "title": "any title",
        "description": "any description",
        "status": "TODO",
        "isLocked": false,
        "color": "#000000"
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
    - body:
      ```json
      {
        "data": [
          {
            "id": 123,
            "title": "any title",
            "description": "any description",
            "status": "TODO",
            "isLocked": false,
            "color": "#000000"
          },
          {
            "id": 123,
            "title": "any title",
            "description": "any description",
            "status": "TODO",
            "isLocked": false,
            "color": "#000000"
          }
        ],
        "page": 1,
        "count": 5,
        "limit": 10,
        "itemsPerPage": 10,
        "hasNext": true
      }
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
    - body:
      ```json
      {
        "id": 123,
        "title": "any title",
        "description": "any description",
        "status": "TODO",
        "isLocked": false,
        "color": "#000000"
      }
      ```
    - https status code: **200 OK**

* **Editar a descrição de uma tarefa:**

  - route: /api/v1/tasks/:id/description
  - method: /PATCH

  - **Request:**

    - headers:
      - Content-Type: application/json
      - Authorization: (JWT)
    - body:
      ```json
      {
        "description": "any description"
      }
      ```

  - **Response:**

    - headers:
      - Content-Type: application/json
    - body:
      ```json
      {
        "id": 123,
        "title": "any title",
        "description": "any description",
        "status": "TODO",
        "isLocked": false,
        "color": "#000000"
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
        "status": "TODO"
      }
      ```

  - **Response:**

    - headers:
      - Content-Type: application/json
    - body:
      ```json
      {
        "id": 123,
        "title": "any title",
        "description": "any description",
        "status": "TODO",
        "isLocked": false,
        "color": "#000000"
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
    - https status code: **204 NO CONTENT**

* **Bloquear uma tarefa do sistema:**

  - route: /api/v1/tasks/:id/unlock
  - method: PATCH

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
    - body:
      ```json
      {
        "id": 123,
        "title": "any title",
        "description": "any description",
        "status": "TODO",
        "isLocked": true,
        "color": "#000000"
      }
      ```
    - https status code: **200 OK**

* **Desbloquear uma tarefa do sistema:**

  - route: /api/v1/tasks/:id/lock
  - method: PATCH

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
    - body:
      ```json
      {
        "id": 123,
        "title": "any title",
        "description": "any description",
        "status": "TODO",
        "isLocked": false,
        "color": "#000000"
      }
      ```
    - https status code: **200 OK**
