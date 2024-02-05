# Base Application in Golang

## ToDos:

- [X] Index EmailType Mongo
- [X] Persist Error in MongoDocument
- [X] Adjust Listener (get message to json)
- [X] Review Listener error handler
- [X] Review Listeners Concurrency (via multiple pods)
- [X] Send Template to Resources
- [ ] Reprocess Emails
- [X] Get Emails By Filter
- [ ] Tests
- [ ] Review Log Profiles
- [ ] DLQ listeners
- [ ] Review Feature Gateway
- [ ] Review Metrics Listeners


## Guides:
- Errors:
  - Log in the gateway layer, after transform do domain error




```
go test ./test/usecases/... -coverprofile=../coverageresults/cover_usecases.out -coverpkg ./main/usecases
go tool cover -html=../coverageresults/cover_usecases.out
```


https://gist.github.com/TsuyoshiUshio/7eca53ff455b67eb08e3f1db8a01640d




```json
{
    "message": {
        "emailTemplateType": "WELCOME_EMAIL",
        "appOwner": "mp-test",
        "requestUserId": "123",
        "to": [
            "gabriel.estmr@gmail.com"
        ],
        "subject": "Test app",
        "bodyParams": {
            "Message": "Mensagem Teste",
            "Name": "Gabriel Rodrigues"
        }
    },
    "eventId": "11c68b42-b4f2-41d5-a392-3f342cb8fe41",
    "eventDate": "2024-01-18T19:10:42.69149448-03:00"
}
```






