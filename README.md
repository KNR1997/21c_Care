goose -dir ./migrations postgres "host=localhost user=myuser password=mypassword dbname=mydatabase port=5432 sslmode=disable" up

swag init -g main.go -d cmd/service,internal
