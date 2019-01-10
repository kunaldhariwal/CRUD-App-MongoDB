# CRUD-App-MongoDB

**you need to import following packages**

`go get github.com/gin-gonic/gin

 go get gopkg.in/mgo.v2

'go get gopkg.in/mgo.v2/bson`

**To run the server**

`go run server.go`


**Commands**

To get user:

curl -v -H "Accept: application/json" -H "Content-type: application/json" -X GET http://localhost:8000/users 

<br>

To get data using Object id:

curl -v -H "Accept: application/json" -H "Content-type: application/json" -X GET http://localhost:8000/users/{oid}

<br>

To Post the data:

curl -v -H "Accept: application/json" -H "Content-type: application/json" -X POST -d '{"BlockId":"01","BlockHash":"Fjkkfvhdfkjgh578t54u4545", "Height":34, "ActivityId":111, "EventId":011}' http://localhost:8000/users

<br>
	
curl -v -H "Accept: application/json" -H "Content-type: application/json" -X PUT -d '{"BlockId":"01","BlockHash":"Fjkkfvhdfkjgh578t54u4545", "Height":34, "ActivityId":111, "EventId":011}' http://localhost:8000/users/{oid}

<br>

Delete the data:

curl -i -H "Accept: application/json" -H "Content-type: application/json" -X DELETE http://localhost:8000/users/{oid}
