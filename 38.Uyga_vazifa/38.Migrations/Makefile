DB_URL=postgres://postgres:abdulaziz1221@localhost:5432/h35?sslmode=disable

run:
	go run cmd/main.go
  
migrate_up:
	migrate -path migrations -database ${DB_URL} -verbose up

migrate_down:
	migrate -path migrations -database ${DB_URL} -verbose down

migrate_force:
	migrate -path migrations -database ${DB_URL} -verbose force 1

migrate_file:
	migrate create -ext sql -dir migrations -seq gin

swag-gen:
	~/go/bin/swag init -g ./api/router.go -o api/docs force 1
