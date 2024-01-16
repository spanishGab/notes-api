# notes-api

## Requisitos

A API deverá permitir:
- cadastrar uma nova conta de usuário
  - route: /signup
  - method: POST
- entrar e sair de uma conta na API 
  - route: /logout
  - method: POST
- criar uma nova nota
  - route: /notes
  - method: POST
* buscar uma pertencente ao usuário nota 
  - route: /
  - method: 
* listar as notas do usuário 
  - route: /notes/:id
  - method: GET
* editar o texto de uma nota 
  - route: /notes/:id
  - method: /PUT
* editar o status de uma nota 
  - route: /notes/:id
  - method: PATCH (atualizacao parcial?)
* apagar uma nota do sistema 
  - route: /notes/:id
  - method: DELETE
