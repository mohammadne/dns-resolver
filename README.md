# DNS Resolver

DNS resolver written in Go

## Testing

### DNS

```bash
echo "Hello, UDP server" | socat - UDP-DATAGRAM:localhost:5354
```

### Health endpoints

```bash
curl localhost:8080/healthz/liveness
curl localhost:8080/healthz/readiness
```
