# 转账操作

1. 在本地恢复账户

创建一个新账户的操作如下：

```
iriscli keys add iris 
```

请按照提示输入密码，这个命令将创建一个代号为iris的本地账户。示例输出：

```
iriscli keys list
NAME: ADDRESS:      PUBKEY:
iris cosmosaccaddr1sn6e9c3vj89ss9f9jhtvyg94m922p69ex9asn5 cosmosaccpub1zcjduc3qnswwrzl26dafq26guswr0pyj7yeyq6xt44nul839c5hqquruegestlsjfq
*Important** write this seed phrase in a safe place.
It is the only way to recover your account if you ever forget your password.
```
⚠️：请保存好助记词，这非常重要！

你也可以通过助记词在本地恢复genesis文件中的账户，助记词就是执行gentx操作时产生的`secret`：

```
iriscli keys add test --recover
```

示例输出

```
iriscli keys list
NAME: ADDRESS:      PUBKEY:
test cosmosaccaddr1sn6e9c3vj89ss9f9jhtvyg94m922p69ex9asn5 cosmosaccpub1zcjduc3qnswwrzl26dafq26guswr0pyj7yeyq6xt44nul839c5hqquruegestlsjfq
```

2. 查询余额

执行以下操作可以得到账户的余额为50iris，因为绑定了100iris

```
iriscli account cosmosaccaddr1sn6e9c3vj89ss9f9jhtvyg94m922p69ex9asn5
```


3. 完成token的转移

通过以下命令可以生成一个新的账户：

```
	iriscli key add alice

```
注意⚠️：密码长度至少为8位。

示例输出：
```
NAME:	ADDRESS:						PUBKEY:
alice	cosmosaccaddr19hwmyksmg7886wkwjxq88spwnr3f8awmk280dz	cosmosaccpub1zcjduc3q2pgjuyw50ezech5j5naxc20srf8tw0xkx4k5rpldrlx2n8xvzw7qj63snp
**Important** write this seed phrase in a safe place.
It is the only way to recover your account if you ever forget your password.

reward spider ethics over news flower crazy ask term solar gesture artwork tonight alter audit abandon
```
成功生成账户后，会产生账户的有关信息。

你需要保存好助记词方便以后恢复账户。

之后你可以向新的账户转账：

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

