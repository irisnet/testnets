# 如何使用一个IRIShub节点

## Start命令

通过以下命令启动全节点，并将日志输出到文件中：
```
    iris start --home {path_to_your_home} > log文件地址 &
```
通过执行以下操作确认节点的运行状态：
```
    iriscli status
```
示例输出：
```json
    {"node_info":{"id":"e43f676fd19ad5f1869d9edd2a6800dc48f40335","listen_addr":"172.20.155.233:26656","network":"fuxi-1000","version":"0.21.0","channels":"40202122233038","moniker":"bianjie","other":["amino_version=0.10.1","p2p_version=0.5.0","consensus_version=v1/0.2.2","rpc_version=0.7.0/3","tx_index=on","rpc_addr=tcp://0.0.0.0:26657"]},"sync_info":{"latest_block_hash":"2A815BA6C6F11991F3BA9B57554617E6646C0D64","latest_app_hash":"38E4313CD4A50513BA0A259D0F86C5845DF9A12C","latest_block_height":"172","latest_block_time":"2018-07-17T14:27:41.023715315Z","catching_up":false},"validator_info":{"address":"F23FF36BD5B90C33CE3A03ED72DBDCF5EC07D6AF","pub_key":{"type":"tendermint/PubKeyEd25519","value":"2JoNf1gavJ1d6XFIumO1Mki5GVMOcg58AioHksU3maE="},"voting_power":"100"}}
```
通过以上命令可以查看状态：

* `“catching_up":false`表示节点与网络保持同步
* `“catching_up":true`表示节点正在同步区块
* `"latest_block_height"` 表示最新的区块高度


## Reset命令

### 重置IRIShub节点流程如下：

1. 关闭iris进程
```
    kill -9 <PID>
```

若Genesis文件有变动，则需要下载新的文件到$IRISHOME/config目录下。

2. 重置iris
```
    iris unsafe_reset_all --home=<home>
```

3. 重新启动

通过以下命令启动全节点，并将日志输出到文件中：
```
    iris start --home <path_to_your_home> > log文件地址 &
```
通过执行以下操作确认节点的运行状态：
```
    gaiacli status
```
示例输出：
```json
    {"node_info":{"id":"e43f676fd19ad5f1869d9edd2a6800dc48f40335","listen_addr":"172.20.155.233:26656","network":"fuxi-1000","version":"0.21.0","channels":"40202122233038","moniker":"bianjie","other":["amino_version=0.10.1","p2p_version=0.5.0","consensus_version=v1/0.2.2","rpc_version=0.7.0/3","tx_index=on","rpc_addr=tcp://0.0.0.0:26657"]},"sync_info":{"latest_block_hash":"2A815BA6C6F11991F3BA9B57554617E6646C0D64","latest_app_hash":"38E4313CD4A50513BA0A259D0F86C5845DF9A12C","latest_block_height":"172","latest_block_time":"2018-07-17T14:27:41.023715315Z","catching_up":false},"validator_info":{"address":"F23FF36BD5B90C33CE3A03ED72DBDCF5EC07D6AF","pub_key":{"type":"tendermint/PubKeyEd25519","value":"2JoNf1gavJ1d6XFIumO1Mki5GVMOcg58AioHksU3maE="},"voting_power":"100"}}
```
通过以上命令可以查看状态：

`“catching_up"`:false表示节点与网络保持同步

`"latest_block_height"`:表示最新的区块高度
