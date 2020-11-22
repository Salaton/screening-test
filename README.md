# Simple Customer API

This is a simple Golang service that allows a customer to make an order

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

This implementation assumes you have Go , Docker and postgresql installed in your machine.
Add a `.env` file at the project root and add the following:

```
user=<Database user>
password=<Database Password>
dbname=<Database name>
port=5432
sslmode=disable
TimeZone=Africa/Nairobi
AFRICASTALKINGUSERNAME = <Your AT's Username>
AFRICASTALKINGAPIKEY = <Your AT's API Key>
```

## Running the service using Docker
Pull the image from docker hub using the following command
```
docker pull salaton/savannahtest:latest
```
Run it using 

```
docker run ---env-file <path-to-env-file-created> --network=host
```
After this, using your browser, access localhost on port `8080` to view the GraphQL playground;

```
http://localhost:8080/
```

## Running the service locally

How to get the development env running

```
go get
```

This will fetch and install the go packages from the source repository

To run the application, move to the project root directory and run:

```
go run server.go
```

After this, using your browser, access localhost on port `8080` to view the GraphQL playground;

```
http://localhost:8080/
```

### Running the mutations

There are 6 mutations to run. `createCustomer` `createProduct` `login` `createOrder` `updateOrder` `deleteOrder`

#### Create Customer

Sample GraphQL Payload for creating a customer. It returns a customer type.

```graphql
mutation {
  createCustomer(
    input: {
      name: "John Doe"
      email: "johndoe@gmail.com"
      phonenumber: "254712345678"
      password: "password12345"
    }
  ) {
    name
    email
    phonenumber
    password
  }
}
```
#### Create Product

Sample GraphQL Payload for creating a product. It returns a product type.

```graphql
mutation {
  createProduct(input: { name: "Product one", price: 65500.50 }) {
    name
    Price
  }
}
```

#### Login

For authentication purposes, the customer needs to log in. The output will be a JWT token string that will be used to authenticate the process of creating, updating and deleting an order

```graphql
mutation {
  login(input: { email: "johndoe@gmail.com", password: "password12" })
}
```

Result:

```graphql
{
  "data": {
    "login": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
  }
}
```

#### Create Order

This payload will be used to create the order

```graphql
mutation {
  createOrder(
    input: {
      customer_id: 1
      item: [{ product_id: 2, quantity: 40 }]
      price:500.00
    }
  ) {
    id
    item {
      product {
        name
        Price
      }
      product_id
      quantity
    }
  }
}

```

For this mutation, the customer needs to be authorized. The authorization header must be set. At the bottom of the graphQL playground, select `HTTP HEADERS` and fill this

```graphql
{
  "Authorization":"" //use the generated token during login here
}
```

#### Update order

This mutation takes the same format as the create order one but with updated values.

#### Delete order

```graphql
mutation {
  deleteOrder(orderID: 1)
}
```

### Running the queries

#### Query Customers

```graphql
query {
  customers {
    id
    name
    phonenumber
    email
  }
}
```

#### Query Orders

You can Query the customers using this payload as well.

```graphql
query {
  orders {
    id
    customer {
      name
    }
    item {
      product_id
      quantity
    }
    totalPrice
    date_order_placed
  }
}
```
#### Query Products
```graphql
query{
  products{
    name
    Price
  }
}

```

## Running the tests

` go test -cover ./...`
