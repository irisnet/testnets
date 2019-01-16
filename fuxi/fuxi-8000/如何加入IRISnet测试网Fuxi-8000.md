### 如何加入IRISnet测试网 : Fuxi-8000

1月16日，IRIShub开发团队发布了最新版本：v0.10.0。 这个版本中对Tendermint依赖也升级到 v0.27.3。以下是该版本的**重大改动**：

- 在解除绑定的同时也取回所有的佣金
- 通过地址作为验证人的index
- 修复在抵押有变化时，验证人被动取回其奖励的问题
- 修复LCD中证明的确实问题
- 修复查询tx返回tags的问题
- 修复未检查账户中第一个Coin的问题
- 修复export-for-zero-height中未重置jailed-validator的信息的问题
- 如果一个验证人未jailed状态，他就收不到抵押奖励
- 不再导出空账户

**新特性**

- 升级Tendermint至v0.27.3
- 增加mint模块，实现抵押获利按块分发的模式
- 在service中增加惩罚机制
- 重新设计基于链上治理的软件升级流程
- 新增gov/slashing/service/stake/distribution模块参数
- 在iriscli客户端默认以broadcast方式广播交易
- 增加代币总量信息查询接口
- 增加代币销毁burn功能接口
- 删掉Record模块
- 删除set-withdraw-addr子命令

### Fuxi-8000的重要意义

从IRIShub v0.9.0开始，IRIS开发团队不断反思，并重新构思一种更可靠的软件升级流程。由于在升级过程中遭遇共识失败，fuxi-7000测试网停滞。
具体原因分析：<https://github.com/irisnet/irishub/issues/999>。

作为IRISnet Betanet上线前倒数第二个测试网，为了模拟Betanet上线后的实际情况，开发团队计划在测试网中模拟一系列的软件升级流程，并且在网络中仿真大量交易。Fuxi-8000将测试以下重要特性：

- 不同类型的软件升级流程
- 验证人节点是否能够在高吞吐量下正常工作

软件升级类型分析：

**自愿升级**: 此类升级将引入状态机的改进。验证人可根据实际情况决定是否升级。升级的比例将不影响共识。

**平滑升级**:此类升级在线软件平滑升级。在对应的软件升级提案中将指定升级高度。Tendermint 在此高度前将不断收集区块头部中的版本信息。如果Tendermint认为指定比例的 voting power已经升级到新版本, 那么该升级提案被成功执行，新功能会被激活。该升级方案类似Bitcoin MASF升级方式

**补丁升级**:当区块链系统遇到可恢复的共识失败问题，例如appHash冲突，验证人可以通过安装新版本软件来恢复网络共识。

**重启升级**:当区块链系统遇到不可恢复的共识失败问题。我们只能通过尝试导出当前网络的状态快照，然后将其作为新的区块链的初始状态来恢复网络。这类升级方式将放弃原有的历史信息。在测试网中，我们可以通过发起一个`SystemHalt`提案，先停止网络，再导出快照

想要了解更多有关链上治理的，请阅读以下[文章](https://www.irisnet.org/docs/features/governance.html)。

![img](https://cdn-images-1.medium.com/max/800/1*FqrVw8a_4UP4S2y39dc1og.png)Fuxi-8000升级流程

### 参与到测试网Genesis文件生成中

> 完成文档在[这里](https://github.com/irisnet/testnets/blob/master/fuxi/docs_CN/Genesis-generation.md)

- 如果你参与了Fuxi-6000测试网，请不要继续使用同一个home目录，无需停止已有**iris**进程
- 参考以下教程安装 **IRIShub v0.10.0**
- 你可以使用已有的账户作为验证人Operator对应的账户，也可以创建新的账户：

```
iriscli keys add {account_name}
```

输入密码后，你可以得到以下账户信息，包括账户地址、账户的公钥地址、助记词。请注意账户地址和账户的公钥地址都是采用`Bech32`编码的。

```
NAME:	TYPE:	ADDRESS:						PUBKEY:
account_name	local	faa13t6jugwm5uu3h835s5d4zggkklz6rpns59keju	fap1addwnpepqdne60eyssj2plrsusd8049cs5hhhl5alcxv2xu0xmzlhphy9lyd5kpsyzu
**Important** write this seed phrase in a safe place.
It is the only way to recover your account if you ever forget your password.
```

```
witness exotic fantasy gaze brass zebra adapt guess drip quote space payment farm argue pear actress garage smile hawk bid bag screen wonder person
```

- 初始化节点

```
iris init --home={path_to_your_home} --chain-id=fuxi-8000 --moniker={validator-name}
```

- 执行Gentx命令

该命令会在home/config/gentx目录下创建相应一个json文件。该文件包含了一个签名的创建验证人交易，收集这些交易将可以生成最终的genesis文件。默认参数如下：

- delegation amount: 100iris
- commission rate: 0.1
- commission max rate: 0.2
- commission max change rate: 0.01

```
iris gentx --name={key_name} --home={path_to_your_home} --ip={node-public-ip}
```

- 提交gentx文件

通过提交PR将gentx文件提交到以下目录：[https://github.com/irisnet/testnets/tree/master/fuxi/fuxi-8000/config/gentx](https://github.com/irisnet/testnets/tree/master/fuxi/fuxi-7000/config/gentx)

**Deadline: 1/16 at 8:00 UTC**

### 测试网激励任务

作为一个共有区块链网络，IRISnet网络的价值需要一组可靠的验证人来维护。Fuxi测试网的目的在于让验证人模拟真实上线后的情况，为将来完成好验证人的工作做好准备。 测试网的奖励将在主网上线后发放给社区成员。你需要用你的Keybase签名一个irishub的地址，然后将其发送给团队。

- Fuxi-8000 任务列表:

[https://github.com/irisnet/testnets/blob/master/fuxi/fuxi-8000/README_CN.md](https://github.com/irisnet/testnets/blob/master/fuxi/fuxi-7000/README_CN.md)

你可以在以下链接中提交完成证明：https://github.com/irisnet/testnets/issues/252.

Fuxi-7000任务完成情况汇总在[这里](https://github.com/irisnet/testnets/issues/251)。