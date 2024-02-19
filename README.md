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
}
```

`LogInfo()`

### GitHub Actions
| Action | Description |
| --- | --- |
|push| Lint and Unit tests are run |
|tag| release created, tagged build pushed to docker hub |

#### How to create new tag:
From command line:
```
git tag 'vX.Y.Z'
git push origin 'vX.Y.Z'
```

or use the UI from within the GitHub repo.