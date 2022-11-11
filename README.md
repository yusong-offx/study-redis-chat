# Redis-chat
Make chat using fiber, redis  
You can test by docker-compose and test.html

```sh
git clone (this repo)
docker compose up -d
## open test.html
```

 \+ you have to check redis-server dns (/app-server/components/redis.go:14),  if you change git clone folder name.
  
<br/><br/>

## Feature
----
Connect chat by websocket with user_id and pub/sub by redis.
 * Separate listen and spread
 * Each part works on each goroutine
```go
// redis.subscribe -> websocket(web) 
func rxChat(...)

// websocket(web) -> redis.publish
func txChat(...)
```
<br />

## Usage
 * Go
 * Redis
 * Docker