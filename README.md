# blog-aggregator
RSS feed - shows the blog name with url of the blog's RSS index.xml page where it has links to each post in the blog with details about it.

go mod vendor --> copies the code into vendor folder

sqlc - is used to write queries
goose - used to do migrations

goose postgres postgres://newuser:password:@localhost:5432/rssagg down
goose postgres postgres://newuser:password:@localhost:5432/rssagg up

users.sql - sqlc
-- name: CreateUser :one
INSERT INTO users (id,created_at,updated_at,name)
VALUES ($1,$2,$3,$4)
It means CreateUser is the function with 4 arguments (parameters) and return one output