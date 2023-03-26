# course

This is mono repo of two course sample project, this course project will contain two mono repo of sample project.

## **Array**

Array is project that implement the functionality that take two array and return them but seperately where one return add value and other include removed value.

### Run array service

Build the binary by running:

```
task array:build
```

Execute the array program :

```
task array:run
```

## Bank Account

Bank account application is the project that is responsible to implement the simple bank functionality simiralities that is implement in `golang`.


### Run bank-account

Download the application packages by run:

```
task tidy
```

Build the binary by running:

```
task account:build
```

Execute the bank acount program :

```
task account:run
```

### Testing

Creating account

```
curl --data account_number=32 POST localhost:8080/create
```

Response

```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 26 Mar 2023 00:21:42 GMT
Content-Length: 47

{
    "message": "Account 32 has been created"
}

```

Deposit to account

```
curl --data "account_number=32&amount=500.400"  POST localhost:8080/deposit -H "Content-Type: application/x-www-form-urlencoded"
```

Response

```
x-www-form-urlencoded" 
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 26 Mar 2023 00:47:59 GMT
Content-Length: 82

{
    "message": "Deposit of 500.400 amount has been deposited to 32 successful."
}

```

Check Balance

```
curl --data account_number=32 GET localhost:8080/balance
```

Response

```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 26 Mar 2023 00:51:52 GMT
Content-Length: 41

{
    "number": "32",
    "balance": 500.4
}

```

WithDraw amount from account

```
curl --data "account_number=32&amount=100"  POST localhost:8080/withdraw -H "Content-Type: application/x-www-form-urlencoded
```

Response

```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 26 Mar 2023 00:55:09 GMT
Content-Length: 83

{
    "message": "Withdrawal of 100 amount has been withdrawn from account of: 32 successful."
}

```

Thank!
