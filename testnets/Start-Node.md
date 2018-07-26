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
    {"node_info":{"id":"3fb472c641078eaaee4a4acbe32841f18967672c","listen_addr":"172.31.0.190:46656","network":"fuxi-1002","version":"0.21.0","channels":"4020212223303800","moniker":"name","other":["amino_version=0.9.9","p2p_version=0.5.0","consensus_version=v1/0.2.2","rpc_version=0.7.0/3","tx_index=on","rpc_addr=tcp://0.0.0.0:46657"]},"sync_info":{"latest_block_hash":"7B1168B2055B19F811773EEE56BB3C9ECB6F3B37","latest_app_hash":"B8F7F8BF18E3F1829CCDE26897DB905A51AF4372","latest_block_height":12567,"latest_block_time":"2018-07-26T02:43:56.757513477Z","syncing":false},"validator_info":{"address":"CAF80DAEC0F4A7036DD2116B56F89B07F43A133E","pub_key":{"type":"AC26791624DE60","value":"Cl6Yq+gqZZY14QxrguOaZqAswPhluv7bDfcyQx2uSRc="},"voting_power":0}}
```
通过以上命令可以查看状态：

* `“syncing":false`表示节点与网络保持同步
* `“syncing":true`表示节点正在同步区块
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
    iriscli status
```
示例输出：
```json
    {"node_info":{"id":"3fb472c641078eaaee4a4acbe32841f18967672c","listen_addr":"172.31.0.190:46656","network":"fuxi-1002","version":"0.21.0","channels":"4020212223303800","moniker":"name","other":["amino_version=0.9.9","p2p_version=0.5.0","consensus_version=v1/0.2.2","rpc_version=0.7.0/3","tx_index=on","rpc_addr=tcp://0.0.0.0:46657"]},"sync_info":{"latest_block_hash":"7B1168B2055B19F811773EEE56BB3C9ECB6F3B37","latest_app_hash":"B8F7F8BF18E3F1829CCDE26897DB905A51AF4372","latest_block_height":12567,"latest_block_time":"2018-07-26T02:43:56.757513477Z","syncing":false},"validator_info":{"address":"CAF80DAEC0F4A7036DD2116B56F89B07F43A133E","pub_key":{"type":"AC26791624DE60","value":"Cl6Yq+gqZZY14QxrguOaZqAswPhluv7bDfcyQx2uSRc="},"voting_power":100}}
```
通过以上命令可以查看状态：

`“syncing"`:false表示节点与网络保持同步

`"latest_block_height"`:表示最新的区块高度
