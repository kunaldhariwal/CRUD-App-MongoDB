# CRUD-App-MongoDB

## **_CRUD App for inserting Blockchain data into the database_**

<br>

**you need to import following packages**

`go get github.com/gin-gonic/gin`

 `go get gopkg.in/mgo.v2`

 `go get gopkg.in/mgo.v2/bson`

**To run the server**

`go run server.go`


**Commands**

_To get user:_

`curl -v -H "Accept: application/json" -H "Content-type: application/json" -X GET http://localhost:8000/users `

<br>

_To get data using Object id:_

`curl -v -H "Accept: application/json" -H "Content-type: application/json" -X GET http://localhost:8000/users/{oid}`

<br>

_To Post the data:_

`curl -v -H "Accept: application/json" -H "Content-type: application/json" -X POST -d '{"BlockId":"01","BlockHash":"Fjkkfvhdfkjgh578t54u4545", "Height":34, "ActivityId":111, "EventId":011}' http://localhost:8000/users`

<br>

_To put the data:_
	
`curl -v -H "Accept: application/json" -H "Content-type: application/json" -X PUT -d '{"BlockId":"01","BlockHash":"Fjkkfvhdfkjgh578t54u4545", "Height":34, "ActivityId":111, "EventId":011}' http://localhost:8000/users/{oid}`

<br>

_Delete the data:_

`curl -i -H "Accept: application/json" -H "Content-type: application/json" -X DELETE http://localhost:8000/users/{oid}`
