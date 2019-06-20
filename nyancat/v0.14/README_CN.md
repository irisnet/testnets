# Nyancat 测试网 V0.14

## 最新版本: [v0.14.1](https://github.com/irisnet/irishub/releases/tag/v0.14.1) [[更新日志](https://github.com/irisnet/irishub/blob/master/CHANGELOG.md#0141)]

## 任务列表

### v0.14.2

我们在 Nyancat 测试网发现了一个小 Bug：`genesis.json` 中 `Gov` 配置相关的 `min_deposit` 没有在系统中做`iris`单位转换，导致 Nyancat 测试网发起提议所需要的抵押金额非常低（配置为 100 iris，而系统则初始化为 100 iris-atto，详见：[genesis.json 第90行](../config/genesis.json#L90)）。不过这个问题在主网是不存在的，因为主网版本的二进制文件在启动时对 genesis 参数都有严格的校验，测试网版本为了方便测试，所以关闭了校验，从而导致这个问题没有被及时发现。于是我们做了如下优化：

- 修复 genesis 对 `iris` 单位的转换
- 开启测试网版本二进制文件对 genesis 参数的校验

感谢这个 Bug 的出现，我们可以在 v0.15 发布前做一些家庭作业了！

我们都知道 IRIS Hub 具有完备的链上[升级](https://www.irisnet.org/docs/zh/features/upgrade.html)和[治理](https://www.irisnet.org/docs/zh/features/governance.html)功能，不过出于安全考量，链上治理无法治理自己，所以对 `Gov` 参数的修改，只能通过 `Class 4` 升级来完成。

这次升级分为如下几个阶段：

#### 投票阶段

由 [Profiler](./README_CN.md "唯一有权限发起Critical级别提议的账户，定义在 genesis.json 中") 发起 `SystemHalt` 提议，指定在未来的某一高度停止出块，验证人参与对该提议的投票

#### 投票阶段任务

| No   | Name                                           | Details                                                      | Criteria                                                     | Points |
| ---- | ---------------------------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ | ------ |
| 1    | 对 `SystemHalt` 提议进行投票 | 使用你的Validator Operator账户对 `SystemHalt` 提议进行投票，强烈建议投出赞成票，以便接下来的任务的顺利执行  | 完成投票交易,并将 `<name>-<pgpid>` 作为 memo，[参考文档](https://www.irisnet.org/docs/zh/cli-client/gov/vote.html)      | 5    |

#### 升级阶段

当 `SystemHalt` 提议通过并到达指定高度后，Nyancat 测试网将会自动进入 Halt 状态无法继续出块，即升级阶段开始。

升级步骤：

1. 停止运行中的节点，导出节点状态

    ```bash
    iris export --for-zero-height --height 0 --output-file nyancat_export.json
    ```

2. 修改新的网络的启动时间（待定）

    ```bash
    sed -i 's/"genesis_time": "2019-06-04T03:19:19.48290918Z"/"genesis_time": "待定"/g' nyancat_export.json | jq -S -c -M '' > new_genesis.json
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

#### 升级阶段任务

| No   | Name                                           | Details                                                      | Criteria                                                     | Points |
| ---- | ---------------------------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ | ------ |
| 2    | 升级节点版本至 v0.14.2 | 按如上步骤完成升级  | 在新的 Nyancat 网络中前10个区块高度内存在有效 Pre Commit 的验证人节点将会获得奖励| 20    |

### v0.14.1

作为 Nyancat 测试网的初始版本，v0.14.1 已经稳定的运行于 IRIS Hub 主网，所以我们没有很多任务需要完成，我们只需要为新进验证人们启动一个测试网，用于熟悉验证人相关的各种操作，避免在主网上面犯错，同时为即将在 v0.15 版本发布的测试网任务做好准备。

| No   | Name                                           | Details                                                      | Criteria                                                     | Points |
| ---- | ---------------------------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ | ------ |
| 1    | 创建一个验证人节点                                | 创建一个全节点，并以 `<name>-<pgp-fingerprint>` 作为 `monkier`, 然后升级该节点为验证人节点 | 提交你的公网IP使我们可以校验你的节点状态      | N/A    |
