## Dependencies

- [gin-gonic/gin](https://github.com/gin-gonic/gin)
- [stretchr/Testify](https://github.com/stretchr/testify)
- [joho/go-dotenv](https://github.com/joho/godotenv)
- [netflix/go-env](https://github.com/Netflix/go-env)
- [golang-jwt/jwt](https://github.com/golang-jwt/jwt)
- [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)

# Migrations

## MySQL
`go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`

`migrate -path src/database/migrations/mysql -database "mysql://admin:admin@tcp(localhost:3306)/admin" --verbose up`

### Connect to MySQL container

- docker exec -it mercafacil_mysql_1 mysql -uroot -p
- Password: admin
- SHOW DATABASES;
- USE admin
- SHOW TABLES;