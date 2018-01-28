# rester-tester
Rest server to test serverless usage patterns




## Run

```
docker build -t rester-tester .
docker run -p 8080:8080 rester-tester:latest
```

### Ping
Basic GET test

```
curl http://localhost:8080/ping
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
