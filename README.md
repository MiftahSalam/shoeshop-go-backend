# Simple Shoe Shop

A small and simple online shoe shop service developed with go.

Developed based on [this](https://www.youtube.com/watch?v=1NWBO8L81J8&list=PLE_Uj9ql8q9_MVzY0bMQPERz4Yulo3We4&index=1) youtube tutorial video

# Feature

- Product browsing and searching
- Review product (rating, comment)
- Cart product
- Order product
- Order payment (paypal)
- Order shipping
- User management (Register, Login, Profile)
- Authentication (JWT)
- Pagination

# Demo

Not available yet

# Stacks

- Router (Echo)
- Graphql (gqlgen)
- ORM (GORM)
- Logger (zap)
- Postgresql
- JWT-GO
- godotenv

# Getting started

## Install Golang

Make sure you have Go 1.13 or higher installed.

https://golang.org/doc/install

## Golang Environment Config

Set-up the standard Go environment variables according to latest guidance (see https://golang.org/doc/install#install).

## Install Dependencies

From the project root, run:

```
go build ./...
or
go mod tidy
```

## Application Environment Config

- Make sure database ready (installed and configured)
- create .env file in projectroot
- create and fill this env var in .env file

```
APPLICATION_HTTPPORT=
APPLICATION_GQLHTTPPORT=
APPLICATION_OPTIONS_SKIPGQLREQBODYLOG=
APPLICATION_OPTIONS_JWTSECRET=
APPLICATION_OPTIONS_JWTEXPIRED=

LOGGER_FILELOCATION=
LOGGER_FILEMAXAGE=
LOGGER_LEVEL=-
LOGGER_STDOUT=

DATABASE_CONNMAXLIFETIME=
DATABASE_AUTOMIGRATE=
DATABASE_DEBUGMODE=
DATABASE_HOST=
DATABASE_MAXOPENCONNECTIONS=
DATABASE_MINIDLECONNECTIONS=
DATABASE_NAME=
DATABASE_PASSWORD=
DATABASE_PORT=
DATABASE_SCHEMA=
DATABASE_USERNAME=
```

## Testing

Not available yet

## Todo

- Create rest api handlers
- Create testing
