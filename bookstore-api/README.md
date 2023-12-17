# Bookstore API
A go application that has a crud api for a bookstore.

## Application
### Compile
```bash
cd cmd &&
go build main.go
```
### Run
```bash
go run main.go
```
or
```bash
./main.exe
```

## Database
This application stores its data in a mysql database.

### Instructions
Make sure to have docker installed on the system and running the following command.
```bash
docker run --name bookstore-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=my-secret-pw -d mysql
```

#### First Time Setup
If this is the first time running/creating the datasource then you will need to create a database.
Login via a database client (DBeaver, Datagrip, etc.) and run.
```sql
create database bookstore
```