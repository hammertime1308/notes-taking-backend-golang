# notes-taking-backend-golang

## API's to perform crud operation for note taking backend

Refer [swagger-file](swagger.json) file for API documentation.

Exposed the below mentioned API's for performing crud operations.

-   Signup as a new user
-   Login using email and password
-   Create a new note
-   Get all notes created
-   Delete a specific note

## Tech stack

-   Golang - for backend service
-   mysql - for storing the data of users and notes
-   Docker - for containerizing the application
-   docker-compose - for running the system locally

## Running locally

For running the system locally you can use the docker-compose file which starts the notes-server and mysql database. It also creates databases and tables.

```
docker-compose up -d
```

-   Refer `init.sql` file for database and table creation.
-   Update the local.json file for database hostname and port.
