#!/bin/bash

# Verfiy signatures by keybase, you need to have keybase installed and running

for f in *.txt
do 
	echo "----------------------------------------------------------"
	echo "Verifying "$f" ..."
	keybase pgp verify -i $f
	echo "Wallet Address:"
	keybase pgp decrypt -i $f
	echo "----------------------------------------------------------"
done
