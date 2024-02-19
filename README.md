# go-micro-log-service

Simple logging service that writes to a mongo database.
Supports http PUT requests and RPC.

### routes

| url | Command | parameters |
| --- | --- | --- |
| `/log` | POST | json: name: string, data: string |


### RPC

```
type RPCPayload struct {
	Name string
	Data string
}```

`LogInfo`