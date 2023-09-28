DB_URL=mysql://root:433205ari@tcp(127.0.0.1:3306)/perumahan2

createmigrate:
	migrate create -ext sql -dir db/migration -seq init_schema

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/ariandi/ppob_go/db/sqlc Store

.PHONY: sqlc test server mock createmigrate migrateup migratedown


# "migrate -path db/migration -database "mysql://root:admin@tcp(127.0.0.1:3303)/perumahan" -verbose down"
# "migrate -path db/migration -database "mysql://root:admin@tcp(127.0.0.1:3303)/perumahan" -verbose up"
# docker run --rm -v "$(pwd):/src" -w /src sqlc/sqlc generate