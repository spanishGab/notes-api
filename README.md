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
