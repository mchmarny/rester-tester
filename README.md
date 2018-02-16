# rester-tester [![Build Status](https://travis-ci.org/mchmarny/rester-tester.svg?branch=master)](https://travis-ci.org/mchmarny/rester-tester)

Simple REST server to test serverless usage patterns

## Run

#### Dependancies 

```
make deps
```

#### Locally 

```
make run
```

#### Locally in Docker

```
docker build -t rester-tester .
docker run -p 8080:8080 rester-tester:latest
```

## Use

### Ping
Basic GET test

##### Request

```
curl http://localhost:8080/ping 
```

##### Response

```
{
     "pong": {
          "service_id": "1f6f4f35-1ef5-4c27-b69c-b34652544229",
          "host_name": "056504ae4039",
          "started_at": "2018-01-28 18:22:32.383411901 +0000 UTC"
     }
}
```

### Prime
Basic max prime number calculator

##### Request

To get max prime using defaults

```
curl http://localhost:8080/prime 
```

To pass max number as argument to prime calculator 

```
curl http://localhost:8080/prime/50000000
```

##### Response

```
{
     "prime": {
          "max": 50000000,
          "host_name": "056504ae4039",
          "val": 49999991
     }
}
```


### Image

Makes PNG thumbnail from video URL. Will test whether the app can execute external ffmpeg command.

```
curl -X POST -H "Content-Type: application/json" http://localhost:8080/image \
     -d '{"src":"https://www.youtube.com/watch?v=DjByja9ejTQ"}'     
```

```
{
    "request_id":"b8845f04-c462-4f3e-91d7-6d512c576e23",
    "created_at":"2018-01-28 18:25:06.475802091 +0000 UTC",
    "status":"Processed",
    "req": {
        "src":"https://www.youtube.com/watch?v=DjByja9ejTQ",
        "width":200,
        "height":200
    },
    "message":"Completed in 1.016136195s",
    "thumb/img_8acd94cd-4035-4c62-b4d7-d6c6cb03a88d.png"
}
```

## TODO

* Push images to distributed object store 
