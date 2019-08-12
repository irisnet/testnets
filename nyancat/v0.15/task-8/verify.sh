#!/bin/bash

echo "Parsing... DO NOT QUIT"

declare -A hash_height_map
declare -A hash_fva_map
declare -A fva_vote_map
declare -A fva_moniker_map

for i in {52001..52301}
do  
    # echo $i;  
    PreCommits=$(curl -s 'http://34.80.154.59:1317/blocks/'$i | jq '.block.last_commit.precommits')
    for row in $(echo "${PreCommits}" | jq -r '.[] | @base64'); do
        _jq() {
            echo ${row} | base64 --decode | jq -r ${1}
        }
        validator_address=$(_jq '.validator_address')
        height=$(_jq '.height')
        # echo $validator_address,$height,${#hash_height_map[@]},${hash_height_map["$validator_address"]}
        if [ "$validator_address" != "null" -a -z "${hash_height_map[$validator_address]}" ]
        then
            hash_height_map["$validator_address"]=$height
        fi
    done
done

for row in $(curl -s 'https://nyancat.irisplorer.io/api/stake/validators?page=1&size=100&type=""&origin=browser' | jq '.Data[] | .proposer_addr+":"+.operator_address' | sed 's/"//g'); do
    # echo $row
    hash=$(echo $row | cut -d':' -f1)
    fva=$(echo $row | cut -d':' -f2)
    hash_fva_map[$hash]=$fva
done

for row in $(curl -s 'https://nyancat.irisplorer.io/api/gov/proposals/1/voter_txs?page=1&size=100' | jq '.items[] | .voter+":"+.moniker+":"+.option' | sed 's/"//g'); do
    # echo $row
    fva=$(echo $row | cut -d':' -f1)
    moniker=$(echo $row | cut -d':' -f2)
    vote=$(echo $row | cut -d':' -f3)
    fva_vote_map[$fva]=$vote
    fva_moniker_map[$fva]=$moniker
done

for key in ${!hash_height_map[@]} 
do
    fva=${hash_fva_map[$key]}
    echo $key": "${hash_height_map[$key]}, $fva, ${fva_vote_map[$fva]}, ${fva_moniker_map[$fva]}
done
