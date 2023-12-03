# Base Application in Golang

## ToDos:

- [X] MongoDB Integration;
- [ ] PostgresDB integration;
- [ ] Kafka Integration;
- [ ] RabbitMQ Integration;
- [ ] FF4J integration;
- [X] YML Properties Integration;
- [X] Swagger
- [X] Msg Bundle
- [ ] Redis Locking
- [ ] OTEL integration
- [X] Request Validation
- [X] Logs
- [ ] How to unmarshall structs with private props
- [ ] Review Redis Repository (With Sentinel)
  - https://redis.uptrace.dev/guide/go-redis.html
- [ ] Review Redis Repository: findByDocumentNumber
- [ ] Get Locale from proxy

## Notes:

### Golang Packages:

Create go modules: go to main path and use the command below:

```
go mod init
```

Build project into one native executable file: (in main's path file)

```
go build
```

Get external libs

```
go get gopkg.in/yaml.v3
```

To use external libs: use last part of the import reference:

```go
package main

import (
	"fmt"
	"modulo/auxiliar"
)

func main() {
	fmt.Print("Some log message")
	auxiliar.Execute("id")
}
```

https://zhwt.github.io/yaml-to-go/



https://stackoverflow.com/questions/64712646/how-to-properly-disconnect-mongodb-client

https://apiux.com/2013/04/25/how-to-localize-your-api/
https://ip2location-go.readthedocs.io/en/latest/quickstart.html


// TODO: check func
// WHen to close connection
// IF Connection failed, how to solve