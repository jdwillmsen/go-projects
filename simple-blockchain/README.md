# Simple Blockchain
This is a simple blockchain project written in golang.

## How to

### Run App
To run the application you can either run the included executable.
```shell
./main.exe
```
or
```shell
./simple-blockchain.exe
```
Or you can use the following command which compiles and runs. (Allow through firewall if necessary)
```shell
go run main.go
```

### Endpoints
You can set up postman, insomnia, some other rest client, or run the curl commands.
#### Get Blockchain
```shell
curl --location 'http://localhost:3000'
```

#### New Book
```shell
curl --location 'http://localhost:3000/new' \
--header 'Content-Type: application/json' \
--data '{
    "title": "Sample Book 1",
    "author": "Hattie McGlynn",
    "isbn": "48227081",
    "publish_date": "Wed Dec 20 2023 15:49:45 GMT-0600 (Central Standard Time)"
}'
```

#### New Block
```shell
curl --location 'http://localhost:3000' \
--header 'Content-Type: application/json' \
--data '{
    "book_id": "fe338d46c1aefbc519f3fc8f967fd69d",
    "user": "Miss Gina Dach",
    "checkout_date": "Wed Dec 20 2023 05:31:47 GMT-0600 (Central Standard Time)"
}'
```