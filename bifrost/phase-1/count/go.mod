module github.com/irisnet/testnets/bifrost/phase-1/count

go 1.15

require (
	github.com/irisnet/irishub-sdk-go v0.0.0-20201106071848-e481ea02c310
	github.com/tidwall/gjson v1.9.3
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4
)
