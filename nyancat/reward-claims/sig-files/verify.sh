#!/bin/bash

# Verfiy signatures by keybase, you need to have keybase installed and running

for f in *.txt
do 
	echo "Verifing "$f" ..."
	cat $f | keybase pgp verify
	echo "----------------------------------------------------------"
done
