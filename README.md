# Simple CRUD user management

#### Demo at ```www.crud.thomasparadis.me```

## How to run
1. `cd build; docker-compose up -d db`
1.  wait a few seconds for db to initialize
1. `docker-compose up -d app`
1. app is now running on localhost:4006 

## Run tests
1. after db is initialized
1. `docker-compose up -d app-test`


## Possible improvements
1. Form validation
1. Popup displaying api after server query (with success or error description ) 
1. Query db one page at a time



#### (upload to server command):
`rsync -av -e ssh --exclude='frontend/node_modules/*' . root@thomasparadis.me:~/simple-crud`