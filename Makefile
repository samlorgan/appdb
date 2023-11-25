.PHONY: runapi runteardown runbuildup

api:
	go run dbapp/cmd/api/main.go
dockerapi:
	docker-compose -f dbapp/docker-compose/docker-compose.yml up
drop:
	cd dbapp/cmd/drop && go run main.go
migrate:
	cd dbapp/cmd/migrate && go run main.go
populate:
	cd dbapp/cmd/populate && go run main.go
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