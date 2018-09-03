# How To Get Prepared

## Install Go

Please follow the following instructions to get Go1.10 installed on your server.

```
$ sudo add-apt-repository ppa:gophers/archive
$ sudo apt-get update
$ sudo apt-get install golang-1.10-go
```

Then, you need to add the GOPATH to system path:

```
echo "export PATH=$PATH:/usr/lib/go-1.10/bin" >> ~/.bash_profile
source ~/.bash_profile

```

Next, you need to set the variables of `$GOPATH`, `$GOBIN`, and `$PATH`:

```
mkdir -p $HOME/go/bin
echo "export GOPATH=$HOME/go" >> ~/.bash_profile
echo "export GOBIN=$HOME/go/bin" >> ~/.bash_profile
echo "export PATH=$PATH:$HOME/go/bin" >> ~/.bash_profile
source ~/.bash_profile
```

To verify you have the proper version of Go installed, you could run:

```
go version
```
and see the following"
```
go version go1.10 linux/amd64
```

## Install IRISHub for Hackathon

### Build from source
After setting up the environment, you could get the correct version of IRISHub installed on your server. Please don't mix the version with which is used for Fuxi testnet.

```
mkdir -p $GOPATH/src/github.com/irisnet
cd $GOPATH/src/github.com/irisnet
git clone https://github.com/irisnet/irishub
cd irishub && git checkout gog/develop
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
make get_vendor_deps && make install
```

You could verify you have the right version installed by running the following commands: 

```
$ iris version
Game Of Genesis
    
$ iriscli version
Game Of Genesis
```

### Download Directly

You could download directrly from our GitHub repo https://github.com/irisnet/irishub/releases/download/gog/irishub_gog_linux_amd64.zip . 


## Configure Your Node

### Initialize Your Node

These instructions are for setting up a brand new full node from scratch.

First, initialize the node and create the necessary config files:
```
iris init --name <your_custom_name> --home=<IRISHOME>
```
Note: Only ASCII characters are supported for the `--name`. Using Unicode characters will render your node unreachable.
The default $IRISHOME is `~/.iris` , You can edit this name later, in the `~/.iris/config/config.toml` file:

### Get Configuration Files

Please download the genesis and the default config file.
```
cd $IRISHOME/config/
rm genesis.json
rm config.toml
wget https://raw.githubusercontent.com/irisnet/testnets/master/Game%20Of%20Genesis/config/genesis.json
wget https://raw.githubusercontent.com/irisnet/testnets/master/Game%20Of%20Genesis/config/config.toml
```

### Change Moniker

After installing irishub, you could customize the configuteration of your node. The config file is located at `<IRISHOME>/config` folder. The `moniker` field is used to idendify this node. In Gema of Genesis, your `moniker` should be composed of name of your team;PGP figer print.

For example, one sentry node of a team 'Satoshi' could be:
```
Satoshi-5AF45A077B12D406-1
```

### Set Seed Node

Seed node is used to help peers find each other. IRIS Foundation will run several official seed nodes. You could put their address in the `seed` field. 

```
seeds="9849b15b61eea1b9ae51d8d42e2b6260fc949c62@54.176.242.8:36656,217b161adf9824114a4f332c5887010c34783ae2@52.8.50.183:36656"
```

## Enable Port

You will need to set `26656` port to get connected with other peers and `26657` to query information of tendermint.

## Start Node

You could start your node with the following command:

```
iris start --home={PATH_TO_HOME}  1>gog.log  2>gog.error & 
```

## Check Node Status

You could use the following command to verify the status of your node:

```
iriscli status 
```

The output should be like this:

```
{"node_info":{"id":"9373fdcd7e8f692928dac70eb672ae52914201e7","listen_addr":"35.165.232.141:26656","network":"game-of-genesis","version":"0.22.6","channels":"4020212223303800","moniker":"jerry","other":["amino_version=0.10.1","p2p_version=0.5.0","consensus_version=v1/0.2.2","rpc_version=0.7.0/3","tx_index=on","rpc_addr=tcp://0.0.0.0:36657"]},"sync_info":{"latest_block_hash":"62F8CE0654C7D3718C743C869ADF288A5597BC80","latest_app_hash":"EB362F02818B8E3B3852F8C954EECDBB87A20711","latest_block_height":"24073","latest_block_time":"2018-08-31T04:00:42.616035333Z","catching_up":false},"validator_info":{"address":"8E6B1A18563EC8F1FCB0A78896CDB7B9250E6F94","pub_key":{"type":"tendermint/PubKeyEd25519","value":"ulGYvYQvcBvtjcOyF6hB3b4Paw5VxjoXLT1d8xgEmnE="},"voting_power":"0"}}
```

If you see the 	`catching_up` is `false`, it means your node is fully synced with the network, otherwise your node is still downloading blocks. Once fully synced, you could upgrade your node to a validator node. The instructions is [here]().	
