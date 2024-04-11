# URLShortener-backend

### Endpoints

- Create new alias for url
```
http://<host>/url

Basic Auth:
  User: "YOUR-USER"
  Password: "YOUR-PASSWORD

Body:
{
  "url": "https://github.com"
  "alias": "gt"
}

Response:
{
  "Status": "OK",
  "alias": "gt"
}

HTTPStatus: 200
```

- Get url by alias
```
http://<host>/{alias}

HTTPStatus: 302
```

- Delete url by alias
```
http://<host>/url/{alias}

Basic Auth:
  User: "YOUR-USER"
  Password: "YOUR-PASSWORD

HTTPStatus: 204
```
