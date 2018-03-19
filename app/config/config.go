package config

// Rpc address
var ETHRPCAddress string = ""

// Temp save last handle block number
var ETHLastBlockNumberFile = "/tmp/go_eth_explorer_handle_block.txt"

// Etherscan API
var (
	// Note: https://ropsten.etherscan.io is test
	ETHERSCANHost          string = "https://ropsten.etherscan.io"
	ETHERSCANApiKeyAppName string = ""
	ETHERSCANApiKeyToken   string = ""
)

// Your mysql connection config
var (
	DBDSN     string = "user:password@tcp(localhost:3306)/eth_explorer?clientFoundRows=false&parseTime=true&loc=Asia%2FShanghai&timeout=5s&collation=utf8mb4_bin&interpolateParams=true"
	DBMaxOpen int    = 100
	DBMaxIdle int    = 10
)
