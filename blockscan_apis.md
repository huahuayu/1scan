# Blockscan API Capabilities

| # | API Method | Tier | Group | Etherscan | BscScan | PolygonScan | zkEVM PolygonScan | BaseScan | Arbiscan | Nova Arbiscan | LineaScan | FTMscan | BlastScan | Optimistic Etherscan | SnowScan | BTTCScan | CeloScan | CronoScan | FraxScan | GnosisScan | KromaScan | MantleScan | Moonbeam | Moonriver | opBNB | ScrollScan | TaikoScan | WemixScan | ZkSync Era | XaiScan |
|---|------------|------|-------|-----------|---------|-------------|-------------------|----------|----------|---------------|-----------|---------|-----------|---------------------|----------|----------|----------|-----------|----------|------------|-----------|------------|----------|-----------|-------|------------|-----------|-----------|------------|----------|
| 1 | Get Ether Balance for a Single Address | Free | Accounts | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y |
| 2 | Get Ether Balance for Multiple Addresses in a Single Call | Free | Accounts | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y |
| 3 | Get a list of 'Normal' Transactions By Address | Free | Accounts | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y |
| 4 | Get 'Internal Transactions' by Transaction Hash | Free | Accounts | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y |
| 5 | Get a list of 'Internal' Transactions by Address | Free | Accounts | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y |
| 6 | Get "Internal Transactions" by Block Range | Free | Accounts | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y |
| 7 | Get a list of 'ERC20 - Token Transfer Events' by Address | Free | Accounts | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y |
| 8 | Get a list of 'ERC721 - Token Transfer Events' by Address | Free | Accounts | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y |
| 9 | Get a list of 'ERC1155 - Token Transfer Events' by Address | Free | Accounts | Y | N | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y |
| 10 | Get list of Blocks Validated/mined by Address | Free | Accounts | Y | Y | Y | Y | N | N | Y | Y | Y | N | Y | Y | Y | Y | Y | Y | Y | N | Y | Y | Y | Y | Y | Y | Y | Y | Y |
| 11 | Get Historical Native Token Balance for a Single Address By BlockNo | Pro | Accounts | Y | Y | Y | N | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y |
| 12 | Get Beacon Chain Withdrawals by Address and Block Range | Pro | Accounts | Y | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N |
| 13 | Get Withdrawal Transactions | Free | Accounts | Y | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N |
| 14 | Get Polygon Plasma Deposit List | Free | Accounts | N | N | Y | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N |
| 15 | Get Estimated Block Countdown Time by BlockNo | Free | Blocks | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y |
| 16 | Get Block Number by Timestamp | Free | Blocks | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y |
| 17 | Get Daily Average Block Size (Pro) | Pro | Blocks | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y |
| 18 | Get Daily Block Rewards (Pro) | Pro | Blocks | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y |
| 19 | Get Block Rewards by BlockNo | Free | Blocks | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y |
| 20 | Get Daily Average Time for A Block to be Included in the Blockchain (Pro) | Pro | Blocks | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y |
| 21 | Get Daily Uncle Block Count and Rewards (Pro) | Pro | Blocks | Y | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N |
| 22 | Get Daily Block Count and Rewards | Pro | Blocks | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y |
| 23 | Get Contract ABI for Verified Contract Source Codes | Free | Contracts | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y |
| 24 | Get Contract Source Code for Verified Contract Source Codes | Free | Contracts | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y |
| 25 | Verify Source Code | Free | Contracts | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y |
| 26 | Check Source Code Verification Submission Status | Free | Contracts | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y |
| 27 | Check Proxy Contract Verification Submission Status | Free | Contracts | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y |
| 28 | Get Contract Creator and Creation Tx Hash | Free | Contracts | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y |
| 29 | Verify Proxy Contract | Free | Contracts | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y | Y |
| 30 | Get Estimation of Confirmation Time | Free | Gas Tracker | Y | N | N | N | N | N | N | Y | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N | N |
| 31 | Get Gas Oracle | Free | Gas Tracker | Y | Y | Y | N | N | N | N | Y | Y | N | N | Y | N | Y | N | N | N | N | N | N | N | N | N | N | N | N | N | 