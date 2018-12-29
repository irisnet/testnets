# Genesis文件生成流程

## 前提条件

每个希望成为验证人的参与者确保按照以下[流程](https://github.com/irisnet/irishub/blob/release0.9/docs/zh/get-started/Install-the-Software.md)安装了对应版本的软件：**iris**


## Step 1: 创建账户  


首先如果你没有现成的账户，或者忘记了原油账户的密码，那么你就需要为自己创建对应的验证人账户：
```bash
iriscli keys add {account_name}
```
得到账户信息，包括账户地址、公钥地址、助记词
```
NAME: TYPE: ADDRESS:            PUBKEY:
account_name  local faa13t6jugwm5uu3h835s5d4zggkklz6rpns59keju  fap1addwnpepqdne60eyssj2plrsusd8049cs5hhhl5alcxv2xu0xmzlhphy9lyd5kpsyzu
**Important** write this seed phrase in a safe place.
It is the only way to recover your account if you ever forget your password.

witness exotic fantasy gaze brass zebra adapt guess drip quote space payment farm argue pear actress garage smile hawk bid bag screen wonder person
```

## Step 2: 初始化节点 

初始化genesis.json和config.toml等配置文件
```bash
iris init --home={path_to_your_home} --chain-id=fuxi-7000 --moniker=key-name
```
该命令会在home目录下创建相应文件

## Step 3: 执行gentx交易

执行gentx交易，并使用刚才创建的验证人账户对交易进行签名
```bash
iris gentx --name={operator_key_name} --home={path_to_your_home} --ip=node-ip
```
这个命令将把交易的结果存储在如下目录：{path_to_your_home}/config/gentx
gentx包含一个签名后的 `CreateValidator` 交易，这个交易将为验证人设置如下默认参数： 
*	delegation amount:           100iris
*	commission rate:             0.1
*	commission max rate:         0.2
*	commission max change rate:  0.01

> IP is your public IP, dont use internal IP

注意⚠️：IP需填写公网IP

生成好的交易数据存放在目录：{path_to_your_home}/config/gentx

## Step 4: 提交gentx文件

将上述提到的json文件以提交Pull Request的形式上传到`https://github.com/irisnet/testnets/tree/master/fuxi/fuxi-7000/config/gentx`目录下：





