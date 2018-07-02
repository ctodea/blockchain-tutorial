# rajols
Simple blockchain in Go, for learning purposes. Forked from mycoralhealth/blockchain-tutorial


Dependencies managed with [dep](https://github.com/golang/dep):
```
dep ensure
```

`ADDR` and `PORT` defined in `config/config.toml`
Add new block using cURL:
```
curl -d '{"BPM":75}' -H "Content-Type: application/json" -XPOST http://$ADDR:$PORT
```
