## Encoding web service

### to compile and install
```
go install github.com/saw2th/encoding-proj/encserv
```

### using the server

run using:
```
bin/encserv
```

### storing Content

```
curl -X POST  http://localhost:8080/store -d '{"Id" : "123", "Content": "abcde"}'
```

example return: 
```
{"Key":"5151d9ef0a6ff067447f72cebd532ca3"}
```

### retrieving Content

then copy the 'Key' value to the POST data 'Key' on the retrieve url
```
curl -X POST  http://localhost:8080/retrieve -d '{"Id" : "123", "Key": "5151d9ef0a6ff067447f72cebd532ca3"}'
```
example return:
```
{"Content":"abcde"}
```

### To run test on client
```
go test github.com/saw2th/encoding-proj/enc-client
```
