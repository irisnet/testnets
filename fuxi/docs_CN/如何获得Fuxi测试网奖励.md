# 如何获得Fuxi测试网奖励

## 重要信息

1 point = 100 IRIS

作为一个公有区块链网络，IRISnet网络的价值需要一组可靠的验证人来维护。Fuxi测试网的目的在于让验证人模拟主网上线后的真是情况，为将来完成好验证人的工作做好准备。
Fuxi-8000是IRISnet Betanet上线前的最后一个测试网，团队在该测试网中测试了复杂的升级计划，也设计了更多的任务 ：list:https://github.com/irisnet/testnets/blob/master/fuxi/fuxi-8000/README.md.
验证人社区也积极参与了测试网的工作，此次提交任务完成的人数也创历史新高： https://github.com/irisnet/testnets/issues/252 .
1. 确认测试网积分奖励:

经过统计，社区一共获得了57820个积分奖励，平均每个人获得1100个积分。测试网积分奖励和iris的兑换比例为1 point = 100 IRIS。测试网奖励积分汇总如下：
https://github.com/irisnet/testnets/issues/290
2. 创建PGP ID

发放给社区成员的iris代币奖励将写入Beta网络的genesis文件中。为了保障奖励发放的准确性，社区成员需要用GitHub账号认证的keybase签名一个irishub的账户地址，然后将其发送给团队。具体教程如下： https://github.com/irisnet/testnets/blob/master/fuxi/%E5%A6%82%E4%BD%95%E4%BD%BF%E7%94%A8keybase.md
3. 生成IRIShub账户地址：

安装IRIShub v0.11.0:https://www.irisnet.org/docs/get-started/Install-the-Software.html
生成你的账户: https://www.irisnet.org/docs/cli-client/keys/add.html
使用keybase命令行工具https://keybase.io/docs/command_lin，e对该地址签名，例如如果你的地址是: faa1mmsm487rqkgktl2qgrjad0z3yaf9n8t5pkp33m
```
keybase pgp sign -m "faa1mmsm487rqkgktl2qgrjad0z3yaf9n8t5pkp33m"
```
4. 将返回信息填写到以下表格中

https://goo.gl/forms/j8jP6Gc6zDVXMYkE3

5. 在团队公布genesis文件后确认账户余额是否与奖励一致。公布的日期大约在2月13日。
