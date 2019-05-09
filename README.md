# Golang Restful API using GORM ORM (MySQL), Gorilla Mux, GWT

## Getting Started

### Folder Structure
This is my folder structure under my `$GOPATH` or `$HOME/your_username/go`.
```
.
|-- bin
+-- src
|   +-- github.com
|       +-- gamorvi
|          +-- restapi2
|             |-- .env
|             |-- main.go
|             |-- server.go
|             |-- .gitignore.go
|             |-- README.md
|             +-- app
|                 +-- controllers
|                     +-- auth
|                         |-- authController.go
|                     |-- usersController.go
|                 +-- models
|                     |-- base.go
|                     |-- user.go
|             +-- routes
|                 |-- api.go
|             +-- utils
|                 |-- utils.go
```
Ensure you create the `gamorvi` directory in your `github.com` directory. `cd` into the `gamorvi` directory before `git clone https://github.com/gamorvi/restful-api-with-golang.git`

## Download the packages used to create this rest API
Run the following Golang commands to install all the necessary packages.

`go get -u github.com/gorilla/mux` for serving the api

`go get -u github.com/jinzhu/gorm` ORM supports MySQL, SQLite, MSSQL, Postgres

`go get -u github.com/go-sql-driver/mysql` MySQL driver to enable SQL connection

`go get -u github.com/joho/godotenv` Loads your environment variables (database, AWS, Redis configurations etc.)

`go get -u github.com/dgrijalva/jwt-go` Loads your environment variables (database, AWS, Redis configurations etc.)

### Running documentation locally (Only documentation of packages your have installed)
For offline documentation on the following packages run `godoc -http :6060` and then visit `http://localhost:6060`. Note that you can change the port to your preferred port number.

## Setting configuration file
Create a .env file in the root of the project and set the following parameters

`db_name = database_name` Name of database

`db_user = user`  # Database username

`db_pass = secret` # Database password

`db_type = mysql`   # MySQL driver

`db_host = localhost` # Database host

`db_port = 3306`  # Database port

`charset = utf8` # Database charset

`parse_time = True` # Database parse time

`web_port = 8085`   # Port to serve api

`prefix = /api/v1`  # API route sub route prefix

# jwt config
`access_token_expire   =  ` # in minutes, default is 15 minutes if no value is passed

# this is not used yet as I haven't implemented refresh tokens yet # TODO
`refresh_token_expire  = 72 ` # in hours, default is 48 minutes if no value is passed

## Running the project

`go run *.go`

## Database Table Creation Statement
Use the following DDL (Data Definition Language) to create the users table.

``` SQL
CREATE TABLE `users` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `username` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8;
```

## API Endpoints & Usage

To be able to login, you need to use the create new user endpoint to set up a user account.

* POST `api/v1/login` login with username and password in your db to get token back
* GET `api/v1/users` retrieve all users
* GET `api/v1/users/1` retrieve user with id = 1
* POST `api/v1/users` create a new user
* PUT `api/v1/users/1` update the record with id = 1
* DELETE `api/v1/users/1` delete the user with id = 1

### To create a new user

1. POST `api/v1/users`

```
{
	"Name": "Joe Bloke",
	"Username": "joe.bloke@fake-domain.com",
  "Password": "secret"
}
```

*** Output ***

```
{
    "message": "success",
    "status": true,
    "user": {
        "ID": 1,
        "CreatedAt": "2019-05-06T00:54:22.09382+01:00",
        "UpdatedAt": "2019-05-06T00:54:22.09382+01:00",
        "DeletedAt": null,
        "Name": "Joe Bloke",
        "Username": "joe.bloke@fake-domain.com"
    }
}
```

2. Get `api/v1/login`

Remember to use `x-www-form-urlencoded`

*** Input ***

Username: joe.bloke@fake-domain.com
Password: secret

*** Output ***

```
{
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjkwMDAwMDAwMDAwMCwicGFzc3dvcmQiOiJzZWNyZXQiLCJ1c2VybmFtZSI6ImFtYXZpQHh5ei5jb20ifQ.WJ5VMnH5ijHQOZhUlrrnrh7NCYfFpww3jBz26EkRsHQ"
    },
    "message": "success",
    "status": true
}
```

Now make all calls pass the token in the header.
