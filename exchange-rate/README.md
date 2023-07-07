# Exchange Rate

It is a system that receives data from [ExchangeRate-API](https://www.exchangerate-api.com/docs/overview) every minute and writes them to related queues. It also has an API that you can access as rest.

## Run

```Bash
# RabbitMQ default username & password : gues
# system requerments
$ docker-compose up -d
# manuel run
$ go run ./er-api/main.go
$ go run ./er-api-producer/main.go
$ go run ./er-rabbit-consumer/main.go TRY
$ go run ./er-rabbit-consumer/main.go USD
$ go run ./er-rabbit-consumer/main.go EUR
```

## Endpoints

 GET `/api/v1/pair/{base}/{target}`

* Respons Payload

```Json
{
    {
 "id": 1,
 "base_code": "EUR",
 "target_code": "GBP",
 "conversion_rate": 0.8412,
 "created_at": 1585270800
    },
    {
        "id": 2,
        "base_code": "EUR",
        "target_code": "TRY",
        "conversion_rate": 0.8412,
        "created_at": 1585270800
    },
}
```

POST `/api/v1/time`

* Request Payload

```Json
{
    "start_date": "2022-03-12",
    "end_date": "2022-04-03"
}
```

* Respons Payload

```Json
{
    {
 "id": 1,
 "base_code": "EUR",
 "target_code": "GBP",
 "conversion_rate": 0.8412,
 "created_at": 1585270800
    },
    {
        "id": 2,
        "base_code": "EUR",
        "target_code": "GBP",
        "conversion_rate": 0.8412,
        "created_at": 1585270800
    },
}
```

* Error Payload

```Json
{
    "result": "Error",
    "error_type": "Could not complete read from request body"
}
```
