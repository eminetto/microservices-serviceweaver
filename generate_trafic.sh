#!/bin/bash

export TOKEN=$(curl -s -X "POST" "http://localhost:12345/auth" \
     -H 'Accept: application/json' \
     -H 'Content-Type: application/json' \
     -d $'{
  "email": "eminetto@gmail.com",
  "password": "1234567"
}' | jq -r .token)

for i in {1..50}
do 
    curl -X "POST" "http://localhost:12345/feedback" \
        -H 'Accept: application/json' \
        -H 'Content-Type: application/json' \
        -H "Authorization:$TOKEN"  \
        -d $'{
    "title": "Feedback test",
    "body": "Feedback body"
    }'
done

for i in {1..50}
do 
    curl -X "POST" "http://localhost:12345/vote" \
        -H 'Accept: application/json' \
        -H 'Content-Type: application/json' \
        -H "Authorization:$TOKEN" \
        -d $'{
    "talk_name": "Go e Microserviços",
    "score": "10"
    }'
done