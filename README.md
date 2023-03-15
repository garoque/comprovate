# Comprovate :closed_lock_with_key:

<p float="left">
    <img src="https://go.dev/blog/go-brand/Go-Logo/SVG/Go-Logo_Aqua.svg" height="50">
    <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/3/38/SQLite370.svg/764px-SQLite370.svg.png?20140602232932" height="50">
    <img src="https://camo.githubusercontent.com/f72d07b7d898f8935d557867df17416a1b430a2572f8ea1bae57d1700f5c754b/68747470733a2f2f63646e2e7261776769742e636f6d2f676f2d6368692f6368692f6d61737465722f5f6578616d706c65732f6368692e737667" height="50">
</p>

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