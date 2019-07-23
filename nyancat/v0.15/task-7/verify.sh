#!/bin/bash

rm -rf /tmp/.iriscli # reset
# get txs by sender from lcd
TXS=$(curl http://34.80.154.59:1317/txs?action=send&sender=faa1282eufkw9qgm55symgqqg38nremslvggpylkht | jq 'map(select(.result.Code == 0))')

echo "Parsing... DO NOT QUIT"

for f in *.json
do 
    pgp="$(cut -d'.' -f1 <<<"$f")" # get PGP_ID from file name
    address=$(echo "1234567890" | iriscli keys add $pgp --recover --keystore $f --home /tmp/.iriscli | awk '{print $3}' | cut -d':' -f2) # import the keystore and get the address
    if [ $? -eq 0 ] # if success
    then 
        # select tx by recipient
        tx=$(jq -r --arg addr $address 'map(select(.tx.value.msg[0].value.outputs[0].address==$addr))' <<< "$TXS")
        complete=$(jq 'length' <<< "$tx")
        height=$(jq -r '.[0].height' <<< "$tx")
        hash=$(jq -r '.[0].hash' <<< "$tx")
        timestamp=$(jq -r '.[0].timestamp' <<< "$tx")
        res='{"PGP_ID": "'$pgp'", "Complete": "'$complete'", "height": "'$height'", "hash": "'$hash'", "timestamp": "'$timestamp'"}'
    else
        res='{"PGP_ID": "'$pgp'", "Complete": "0"}'
    fi
    RESULT+=","$res
done
RESULT=$(echo $RESULT | sed 's/,//') # remove the 1st ","
echo $RESULT
echo $(jq 'sort_by(.timestamp)' <<< "[$RESULT]") | jq . # sort by timestamp