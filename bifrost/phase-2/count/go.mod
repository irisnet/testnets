module github.com/irisnet/testnets/bifrost/phase-2/count

go 1.15

require (
	github.com/cosmos/cosmos-sdk v0.41.3
	github.com/irisnet/irishub-sdk-go v0.0.0-20210119021152-78fac3ec7b9b
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4
)
