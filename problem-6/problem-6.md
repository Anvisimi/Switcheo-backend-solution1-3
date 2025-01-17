# Problem 6: Transaction Broadcaster Service

<aside>
⏰ Duration: You should not spend more than **3 hours** on this problem.
*Time estimation is for internship roles, if you are a software professional you should spend significantly less time.*

</aside>

# Task

This is a system design question. Describe in detail (~500-1000 words) the specifications on how you would design a transaction broadcaster service. You may additionally attach a drawings/diagrams/illustrations if you wish.

- You should focus on designing the software abstractions and architectural flow required to fulfill the service requirements, instead of choosing cloud services or software packages.
- Your submission will be graded on correctness, scalability, and robustness.

1. The broadcast service exposes an internal api for broadcast requests that will be used by other services.
It returns HTTP `200`, or HTTP `4xx`-`5xx` .
    
    ```jsx
    POST /transaction/broadcast 
    
    {"message_type": "add_weight(address _addr, uint256 _weight)", "data": "0xd71363280000000000000000000000005eb715d601c2f27f83cb554b6b36e047822fb70a00000000000000000000000000000000000000000000000000000000000000fa"}
    ```
    
2. Using the post request parameters, the broadcaster service signs the `data` and outputs a `signed transaction`. Next, it broadcasts the `signed transaction` to an EVM-compatible blockchain network.
    1. A broadcasted transaction might fail and if it fails, it should be retried automatically.
    2. To broadcast a signed transaction, you make a RPC request to a blockchain node. 
        1. 1% of the time, it does not respond earlier than 30 seconds. 
        2. 95% of the time it responds with a success code within 20-30 seconds. 
        3. The rest of the time it returns a failure code.
    3. There should also be a page that shows the list of transactions that passed or failed.

Additional Requirements

1. If `POST /transaction/broadcast` returns HTTP `200 OK`, it is assumed that the transaction will eventually be broadcasted successfully. If the broadcaster service restarts unexpectedly, it should still fulfill them.
2. An admin is able to, at any point in time, retry a failed broadcast.

## Solution

refer to the solution in TransactionBroadcasterService.pdf file in the problem-6 directory