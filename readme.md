# Ent Setup
install ent
<!-- go get -d entgo.io/ent/cmd/ent -->
go install entgo.io/ent/cmd/ent@latest


generate schema files
go run  entgo.io/ent/cmd/ent new Country

run code gen from schema files
go install entgo.io/ent/cmd/ent@latest
ent generate ./ent/schema/
