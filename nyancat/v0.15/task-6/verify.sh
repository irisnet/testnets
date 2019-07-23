#!/bin/bash
# select txs from lcd where recipient = `public_account` and pub_key.type = "tendermint/PubKeyMultisigThreshold" order by timestamp asc
TXS=$(curl -N 'http://34.80.154.59:1317/txs?action=send&recipient=faa1dhwf97xsdjy8xdrxqac5f6zvl2f664dsujmrl5' | jq 'sort_by(.timestamp) | map(select(.result.Code == 0)) | map(select(.tx.value.signatures[].pub_key.type == "tendermint/PubKeyMultisigThreshold"))')

echo "Parsing... DO NOT QUIT"

for row in $(echo "${TXS}" | jq -r '.[] | @base64'); do
    _jq() {
        echo ${row} | base64 --decode | jq -r ${1}
    }

    hash=$(_jq '.hash')
    height=$(_jq '.height')
    timestamp=$(_jq '.timestamp')
    pgp=$(_jq '.tx.value.memo')
    RESULT+=',{"PGP_ID": "'$pgp'", "height": "'$height'", "hash": "'$hash'", "timestamp": "'$timestamp'"}'
done

RESULT=$(echo $RESULT | sed 's/,//') # remove the 1st ","
echo [$RESULT] | jq . 