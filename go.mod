module github.com/ownesthq/cosmos-sdk

require (
	github.com/aws/aws-sdk-go v1.20.11
	github.com/bartekn/go-bip39 v0.0.0-20171116152956-a05967ea095d
	github.com/bgentry/speakeasy v0.1.0
	github.com/btcsuite/btcd v0.0.0-20190115013929-ed77733ec07d
	github.com/cosmos/cosmos-sdk v0.35.0 // indirect
	github.com/cosmos/go-bip39 v0.0.0-20180618194314-52158e4697b8
	github.com/cosmos/ledger-cosmos-go v0.10.3
	github.com/gogo/protobuf v1.2.1
	github.com/golang/mock v1.3.1-0.20190508161146-9fa652df1129
	github.com/golang/protobuf v1.3.0
	github.com/google/pprof v0.0.0-20190515194954-54271f7e092f // indirect
	github.com/gorilla/mux v1.7.0
	github.com/lusis/go-slackbot v0.0.0-20180109053408-401027ccfef5 // indirect
	github.com/lusis/slack-test v0.0.0-20190426140909-c40012f20018 // indirect
	github.com/mattn/go-isatty v0.0.6
	github.com/nlopes/slack v0.5.0
	github.com/pelletier/go-toml v1.2.0
	github.com/pkg/errors v0.8.1
	github.com/rakyll/statik v0.1.4
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.3
	github.com/spf13/viper v1.3.2
	github.com/stretchr/testify v1.3.0
	github.com/tendermint/btcd v0.1.1
	github.com/tendermint/go-amino v0.15.0
	github.com/tendermint/iavl v0.12.2
	github.com/tendermint/tendermint v0.32.0
	golang.org/x/arch v0.0.0-20190312162104-788fe5ffcd8c // indirect
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2
	gopkg.in/yaml.v2 v2.2.2
)

replace golang.org/x/crypto => github.com/tendermint/crypto v0.0.0-20180820045704-3764759f34a5
