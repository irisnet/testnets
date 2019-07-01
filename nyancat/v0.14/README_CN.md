# Nyancat 测试网 V0.14

## 最新版本: [v0.14.3](https://github.com/irisnet/irishub/releases/tag/v0.14.3)

**注**：积分兑换比例会在测试网活动结束时公布

## v0.14.3

### 投票任务

| No   | Name                                           | Details                                                      | Criteria                                                     | Points |
| ---- | ---------------------------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ | ------ |
| 3    | 参与链上参数治理 | `Bianjie` 将会发起一个修改 `Slashing` 参数的提议，请大家对该提议进行投票  | 以 `nyancat-2` 链上对[提议1](https://nyancat.irisplorer.io/#/ProposalsDetail/1)的有效投票为准| 50    |

这次参数修改的提议将会是主网参数修改的演练，大家可以通过本次演练了解投票的的基本流程，以及参数修改生效后对验证人节点的影响。对于参数修改有不同的意见，也欢迎在论坛中发表：<https://forum.irisnet.org/t/parameter-changes-raising-the-difficulty-level-a-little-bit-for-validators/127>

投票参考文档：<https://www.irisnet.org/docs/cli-client/gov/vote.html>

### ~~升级任务（已结束）~~

| No   | Name                                           | Details                                                      | Criteria                                                     | Points |
| ---- | ---------------------------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ | ------ |
| 2    | 升级节点版本至 v0.14.3 | 按如上步骤完成升级  | 在新的 Nyancat 网络中前100个区块高度内，存在有效 Pre Commit 的验证人节点将会获得奖励| 200    |

这个新版本是为了解决`v0.14.1`中的一个缺点，当治理参数中的 Token 单位设置错误时，会导致链上治理在测试网络中无法使用。
我们尝试在 Nyancat 执行链上治理时发现了这个问题。

### v0.14.1 Bug 详情

[Testnet Genesis 参数校验](https://github.com/irisnet/irishub/blob/v0.14.1/modules/gov/params.go#L362)

[Nyancat Genesis 配置](../config/genesis.json#L90)

从上面的代码中我们可以看到，对于`测试网`版本的软件，为了方便测试，在加载 Genesis 配置时跳过了验证，同时，也没有对 `Gov` 配置相关的 `min_deposit` 的 denom 进行转换，从而触发了这个Bug。不过这个问题在主网是不存在的，因为主网版本的软件在启动时对 genesis 参数进行了严格的校验。

为了解决这个问题，我们在`v0.14.3` 中做了如下改进：

- 开启测试网版本软件对 genesis 参数的校验

感谢这个 Bug 的出现，我们可以在 v0.15 发布前做一些家庭作业了！

我们都知道 IRIS Hub 具有完备的[软件升级](https://www.irisnet.org/docs/zh/features/upgrade.html)和[治理](https://www.irisnet.org/docs/zh/features/governance.html)功能，不过出于安全考量，我们有目的地选择不允许[治理](https://www.irisnet.org/docs/zh/features/governance.html)本身受到管理。因此，要更新错误设置的治理参数，我们必须通过硬分叉。

对于 Nyancat testnet，我们将按以下步骤执行fork：

1. 我们会将“Bianjie”节点的投票权提高到2/3以上

2. 我们将把“Bianjie”节点下线以停止 Nyancat 网络，我们会在这之前做出公告

3. 导出高度为 `287373` 的区块链状态

    ```bash
    iris export --for-zero-height --height 287373 --output-file nyancat_export.json
    ```

4. 修改新的网络的参数及启动时间

    你需要安装 [jq](https://stedolan.github.io/jq/) 来执行以下命令，[[在线测试](https://jqplay.org/s/jTO8nHeCGb)]

    ```bash
    jq -S -c -M '.genesis_time="2019-06-26T13:00:00.000000000Z"|.chain_id="nyancat-2"|.app_state.gov.params = (.app_state.gov.params | .critical_min_deposit[0] = {"denom": "iris-atto", "amount": "100000000000000000000"}|.important_min_deposit[0] = {"denom": "iris-atto", "amount": "100000000000000000000"}|.normal_min_deposit[0] = {"denom": "iris-atto", "amount": "50000000000000000000"})' nyancat_export.json > new_genesis.json
    ```

    校验 `new_genesis.json` 的 SHA256

    ```bash
    sha256sum new_genesis.json
    0ad7da96167c22ba7acb505bbfdb6a83de1e87cf6be2f7013e6cfb766a0e9953  new_genesis.json
    ```

5. 使用 `new_genesis.json` 覆盖原 `genesis.json`

    ```bash
    cp new_genesis.json <your_home>/config/genesis.json
    ```

6. 编译安装新版本 v0.14.3

    ```bash
    git clone https://github.com/irisnet/irishub
    cd irishub && git checkout v0.14.3
    source scripts/setTestEnv.sh
    make get_tools && make install
    ```

    校验版本信息

    ```bash
    iris version
    0.14.3-c9616441-0
    ```

7. 重置原 Nyancat 测试网数据

    ```bash
    iris unsafe-reset-all --home=<your_home>
    ```

8. 启动节点

    ```bash
    iris start --home=<your_home>
    ```

## v0.14.1

作为 Nyancat 测试网的初始版本，v0.14.1 已经稳定的运行于 IRIS Hub 主网，所以我们没有很多任务需要完成，我们只需要为新进验证人们启动一个测试网，用于熟悉验证人相关的各种操作，避免在主网上面犯错，同时为即将在 v0.15 版本发布的测试网任务做好准备。

| No   | Name                                           | Details                                                      | Criteria                                                     | Points |
| ---- | ---------------------------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ | ------ |
| 1    | 创建一个验证人节点                                | 创建一个全节点，并以 `<name>-<pgp-fingerprint>` 作为 `monkier`, 然后升级该节点为验证人节点 | 提交你的公网IP使我们可以校验你的节点状态      | N/A    |
