# Fetch Receipt Processor

# Introduction

- This document guide you how to setup the receipt processor project and how to run on docker and without docker

> **_NOTE:_**  This is for fetch code reviewer only: I am beginner in go language but I am professional expert in Java and I have developed one project with java as well here it is: https://github.com/DashrathGelot/receipt-processor 

## Technologies

The following technologies stack used as part of developing this project:

* Backend: Go 1.21.0, Gin 1.9.1
* API interceptor : Postman

## Rule Assumption
- For Rule 4, consider pair as lower bound value, e.g. if you have 5 items than it will consider 2 pair 
- To valid Receipt did validation on only that fields which being used to calculate points, e.g.  
    Item {
      "shortDescription": "Mountain Dew 12PK",
      "price": "6.49s"
    }
    
    here, price is invalid but still going to process this without throwing error because price is not being used due to description rule.


# Getting Started

### How to run
- you can run application using docker or local without docker

#### Using Docker
* go to main directory

```
cd /receipt-processor
```

* build go application image in docker using below command
```
docker build -t receipt-processor .
```

* run docker file
```
docker run -p 8080:8080 receipt-processor
```

#### Using Local Without Docker
- for this you need Go v1.21.0 and Gin on your machine

* go to main directory
```
cd /receipt-processor
```

* download packages
```
go mod download
```

* build project
```
go build -o /receipt-processor
```

* run server
```
./receipt-processor
```

### Guides

#### Process Receipt
```curl
POST http://localhost:8080/receipts/process
```
body:
```json
{
  "retailer": "M&M Corner Market",
  "purchaseDate": "2022-03-20",
  "purchaseTime": "14:33",
  "items": [
    {
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    }
  ],
  "total": "9.00"
}
```
Note: here in json object all fields are required otherwise it will send error message
- here handled validation as well

Console:
```json
{
    "id": "7cefe133-c027-4bc2-98c7-fa4a1f189ba5"
}
```

#### Get Points
```curl
GET http://localhost:8080/receipts/7cefe133-c027-4bc2-98c7-fa4a1f189ba5/points
```
Console:
```json
{
    "points": 109
}
```
### Test Cases
- You can find test cases in the directory below
```
/src/test/java/com/fetch/receiptprocessor
```
- here I did integration testing in java spring boot using junit5


### API Interceptor Output
