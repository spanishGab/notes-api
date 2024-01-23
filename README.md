# notes-api

## Requisitos

A API deverá permitir:

- **Cadastrar uma nova conta de usuário:**

  - route: /api/v1/signup
  - method: POST

  - **Request:**

    - headers:
      - Content-Type: application/json
      - Authorization: (JWT)
    - body:
      ```javascript
      {
        "firstname": "firstname",
        "lastname": "lastname",
        "nickname": "nickname",
        "email": "email@example.com",
        "password": "password"
      }
      ```

  - **Response:**
    - https status code: **201 Created**
      - Location: /api/v1/users/:id
      - Content-Type: application/json
      - body:
        ```javascript
        {
          "message": "User created successfully!",
          "user": {
            "userId": 123,
            "firstname": "firstname",
            "lastname": "lastname",
            "nickname": "nickname",
            "email": "firstname.lastname@example.com"
          }
        }
        ```

- **Entrar em uma conta na API:**

  - route: /api/v1/login
  - method: POST

  - **Request:**

    - headers:
      - Content-Type: application/json
      - Authorization: (JWT)
    - body:
      ```javascript
      {
        "email": "firstname.lastname@example.com",
        "password": "password"
      }
      ```

  - **Response:**
  - https status code: **200 OK**
    - Location: /api/v1/users/:id
    - Content-Type: application/json
    - body:
      ```javascript
      {
        "token": "token",
        "userId": 123,
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
      ```javascript
      {
        "userId": 123,
      }
      ```

  - **Response:**

    - headers:
      - Content-Type: application/json
    - https status code: **204 No Content**

- **Criar uma nova nota:**

  - route: /api/v1/notes
  - method: POST

  - **Request:**

    - headers:
      - Content-Type: application/json
      - Authorization: (JWT)
    - body:
      ```javascript
      {
        "title": "any title",
        "content": "any content",
      }
      ```

  - **Response:**

    - headers:
      - Content-Type: application/json
    - body:
      ```javascript
      {
        "id": 123,
        "title": "any title",
        "content": "any content",
        "status": "PENDING",
      }
      ```
    - https status code: **201 Created**

* **Buscar nota pertencente ao usuário:**

  - route: /api/v1/notes/:id
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
      ```javascript
      "data": {
          "id": 123,
          "title": 'any title',
          "content": 'any content',
        }
      ```
    - https status code: **200 OK**

* **Listar as notas do usuário:**

  - route: /api/v1/notes
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
      ```javascript
      [
        {
          id: 123,
          title: 'any title',
          content: 'any content',
        },
        {
          id: 123,
          title: 'any title',
          content: 'any content',
        },
      ];
      ```
    - https status code: **200 OK**

* **Editar o titulo de uma nota:**

  - route: /api/v1/notes/:id/title
  - method: /PATCH

  - **Request:**

    - headers:
      - Content-Type: application/json
      - Authorization: (JWT)
    - body:
      ```javascript
      {
        "title": 'any title',
      }
      ```

  - **Response:**

    - headers:
      - Content-Type: application/json
      - Authorization: (JWT)
    - body:
      ```javascript
      {
        "id": 123,
        "title": 'any title',
        "content": 'any content',
        "status": "FINISHED",
      }
      ```
    - https status code: **200 OK**

* **Editar o texto de uma nota:**

  - route: /api/v1/notes/:id/text
  - method: /PATCH

  - **Request:**

    - headers:
      - Content-Type: application/json
      - Authorization: (JWT)
    - body:
      ```javascript
      {
        "content": 'any content',
      }
      ```

  - **Response:**

    - headers:
      - Content-Type: application/json
      - Authorization: (JWT)
    - body:
      ```javascript
      {
        "id": 123,
        "title": 'any title',
        "content": 'any content',
        "status": "FINISHED",
      }
      ```
    - https status code: **200 OK**

* **Editar o status de uma nota**

  - route: /api/v1/notes/:id/status
  - method: PATCH

  - **Request:**

    - headers:
      - Content-Type: application/json
      - Authorization: (JWT)
    - body:
      ```javascript
      {
        "status": 'FINISHED',
      }
      ```

  - **Response:**

    - headers:
      - Content-Type: application/json
      - Authorization: (JWT)
    - body:
      ```javascript
      {
        "id": 123,
        "title": 'any title',
        "content": 'any content',
        "status": "FINISHED",
      }
      ```
    - https status code: **200 OK**

* apagar uma nota do sistema

  - route: /api/v1/notes/:id
  - method: DELETE

  - **Request:**

    - headers:
      - Content-Type: application/json
      - Authorization: (JWT)
    - body:
      ```javascript
      {
       "id": 123,
      }
      ```

  - **Response:**

    - headers:
      - Content-Type: application/json
      - Authorization: (JWT)
    - https status code: **200 OK**
