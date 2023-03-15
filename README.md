# Comprovate :closed_lock_with_key:

## :mag: Sobre o projeto

É uma simples API que possui dois endpoints, um para se autenticar e outro para listar os produtos. Sendo pública a rota de autenticação e protegida a para listar os produtos, ou seja, é necessário um token JWT gerado pelo primeiro endpoint para fazer a requisição.

**Os endpoints:**
1. `/user/auth` [POST]: rota pública, utilizada para se autenticar e obter o token JWT com expiração.
Exemplo de payload da request para teste:
```
{
    "email": "g@g.com",
    "password": "123456"
}
```

2. `/products` [GET]: rota protegida, sendo necessário informar o token no header da requisição (bearer token).

> Na pasta ./test, há um exemplo de request para cada um dos endpoints.

## :racing_car: Como rodar?

Utilizando o comando padrão, `go run main.go`. O comando irá levantar o banco de dados, inserir os dados mockados dos produtos e de um user para teste.

## Stacks:
[![CircleCI](https://dl.circleci.com/status-badge/img/gh/marcopollivier/go-simple-icons/tree/main.svg?style=svg)]
