#!/bin/bash

curl -X POST -H "Content-Type: application/json" http://10.1.1.150:8080/pull -d '{"url":"https://github.com/codecowboydotio/swapi-json-server", "branch":"dev"}'
