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
- [spf13/viper](https://github.com/spf13/viper)
- [github.com/go-playground/validator/v10](https://github.com/go-playground/validator)
- [stretchr/Testify](https://github.com/stretchr/testify)
- [netflix/go-env](https://github.com/Netflix/go-env)
- [golang-jwt/jwt](https://github.com/golang-jwt/jwt)
- [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) MySQL DB driver
- [github.com/lib/pq](https://github.com/lib/pq) Postgres DB driver

# Containers connection
### Application
`docker exec -it app sh`

### MySQL
`docker exec -it mercafacil_mysql_1 mysql -uroot -p`
- Password: admin

### Postgres
`docker exec -it mercafacil_postgresql_1 psql -d admin -U admin`
