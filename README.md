# DNS Resolver

DNS resolver written in Go

## Deployment

```bash
# 
make deploy

# 
make undeploy
```

## Request

### DNS

```bash
# testing locally
dig @localhost -p 5354 example19.com A

# test from inside the cluster
kubectl run -it --image=nicolaka/netshoot --rm=true --restart=Never tshoot -- bash
dig @dns-resolver-dns.dns-resolver -p 5354 example19.com

# test from outside the cluster via service loadbalancer
ip=$(k get svc/dns-resolver-dns -n dns-resolver | awk '{print $4}' | grep -v EXTERNAL-IP | cut -d, -f2 )
dig @$ip -p 5354 example19.com
```

### Health endpoints

```bash
# testing locally
curl localhost:8080/healthz/liveness
curl localhost:8080/healthz/readiness

# test from inside the cluster
kubectl run -it --image=nicolaka/netshoot --rm=true --restart=Never tshoot -- bash
curl dns-resolver-server-http/healthz/readiness
```
