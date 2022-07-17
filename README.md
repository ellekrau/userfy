## Dependencies

- [gin-gonic/gin](https://github.com/gin-gonic/gin)
- [stretchr/Testify](https://github.com/stretchr/testify)
- [joho/go-dotenv](https://github.com/joho/godotenv)
- [netflix/go-env](https://github.com/Netflix/go-env)
- [golang-jwt/jwt](https://github.com/golang-jwt/jwt)
- [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) MySQL DB driver
- [github.com/lib/pq](https://github.com/lib/pq) Postgres DB driver

# Migrations

`migrate create -ext sql -dir src\database\migrations\mysql -seq create_table_contacts`

## MySQL
`go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`

`migrate -path src/database/migrations/mysql -database "mysql://admin:admin@tcp(localhost:3306)/admin" --verbose up`

### Connect to MySQL container
`docker exec -it mercafacil_mysql_1 mysql -uroot -p`
- Password: admin
- SHOW DATABASES;
- USE admin
- SHOW TABLES;

## Postgres

`go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`

`migrate -path src/database/migrations/postgres -database "postgres://admin:admin@localhost:5432/admin?sslmode=disable" --verbose up`

## Connect to Postgres container

`docker exec -it mercafacil_postgresql_1 psql -d admin -U admin`