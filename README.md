## Project setup and run instructions
1. Create .env file in root:
  ```
    APP_MODE=debug

    DATABASE.USERNAME=myuser
    DATABASE.PASSWORD=mypassword
    DATABASE.NAME=mydatabase
    DATABASE.HOST=db
    DATABASE.PORT=5432

    APP.HTTP_PORT=8080
  ```
2. Run docker and execute command
  ```
    docker-compose up
  ```