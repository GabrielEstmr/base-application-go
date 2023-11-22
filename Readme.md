# Base Application in Golang

## ToDos:

- [ ] MongoDB Integration;
- [ ] PostgresDB integration;
- [ ] Kafka Integration;
- [ ] RabbitMQ Integration;
- [ ] FF4J integration;
- [ ] YML Properties Integration;
- [ ] Swagger
- [ ] Msg Bundle
- [ ] 
- [ ] 
- [ ] 
- [ ] 

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