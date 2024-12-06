#!/bin/bash

URL="http://localhost:8080/1/?module=account&action=balancemulti&address=0xddbd2b932c763ba5b1b7ae3b362eac3e8d40121a,0x63a9975ba31b0b9626b34300f7f627147df1f526,0x198ef1ec325a96cc354c7266a038be8b5c558f67&tag=latest"

start_time=$(date +%s.%N)

for i in {1..12}; do
    echo "Request $i:"
    response=$(curl -s -w "\n%{http_code}" "$URL")
    http_code=$(echo "$response" | tail -n1)
    body=$(echo "$response" | sed '$d')

    echo "Status Code: $http_code"
    echo "Response: $body"
    echo "-------------------"

    # Small delay to ensure all requests happen within 1 second
    sleep 0.05
done

end_time=$(date +%s.%N)
duration=$(echo "$end_time - $start_time" | bc)
echo "Total duration: $duration seconds"
