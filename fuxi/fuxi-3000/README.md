# How to Finish A Testnet Task

## What is IRISnet incentivized testnet game?

IRISnet development team will publish several task for each iteration of Fuxi testnet.

You need to use keybase to generate your own pop ID first. 

The instructions are the following: 

## Tasks for Fuxi-3000

No|Name|Details|Criteria|Reward
----- | ----- | ----- | ----- |-----
1|Setup a Full node|set up a full node and use `name-pgpid` as `monkier`|Submit your IP and thr team could check the configutration of your node| 100 iris
2|Complete one token transaction| Create an account, and send this new account some **iris** |Run this send transaction with your `name-pgpid` as memo. Submit your transaction hash and the team could verify the details of transaction| 20 iris
3|Delegate some **iris** to a validator|Get some **iris** from faucet and delegate to a validator|Run this send transaction with your `name-pgpid` as memo. Submit your transaction hash and the team could verify the details of transaction| 100 iris
4|Unbond some **iris** from a validator|Complete an unbond transaction|Run `unbond begin` and `unbond complete` transaction with your `name-pgpid` as memo. Submit your transaction hash and the team could verify the details of transaction|  100 iris
Example:

1. Full node info

2. **Details of a transaction:**

   Command:

   ```
   iriscli advanced tendermint tx F6627E89466151C95C4381D5AB4EB5AB5E81D78C
   ```
   The output:
   ```
   {
     "hash": "F6627E89466151C95C4381D5AB4EB5AB5E81D78C",
     "height": "15499",
     "tx": {
       "type": "auth/StdTx",
       "value": {
         "msg": [
           {
             "type": "cosmos-sdk/Send",
             "value": {
               "inputs": [
                 {
                   "address": "faa1893x4l2rdshytfzvfpduecpswz7qtpstpr9x4h",
                   "coins": [
                     {
                       "denom": "iris",
                       "amount": "10"
                     }
                   ]
                 }
               ],
               "outputs": [
                 {
                   "address": "faa1uj5gfwrklrdzgqjwh0mnayclamkcy3jumg8g06",
                   "coins": [
                     {
                       "denom": "iris",
                       "amount": "10"
                     }
                   ]
                 }
               ]
             }
           }
         ],
         "fee": {
           "amount": [
             {
               "denom": "",
               "amount": "0"
             }
           ],
           "gas": "200000"
         },
         "signatures": [
           {
             "pub_key": {
               "type": "tendermint/PubKeySecp256k1",
               "value": "A3eTe0+O8flcd9rEDgC/DVmPCWayZBuJxf6AdVnG+IGY"
             },
             "signature": {
               "type": "tendermint/SignatureSecp256k1",
               "value": "MEUCIQCsNb2ylR+kefwqrlAz0WA1azbcq1B7LH0ZD/bop0DQuQIgTQ2LwWPP5sr9ObjU5dUiSlg8TZmQztORKyCPXEiRqQA="
             },
             "account_number": "15",
             "sequence": "56"
           }
         ],
         "memo": "suyu-5AF45A077B12D406"
       }
     },
     "result": {
       "log": "Msg 0: ",
       "gas_used": "3507",
       "tags": [
         {
           "key": "c2VuZGVy",
           "value": "ZmFhMTg5M3g0bDJyZHNoeXRmenZmcGR1ZWNwc3d6N3F0cHN0cHI5eDRo"
         },
         {
           "key": "cmVjaXBpZW50",
           "value": "ZmFhMXVqNWdmd3JrbHJkemdxandoMG1uYXljbGFta2N5M2p1bWc4ZzA2"
         }
       ],
       "fee": {}
     }
   }
   ```

