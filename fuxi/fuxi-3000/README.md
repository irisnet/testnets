# How to Upgrade to Fuxi-3001

Fuxi-3000 is short lived and validator FDC8E91E2F361F4BFCCEFA80C1ABAD138E5175AE double-sign genesis block triggeres a bug in slashing module. More analysis is here: https://github.com/irisnet/cosmos-sdk/pull/109
S
The team released a hot fix that should fix the problem: https://github.com/irisnet/irishub/releases/tag/v0.4.
Please follow the instructions to reset your node and join Fuxi-3001!

1. Reset your node
```
iris unsafe_reset_all --home
rm genesis.json
rm addrbook.json
```

2. Install latest irishub
```
cd $GOPATH/src/github.com/irisnet/irishub
git fetch -a origin
git checkout v0.4.2
make get_vendor_deps
make install
```

3.Download genesis file for fuxi-3001
```
wget https://raw.githubusercontent.com/irisnet/testnets/master/fuxi/fuxi-3000/config/genesis.json

```

4. Start 