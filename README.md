# notes-api

## Requisitos

A API deverá permitir:
- cadastrar uma nova conta de usuário
  - route: /api/v1/signup
  - method: POST
- entrar em uma conta na API 
  - route: /api/v1/logout
  - method: POST
- sair de uma conta na API
  - route: /api/v1/login
  - method: POST
- criar uma nova nota
  - route: /api/v1/notes
  - method: POST
* buscar nota pertencente ao usuário 
  - route: /api/v1/notes/:id
  - method: GET
* listar as notas do usuário 
  - route: /api/v1/notes
  - method: GET
* editar o texto de uma nota 
  - route: /api/v1/notes/:id/text
  - method: /PATCH
* editar o status de uma nota 
  - route: /api/v1/notes/:id/status
  - method: PATCH 
* apagar uma nota do sistema 
  - route: /api/v1/notes/:id
  - method: DELETE


## Descrição técnica das rotas

Nessa parte cada rota deverá ser descrita no seguinte formato:

## path da rota
### request
headers: (se houver)

body: (se houver)

query parameters: (se houver)

path parameters: (se houver)

### response
headers: (se houver)

body: (se houver)

HTTP status code: