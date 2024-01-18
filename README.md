# notes-api

## Requisitos

A API deverá permitir:
- Cadastrar uma nova conta de usuário:
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
            "id": 123,
            "firstname": "firstname",
            "lastname": "lastname",
            "nickname": "nickname",
            "email": "novo.usuario@example.com"
          }
        }
        ```
    
- Entrar em uma conta na API:
  - route: /api/v1/logout
  - method: POST
  
  - **Request:**
    - headers:
      - Content-Type: application/json
    - body:
      ```javascript
      {
        "email": "novo.usuario@example.com",
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
    
- Sair de uma conta na API:
  - route: /api/v1/login
  - method: POST
  
  - **Request:**

    - headers:
      - Content-Type: application/json
      
    - body:
    - query parameters: 
    - path parameters: 

  - **Response:**

    - headers:
      - Content-Type: application/json
       
    - body:
    - https status code:
    
- criar uma nova nota
  - route: /api/v1/notes
  - method: POST
  
  - **Request:**

    - headers:
      - Content-Type: application/json
      
    - body:
    - query parameters: 
    - path parameters: 

  - **Response:**

    - headers:
      - Content-Type: application/json
       
    - body:
    - https status code:
    
* buscar nota pertencente ao usuário 
  - route: /api/v1notes/:id
  - method: GET
  
  - **Request:**

    - headers:
      - Content-Type: application/json
      
    - body:
    - query parameters: 
    - path parameters: 

  - **Response:**

    - headers:
      - Content-Type: application/json
       
    - body:
    - https status code:
    
* listar as notas do usuário 
  - route: /api/v1/notes
  - method: GET
  
  - **Request:**

    - headers:
      - Content-Type: application/json
      
    - body:
    - query parameters: 
    - path parameters: 

  - **Response:**

    - headers:
      - Content-Type: application/json
       
    - body:
    - https status code:
    
* editar o texto de uma nota 
  - route: /api/v1/notes/:id/text
  - method: /PATCH
  
  - **Request:**

    - headers:
      - Content-Type: application/json
      
    - body:
    - query parameters: 
    - path parameters: 

  - **Response:**

    - headers:
      - Content-Type: application/json
       
    - body:
    - https status code:
    
* editar o status de uma nota 
  - route: /api/v1/notes/:id/status
  - method: PATCH 
  
  - **Request:**

    - headers:
      - Content-Type: application/json
      
    - body:
    - query parameters: 
    - path parameters: 

  - **Response:**

    - headers:
      - Content-Type: application/json
       
    - body:
    - https status code:
    
* apagar uma nota do sistema 
  - route: /api/v1/notes/:id
  - method: DELETE
  
  - **Request:**

    - headers:
      - Content-Type: application/json
      
    - body:
    - query parameters: 
    - path parameters: 

  - **Response:**

    - headers:
      - Content-Type: application/json
       
    - body:
    - https status code:
    
