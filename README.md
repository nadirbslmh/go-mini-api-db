# go-mini-api-db

Simple REST API application written in Go with MySQL database.

## How to Use

1. Clone this repository.

2. Create the database.

```sql
CREATE DATABASE learndocker;
```

3. Copy the configuration file for the database. After that, fill the configuration inside `.env` file.

```sh
cp .env.example .env
```

4. Run the application.

```sh
go run main.go
```

5. Test the application. You can use REST API client like Postman if you want.

```sh
# create new posts
curl -XPOST -H "Content-type: application/json" -d '{"title":"first blog","content":"this is the first blog"}' 'http://localhost:1323/posts'

curl -XPOST -H "Content-type: application/json" -d '{"title":"second blog","content":"this is the second blog"}' 'http://localhost:1323/posts'

# get all posts
curl -XGET 'http://localhost:1323/posts'
```
