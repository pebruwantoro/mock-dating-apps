**WELCOME TO THE MOCK PROJECT DATING APPS**
![Illustrations](assets/img.png)

**NON-FUNCTIONALITY**
1. Go Programming Language
2. Postgres
3. Docker
4. Makefile

**FUNCTIONALITY**
1. User is able to sign up or register
2. User is able to login
3. User is able to swipe left/right to others, if user is premium the swipe will be unlimited (User Default is not premium).
4. User is able to purchase premium package

**HOW TO USE THIS APPS**
1. Clone this repository
2. Download the dependencies by running `make init`
3. Run the unit test by running `make test`
4. Run the apps on the docker by running `make docker-up`
5. Shut down the apps by running `make docker-down`

**API ENDPOINT**
1. POST /users/sign-up

* Request Body
```bash
curl --location 'http://localhost:8080/users/sign-up' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "user1@yahoo.com",
    "username": "user",
    "password": "user"
}'
```

* Response Body
```bash
{
    "success": true,
    "message": "success sign up",
    "data": {
        "id": "a6003dcb-c76f-4244-bbb5-993e21fcc458"
    }
}
```

2. POST /users/login

* Request Body
```bash
curl --location 'http://localhost:8080/users/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "identity": "user1@yahoo.com",
    "password": "user"
}'
```

* Response Body
```bash
{
    "success": true,
    "message": "success login",
    "data": {
        "id": "a0b2a1b6-cf9f-42f4-8055-324d87449725",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1dWlkIjoiYTBiMmExYjYtY2Y5Zi00MmY0LTgwNTUtMzI0ZDg3NDQ5NzI1IiwiZW1haWwiOiJ1c2VyMUB5YWhvby5jb20iLCJ1c2VybmFtZSI6InVzZXIxIiwiaXNfcHJlbWl1bSI6dHJ1ZSwiaXNzIjoiZGF0aW5nLWFwcHMuY29tIiwiZXhwIjoxNzM3NDczMTA0fQ.mbusDvcEebY6hOC8xlZBPWxW1NwPEDFp_edR2FRpg6A"
    }
} 
```

3. POST /users/swipe

* Request Body
```bash
curl --location 'http://localhost:8080/users/swipe' \
--header 'Authorization: Bearer ${TOKEN}' \
--header 'Content-Type: application/json' \
--data '{
    "target_id": "49fbe501-2c61-4530-a1a5-dcb1499029e9",
    "direction": "right"
}'
```

* Response Body
```bash
{
    "success": true,
    "message": "success swipe",
    "data": {
        "id": "45297b2e-251a-4424-9ce9-5934c7a888c0"
    }
} 
```

4. POST /users/purchase

* Request Body
```bash
curl --location 'http://localhost:8080/users/purchase' \
--header 'Authorization: Bearer ${TOKEN}' \
--header 'Content-Type: application/json' \
--data '{
    "amount": "100000"
}'
```

* Response Body
```bash
{
    "success": true,
    "message": "success purchase premium package",
    "data": {
        "id": "6f3ff8e4-036f-4f9a-853f-bb7f8580e1de"
    }
}
```