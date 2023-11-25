launch db via 
`docker compose up`
ports will be available on 8080 for the adminer app and 3306 for direct connection
the initdb.sql script is injected into the container at creation, and is autorun by the db at initial startup
to force the rerun of the init script
`docker compose down --volumes`
will reset and clear storage after which 
`docker compose up`
can be used again