# rester-tester
Rest server to test serverless usage patterns




## Run

```
docker build -t rester-tester .
docker run -p 8080:8080 rester-tester:latest
```

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

### Image

Makes PNG thumbnail from video. To test, pass URL using Curl

```
curl -X POST -H "Content-Type: application/json" http://localhost:8080/image \
     -d '{"src":"https://www.youtube.com/watch?v=DjByja9ejTQ"}'     
```

## TODO

* Remap resulting thumbnail to external URL
* Push images to distributed object store 
