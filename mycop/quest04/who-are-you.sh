#!/bin/bash

curl -s https://01.alem.school/assets/superhero/all.json | jq .[52] | jq .name