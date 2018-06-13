# AWS SAM Swagger Golang Example

## Deploy

```
$ make deploy
...
...
...
--------------------------------------------------------------------------------------------------------------
|                                               DescribeStacks                                               |
+---------------------------+------------+-------------------------------------------------------------------+
|        Description        | OutputKey  |                            OutputValue                            |
+---------------------------+------------+-------------------------------------------------------------------+
|  URL of your API endpoint |  ApiUrl    |  https://xxxxxxxxxx.execute-api.ap-northeast-1.amazonaws.com/dev  |
+---------------------------+------------+-------------------------------------------------------------------+
```

## Deploy to specific stage 

```
$ make deploy STAGE=prod
```

## Local Test

```
$ make local
$ curl http://localhost:3000/
Processing request data for request c6af9ac6-7b61-11e6-9a41-93e8deadbeef.
Body size = 0.
Headers:
    X-Forwarded-Port: 3000
    Host: localhost:3000
    X-Forwarded-Proto: http
    Accept: */*
    User-Agent: curl/7.58.0
```