3. Delegate some tokens

   Command:

   ```
   iriscli advanced tendermint tx FF3E52C1CB69C8050E93FAF7165EDC52502710F9
   ```

   The output:

   ```
   iriscli advanced tendermint tx FF3E52C1CB69C8050E93FAF7165EDC52502710F9
   {
     "hash": "FF3E52C1CB69C8050E93FAF7165EDC52502710F9",
     "height": "15621",
     "tx": {
       "type": "auth/StdTx",
       "value": {
         "msg": [
           {
             "type": "cosmos-sdk/MsgDelegate",
             "value": {
               "delegator_addr": "faa1uj5gfwrklrdzgqjwh0mnayclamkcy3jumg8g06",
               "validator_addr": "faa1uj5gfwrklrdzgqjwh0mnayclamkcy3jumg8g06",
               "delegation": {
                 "denom": "iris",
                 "amount": "10"
               }
             }
           }
         ],
         "fee": {
           "amount": [
             {
               "denom": "",
               "amount": "0"
             }
           ],
           "gas": "200000"
         },
         "signatures": [
           {
             "pub_key": {
               "type": "tendermint/PubKeySecp256k1",
               "value": "At4WnZ3aLJsvCUb59rhPRJkYWTFNTJrPjLn4jUT7VN1F"
             },
             "signature": {
               "type": "tendermint/SignatureSecp256k1",
               "value": "MEQCIFP1/kZxjD9KYM0B6b9pIUarVO0blPtsC2vj5lWKkYg6AiA0blXXJ5M/gD3EUMixzTM0sDbn65+KmgBijsL5hX3pQw=="
             },
             "account_number": "10",
             "sequence": "2"
           }
         ],
         "memo": "suyu-5AF45A077B12D406"
       }
     },
     "result": {
       "log": "Msg 0: ",
       "gas_used": "14796",
       "tags": [
         {
           "key": "YWN0aW9u",
           "value": "ZGVsZWdhdGU="
         },
         {
           "key": "ZGVsZWdhdG9y",
           "value": "ZmFhMXVqNWdmd3JrbHJkemdxandoMG1uYXljbGFta2N5M2p1bWc4ZzA2"
         },
         {
           "key": "ZGVzdGluYXRpb24tdmFsaWRhdG9y",
           "value": "ZmFhMXVqNWdmd3JrbHJkemdxandoMG1uYXljbGFta2N5M2p1bWc4ZzA2"
         }
       ],
       "fee": {}
     }
   }
   ```

4. Unbond your delegation

Command:

```
iriscli stake unbond begin --address-delegator=<your_address> --address-validator=<bonded_validator_address>  --shares-amount=10 --from=<key_name> --chain-id=<name_of_testnet_chain>
```


```
{
  "hash": "B9A1A64BBCAFE31D13D0D9DE00BB4DC883BC4F42",
  "height": "15676",
  "tx": {
    "type": "auth/StdTx",
    "value": {
      "msg": [
        {
          "type": "cosmos-sdk/BeginUnbonding",
          "value": {
            "delegator_addr": "faa1uj5gfwrklrdzgqjwh0mnayclamkcy3jumg8g06",
            "validator_addr": "faa1uj5gfwrklrdzgqjwh0mnayclamkcy3jumg8g06",
            "shares_amount": "10"
          }
        }
      ],
      "fee": {
        "amount": [
          {
            "denom": "",
            "amount": "0"
          }
        ],
        "gas": "200000"
      },
      "signatures": [
        {
          "pub_key": {
            "type": "tendermint/PubKeySecp256k1",
            "value": "At4WnZ3aLJsvCUb59rhPRJkYWTFNTJrPjLn4jUT7VN1F"
          },
          "signature": {
            "type": "tendermint/SignatureSecp256k1",
            "value": "MEQCICf5W4CrVDSEW1D7cRIEyM0Eaye5wScN3UQwMTs2Yeo+AiBG23eLvIDeB2rOyhm3FIH6HhBmM1pFfecDnQbXXr828g=="
          },
          "account_number": "10",
          "sequence": "3"
        }
      ],
      "memo": "suyu-5AF45A077B12D406"
    }
  },
  "result": {
    "log": "Msg 0: ",
    "gas_used": "16695",
    "tags": [
      {
        "key": "YWN0aW9u",
        "value": "YmVnaW4tdW5ib25kaW5n"
      },
      {
        "key": "ZGVsZWdhdG9y",
        "value": "ZmFhMXVqNWdmd3JrbHJkemdxandoMG1uYXljbGFta2N5M2p1bWc4ZzA2"
      },
      {
        "key": "c291cmNlLXZhbGlkYXRvcg==",
        "value": "ZmFhMXVqNWdmd3JrbHJkemdxandoMG1uYXljbGFta2N5M2p1bWc4ZzA2"
      }
    ],
    "fee": {}
  }
}
```

### How to submit evidence

You could submit the transaction Hash to prove you have finished the tasks above by replying this issue:

```
GitHub ID: XXXX
pgp ID: XXX
Task 1: Hash
Task 2: Hash
Task 3: Hash
Task 4: Hash

```