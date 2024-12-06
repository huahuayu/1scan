#!/bin/bash

# URL to test
URL="http://localhost:8080/1/?module=account&action=balancemulti&address=0xddbd2b932c763ba5b1b7ae3b362eac3e8d40121a,0x63a9975ba31b0b9626b34300f7f627147df1f526,0x198ef1ec325a96cc354c7266a038be8b5c558f67&tag=latest"

# Function to make request and extract API key
make_request() {
    response=$(curl -s -w "\n%{http_code}" "$URL")
    http_code=$(echo "$response" | tail -n1)
    body=$(echo "$response" | sed '$d')

    # Print request number, HTTP status code, and response
    echo "Request $1:"
    echo "Status Code: $http_code"
    echo "Response: $body"
    echo "-------------------"
}

# Start time
start_time=$(date +%s.%N)

# Make 12 requests in parallel
for i in {1..12}; do
    make_request $i &
done

# Wait for all requests to complete
wait

# End time
end_time=$(date +%s.%N)

# Calculate duration
duration=$(echo "$end_time - $start_time" | bc)
echo "Total duration: $duration seconds"
