# Nyancat 测试网奖励

## 兑换比例

1 point = 待定 IRIS

## 积分确认

Nyancat 积分汇总: <https://github.com/irisnet/testnets/issues/406>

*如果您对此有所疑问，请和我们的工作人员联系，或者在该 Issue 中进行评论*

## 提交签名

您需要提供一个IRISHub主网地址，用以接收 Nyancat 测试网奖励。

### 使用 Keybase 签名

请使用您在 Nyancat 测试网报名表单中提交的 Keybase PGP 对您接收奖励的主网地址进行签名

以下两种签名方法任选其一：

- 使用[Keybase命令行工具](https://keybase.io/docs/command_line)签名

    ```bash
    keybase pgp sign -m "您的IRISHub主网地址"
    ```

- 或者使用[在线工具](https://keybase.io/sign)完成签名

### 提交签名信息

您需要将以上签名后的内容保存于 `[github-user-name]-[keybase fingerprint].txt`

例如：`irisnetvalidator-6763B2C7947A9363.txt`

然后使用您在 Nyancat 测试网报名表单中提交的 Github 账号，以`Pull Request`的方式，提交该签名文件到[sig-files](https://github.com/irisnet/testnets/tree/master/nyancat/v0.16/reward-claims/sig-files)文件夹下.

团队收到您的 PR 后，视为您已对 Nyancat 测试网的积分和奖励完成了确认，并将尽快给您发放主网奖励。
