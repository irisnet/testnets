# Setup A Validator Node

Before setting up your validator node, make sure you've already installed  **Iris** by this [guide](https://github.com/irisnet/testnets/blob/master/Game%20Of%20Genesis/get-started/Node-Setup.md)


## Get IRIS Token

### create account
You could only get some IRIS token to play the game of genesis after you have finished the registration process.

You need to get `iris` and `iriscli` installed first. Then, follow the instructions below to create a new account:

```
iriscli keys add <NAME_OF_KEY>
```

Then, you should print a password of at least 8 characters.

The output will look like the following:
```
NAME:	TYPE:	ADDRESS:						PUBKEY:
tom	local	faa1arlugktm7p64uylcmh6w0g5m09ptvklxm5k69x	fap1addwnpepqvlmtpv7tke2k93vlyfpy2sxup93jfulll6r3jty695dkh09tekrzagazek
**Important** write this seed phrase in a safe place.
It is the only way to recover your account if you ever forget your password.

blast change tumble toddler rival ordinary chicken dirt physical club few language noise oak moment consider enemy claim elephant cruel people adult peanut garden
```

You could see the address and public key of this account. Please node that account address in IRISnet will start with `faa` and public key of account will start with `fap`.

The seed phrase of this account will also be displayed. You could use these 24-character phrase to recover this account in another server. The recover command is:
```
iriscli keys add <NAME_OF_KEY> --recover
```

Once you have created your own address, please comment this issue with your address and you could receive 1,000iris soon. Each team will receive once, then you could use these tokens to stake as a validator. The following command is used to check the balance of your account:
```
iriscli bank account <ACCOUNT> 
```
### Claim Tokens


Please reply this [issue](https://github.com/irisnet/testnets/issues/69) with the generated address. The team will send you 1000*10^18iris after verification.


## Create A Validator


You need to get the public key of your node before upgrade your node to a validator node. The public key of will start `fvp`, it can be used to create a new validator by staking tokens. 

You can find your validator pubkey by running:

```
iris tendermint show-validator --home=<IRIS-HOME>
```
Example output:
```
fvp1zcjduepqv7z2kgussh7ufe8e0prupwcm7l9jcn2fp90yeupaszmqjk73rjxq8yzw85
```
Next, use the output as `<pubkey>` field for `iriscli stake create-validator` command:


```
iriscli stake create-validator --amount=<amount>iris --pubkey=<pubkey> --address-validator=<val_addr> --moniker=<moniker> --chain-id=game-of-genesis --name=<key_name> 
```
Please note the **amount** needs to be more than 10^18iris.

Please you need to specify the name of your validator to **name of your team**. 

In this way, to stake 1IRIS, you need to do:

```
iriscli stake create-validator --chain-id=<chain-id> --from=<key name> --fee=0.004iris \
    --pubkey=<validator public key> --amount=10iris --moniker=<validator name> \
    --commission-max-change-rate=0.1 --commission-max-rate=0.5 --commission-rate=0.1
```
Don't forget the `fee` and `gas` field.

### View Validator Info

View the validator's information with this command:

```
iriscli stake validator <account> --chain-id=game-of-genesis
```

The `<account>` is start with config.toml

### Confirm Your Validator is Running

Your validator is active if the following command returns anything:

```
iriscli status
```

You should also be able to see your power is above 0. Also, you should see validator on the [Explorer](https://testnet.irisplorer.io).


### Edit Validator Description

You can edit your validator's public description. This info is to identify your validator, and will be relied on by delegators to decide which validators to stake to. Make sure to provide input for every flag below, otherwise the field will default to empty (`--moniker`defaults to the machine name).

You should put your name of your team in `details`. 

```
iriscli stake edit-validator --from=<key> --chain-id=game-of-genesis \
    --moniker="choose a moniker" --identity=<uport_or_keybase_fingerprint> \
    --website="https://irisnet.org" --details="team"
```
