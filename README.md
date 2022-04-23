# dnscheck

 dnscheck is for querying the IP, CName, NS, MX and TXT records of the URL.
 It uses gin-gonic, net and basic goroutines and channel(just for experimentation purpose). Goroutines and channel can be removed if needed.

## How to run

### To run it locally

```sh
go run main.go
```

### To run on docker

```sh
docker-compose up
```

## Use the cURL commands below to interact with the endpoint

```sh
curl "http://localhost:3000/dns/{url-to-query}"
```

eg: `curl "http://localhost:3000/dns/google.com"`
