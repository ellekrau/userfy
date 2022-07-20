Complete documentation (in Portuguese):
[here](https://ellekrau91.gitbook.io/userfy/)

# .env

Put it on projects root

````
PORT=
JWT_KEY=
````


# Dependencies

- [gin-gonic/gin](https://github.com/gin-gonic/gin) 
- [spf13/viper](https://github.com/spf13/viper)
- [github.com/go-playground/validator/v10](https://github.com/go-playground/validator)
- [stretchr/Testify](https://github.com/stretchr/testify)
- [golang-jwt/jwt](https://github.com/golang-jwt/jwt)
- [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)
- [github.com/lib/pq](https://github.com/lib/pq)

# Setup

### Build
`make build`

### Initial seed
`make seed`

### Up
`make up`

### Down
`make down`

# Containers connection
### Application
`make enter_userfy`

### MySQL
`make enter_mysql`

### Postgres
`make enter_postgres`
