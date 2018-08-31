# Setup A Validator Node

Before setting up your validator node, make sure you've already installed  **Iris** by this [guide](https://github.com/irisnet/testnets/blob/master/Game%20Of%20Genesis/get-started/Node-Setup.md)


## Get IRIS Token

You could only get some IRIS token to play the game of genesis after you have finished the reginsteration process. 

You need to get `iris` and `iriscli` installed first. Then, follow the instructions below to create a new account:

```

```

## Create A Validator

Your `fvp` can be used to create a new validator by staking tokens. You can find your validator pubkey by running:

```
iris tendermint show_validator --home=<IRIS-HOME>
```

Next, craft your `iriscli stake create-validator` command:

You can always get some `IRIS`  by using the [Faucet](https://testnet.irisplorer.io/#/faucet). Please don't abuse it. 

```
iriscli stake create-validator --amount=100000000000000000000iris --pubkey=<pubkey> --address-validator=<val_addr> --moniker=<moniker> --chain-id=game-of-genesis --name=<name>
```



### View Validator Description

View the validator's information with this command:

```
iriscli stake validator \
  --address-validator=<account_cosmosaccaddr> \
  --chain-id=game-of-genesis 
```

### Confirm Your Validator is Running

Your validator is active if the following command returns anything:

```
iriscli advanced tendermint validator-set | grep "$( tendermint show_validator)"
```

You should also be able to see your validator on the [Explorer](https://testnet.irisplorer.io). You are looking for the `bech32` encoded `address` in the `~/.iris/config/priv_validator.json` file.


### Edit Validator Description

You can edit your validator's public description. This info is to identify your validator, and will be relied on by delegators to decide which validators to stake to. Make sure to provide input for every flag below, otherwise the field will default to empty (`--moniker`defaults to the machine name).

```
iriscli stake edit-validator
  --address-validator=<account>
  --moniker="choose a moniker" \
  --website="https://irisnet.org" \
  --details=""
  --chain-id=game-of-genesis  \
  --name=<key_name>
```
