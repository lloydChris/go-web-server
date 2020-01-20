# go-web-server

## Starting Database
stack.yml contains a postgres image and admin interface.  Launch them with:
`docker-compose -f stack.yml up`

login to the admin interface at localhost:8080.  (confirm the following at main.go):
user: postgres
password: devpwd
dbname: postgres

Load the schema from file dbSchema.sql.  The UI provides an Import function that will work with this.
