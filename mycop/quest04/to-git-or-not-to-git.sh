#!/bin/bash

curl -s https://01.alem.school/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{githubLogin:{_eq:\"Qarapaiym\"}}){id}}"}' | jq .data | jq .user | jq .[0] |jq .[]
