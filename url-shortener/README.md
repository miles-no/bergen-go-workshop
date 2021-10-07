# The best URL shortener in the world

## Shortening a URL

```
curl http://localhost:8080/urls -X POST -d '{"url": "https://www.miles.no"}'

Response:
HTTP status: 201
Body: {"url":"https://www.miles.no","short_url":"/urls/691a9c872085"}
```

## Resolving a shortened URL

```
curl -v http://localhost:8080/urls/691a9c872085

Response:
HTTP redirect to https://www.miles.no
```
