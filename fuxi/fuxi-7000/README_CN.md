# Fuxi-7000测试网激励任务

作为一个共有区块链网络，IRISnet网络的价值需要一组可靠的验证人来维护。Fuxi测试网的目的在于让验证人模拟真实上线后的情况，为将来完成好验证人的工作做好准备。
测试网的奖励将在主网上线后发放给社区成员。你需要用你的Keybase签名一个irishub的地址，然后将其发送给团队。

## Fuxi-7000测试网激励任务


| **编号** | **名称**                                           | **详情**                                                     | **证明**                                                     | **积分** |
| -------- | -------------------------------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ | -------- |
| 1        | 参与到Genesis文件的生成中                          | 提交gen-tx.json文件， 在验证人name字段中注明为name-pgp-fingerprint | 提供提交PR的链接                                             | 100      |
| 2        | 运行一个全节点                                     | 运行一个全节点，将该节点的monkier设置为name-pgp-fingerprint  | 提供节点的IP，并保证相应端口可访问                           | 100      |
| 3        | 从水龙头获得一些iris，然后将它委托给某个验证人     | 从水龙头获得一定的iris代币，然后将执行delegate操作将这部分代币委托给某个验证人 | 完成delegate操作，在交易的memo中注明 name-pgpid，便于团队确认任务完成情况 | 50       |
| 4    | 参与到链上治理通过软件升级提案的过程中                               | Complete an `vote `transaction                               | 完成vote交易,并将 `name-pgpid` 作为 memo. 提交交易哈希值作为验证。 可以参考以下[文档](https://www.irisnet.org/docs/cli-client/gov/vote.html) 了解如何完成投票  | 50     |
| 5    | 参与到软件升级过程中，你需要在软件升级提案通过后，将irishub软件升级到v0.9.1。| 参与软件升级过程 | 根据以下 [文档](https://www.irisnet.org/docs/features/upgrade.html#interaction-process) 完成节点的升级| 200    |




### 如何提交完成证明

请在以下 [issue](https://github.com/irisnet/testnets/issues/211)下提交完成证明，格式如下：

```
GitHub ID: XXXX
pgp ID: XXX
Task 1: Link to your PR
Task 2: Node URL
Task 3: Tx Hash
Task 4: Tx Hash
Task 5: 在升级完成后，浏览器中查看你的验证人节点信息的链接
...

```
只有前40名完成任务的人都能获得积分奖励！
