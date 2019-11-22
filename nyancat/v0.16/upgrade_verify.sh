#!/bin/bash

echo "Parsing... DO NOT QUIT"

declare -A hash_height_map
declare -A fva_hash_map
declare -A fva_vote_map
declare -A fva_moniker_map
declare -A testnet_identity_map
declare -A mainnet_identity_map

### acquire validator_address of validator who signed one of blocks between 89500 and 89800 from lcd to save to hash_height_map, the height of software upgrade switched to v0.16.0-rc1 is 89500.
for i in {89501..89801}
do
    # echo $i;
    PreCommits=$(curl -s 'https://lcd.nyancat.irisnet.org/blocks/'$i | jq '.block.last_commit.precommits')
    for row in $(echo "${PreCommits}" | jq -r '.[] | @base64'); do
        _jq() {
            echo ${row} | base64 --decode | jq -r ${1}
        }
        validator_address=$(_jq '.validator_address')
        height=$(_jq '.height')
        ### filter validator_address of validator who already has signed last block.
        if [ "$validator_address" != "null" -a -z "${hash_height_map[$validator_address]}" ]
        then
            hash_height_map["$validator_address"]=$height
        fi
    done
done
echo "count of hash_height_map: " ${#hash_height_map[@]}

### acquire info of validators from api of explorer of nyancat and save it to fva_hash_map and testnet_identity_map.
curl -s 'https://nyancat.irisplorer.io/api/stake/validators?page=1&size=100&type=""&origin=browser' | jq '.Data[] | .proposer_addr+":"+.operator_address+":"+.description.identity' | sed 's/"//g' > nyancat-task-temp
while read row; do
    # echo $row
    hash=$(echo $row | cut -d':' -f1)
    fva=$(echo $row | cut -d':' -f2)
    identity=$(echo $row | cut -d':' -f3)
    fva_hash_map[$fva]=$hash
    testnet_identity_map[$fva]=$identity
done < nyancat-task-temp

### acquire info of validators from api of explorer of mainnet and save it to mainnet_identity_map.
curl -s 'https://www.irisplorer.io/api/stake/validators?page=1&size=100&type=""&origin=browser' | jq '.Data[] | .proposer_addr+":"+.operator_address+":"+.description.identity' | sed 's/"//g' > nyancat-task-temp
while read row; do
    identity=$(echo $row | cut -d':' -f3)
    # echo $identity
    [ -n "$identity" ] && mainnet_identity_map[$identity]="ok"
done < nyancat-task-temp

### acquire voted info in proposal #1 & #3 from api of explorer of nyancat ordered by time.
curl -s 'https://nyancat.irisplorer.io/api/gov/proposals/1/voter_txs?page=1&size=100' | jq '.items | reverse | .[] | .voter+":"+.moniker+":"+.option+"::"+.timestamp' | sed 's/"//g' > nyancat-task-temp
curl -s 'https://nyancat.irisplorer.io/api/gov/proposals/3/voter_txs?page=1&size=100' | jq '.items | reverse | .[] | .voter+":"+.moniker+":"+.option+"::"+.timestamp' | sed 's/"//g' >> nyancat-task-temp
index=0
while read row; do
    # echo $row
    fva=$(echo $row | cut -d':' -f1)
    moniker=$(echo $row | cut -d':' -f2)
    vote=$(echo $row | cut -d':' -f3)
    timestamp=$(echo $row | awk -F:: '{print $2}')
    ### filter validator of duplicate vote and unsigned one of 300 blocks after software upgrade.
    if [ -z "${fva_vote_map[$fva]}" -a -n "${fva_hash_map[$fva]}" ] && [ -n "${hash_height_map[${fva_hash_map[$fva]}]}" ]; then
        if [ $index -lt 40 ]; then
            index=$[$index + 1]
            fva_vote_map[$fva]=$vote
            echo -e $index"\t"$fva"\t"$timestamp"\t"$moniker"\t"${testnet_identity_map[$fva]}
        ### validator of mainnet beyond top 40 is passed.
        elif [ -n "${testnet_identity_map[$fva]}" ] && [ "${mainnet_identity_map[${testnet_identity_map[$fva]}]}" == "ok" ]; then
            index=$[$index + 1]
            fva_vote_map[$fva]=$vote
            echo -e $index"\t"$fva"\t"$timestamp"\t"$moniker"\t"${testnet_identity_map[$fva]}"\t"ok
        fi
    fi
    # echo $fva,$moniker,$vote
done < nyancat-task-temp
rm -f nyancat-task-temp
