# How To Get Prepared

## Install Go

Please follow the following instructions to get Go1.10 installed on your server.

```
$ sudo add-apt-repository ppa:gophers/archive
$ sudo apt-get update
$ sudo apt-get install golang-1.10-go
```

Then, you need to add the GOPATH to system path:

```
echo "export PATH=$PATH:/usr/lib/go-1.10/bin" >> ~/.bash_profile
source ~/.bash_profile

```

Next, you need to set the variables of `$GOPATH`, `$GOBIN`, and `$PATH`:

```
mkdir -p $HOME/go/bin
echo "export GOPATH=$HOME/go" >> ~/.bash_profile
echo "export GOBIN=$HOME/go/bin" >> ~/.bash_profile
echo "export PATH=$PATH:$HOME/go/bin" >> ~/.bash_profile
source ~/.bash_profile
```

To verify you have the proper version of Go installed, you could run:

```
go version
```
and see the following"
```
go version go1.10 linux/amd64
```

## Install IRISHub for Hackathon

After setting up the environment, you could get the correct version of IRISHub installed on your server. Please don't mix the version with which is used for Fuxi testnet.

```
mkdir -p $GOPATH/src/github.com/irisnet
cd $GOPATH/src/github.com/irisnet
git clone https://github.com/irisnet/irishub
cd irishub && git checkout gog/develop
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
make get_vendor_deps && make install
```

You could verify you have the right version installed by running the following commands: 

```
$ iris version
vGame Of Genesis
    
$ iriscli version
vGame Of Genesis
```

## Configure Your Node

### Change Moniker

After installing irishub, you could customize the configuteration of your node. 

### Set Seed Node

## Start Node


