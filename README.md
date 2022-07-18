Complete documentation (in Portuguese):
[here](https://ellekrau91.gitbook.io/userfy/)

# .env

Put it on projects root

````
PORT=

DB_HOST=
DB_USER=
DB_PASSWORD=
DB_NAME=

CELLPHONE_PATTERN=
NAME_PATTERN=

JWT_KEY=
DATA_SECURITY_KEY=
````


# Dependencies

- [gin-gonic/gin](https://github.com/gin-gonic/gin) 
- [github.com/go-playground/validator/v10](https://github.com/go-playground/validator)
- [stretchr/Testify](https://github.com/stretchr/testify)
- [joho/go-dotenv](https://github.com/joho/godotenv)
- [netflix/go-env](https://github.com/Netflix/go-env)
- [golang-jwt/jwt](https://github.com/golang-jwt/jwt)
- [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) MySQL DB driver
- [github.com/lib/pq](https://github.com/lib/pq) Postgres DB driver

Migration package:

`go install -tags 'mysql,postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`

# Migrations

Run MySQL initial migration:

`docker exec -it app migrate -path src/database/migrations/mysql -database "mysql://admin:admin@tcp(mysql:3306)/admin" --verbose up`

Run Postgres initial migration:

`docker exec -it app migrate -path src/database/migrations/postgres -database "postgres://admin:admin@postgresql:5432/admin?sslmode=disable" --verbose up`

# Containers connection
### Application
`docker exec -it app sh`

### MySQL
`docker exec -it mercafacil_mysql_1 mysql -uroot -p`
- Password: admin

### Postgres
`docker exec -it mercafacil_postgresql_1 psql -d admin -U admin`
