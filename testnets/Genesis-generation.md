# Genesis文件生成流程


1. 每个希望成为验证人的参与者确保安装了对应版本的软件：iris v0.2.0
2. 执行gentx命令，获得一个node-id.json的文件
```
       iris init gen-tx --name=<your_name> --home=<path_to_home>
```
   代码示例：
       iris init gen-tx --name=alice --home=/Users/user/Documents/bianjie/fuxi/1002
       {
        "app_message": {
          "secret": "similar spread grace kite security age pig easy always prize salon clip exhibit electric art abandon"
        },
        "gen_tx_file": {
          "node_id": "3385a8e3895b169eab3024079d987602b4d2b383",
          "ip": "192.168.1.7",
          "validator": {
            "pub_key": {
              "type": "AC26791624DE60",
              "value": "RDxXckkpTc35q9xlLNXjzUAov6xMkGJlwtWg2IqAkD8="
            },
            "power": 100,
            "name": ""
          },
          "app_gen_tx": {
            "name": "alice",
            "address": "8D3B5761BC2B9048E2A7745B14E62D51C82E0B7C",
            "pub_key": {
              "type": "AC26791624DE60",
              "value": "RDxXckkpTc35q9xlLNXjzUAov6xMkGJlwtWg2IqAkD8="
            }
          }
        }
       }
  ```
   然后你可以发现在$IRISHOME/config目录下生成了一个gentx文件夹。里面存在一个gentx-<node-ID>.json文件。这个文件包含了如下信息：
   ```
       {
        "node_id": "3385a8e3895b169eab3024079d987602b4d2b383",
        "ip": "192.168.1.7",
        "validator": {
          "pub_key": {
            "type": "AC26791624DE60",
            "value": "RDxXckkpTc35q9xlLNXjzUAov6xMkGJlwtWg2IqAkD8="
          },
          "power": 100,
          "name": ""
        },
        "app_gen_tx": {
          "name": "alice",
          "address": "8D3B5761BC2B9048E2A7745B14E62D51C82E0B7C",
          "pub_key": {
            "type": "AC26791624DE60",
            "value": "RDxXckkpTc35q9xlLNXjzUAov6xMkGJlwtWg2IqAkD8="
          }
        }
       }
  ```
   validator字段对应了home/config下的节点信息

   `app_gen_tx`中说明了拥有这个节点的账户信息。这个账户的助记词就是刚刚的secret

3. 将上述提到的json文件以提交Pull Request的形式上传到`https://github.com/irisnet/testnets/master/testnets/fuxi-1002/gentx`目录下：

   注意⚠️：json文中的IP改成公网IP

4. 在完成提交后，所有人下载所有的gentx.json文件到本地

   将配置文件拷贝到默认路径下: `$IRISHOME/config/gentx`

5. 执行生成genesis文件的操作

   通过执行以下命令来生成genesis.json 和 config.toml 文件:
   ```
       iris init --gen-txs -o --chain-id=fuxi-1001 --home=<pat_to_home>
   ```




