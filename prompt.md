There are many etherscan-like blockchain explorers, they api endpoints are different, and each explorer has its own apikeys.

I have a default config file to store those apikeys, and the config file is like this:

```json
{
  "1": {
    "endpoint": "api.etherscan.io",
    "keys": {
      "yourAPIKey1GetFromEtherscan": 5,
      "yourAPIKey2GetFromEtherscan": 10
    }
  },
  "56": {
    "endpoint": "api.bscscan.com",
    "keys": {
      "yourAPIKey1GetFromBscscan": 5,
      "yourAPIKey2GetFromBscscan": 10
    }
  }
}
```

In the example above, I have two chainIDs, 1 and 56, and the corresponding explorer api endpoints are `api.etherscan.io` and `api.bscscan.com`. For chainID 1, I have two apikeys, `apikeyStr1` and `apikeyStr2`, and the rate limit for each apikey is 5 and 10 respectively.

With those apikeys, I can send requests to the blockchain explorer, and the rate limit is the maximum number of requests that I can send in one second, it should be considered when sending requests.

Use a load balancer to send requests to the blockchain explorer, and the load balancer should consider the rate limit of each apikey.

When interacting with the blockchain explorer, the request should be send to

```
https://$endpoint/api
   ?param1=paramVal1
   &param2=paramVal2
   &apikey=$oneOfApiKey
```

The request maybe get or post request, but what they have in common is that they all have an apikey parameter in URL and you can find the $endpoint by mapping the chainID.

So in the end what I want is a unified api endpoint that I can send requests to, the unified api endpoint be like: http://localhost:8080/$chainID, for example, when I send a request to http://localhost:8080/1, the request should be forwarded to `https://api.etherscan.io/api`, and the apikey should be selected from the apikey pool, and the rate limit should be considered.

User can provide their own apikey by http://localhost:8080/$chainID?apikey=$apikey&otherParam=someVal, and the user provided apikey should be used and do not consider the rate limit.

Add a logger to log important info, e.g. request & response, but by default in debug model, it should enable to see the log, so there should be a command flag to control the log level

Help me to build the whole app.

This is an opensource project that other developer can use.

Based promote.md and the whole app code, generate the README.md for me
