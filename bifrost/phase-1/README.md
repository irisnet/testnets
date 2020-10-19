# Bifrost-1

## How to Join

It'll be the same as joining [IRIS Hub Mainnet](https://bifrost.irisnet.org/docs/get-started/mainnet.html), except that:

- Go 1.14.1+ is required for building and installing the IRIShub software.
- [version](https://github.com/irisnet/irishub/releases/tag/v1.0.0-beta)
- [genesis file](./genesis.json)

Please don't forget fill in this [form](http://nyancat-irisnet.mikecrm.com/JsLCZ4H) to complete your sign-up process, and remember to set your validator moniker to `<your-name>-<pgp-id>`, identity to `<pgp-id>`. You need to complete the tasks use your validator account.

## Peers

Persistent Peers:

```bash
a246f5ea3055f7507c63b9ea50dbf5d401537b3a@34.80.22.255:26656
```

## Faucet

Unfortunately we do not have a public faucet for bifrost, however you are welcome to ask for the test tokens in our validator communities:

- [Riot English community](https://matrix.to/#/!bmimZgJrUWSmxqQEmG:matrix.org?via=matrix.org&via=t2bot.io)
- [QQ Chinese community](https://jq.qq.com/?_wv=1027&k=5BeP3tJ)

## Explorer

https://bifrost.irisplorer.io

## Task list and reward

### Service task list

| Task                                        | Details                                                              | Completion standard                         | Badge     |
| ------------------------------------------- | -------------------------------------------------------------------- | ------------------------------------------- | --------- |
| Create service definition                   | Create a service definition                                          | Successfully created the service definition | Silver *1 |
| Bind service definition created by yourself | Bind the service definition you created                              | Successfully created the service binding    | Bronze *1 |
| Bind service definition created by others   | Bind the service definition others created                           | Successfully created the service binding    | Bronze *1 |
| Call service once                           | Request a bound service once                                         | Successfully initiated the request          | Bronze *1 |
| Call service repeatedly                     | Request a bound service repeatedly                                   | Successfully created request-context        | Bronze *1 |
| Respond service request                     | Respond a service you bound                                          | Successfully responded the service request  | Gold *1   |
| Set withdrawal address and withdraw fees    | Set the fees withdrawal address and withdraw the earned service fees | Successfully withdrew fees                  | Silver *1 |

### Record & Oracle task list

Each participant’s address will be assigned a serial number, which corresponds to the value of the feedname specified in oracle

| Task                                          | Details                                                                                                                                          | Completion standard                    | Badge     |
| --------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------ | -------------------------------------- | --------- |
| Query oracle response value and create record | According to the assigned serial number, read the value at the specified position from Oracle, and create a record with the value as the content | Rrecord content meets the requirements | Silver *1 |

### NFT task list

| Task         | Details                                                              | Completion standard                   | Badge     |
| ------------ | -------------------------------------------------------------------- | ------------------------------------- | --------- |
| Issue NFT    | Issue an NFT with custom denom                                       | Successfully issued the NFT           | Silver *1 |
| Mint NFT     | Mint the NFT you issued, create at least two tokens                  | Successfully minted the NFT tokens    | Bronze *1 |
| Edit NFT     | Edit the first token's uri to `bifrost-[denom]`                      | Successfully edited the NFT token     | Bronze *1 |
| Transfer NFT | Create a new account, then transfer the second token to this account | Successfully trasnfered the NFT token | Bronze *1 |
| Burn NFT     | Burn the second token                                                | Successfully burnt the NFT token      | Bronze *1 |

### Random task list

| Task                              | Details                         | Completion standard            | Badge     |
| --------------------------------- | ------------------------------- | ------------------------------ | --------- |
| Request a random (without oracle) | Request a random without oracle | Successfully get random number | Bronze *1 |
| Request a random (with oracle)    | Request a random with oracle    | Successfully get random number | Bronze *1 |

### Reference documents

https://bifrost.irisnet.org/docs/
