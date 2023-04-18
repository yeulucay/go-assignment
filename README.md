# Go - Case Study

This is a case study application that runs as an API.

```
https://getir-case.herokuapp.com/
```

It has 3 endpoints which described below.

| path       | method | description                        |
|------------|--------|------------------------------------|
| /in-memory | GET    | brings key-value pair from memory  |
| /in-memory | POST   | puts key-value pair into memory    |
| /records   | POST   | filters and brings records from db |

## Running Locally

If PORT environment variable is not set, the application runs on 8000 port. In the root level of the project, run the command below

```
go run .
```
The application will start to serve on 8000 (or the port number exposed by env PORT)

## Example Requests

### Get data from Mongo DB
```
curl --location 'https://getir-case.herokuapp.com/records' \
--header 'Content-Type: application/json' \
--data '{
    "startDate": "2016-01-26",
    "endDate": "2018-02-02",
    "minCount": 2700,
    "maxCount": 3000
}'
```
### Put key-value pair into in-memory storage
```
curl --location 'https://getir-case.herokuapp.com/in-memory' \
--header 'Content-Type: application/json' \
--data '{
    "key": "active-tabs",
    "value": "getir"
}'
```
### Get key-value pair by key from in-memory storage
```
curl --location 'https://getir-case.herokuapp.com/in-memory?key=active-tabs'
```