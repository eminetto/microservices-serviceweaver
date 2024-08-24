#!/bin/bash
export IP=$(kubectl -n servicewaver-example get service -l serviceweaver/app=microservices -o jsonpath="{.items[0].status.loadBalancer.ingress[0].ip}")

export TOKEN=$(curl -s -X "POST" "http://$IP/auth" \
     -H 'Accept: application/json' \
     -H 'Content-Type: application/json' \
     -d $'{
  "email": "eminetto@gmail.com",
  "password": "1234567"
}' | jq -r .token)

for i in {1..50}
do 
    curl -X "POST" "http://$IP/feedback" \
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
    curl -X "POST" "http://$IP/vote" \
        -H 'Accept: application/json' \
        -H 'Content-Type: application/json' \
        -H "Authorization:$TOKEN" \
        -d $'{
    "talk_name": "Go e Microserviços",
    "score": "10"
    }'
done
