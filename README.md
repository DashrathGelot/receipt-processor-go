# Fetch Receipt Processor

# Introduction

- This document guides you on how to set up the Receipt Processor project and how to run it with or without Docker.

> **_NOTE:_** This information is intended for Fetch code reviewers. I developed this project in Go, and I am a beginner in the Go language, but I am a professional expert in Java. I have also developed this project in Java, which you can find here: [Receipt Processor Java Project](https://github.com/DashrathGelot/receipt-processor) 

## Technologies

The following technology stack was used in developing this project:

* Backend: Go 1.21.0, Gin 1.9.1
* API Interceptor : Postman

## Rule Assumption
The following assumption I made in the rules:
- For Rule 4, consider a pair as a lower bound value. For example, if you have 5 items, it will consider 2 pairs.
- To validate a receipt, perform validation only on fields that are being used to calculate points. For example:
  Item {"shortDescription": "Mountain Dew 12PK", "price": "6.49s"} Here, the price is invalid, but the processing will continue without throwing an error because the price is not being consider due to the description Rule 5.
- Consider that the total and price are given as valid positive numbers.
- Consider the total as the sum of all item prices.
- 0 rewarded points If the item price is 0.
- 0 rewarded points if the value of total is 0.
- Consider the purchase date to be in a valid yyyy-mm-dd format.

# Getting Started

### How to run
- You can run the application using Docker or locally without Docker.

#### Using Docker
* clone the repository
```
git clone git@github.com:DashrathGelot/receipt-processor-go.git
```

* Navigate to the main directory:
```
cd /receipt-processor-go
```

* Build the Go application image in Docker using the following command:
```
docker build -t receipt-processor .
```

* Run the Docker container:
```
docker run -p 8080:8080 receipt-processor
```

#### Using Local Without Docker
- To run the application locally without Docker, make sure you have Go v1.21.0 and Gin installed on your machine.

* Navigate to the main directory:
```
cd /receipt-processor-go
```

* Download the required packages:
```
go mod download
```

* Build the project:
```
go build -o receipt-processor
```

* Run the server:
```
./receipt-processor
```

#### Test Cases
- You can find the test cases in the `main_test.go` file. To run the test cases, execute the following command:
```
go test
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
> **_NOTE:_** All fields in the JSON object are required; otherwise, an error message will be sent. Validation is also handled.

Console:
```json
{
    "id": "3faabe82-92d9-4e0e-9c9d-24b333a63e3e"
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

### API Interceptor Output

![Screenshot 2023-08-18 at 10.01.42 AM.png](resources%2FScreenshot%202023-08-18%20at%2010.01.42%20AM.png)

![Screenshot 2023-08-17 at 11.50.42 PM.png](resources%2FScreenshot%202023-08-17%20at%2011.50.42%20PM.png)
