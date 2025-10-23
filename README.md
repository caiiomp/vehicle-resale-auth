# Serviço de Autenticação - JWT

Este repositório contém o serviço de autenticação que gerencia o registro e a autorização dos compradores utilizando **JWT (JSON Web Tokens)**. O serviço é separado da API de veículos para garantir a segurança e a separação dos dados de autenticação dos dados transacionais.

## Funcionalidades

- **Cadastro de usuários (compradores):** Permite o registro de novos compradores.
- **Login de usuários:** Permite o login dos compradores e gera um token JWT para autenticação.
- **Autenticação:** Protege os endpoints da API de veículos, validando o token JWT nas requisições.

## Tecnologias Utilizadas

- **Go (Golang):** Para o desenvolvimento do serviço de autenticação.
- **JWT (JSON Web Tokens):** Para a geração e validação de tokens de autenticação.
- **MongoDB:** Para armazenar os dados dos usuários autenticados.
- **Gin:** Framework web para o desenvolvimento da API.

## Como Rodar o Projeto Localmente

### 1. Pré-requisitos

Antes de rodar o serviço localmente, verifique se você tem as seguintes dependências instaladas:

- **Go (Golang)** versão 1.18 ou superior.
- **Git** para clonar o repositório.
- **Docker** e **Docker Compose** (caso queira rodar o MongoDB via Docker).

### 2. Configuração do MongoDB com Docker Compose

1. Clone o repositório:

    ```bash
    git clone git@github.com:caiiomp/vehicle-resale-auth.git
    ```

2. Na raiz do projeto, inicie o MongoDB com o `docker`:

    ```bash
    docker compose up -d
    ```

    Isso irá iniciar o MongoDB em um contêiner.

### 3. Configuração da API de Autenticação

1. Na raiz do projeto instale as dependências do Go:

    ```bash
    go mod tidy
    ```

2. Inicie o servidor de autenticação:

    ```bash
    go run src/main.go
    ```

    A API de autenticação estará disponível em `http://localhost:8080`.

### 4. Testando a API de Autenticação

Use **Postman**, **Insomnia**, **cURL** ou qualquer outro cliente **HTTP** para testar os endpoints:

app.POST("/users", service.create)
	app.GET("/users", service.search)
	app.GET("/users/:user_id", service.get)

- `POST /users` - Registrar um novo usuário (comprador).
- `POST /login` - Realizar login e obter um token JWT.
- `GET /users` - Listar todos os usuários.
- `GET /users/:user_id` - Buscar um usuário pelo id.

Os testes unitários e os testes de integração podem ser executados da seguinte forma respectivamente:
```bash
    go test ./... -v
    go test -tags=integration -v ./...
```

### 5. Protegendo Endpoints com JWT

O token JWT gerado no login pode ser utilizado para autenticar o comprador em outros serviços (como o serviço de API de veículos). Para isso, inclua o token no cabeçalho de autorização das requisições:

```
Authorization: Bearer <token-jwt>
```

## Documentação (Swagger)

Para acessar a documentação do serviço, acessar o seguinte endpoint: 
```
http://localhost:8080/swagger/index.html
```
