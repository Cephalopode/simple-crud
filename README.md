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
1. Form validation (display fields in red if inputted values are invalid)
1. Popup displaying api response after a query (display 'success' or describe error ) 
1. Query db one page at a time (Instead of downloding all the table entries at once)
1. Wrap "users" methods into an interface so the same methods can be reused to manipulate other tables.
1. Authentification system. This would require : login page and login method, admin table (container username and password hashes), store session_id as a cookie or in local storage.

## Architechture choices
1. UI framework : Vuetify. I'm more familiar with Element UI but it's not reactive, so I chose to experiment with a new framework for this projet. Vuetify is currently the leading UI framework for Vue.js (along with Quasar)
1. Webserver : Go + gin + gorm. Gorm allows to easily manipulate mysql databases.
1. Deployment : Docker-compose. This wrapper over docker allows to easily manage multi-container apps. Also, using the parameter `restart: alway`, the containers are automatically restarted if an error appears. A more robust way would be to use kubernetes to deploy the containers in multiple nodes (servers) so there is redundancy if one of the servers is unreachable.



#### (upload to server command):
`rsync -av -e ssh --exclude='frontend/node_modules/*' . root@thomasparadis.me:~/simple-crud`
