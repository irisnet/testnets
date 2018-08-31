# Setup A Validator Node

Before setting up your validator node, make sure you've already installed  **Iris** by this [guide]()




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

### Edit Validator Description

You can edit your validator's public description. This info is to identify your validator, and will be relied on by delegators to decide which validators to stake to. Make sure to provide input for every flag below, otherwise the field will default to empty (`--moniker`defaults to the machine name).

```
iriscli stake edit-validator
  --address-validator=<account_cosmosaccaddr>
  --moniker="choose a moniker" \
  --website="https://cosmos.network" \
  --details=""
  --chain-id=game-of-genesis  \
  --name=<key_name>
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

