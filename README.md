# rajols
Simple blockchain in Go, for learning purposes. Forked from mycoralhealth/blockchain-tutorial

```
go get github.com/davecgh/go-spew/spew
go get github.com/gorilla/mux
go get github.com/joho/godotenv
```

Add new block using cURL:
```
curl -d '{"BPM":75}' -H "Content-Type: application/json" -XPOST http://localhost:$PORT
```
