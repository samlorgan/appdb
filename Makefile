.PHONY: runapi runteardown runbuildup

api:
	go run dbapp/cmd/api/main.go
teardown:
	go run dbapp/cmd/drop/main.go
migrate:
	go run dbapp/cmd/migrate/main.go
populate:
	go run dbapp/cmd/populate/main.go
entinit:
	go run  entgo.io/ent/cmd/ent new Country	
entgen:
	ent generate dbapp/ent/schema
dbstart:
	docker-compose -f mariadb/docker-compose.yml up -d
dbstop:
	docker-compose -f mariadb/docker-compose.yml down
dblogs:
	docker-compose -f mariadb/docker-compose.yml logs -f
dbwipe:
	docker-compose -f mariadb/docker-compose.yml down --volumes