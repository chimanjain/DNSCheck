# dnscheck

 dnscheck is for querying the IP, CName, NS, MX and TXT records of the URL.
 It uses fiber and basic goroutines and channel(just for experimentation purpose). Goroutines and channel can be removed if needed.

 It also uses redis for caching mechanism.

 The structure is kept simple and straightforward following the principle of **KISS**(Keep It Stupid Simple).

## How to run

```sh
docker compose up
```

## Use the cURL commands below to interact with the endpoint

```sh
curl "http://localhost:3000/dns/{url-to-query}"
```

eg: `curl "http://localhost:3000/dns/google.com"`
