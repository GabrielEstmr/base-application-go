seq 1 8 | xargs -n1 -P8 curl --request POST \
  --url http://localhost:8081/api/v1/users \
  --header 'Accept-Language: en-US' \
  --header 'Authorization: aushaihsuahs' \
  --header 'Content-Type: application/json' \
  --header 'accept: application/json' \
  --data '{
    "documentId": "3333332",
    "userName": "jujuLinda2",
    "firstName": "Pepeta",
    "lastName": "Peta",
    "email": "takona1969@seosnaps.com",
    "birthday": "2024-01-28T11:22:19.753Z",
    "password": "defaultPassword",
    "phoneContacts": [
        {
            "phoneCountry": "US",
            "phoneType": "MOBILE",
            "phoneNumber": "+5534991819370"
        }
    ]
}'