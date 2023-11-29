.PHONY: runapi runteardown runbuildup

localapi:
	go run dbapp/cmd/api/main.go
apiup:
	docker-compose -f dbapp/docker-compose/docker-compose.yml up
apidown:
	docker-compose -f dbapp/docker-compose/docker-compose.yml down
apilog:
	docker-compose -f dbapp/docker-compose/docker-compose.yml logs -f
apibuild:
	docker-compose -f dbapp/docker-compose/docker-compose.yml build --no-cache
apishell:
	docker-compose -f dbapp/docker-compose/docker-compose.yml exec api /bin/bash
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
dbup:
	docker-compose -f mariadb/docker-compose.yml up -d
dbdown:
	docker-compose -f mariadb/docker-compose.yml down
dblogs:
	docker-compose -f mariadb/docker-compose.yml logs -f
dbwipe:
	docker-compose -f mariadb/docker-compose.yml down --volumes
tsup:
	docker-compose -f troubleshooter/docker-compose.yml up -d
tsdown:
	docker-compose -f troubleshooter/docker-compose.yml down
tsshell:
	docker-compose -f troubleshooter/docker-compose.yml exec troubleshooter /bin/bash
tsbuild:
	docker-compose -f troubleshooter/docker-compose.yml build --no-cache
