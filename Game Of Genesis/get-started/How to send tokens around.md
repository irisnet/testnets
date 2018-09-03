# Token Transfer

## Send tokens

You can send some coins to your newly created account:

```
        iriscli send --amount=1000000000000000000iris --chain-id=<name_of_testnet_chain> --name=<key_name> --to=<destination_address> --fee=400000000000000iris --gas=20000  --node=localhost:36657
```
⚠️： You also need --node if you haven't ran a full node locally.

After executing the send command, you can use the following command to check if you send your coins successfully:

```
        iriscli account <origin_address>
        iriscli account <destination_address>
```

## Query Balance


Execute the following command and you will find that there are some iris tokens in this account.
```
    iriscli account faa1rdnxluhklx4k3s34d9fjg0vaacyfa3ml20ftgm --node=localhost:36657
```
⚠️：  --node is the address of a full node in game-of-genesis. You need add this parameter if you haven't ran a full node locally.

An newly created account will be stored in iris network only after there is at least one transaction related to this account in iris network. So, if your account doesn't have any transaction, executing this command will give you an error message.


```
ERROR: No account with address faa1689na4jtu0c236h9vhq9w292q3xdjt2tmu4pjv was found in the state.
Are you sure there has been a transaction involving it?
```
