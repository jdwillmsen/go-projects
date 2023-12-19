# Fiber CRM Basic
A fiber crm basic implementation project.

## Database
This application stores its data in a postgres database.

### Instructions
Make sure to have docker installed on the system and run the following command.
```bash
docker run --name fiber-crm-basic-postgres -p 9500:5432 -e POSTGRES_PASSWORD=my-secret-pw -d postgres
```

#### First Time Setup
If this is the first time running/creating the datasource then you will need to create a database.
Login via a database client (DBeaver, Datagrip, etc.) and run.
```sql
create database leads
```