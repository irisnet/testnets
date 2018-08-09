# IRISnet Testnet

## What is IRISnet

IRISnet is named after Greek goddess **Iris**, said to be the personification of the rainbow and the faithful messenger between heaven and humanity. IRIS network aims to build the foundation which facilitates construction of distributed business applications. IRIS hub will provide iServices, which allow resources such as data service and computing services being invoked across blockchains. To know more about IRISnet, please read this blog.

## How to Join Fuxi-2000 Testnet

### Step 1: Install Iris on Your Server

Please follow this [instruction](https://github.com/kidinamoto01/testnets-1/blob/master/testnets/docs/install%20iris.md) to get **Iris** installed locally.


### Step 2: Setting Up Your Node

These instructions are for setting up a brand new full node from scratch.

First, initialize the node and create the necessary config files:

```
iris init --name <your_custom_name> --home=<IRISHOME>
```

> Note: Only ASCII characters are supported for the `--name`. Using Unicode characters will render your node unreachable. 

The default \$IRISHOME is `~/.iris` , You can edit this `name` later, in the `~/.iris/config/config.toml` file:

```
# A custom human readable name for this node
moniker = "<your_custom_name>"
external_address = "<your-public-IP>"
```


Optional:
Set `addr_book_strict` to `false` to make peering more easily. 

```
addr_book_strict = false
```
Your full node has been initialized! 

### Get Configuration Files

If you want to be part of the genesis file geneartion process, please follow this [guide](https://github.com/kidinamoto01/testnets-1/blob/master/testnets/docs/Genesis%20Generation%20Process.md) to sumbmit a json file. Otherwise, you could always send related transaction to become a validator later. 

After the genesis file generation process is finished, please download the genesis and the default config file. 

```
cd $IRISHOME/config/
rm genesis.json
rm config.toml
wget https://raw.githubusercontent.com/irisnet/testnets/master/testnets/fuxi-2000/config/config.toml
wget https://raw.githubusercontent.com/irisnet/testnets/master/testnets/fuxi-2000/config/genesis.json
```

### Add Seed Nodes

Your node needs to know how to find peers. You'll need to add healthy seed nodes to `$IRISHOME/config/config.toml`. Here are some seed nodes you can use:

```
3fb472c641078eaaee4a4acbe32841f18967672c@35.165.232.141:46657
```

Meanwhile, you could add some known full node as `Persistent Peer`. Your node could connect to `sentry node` as `persistent peers`. 



### Run a Full Node

Start the full node with this command:

```
iris start --home=$IRISHOME > iris.log
```

Check that everything is running smoothly:

```
iriscli status
```
You could see the following 
```
{"node_info":{"id":"71b188e9fdefd939453b3cd10c0eae45a8d02a2b","listen_addr":"172.31.0.190:26656","network":"fuxi-2000","version":"0.22.6","channels":"4020212223303800","moniker":"name","other":["amino_version=0.10.1","p2p_version=0.5.0","consensus_version=v1/0.2.2","rpc_version=0.7.0/3","tx_index=on","rpc_addr=tcp://0.0.0.0:26657"]},"sync_info":{"latest_block_hash":"CC9BBE0B38643DAF3D9B78D928E2ACA654E5A39C","latest_app_hash":"56B9228A97D5B85BFDBEE020E597D45D427ABC43","latest_block_height":"30048","latest_block_time":"2018-08-02T08:23:44.566550056Z","catching_up":false},"validator_info":{"address":"F638F7EA8A8E4DA559A346E1C404F83941749713","pub_key":{"type":"tendermint/PubKeyEd25519","value":"oI16LfBmnP8CefSGwIjAIO3QZ05xwB1+s4oPIQ3Yaag="},"voting_power":"10"}}
```
When you see `catching_up` is `false`, it means the node is synced with the rest of testnet, otherwise it means it's still syncing.


### Upgrade to Validator Node

You now have an active full node. What's the next step? 

If you have participated in the genesis file generation process, you should be a validator once you are fully synced. 

If you miss the genesis file generation process, you can still upgrade your full node to become a IRISnet Validator. The top 100 validators have the ability to propose new blocks to the IRIS Hub. Continue onto [the Validator Setup](https://github.com/kidinamoto01/testnets-1/blob/master/testnets/docs/Setup%20A%20Validator%20Node.md).


##  Useful Links

* Riot chat: #irisvalidators:matrix.org
