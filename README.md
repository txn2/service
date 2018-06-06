# Service

[TXN2] Golang service libraries. Check out [Go Microservices Boilerplate] for examples and a refrence implementation.

### Install

```bash
go get github.com/txn2/service
```

## Ack
HTTP service acknowledgement structure.

The Ack data structure is used for providing a common wrapper around HTTP JSON api calls developed in [go].

```go
// Ack
type Ack struct {
	Version     int         `json:"ack_version"`
	Agent       string      `json:"agent"`
	Uuid        string      `json:"ack_uuid"`
	RequestUuid string      `json:"req_uuid"`
	DateTime    string      `json:"date_time"`
	Success     bool        `json:"success"`
	ServerCode  int         `json:"server_code"`
	Location    string      `json:"location"`
	PayloadType string      `json:"payload_type"`
	Payload     interface{} `json:"payload"`
}
```

**Example JSON:**

```json
{
    "ack_version": 5,
    "agent": "",
    "ack_uuid": "f8b0d2ca-5a7c-48fd-ba26-5b53946d741c",
    "req_uuid": "",
    "date_time": "2018-06-05T10:36:38-07:00",
    "success": true,
    "server_code": 200,
    "location": "/",
    "payload_type": "Message",
    "payload": {
        "message": "service boilerplate"
    }
}
```

#### Use with [gin-gonic] web framework:

```bash
go get github.com/txn2/service/ginack
```

**Example implementation:**

```go
import (
	"github.com/gin-gonic/gin"
	"github.com/txn2/service/ginack"
)

...
    // route handler
    func(c *gin.Context) {
        ack := ginack.Ack(c)
        ack.SetPayload(gin.H{"message":"service boilerplate"})
    
        // return
        c.JSON(ack.ServerCode, ack)
    }
...
```


## Cassandra

See **examples** folder.


[gin-gonic]: https://github.com/gin-gonic
[Go Microservices Boilerplate]: https://github.com/txn2/boilerplate-go
[TXN2]: https://txn2.com
[go]: https://golang.org/