# Test that echo method is working/echo
GET http://localhost:8080/HelloService/echo
{
    "message": "Hello!"
}

HTTP 200
content-type: application/json
[Asserts]
jsonpath "$.response" == "Hello!"

# Test that empty message request returns 400
GET http://localhost:8080/HelloService/echo
{
    "message": ""
}

HTTP 400

# Missing message property returns 400
GET http://localhost:8080/HelloService/echo
{
    "notMessage": "Hello!"
}

HTTP 400

# Invalid JSON body returns 400
GET http://localhost:8080/HelloService/echo
```json
{
    "": "Hello!"
}
```

HTTP 400

# Invalid service returns 404
GET http://localhost:8080/noservice
{
    "message": "Hello!"
}

HTTP 404

# Invalid method for valid service returns 404
GET http://localhost:8080/noservice
{
    "message": "Hello!"
}

HTTP 404



