# Synapsis Test Backend

## How to Install
Step 1 
create database run query : 
```
    CREATE database =synapsis-test-backend;
```

create table and insert data dummy :
```
 run db in file file\db.sql
```

Step 2
- configuration env 
example : 
env.example

Step 3
```
- Run the server
```
go mod vendor
cd server
go run main.go
```

Step 4
- Configuration Postman
Postman collection : 
    in file synapsis-test-backend.postman_collection.json
Postman Env : 
    in file synapsis-test-backend.postman_environment.json


## Folder structure
    .
    ├── files                       # This folder contains the sql migration table
    ├── helper                      # Helper function that usually called in usecase 
    ├── log                         # Log file
    ├── model                       # SQL query function
    ├── pkg                         # 3rd party & global function
    ├── server                      # Main service
    │   ├── bootstrap               # Init middleware and routes
    │   └── handler                 # Handler function to validate parameter inputed and handle response body
    │   └── middleware              # Route middleware
    │   └── request                 # Request body struct
    ├── static                      # Static folder contains assets in the form of invoice, images and uploaded csv files
    ├── usecase                     # API logic flow
    │   └── viewmodel               # Struct of usecase response body             

### Description ENV variable  
    - APP_DEBUG=false : true/false, flag to debug app if panic happen  
    - APP_HOST=0.0.0.0:3000 : default port that app will running  
    - APP_LOCALE=en : default validator v9 default language

    - TOKEN_SECRET=jwtsecret : jwt string secret  
    - TOKEN_REFRESH_SECRET=jwtsecretrefresh : jwt string refresh secret  
    - TOKEN_EXP_SECRET=72 : jwt secret lifetime in hours  
    - TOKEN_EXP_REFRESH_SECRET=720 : jwt refresh secret lifetime in hours  

    - REDIS_HOST=127.0.0.1:6379 : redis connection  
    - REDIS_PASSWORD= : redis password  

    - DATABASE_HOST=127.0.0.1 : msql ip host  
    - DATABASE_DB=synapsis-test-backend : msql db name  
    - DATABASE_USER=msql : msql username  
    - DATABASE_PASSWORD= : msql password  
    - DATABASE_PORT=5432 : msql port  
    - DATABASE_SSL_MODE=disable : ssl mode, disable means no private key required for conenction  

    - LOG_DEFAULT=system : file/system, need to fill file path if log default is file  
    - LOG_FILE_PATH=../log/system.log : log file path  

    - FILE_MAX_UPLOAD_SIZE=10000000 : max upload size in bytes 
    - FILE_STATIC_FILE=../static : local public directory  
    - FILE_PATH=/synapsis-test-backend : subpath to access public directory