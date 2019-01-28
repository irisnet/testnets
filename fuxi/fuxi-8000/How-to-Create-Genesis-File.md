# How to Create Genesis File

1. Backup your `irishome`
2. Export Snapshot of height #157491 from the halted network
  Using iris `v0.10.2-patch`, run the following command:
```
iris export  --home= --for-zero-height
```
This command should export a new genesis file in your current path.
Md5sum should be: `3d140361770cd4f18618f1629eead49e`
2. Modify fileds in genesis.json
* Add `tx_size` param in `auth`
```json
"params": {
        "gas_price_threshold": "20000000000000",
        "tx_size": "1000"
      }
```
* Modify `Upgrade` field:
```json
"GenesisVersion": {
        "Success": true,
        "UpgradeInfo": {
          "ProposalID": "0",
          "Protocol": {
            "height": "1",
            "software": "https://github.com/irisnet/irishub/releases/tag/v0.11.0",
            "threshold": "0.9000000000",
            "version": "0"
          }
        }
      }
    }
```
* Add  `slash-fraction-censorship`and `censorship-jail-duration` params in `slashing`
```
"params": {
        "censorship-jail-duration": "604800000000000",
        "double-sign-unbond-duration": "432000000000000",
        "downtime-unbond-duration": "172800000000000",
        "max-evidence-age": "86400000000000",
        "min-signed-per-window": "0.5000000000",
        "signed-blocks-window": "20000",
        "slash-fraction-censorship": "0.0200000000",
        "slash-fraction-double-sign": "0.0100000000",
        "slash-fraction-downtime": "0.0050000000"
      }
```
*  Add `tx_size` param in `service`
```
"tx_size_limit": "4000"
```
Md5sum for new genesis file is: 
3d140361770cd4f18618f1629eead49e
Example output is [here](https://raw.githubusercontent.com/irisnet/testnets/master/fuxi/fuxi-8000/config/genesis.json)

3. Start with new genesis file and iris version v0.11.0 
* Verify `iris version`
```
iris version
0.11.0-ae1e491-0
```

* Clear previous data
```
iris unsafe-reset-all --home
```
* Start iris
```
iris start --home
```

