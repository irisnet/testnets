# 转账操作

通过执行以下命令，可以在在账户之间完成iris token的转移
```

    iriscli send --amount=10iris --chain-id=<name_of_testnet_chain> --name=<key_name> --to=<destination_address>
```   

在完成转账后，通过以下命令可以验证转账是否成功：
```
    iriscli account <origin_address>
    iriscli account <destination_address>
```   

同时，你也可以查询在某一个区块上的账户余额。
```
gaiacli account <your_address> --block=<block_height>
```

