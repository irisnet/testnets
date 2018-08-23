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
    iris	local	faa1689na4jtu0c236h9vhq9w292q3xdjt2tmu4pjv	fap1addwnpepq0h3wad9ps95nwwfh94pnl6y7kafgfd5xu5z2atk8ts6zuqe0zeggm0euh5
```
假设你需要使用bob的账户来绑定成为一个验证人。

查询bob账户的余额
```
    iriscli account faa1689na4jtu0c236h9vhq9w292q3xdjt2tmu4pjv
```

若账户余额大于0，则可以执行以下操作。

2.  获得节点的公钥

执行以下命令：
```
iris tendermint show_validator --home=<home>
```

示例输出：

```
iris tendermint show_validator --home=$IRISHOME
fvp1zcjduepq5zxh5t0sv6w07qne7jrvpzxqyrkaqe6ww8qp6l4n3g8jzrwcdx5qr8qdx7
```

以上的返回值就是节点的公钥

3. 创建验证人候选人

执行以下命令：

```
iriscli stake create-validator --amount=100iris --pubkey=<pubkey> --address-validator=<val_addr> --moniker=<moniker> --chain-id=fuxi-2000 --from=<name>
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
iriscli stake edit-validator --details="details!" --website="https://irisnet.org"  --from=name  --address-validator=<address > --chain-id=fuxi-2000

```

通过以下命令查看验证人信息：

```
iriscli stake validator <your_address> --chain-id=<name_of_the_testnet_chain>
```

可以看到信息有所变化。


5. 委托给验证人

在测试网中，你可以将iris代币委托给验证人，这个验证人也可以是你自己。

查询所有的验证人信息：

```
iriscli stake validators
```
示例输出：
```
Validator
Owner: faa1kql3yp6yl4pswc9svrh5yrq0kk7ymppdu2f9m0
Validator: fvp1zcjduepq2nkjmgyu8ccwq8tz5wejph7eh2n00mectch4gtxwf2xwrkdnst5qs7at3l
Revoked: false
Status: Bonded
Tokens: 100.0000000000
Delegator Shares: 100.0000000000
Description: {yelong   }
Bond Height: 0
Proposer Reward Pool:
Commission: 0/1
Max Commission Rate: 0/1
Commission Change Rate: 0/1
Commission Change Today: 0/1
Previous Bonded Tokens: 0/1

Validator
Owner: faa1k2atwc52n9mwstexn4rqatsgm59v0pg9qe5zs5
Validator: fvp1zcjduepqm2f0ganwqsqk6dt3xs2r20mzk9lazzkfqm7kynr4e83grgzk8lcq82cy87
Revoked: false
Status: Bonded
Tokens: 100.0000000000
Delegator Shares: 100.0000000000
Description: {silei   }
Bond Height: 0
Proposer Reward Pool:
Commission: 0/1
Max Commission Rate: 0/1
Commission Change Rate: 0/1
Commission Change Today: 0/1
Previous Bonded Tokens: 0/1
```


命令如下：
```
iriscli stake delegate --amount=10iris --address-delegator=<your_address> --address-validator=<bonded_validator_owner_address> --name=<key_name> --chain-id=<name_of_testnet_chain>
```

命令举例：

```
iriscli stake delegate --amount=10iris --address-delegator=cosmosaccaddr12njc8n8e7fn8n2cf38qxdu44yajfch6zydhf2g --address-validator=cosmosaccaddr15pzxwa0ekcj9a2xpngqxazef5lkyhhskmru36f --name=test --chain-id=fuxi-1002
```

通过以下命令查询验证人信息的变化。
```
iriscli stake validator <bonded_validator_owner_address>
```

命令举例如下，可以看到验证人的绑定数量增加了。

```
Validator
Owner: faa1kql3yp6yl4pswc9svrh5yrq0kk7ymppdu2f9m0
Validator: fvp1zcjduepq2nkjmgyu8ccwq8tz5wejph7eh2n00mectch4gtxwf2xwrkdnst5qs7at3l
Revoked: false
Status: Bonded
Tokens: 110.0000000000
Delegator Shares: 110.0000000000
Description: {yelong   }
Bond Height: 0
Proposer Reward Pool:
Commission: 0/1
Max Commission Rate: 0/1
Commission Change Rate: 0/1
Commission Change Today: 0/1
Previous Bonded Tokens: 0/1
```

同时通过查看账户余额的变化，可以看到余额减少了。



6. 将验证人解绑

通过以下命令可以完成对一个验证人的解绑，解绑的单位为share，share可以是一个小数，例如12.1,也可以使用MAX表示完全解绑。

```
iriscli stake unbond begin --address-delegator=<your_address> --address-validator=<bonded_validator_address>  --shares-amount=10 --from=<key_name> --chain-id=<name_of_testnet_chain>
```

在解绑完成后，你可以通过查询账户余额的方式验证解绑是否成功。若解绑成功，余额将增加。

```
iriscli account <your_address>
```


### 验证人被处罚后如何处理

若节点不能保证持续在线，则会被罚没一部分抵押的iris。在fuxi-2000中，如果前200块中，某个验证人错过了超过100个投票，那么验证人将被惩罚，即被slash。

这时，如果查询本地状态会发现`power`变为0.

```json
{"node_info":{"id":"e43f676fd19ad5f1869d9edd2a6800dc48f40335","listen_addr":"172.20.155.233:26656","network":"fuxi-1000","version":"0.21.0","channels":"40202122233038","moniker":"bianjie","other":["amino_version=0.10.1","p2p_version=0.5.0","consensus_version=v1/0.2.2","rpc_version=0.7.0/3","tx_index=on","rpc_addr=tcp://0.0.0.0:26657"]},"sync_info":{"latest_block_hash":"2A815BA6C6F11991F3BA9B57554617E6646C0D64","latest_app_hash":"38E4313CD4A50513BA0A259D0F86C5845DF9A12C","latest_block_height":"172","latest_block_time":"2018-07-17T14:27:41.023715315Z","catching_up":false},"validator_info":{"address":"F23FF36BD5B90C33CE3A03ED72DBDCF5EC07D6AF","pub_key":{"type":"tendermint/PubKeyEd25519","value":"2JoNf1gavJ1d6XFIumO1Mki5GVMOcg58AioHksU3maE="},"voting_power":"0"}}
```

通过执行以下命令，恢复验证人的身份。

```
iriscli stake unrevoke <address> --name=<name>--chain-id=fuxi-2000
```
