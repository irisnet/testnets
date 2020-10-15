# Bifrost-1

## Task list and reward

### Service task list

| Task                                        | Details                                                              | Completion standard                         | Materials need to be submitted | Badge     |
| ------------------------------------------- | -------------------------------------------------------------------- | ------------------------------------------- | ------------------------------ | --------- |
| Create service definition                   | Create a service definition                                          | Successfully created the service definition | service-name                   | Silver x1 |
| Bind service definition created by yourself | Bind the service definition you created                              | Successfully created the service binding    | service-name                   | Bronze x1 |
| Bind service definition created by others   | Bind the service definition others created                           | Successfully created the service binding    | service-name                   | Bronze x1 |
| Call service once                           | Request a bound service once                                         | Successfully initiated the request          | request-id                     | Bronze x1 |
| Call service repeatedly                     | Request a bound service repeatedly                                   | Successfully created request-context        | request-context-id             | Bronze x1 |
| Respond service request                     | Respond a service you bound                                          | Successfully responded the service request  | request-id                     | Gold x1   |
| Set withdrawal address and withdraw fees    | Set the fees withdrawal address and withdraw the earned service fees | Successfully withdrew fees                  | two tx hash                    | Silver x1 |

### Record & Oracle task list

Each participant’s address will be assigned a serial number, which corresponds to the value of the feedname specified in oracle

| Task                                          | Details                                                                                                                                          | Completion standard                    | Materials need to be submitted | Badge     |
| --------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------ | -------------------------------------- | ------------------------------ | --------- |
| Query oracle response value and create record | According to the assigned serial number, read the value at the specified position from Oracle, and create a record with the value as the content | Rrecord content meets the requirements | record-id                      | Silver x1 |

### NFT task list

| Task         | Details                                                              | Completion standard                   | Materials need to be submitted | Badge     |
| ------------ | -------------------------------------------------------------------- | ------------------------------------- | ------------------------------ | --------- |
| Issue NFT    | Issue an NFT with custom denom                                       | Successfully issued the NFT           | denom                          | Silver x1 |
| Mint NFT     | Mint the NFT you issued, create at least two tokens                  | Successfully minted the NFT tokens    | two token-id                   | Bronze x1 |
| Edit NFT     | Edit the first token's uri to `bifrost-[denom]`                      | Successfully edited the NFT token     | tx hash                        | Bronze x1 |
| Transfer NFT | Create a new account, then transfer the second token to this account | Successfully trasnfered the NFT token | tx hash                        | Bronze x1 |
| Burn NFT     | Burn the second token                                                | Successfully burnt the NFT token      | tx hash                        | Bronze x1 |

### Random task list

| Task                              | Details                         | Completion standard            | Materials need to be submitted | Badge     |
| --------------------------------- | ------------------------------- | ------------------------------ | ------------------------------ | --------- |
| Request a random (without oracle) | Request a random without oracle | Successfully get random number | request-id                     | Bronze x1 |
| Request a random (with oracle)    | Request a random with oracle    | Successfully get random number | request-id                     | Bronze x1 |

### Reference documents

https://bifrost.irisnet.org/docs/
