# Nyancat 测试网 V0.14

## 最新版本: [v0.14.2](https://github.com/irisnet/irishub/releases/tag/v0.14.2)

## v0.14.2

我们在 Nyancat 测试网 `v0.14.1` 版本发现了一个 Bug：对于链上的 Proposal，我们需要质押一定的金额才能进入投票阶段，`genesis.json` 中 `Gov` 配置相关的 `min_deposit` 没有在系统中做`iris`单位转换，导致 Nyancat 测试网发起提议所需要的质押金额为 100 `iris`，而系统中其实并不存在 `iris` 这个 Token，只有 `iris-atto`，我们在 UI 层所看到的 `iris` 其实是通过公式转换得来的：1 iris = 10^18 iris-atto。也就是说我们只能质押`iris-atto`，而系统却只认系统中并不存在的 `iris`，导致质押金额永远无法满足要求，最终提议失败。

### Bug 详情

[Testnet Genesis 参数校验](https://github.com/irisnet/irishub/blob/v0.14.1/modules/gov/params.go#L362)

[Nyancat Genesis 配置](../config/genesis.json#L90)

从上面的代码中我们可以看到，对于`测试网`版本的软件，为了方便测试，在加载 Genesis 配置时跳过了验证，同时，也没有对 `Gov` 配置相关的 `min_deposit` 单位进行转换，从而触发了这个Bug。不过这个问题在主网是不存在的，因为主网版本的软件在启动时对 genesis 参数进行了严格的校验。

于是我们做了如下优化：

- 开启测试网版本软件对 genesis 参数的校验
- Genesis 中只接受 `iris-atto` 作为 `denom`（account balance 除外）

### 升级任务

感谢这个 Bug 的出现，我们可以在 v0.15 发布前做一些家庭作业了！

我们都知道 IRIS Hub 具有完备的链上[升级](https://www.irisnet.org/docs/zh/features/upgrade.html)和[治理](https://www.irisnet.org/docs/zh/features/governance.html)功能，不过出于安全考量，链上治理无法治理自己，所以对 `Gov` 参数的修改，只能通过 `Class 4` 升级来完成。并且由于 Bug 刚好出现在链上治理模块，我们甚至无法通过 `SystemHalt Proposal` 使我们的 Nyancat 优雅的停止，所以对于这次升级，我们只能采用传统方式进行升级。

首先，我们需要确定一个具体的区块高度，用来导出当前区块链的最终状态，当区块到达目标高度时，验证人即可按如下步骤操作：

1. 停止运行中的节点，导出指定高度的区块链状态

    ```bash
    iris export --for-zero-height --height <待定> --output-file nyancat_export.json
    ```

2. 修改新的网络的参数及启动时间（待定）

    你需要安装 [jq](https://stedolan.github.io/jq/) 来执行以下命令，[[在线测试](https://jqplay.org/s/9QSR4xq_TX)]

    ```bash
    jq -S -c -M '.genesis_time="待定" |.app_state.gov.params = (.app_state.gov.params | .critical_min_deposit[0] = {"denom": "iris-atto", "amount": "100000000000000000000"}|.important_min_deposit[0] = {"denom": "iris-atto", "amount": "100000000000000000000"}|.normal_min_deposit[0] = {"denom": "iris-atto", "amount": "50000000000000000000"})' nyancat_export.json > new_genesis.json
    ```

    校验 `new_genesis.json` 的 SHA256

    ```bash
    sha256sum new_genesis.json
    xxx(待定)
    ```

3. 使用 `new_genesis.json` 覆盖原 `genesis.json`

    ```bash
    cp new_genesis.json <your_home>/config/genesis.json
    ```

4. 安装新版本

    ```bash
    git clone https://github.com/irisnet/irishub
    cd irishub && git checkout v0.14.2
    source scripts/setTestEnv.sh
    make get_tools && make install
    ```

    校验版本信息

    ```bash
    iris version
    0.14.2-1241d9d-0
    ```

5. 重置原 Nyancat 测试网数据

    ```bash
    iris unsafe-reset-all --home=<your_home>
    ```

6. 启动节点

    ```bash
    iris start --home=<your_home>
    ```

| No   | Name                                           | Details                                                      | Criteria                                                     | Points |
| ---- | ---------------------------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ | ------ |
| 2    | 升级节点版本至 v0.14.2 | 按如上步骤完成升级  | 在新的 Nyancat 网络中前50个区块高度内，存在有效 Pre Commit 的验证人节点将会获得奖励| 200    |

## v0.14.1

作为 Nyancat 测试网的初始版本，v0.14.1 已经稳定的运行于 IRIS Hub 主网，所以我们没有很多任务需要完成，我们只需要为新进验证人们启动一个测试网，用于熟悉验证人相关的各种操作，避免在主网上面犯错，同时为即将在 v0.15 版本发布的测试网任务做好准备。

| No   | Name                                           | Details                                                      | Criteria                                                     | Points |
| ---- | ---------------------------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ | ------ |
| 1    | 创建一个验证人节点                                | 创建一个全节点，并以 `<name>-<pgp-fingerprint>` 作为 `monkier`, 然后升级该节点为验证人节点 | 提交你的公网IP使我们可以校验你的节点状态      | N/A    |
