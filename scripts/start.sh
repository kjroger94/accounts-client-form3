#!/bin/sh
apk add curl
count=0

until $(curl --output /dev/null --silent --head --fail -X GET http://localhost:8080/v1/organisation/accounts); 
do
    if [ ${count} -eq 3 ]; then
        echo "Exiting, api server is not up in time"
        exit 1
    fi
    count=$(($count+1))
    sleep 5
done

echo "Executing test now......."
go test -v -cover ./...
echo "Tests completed!"