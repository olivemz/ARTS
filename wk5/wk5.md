## Algorithm

[wk5_2.go](wk5_2.go)

## Review
N/A

## Tips
### redis outage
1. when redis memory full, it could cause the ec2 down, and cause more redis request and result in more retry on 
server which causes DDos and may result in a masive server down in global region.
^^ solution, redis shall not become the single failure point, and shall set a circuit breaker to protect server depending on any other service.

## Share
N/A
