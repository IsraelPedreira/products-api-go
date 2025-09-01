# 🛒 Products API - Go

Uma API REST construída com **Go**, **Gin Framework** e **PostgreSQL**.

## 🚀 Funcionalidades

- ✅ **CRUD completo** de produtos
- 📖 **Documentação Swagger** interativa
- 🐘 **Banco PostgreSQL** com Docker
- 🔄 **Migrations automáticas** com GORM
- 🏗️ **Arquitetura MVC** organizada
- ✨ **Validação de dados** automática

## 🛠️ Tecnologias

- **Go**
- **Gin Web Framework** - HTTP router
- **GORM** - ORM para Go
- **PostgreSQL** - Banco de dados
- **Swagger** - Documentação da API
- **Docker** - Containerização
- **NGINX** - Proxy reverso e load balancer

## 📋 Pré-requisitos

- [Go 1.24+](https://golang.org/dl/)
- [Docker](https://docs.docker.com/get-docker/) e [Docker Compose](https://docs.docker.com/compose/install/)
- [Git](https://git-scm.com/)

## 🏃‍♂️ Como executar

### 1. Clone o repositório
```bash
git clone https://github.com/IsraelPedreira/api-go.git
cd marketplace-go
```

### 2. Inicie o banco de dados
```bash
docker-compose up -d
```

### 3. Configure as variáveis de ambiente
Crie um arquivo `.env` na raiz do projeto:
```env
DB_HOST=
DB_PORT=
DB_USER=
DB_PASSWORD=
DB_NAME=
DB_SSLMODE=
```

### 4. Instale as dependências
```bash
go mod tidy
```

### 5. Gere a documentação Swagger
```bash
swag init --parseDependency --parseInternal
```

### 6. Execute a aplicação
```bash
go run main.go
```

A API estará disponível em: `http://localhost:8080`

## 📚 Documentação

### Swagger UI
Acesse a documentação interativa em: `http://localhost:8080/swagger/index.html`

### Endpoints

| Método | Endpoint | Descrição |
|--------|----------|-----------|
| `POST` | `/product` | Criar novo produto |
| `GET` | `/product` | Listar todos os produtos |
| `GET` | `/product/{id}` | Buscar produto por ID |
| `PUT` | `/product/{id}` | Atualizar produto |
| `DELETE` | `/product/{id}` | Deletar produto |

### Exemplo de Produto
```json
{
  "name": "iPhone 14",
  "description": "Smartphone Apple",
  "price": 4999.99
}
```

## 🏗️ Estrutura do Projeto

```
marketplace-go/
├── docs/                   # Documentação Swagger gerada
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── src/
│   ├── configuration/      # Configurações
│   ├── controllers/        # Controllers da API
│   │   └── ProductController.go
│   ├── database/          # Conexão com banco
│   │   └── database.go
│   ├── models/            # Modelos/Entidades
│   │   └── Product.go
│   └── routes/            # Rotas da API
│       └── ProductRoutes.go
├── docker-compose.yaml    # Docker Compose para PostgreSQL
├── go.mod                 # Dependências Go
├── go.sum
├── main.go               # Arquivo principal
└── README.md
```

## 🧪 Testando a API

### Criar produto
```bash
curl -X POST http://localhost:8080/product \
  -H "Content-Type: application/json" \
  -d '{
    "name": "iPhone 14",
    "description": "Smartphone Apple",
    "price": 4999.99
  }'
```

### Listar produtos
```bash
curl http://localhost:8080/product
```

### Buscar produto por ID
```bash
curl http://localhost:8080/product/1
```

## 🐛 Comandos Úteis

### Gerar documentação Swagger
```bash
swag init --parseDependency --parseInternal
```

### Ver logs do banco
```bash
docker-compose logs go_db
```

### Parar containers
```bash
docker-compose down
```

### Resetar banco de dados
```bash
docker-compose down -v
docker-compose up -d
```

## 📦 Dependências Principais

```go
require (
    github.com/gin-gonic/gin
    github.com/joho/godotenv
    github.com/swaggo/gin-swagger
    github.com/swaggo/files
    gorm.io/gorm
    gorm.io/driver/postgres
)
```

## 🤝 Como Contribuir

1. Faça um fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## 👨‍💻 Autor

**Israel Pedreira**
- GitHub: [@IsraelPedreira](https://github.com/IsraelPedreira)
- LinkedIn: [Israel Pedreira](https://www.linkedin.com/in/israel-pedreira/)


