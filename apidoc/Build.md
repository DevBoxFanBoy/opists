# Go API Server for opists

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This server was generated by the [openapi-generator]
(https://openapi-generator.tech) project.
By using the [OpenAPI-Spec](https://github.com/OAI/OpenAPI-Specification) from a remote server, you can easily generate a server stub.  
-

To see how to make this your own, look here:

[README](https://openapi-generator.tech)

- API version: 1.0.0
- Build date: 2020-11-13T07:30:51.058913100+01:00[Europe/Berlin]


### Running the server
To run the server, follow these simple steps:

```
go run main.go
```

To run the server in a docker container
```
docker build --network=host -t opists .
```

Once image is built use
```
docker run --name opists -P 8080:8080 --rm -it opists 
```

Testing
```
curl -v http://localhost:8080/v1/projects
curl -v http://localhost:8080/v1/projects/DF
curl -v http://localhost:8080/v1/projects/DF/issues
curl -v http://localhost:8080/v1/projects/DF/issues/0

curl -X POST http://localhost:8080/v1/projects/DF/issues -H  "Content-Type: application/json" -d "{\"name\":\"New Bug\",\"description\":\"An error raise when...\",\"status\":\"open\",\"priority\":\"Highest\",\"components\":[\"DrinkOwnChampagne\",\"EatMyOwnApplication\"],\"sprints\":[\"Sprint2\"],\"estimatedPoints\":0,\"estimatedTime\":\"0h\",\"affectedVersion\":\"1.2.3\",\"fixedVersion\":\"1.2.4\"}"

curl -X PUT https://localhost:8080/v1/projects/DF/issues -H  "Content-Type: application/json" -d "{\"name\":\"New Bug\",\"description\":\"An error raise when...\",\"status\":\"refinable\",\"priority\":\"Highest\",\"components\":[\"DrinkOwnChampagne\",\"EatMyOwnApplication\"],\"sprints\":[\"Sprint2\"],\"estimatedPoints\":0,\"estimatedTime\":\"0h\",\"affectedVersion\":\"1.2.3\",\"fixedVersion\":\"1.2.4\"}"

curl -X DELETE https://localhost:8080/v1/projects/DF/issues/0
```