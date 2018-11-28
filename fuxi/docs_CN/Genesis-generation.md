# Genesis文件生成流程

## 前提条件

每个希望成为验证人的参与者确保按照以下[流程](https://www.irisnet.org/docs/zh/get-started/Join-the-Testnet.html#如何加入fuxi测试网)安装了对应版本的软件：**iris**


## Step 1: 创建账户  


首先需要为自己创建对应的验证人账户
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
iris init --home={path_to_your_home} --chain-id={your_chain_id}
```
该命令会在home目录下创建相应文件

## Step 3: 执行gentx交易

执行gentx交易，并使用刚才创建的验证人账户对交易进行签名
```bash
iris gentx --name={account_name} --home={path_to_your_home} --IP
```

注意⚠️：IP需填写公网IP

生成好的交易数据存放在目录：{path_to_your_home}/config/gentx

## Step 4: 提交gentx文件

将上述提到的json文件以提交Pull Request的形式上传到`https://github.com/irisnet/testnets/tree/master/testnets/fuxi-5000/gentx`目录下：





