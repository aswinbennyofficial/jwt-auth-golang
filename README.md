# jwt-auth-golang

A secure user authentication service built with Golang and JWT for seamless signup and signin experiences.

## Dependencies

- [godotenv](https://github.com/joho/godotenv): Used for loading environment variables from a `.env` file.
- [mongo-driver](https://go.mongodb.org/mongo-driver/mongo): MongoDB driver for Golang.
- [jwt](https://github.com/golang-jwt/jwt/v5): Golang implementation of JSON Web Tokens (JWT).
- [bcrypt](https://golang.org/x/crypto/bcrypt): A library for hashing and comparing passwords using bcrypt algorithm.

## Installation

1. Clone the repository:

```bash
git clone https://github.com/aswinbennyofficial/jwt-auth-golang.git
```


2. Install dependencies:

```bash
go get github.com/joho/godotenv
```
```bash
go get go.mongodb.org/mongo-driver/mongo
```

```bash
go get github.com/golang-jwt/jwt/v5
```

```bash
go get golang.org/x/crypto/bcrypt
``` 



3. Configure your environment variables by renaming `.env.example` into `.env`


## Usage
### Running the application

```bash
go run ./cmd/main/
```
By default, the server will start on port 8080.