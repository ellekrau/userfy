build:
	docker-compose up -d --build

up:
	docker-compose up -d

down:
	docker-compose down

seed:
	docker exec -i mercafacil_postgresql_1 psql -U admin < sql/create-table-varejao.sql
	docker exec -i mercafacil_mysql_1 mysql -u root --password=admin admin < sql/create-table-macapa.sql


