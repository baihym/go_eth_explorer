## go eth explorer

Get eth and tokens transactions for explorer.

## Usage

1. Setup config in file 'app/config/config.go'

2. Add tokens(symbol, contract) in table `tokens` what you want to support

3. Build

```
cd $GOPATH/src/github.com/baihym/go_eth_explorer && go build -o bin_eth_explorer ./app
```

4. Run

```
./bin_eth_explorer {lastHandleBlockNumber: default 0} // example: ./bin_eth_explorer 0
```