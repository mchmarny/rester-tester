# rester-tester
Rest server to test serverless usage patterns


## Image

Makes PNG thumbnail from video. To test, pass URL using Curl

```
curl -X POST -H "Content-Type: application/json" http://localhost:8888/image \
     -d '{"src":"https://www.youtube.com/watch?v=DjByja9ejTQ"}'     
```
