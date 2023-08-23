# cushon-techincal-solution


---

This API is responsible for updating customer's investment funds.

---
## Get Available Funds

### Request
POST
https://localhost:8080/funds

### Response

> HTTP Status 200 OK

#### body

content-type: application/json

| parameter | type | description           | example |
|-----------|------|-----------------------|---------|
|  ID       | int  | ID of investment fund |  1      |
|  Name     |string|name of investment fund|  Fund1  |


##### example

```json
[
    {
        "ID":1,
        "Name":"Fund1"
    },
    {
        "ID":2,
        "Name":"Fund2"
    }
]
```

## Get Customer Account

### Request
POST
https://localhost:8080/customeraccount/get

#### body

content-type: application/json

| parameter        | type | description           | example |
|------------------|------|-----------------------|---------|
| CustomerID       | int  | ID of customer        |  1      |
| InvestmentFundID | int  | ID of investment fund |  1      |


##### example

```json
{
    "CustomerID": 1,
    "InvestmentFundID": 1
}
```

### Response

> HTTP Status 200 OK

#### body

content-type: application/json

| parameter        | type | description           | example |
|------------------|------|-----------------------|---------|
|  ID              | int  | ID of customer account|  0      |
| CustomerID       |int   | ID of customer        |  1      |
| InvestmentFundID |int   | ID of investment fund |  1      |
| InvestmentAnmount|int   | ID of customer        |  22.7   |


##### example

```json
{
    "ID": 0,
    "CustomerID": 1,
    "InvestmentFundID": 1,
    "InvestmentAmount": 22.7
}
```


## Add Customer Account Investment Funds

### Request
POST
https://localhost:8080/customeraccount/investmentfund/add

#### body

content-type: application/json

| parameter        | type | description                               | example |
|------------------|------|-------------------------------------------|---------|
| CustomerID       | int  | ID of customer                            |  1      |
| InvestmentFundID | int  | ID of investment fund                     |  1      |
| InvestmentAmount |float | Amount of money to add to investment fund |  5.30   |


##### example

```json
{
    "CustomerID": 1,
    "InvestmentFundID": 1,
    "InvestmentAmount": 5.50
}
```

### Response

> HTTP Status 200 OK

##### example

"customer account updated succesfully"
