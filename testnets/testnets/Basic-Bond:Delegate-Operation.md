# 在IRIShub中的基本绑定/委托操作

## 成为验证人候选人

想要绑定成为一个验证人候选人，需要执行以下步骤：

1. 确认账户的余额是否大于0

执行以下命令：
```
    iriscli keys list
```
示例输出
```
    iriscli keys list
    NAME: ADDRESS:      PUBKEY:
    bob cosmosaccaddr1sn6e9c3vj89ss9f9jhtvyg94m922p69ex9asn5 cosmosaccpub1zcjduc3qnswwrzl26dafq26guswr0pyj7yeyq6xt44nul839c5hqquruegestlsjfq
    tom cosmosaccaddr1nexsact44a7e5uvvlxvvad9j2x596mphse7kd6 cosmosaccpub1zcjduc3q5je66e6gntchys3cf67hr9xv57ry3m8rtdu8clptxejav2t70t0sgnaex0
```
假设你需要使用bob的账户来绑定成为一个验证人。

查询bob账户的余额
```
    iriscli account cosmosaccaddr1sn6e9c3vj89ss9f9jhtvyg94m922p69ex9asn5
```

若账户余额大于0，则可以执行以下操作。

2.  获得节点的公钥

执行以下命令：
```
iris tendermint show_validator --home=<home>
```

示例输出：

```
iris tendermint show_validator --home=/Users/suyu/Documents/bianjie/fuxi/1000
cosmosvalpub1zcjduc3qlz3chwnlqz9sn2mv0yhjfxzg38anh0emtc4weuttqhcykj4xphhqtdclw8
```

以上的返回值就是节点的公钥

3. 创建验证人候选人

执行以下命令：

```
iriscli stake create-validator --amount=100iris --pubkey=<pubkey> --address-validator=<val_addr> --moniker=<moniker> --chain-id=fuxi-1002 --name=<name>
```

示例输出：

```
Defaulting to next sequence number: 10
Committed at block 2286. Hash: CAE473BC392A80000802CEEB14682F9DEA06AF4D
```

通过执行以下操作确认节点的运行状态：

```
iriscli status
```

示例输出：

```
{"node_info":{"id":"e43f676fd19ad5f1869d9edd2a6800dc48f40335","listen_addr":"172.20.155.233:26656","network":"fuxi-1000","version":"0.21.0","channels":"40202122233038","moniker":"bianjie","other":["amino_version=0.10.1","p2p_version=0.5.0","consensus_version=v1/0.2.2","rpc_version=0.7.0/3","tx_index=on","rpc_addr=tcp://0.0.0.0:26657"]},"sync_info":{"latest_block_hash":"2A815BA6C6F11991F3BA9B57554617E6646C0D64","latest_app_hash":"38E4313CD4A50513BA0A259D0F86C5845DF9A12C","latest_block_height":"172","latest_block_time":"2018-07-17T14:27:41.023715315Z","catching_up":false},"validator_info":{"address":"F23FF36BD5B90C33CE3A03ED72DBDCF5EC07D6AF","pub_key":{"type":"tendermint/PubKeyEd25519","value":"2JoNf1gavJ1d6XFIumO1Mki5GVMOcg58AioHksU3maE="},"voting_power":"100"}}
```

若观察到`voting_power`的权重不为0,则说明绑定成功。



4. 编辑验证人信息

你可以通过以下标签添加有关验证人的信息： `--website`, `--keybase-sig`,或者`-details` 

执行以下命令：

```
iriscli stake edit-validator --details="To the cosmos !" --website="https://cosmos.network"
```

通过以下命令查看验证人信息：

```
iriscli stake validator --address-validator=<your_address> --chain-id=<name_of_the_testnet_chain>
```
可以看到信息有所变化。


5. 委托给验证人

在测试网中，你可以将iris代币委托给验证人，这个验证人也可以是你自己。

命令如下：
```
iriscli stake delegate --amount=10iris --address-delegator=<your_address> --address-validator=<bonded_validator_address> --name=<key_name> --chain-id=<name_of_testnet_chain>
```

通过查看账户余额的变化，可以看到余额减少了。



6. 将验证人解绑

通过以下命令可以完成对一个验证人的解绑，解绑的单位为share，share可以是一个小数，例如12.1,也可以使用MAX表示完全解绑。

```
iriscli stake unbond --address-delegator=<your_address> --address-validator=<bonded_validator_address> --shares=MAX --name=<key_name> --chain-id=<name_of_testnet_chain>
```

在解绑完成后，你可以通过查询账户余额的方式验证解绑是否成功。若解绑成功，余额将增加。

```
iriscli account <your_address>
```