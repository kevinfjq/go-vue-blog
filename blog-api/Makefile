#Docker run
#docker run -d --name go_blog_db -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=mypasswordsecret -e POSTGRES_DB=blogdb -p 5432:5432 -v go_blog_pg_data:/var/lib/postgresql/data postgres:latest

#Create migration
#migrate create -ext sql -dir ./migrations -seq migration_name

# Run migration
# migrate -path ./migrations -database "postgres://postgres:password@localhost:5432/postgres?sslmode=disable" up

# Run migration down
# migrate -path ./migrations -database "postgres://postgres:password@localhost:5432/postgres?sslmode=disable" down