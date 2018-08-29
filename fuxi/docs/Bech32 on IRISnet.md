# Bech32 on IRISnet

Bech32 is a new Bitcoin address format proposed by Pieter Wuille and Greg Maxwell. Besides Bitcoin addresses, Bech32 can encode any short binary data. In the IRIS network, keys and addresses may refer to a number of different roles in the network like accounts, validators etc. The IRIS network is designed to use the Bech32 address format to provide robust integrity checks on data. The human readable part(HRP) makes it more efficient to read and the users could see error messages.


## Human Readable Part Table


| HRP        | Definition |
| ------------- |:-------------:|
|faa	|IRISnet Account Address|
|fap|	IRISnet Account Public Key|
|fva	|IRISnet Consensus Address|
|fvp|	IRISnet Consensus Public Key|

## Encoding

Not all interfaces to users IRISnet should be exposed as bech32 interfaces. Many address are still in hex or base64 encoded form.

To covert between other binary reprsentation of addresses and keys, it is important to first apply the Amino enocoding process before bech32 encoding.


