build:
	docker-compose up -d --build

up:
	docker-compose up -d

down:
	docker-compose down

seed:
	docker exec -i mercafacil_postgresql psql -U admin < sql/create-table-varejao.sql
	docker exec -i mercafacil_mysql mysql -u root --password=admin admin < sql/create-table-macapa.sql

enter_mysql:
	docker exec -it mercafacil_mysql mysql -u root --password=admin admin

enter_postgres:
	docker exec -it mercafacil_postgresql psql -d admin -U admin

enter_userfy:
	docker exec -it userfy sh


