# Fuxi测试网

作为一个共有区块链网络，IRISnet网络的价值需要一组可靠的验证人来维护。Fuxi测试网的目的在于让验证人模拟真实上线后的情况，为将来完成好验证人的工作做好准备。
测试网的奖励将在主网上线后发放给社区成员。你需要用你的Keybase签名一个irishub的地址，然后将其发送给团队。

## Fuxi-4000测试网激励任务


| **编号** | **名称**                                           | **详情**                                                     | **证明**                                                     | **积分** |
| -------- | -------------------------------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ | -------- |
| 1        | 参与到Genesis文件的生成中                          | 提交gen-tx.json文件， 在验证人name字段中注明为name-pgp-fingerprint | 提供提交PR的链接                                             | 100      |
| 2        | 运行一个全节点                                     | 运行一个全节点，将该节点的monkier设置为name-pgp-fingerprint  | 提供节点的IP，并保证相应端口可访问                           | 100      |
| 3        | 从水龙头获得一些iris，然后将它委托给某个验证人     | 从水龙头获得一定的iris代币，然后将执行delegate操作将这部分代币委托给某个验证人 | 完成delegate操作，在交易的memo中注明 name-pgpid，便于团队确认任务完成情况 | 50       |
| 4        | 完成委托后通过Redelegate命令将iris委托给其他验证人 | 完成redelegate 交易                                          | 完成 redelegate begin 和 redelegat complete 交易，并在交易的memo中注明 name-pgpid，便于团队确认任务完成情况 | 50       |
| 5        | 参与到链上治理通过软件升级提案的过程中             | 完成software-upgrade操作，将软件升级到v0.6.1                 | 根据以下文档完成情况软件升级                                 | 200      |
| 6        | 对某个提案投票                                     | 完成vote交易                                                 | 完成vote操作，在交易的memo中注明 name-pgpid，便于团队确认任务完成情况 | 20       |
| 7        | 对某个提案添加押金                                 | 完成deposit交易                                              | 完成deposit操作，在交易的memo中注明 name-pgpid，便于团队确认任务完成情况 | 20       |


### 如何提交完成证明

请在以下 [issue](https://github.com/irisnet/testnets/issues/129)下提交完成证明，格式如下：

```
GitHub ID: XXXX
pgp ID: XXX
Task 1: Link to your PR
Task 2: Node URL
Task 3: Tx Hash
Task 4: Tx Hash
Task 5: Proof that you stay as a validator in newer version
Task 6: Tx Hash
Task 7: Tx Hash
...

```
所有完成任务的人都能获得积分奖励！
