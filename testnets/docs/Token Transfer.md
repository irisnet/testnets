# Send Tokens 

1.  Create an account locally:
```
    iriscli keys add iris 
```

Please follow the prompt to set your password. Caution⚠️: Your password must contain at least 8 characters! This command will create an iris account in your machine.

Example Output:
```
    iriscli keys list
NAME:   TYPE:   ADDRESS:                        PUBKEY:
iris    local   faa1689na4jtu0c236h9vhq9w292q3xdjt2tmu4pjv  fap1addwnpepq0h3wad9ps95nwwfh94pnl6y7kafgfd5xu5z2atk8ts6zuqe0zeggm0euh5
**Important** write this seed phrase in a safe place.
It is the only way to recover your account if you ever forget your password.
```

⚠️：Please keep your seed phrase carefully, this is very imortant！

You can recover your account in genesis file locally.Seed phrase is the secret field in the output after executing keys add command or gentx command.

```
    iriscli keys add test --recover
```

Example Output:

```
    iriscli keys list
    NAME: ADDRESS:      PUBKEY:
    test local   faa1689na4jtu0c236h9vhq9w292q3xdjt2tmu4pjv  fap1addwnpepq0h3wad9ps95nwwfh94pnl6y7kafgfd5xu5z2atk8ts6zuqe0zeggm0euh5
```

2. Get Some Coins

You can visit this faucet page to get some **iris**: https://testnet.irisplorer.io/#/faucet

Input your account address and do the bot authentication. After the authentication, you can click the Send me IRIS button to get 10 iris tokens.


3. Query Balance


Execute the following command and you will find that there are some iris tokens in this account.
```
    iriscli account faa1rdnxluhklx4k3s34d9fjg0vaacyfa3ml20ftgm
```

An account will be stored in iris network only after there is at least one transaction related to this account in iris network. So, if your account doesn't have any transaction, executing this command will give you an error message.


```
ERROR: No account with address faa1689na4jtu0c236h9vhq9w292q3xdjt2tmu4pjv was found in the state.
Are you sure there has been a transaction involving it?
```

4. Send tokens around

You can send some coins to your newly created account:

```
        iriscli send --amount=10iris --chain-id=<name_of_testnet_chain> --name=<key_name> --to=<destination_address>
```

After executing the send command, you can use the following command to check if you send your coins successfully:

```
        iriscli account <origin_address>
        iriscli account <destination_address>
```