// Code generated by go-openrpc. DO NOT EDIT.
package web3

type Service interface {
	// Returns the version of the current client
	Web3ClientVersion() (*Web3ClientVersionResult, error)
	// Hashes data using the Keccak-256 algorithm
	Web3Sha3(*Web3Sha3Params) (*Web3Sha3Result, error)
	// Determines if this client is listening for new network connections.
	NetListening() (*NetListeningResult, error)
	// Returns the number of peers currently connected to this client.
	NetPeerCount() (*NetPeerCountResult, error)
	// Returns the chain ID associated with the current network.
	NetVersion() (*NetVersionResult, error)
	// Returns the number of most recent block.
	EthBlockNumber() (*EthBlockNumberResult, error)
	// Executes a new message call (locally) immediately without creating a transaction on the block chain.
	EthCall(*EthCallParams) (*EthCallResult, error)
	// Returns the currently configured chain id, a value used in replay-protected transaction signing as introduced by [EIP-155](https://github.com/ethereum/EIPs/blob/master/EIPS/eip-155.md).
	EthChainId() (*EthChainIdResult, error)
	// Returns the client coinbase address.
	EthCoinbase() (*EthCoinbaseResult, error)
	// Generates and returns an estimate of how much gas is necessary to allow the transaction to complete. The transaction will not be added to the blockchain. Note that the estimate may be significantly more than the amount of gas actually used by the transaction, for a variety of reasons including EVM mechanics and node performance.
	EthEstimateGas(*EthEstimateGasParams) (*EthEstimateGasResult, error)
	// Returns the current price per gas in wei
	EthGasPrice() (*EthGasPriceResult, error)
	// Returns Ether balance of a given or account or contract
	EthGetBalance(*EthGetBalanceParams) (*EthGetBalanceResult, error)
	// Gets a block for a given hash
	EthGetBlockByHash(*EthGetBlockByHashParams) (*EthGetBlockByHashResult, error)
	// Gets a block for a given number salad
	EthGetBlockByNumber(*EthGetBlockByNumberParams) (*EthGetBlockByNumberResult, error)
	// Returns the number of transactions in a block from a block matching the given block hash.
	EthGetBlockTransactionCountByHash(*EthGetBlockTransactionCountByHashParams) (*EthGetBlockTransactionCountByHashResult, error)
	// Returns the number of transactions in a block from a block matching the given block number.
	EthGetBlockTransactionCountByNumber(*EthGetBlockTransactionCountByNumberParams) (*EthGetBlockTransactionCountByNumberResult, error)
	// Returns code at a given contract address
	EthGetCode(*EthGetCodeParams) (*EthGetCodeResult, error)
	// Polling method for a filter, which returns an array of logs which occurred since last poll.
	EthGetFilterChanges(*EthGetFilterChangesParams) (*EthGetFilterChangesResult, error)
	// Returns an array of all logs matching filter with given id.
	EthGetFilterLogs(*EthGetFilterLogsParams) (*EthGetFilterLogsResult, error)
	// Returns raw transaction data of a transaction with the given hash.
	EthGetRawTransactionByHash(*EthGetRawTransactionByHashParams) (*EthGetRawTransactionByHashResult, error)
	// Returns raw transaction data of a transaction with the given hash.
	EthGetRawTransactionByBlockHashAndIndex(*EthGetRawTransactionByBlockHashAndIndexParams) (*EthGetRawTransactionByBlockHashAndIndexResult, error)
	// Returns raw transaction data of a transaction with the given hash.
	EthGetRawTransactionByBlockNumberAndIndex(*EthGetRawTransactionByBlockNumberAndIndexParams) (*EthGetRawTransactionByBlockNumberAndIndexResult, error)
	// Returns an array of all logs matching a given filter object.
	EthGetLogs(*EthGetLogsParams) (*EthGetLogsResult, error)
	// Gets a storage value from a contract address, a position, and an optional blockNumber
	EthGetStorageAt(*EthGetStorageAtParams) (*EthGetStorageAtResult, error)
	// Returns the information about a transaction requested by the block hash and index of which it was mined.
	EthGetTransactionByBlockHashAndIndex(*EthGetTransactionByBlockHashAndIndexParams) (*EthGetTransactionByBlockHashAndIndexResult, error)
	// Returns the information about a transaction requested by the block hash and index of which it was mined.
	EthGetTransactionByBlockNumberAndIndex(*EthGetTransactionByBlockNumberAndIndexParams) (*EthGetTransactionByBlockNumberAndIndexResult, error)
	// Returns the information about a transaction requested by transaction hash.
	EthGetTransactionByHash(*EthGetTransactionByHashParams) (*EthGetTransactionByHashResult, error)
	// Returns the number of transactions sent from an address
	EthGetTransactionCount(*EthGetTransactionCountParams) (*EthGetTransactionCountResult, error)
	// Returns the receipt information of a transaction by its hash.
	EthGetTransactionReceipt(*EthGetTransactionReceiptParams) (*EthGetTransactionReceiptResult, error)
	// Returns information about a uncle of a block by hash and uncle index position.
	EthGetUncleByBlockHashAndIndex(*EthGetUncleByBlockHashAndIndexParams) (*EthGetUncleByBlockHashAndIndexResult, error)
	// Returns information about a uncle of a block by hash and uncle index position.
	EthGetUncleByBlockNumberAndIndex(*EthGetUncleByBlockNumberAndIndexParams) (*EthGetUncleByBlockNumberAndIndexResult, error)
	// Returns the number of uncles in a block from a block matching the given block hash.
	EthGetUncleCountByBlockHash(*EthGetUncleCountByBlockHashParams) (*EthGetUncleCountByBlockHashResult, error)
	// Returns the number of uncles in a block from a block matching the given block number.
	EthGetUncleCountByBlockNumber(*EthGetUncleCountByBlockNumberParams) (*EthGetUncleCountByBlockNumberResult, error)
	// Returns the account- and storage-values of the specified account including the Merkle-proof.
	EthGetProof(*EthGetProofParams) (*EthGetProofResult, error)
	// Returns the hash of the current block, the seedHash, and the boundary condition to be met ('target').
	EthGetWork() (*EthGetWorkResult, error)
	// Returns the number of hashes per second that the node is mining with.
	EthHashrate() (*EthHashrateResult, error)
	// Returns true if client is actively mining new blocks.
	EthMining() (*EthMiningResult, error)
	// Creates a filter in the node, to notify when a new block arrives. To check if the state has changed, call eth_getFilterChanges.
	EthNewBlockFilter() (*EthNewBlockFilterResult, error)
	// Creates a filter object, based on filter options, to notify when the state changes (logs). To check if the state has changed, call eth_getFilterChanges.
	EthNewFilter(*EthNewFilterParams) (*EthNewFilterResult, error)
	// Creates a filter in the node, to notify when new pending transactions arrive. To check if the state has changed, call eth_getFilterChanges.
	EthNewPendingTransactionFilter() (*EthNewPendingTransactionFilterResult, error)
	// Returns the pending transactions list
	EthPendingTransactions() (*EthPendingTransactionsResult, error)
	// Returns the current ethereum protocol version.
	EthProtocolVersion() (*EthProtocolVersionResult, error)
	// The sign method calculates an Ethereum specific signature.
	EthSign(*EthSignParams) (*EthSignResult, error)
	// Returns a list of addresses owned by client.
	EthAccounts() (*EthAccountsResult, error)
	// Creates new message call transaction or a contract creation, if the data field contains code.
	EthSendTransaction(*EthSendTransactionParams) (*EthSendTransactionResult, error)
	// Creates new message call transaction or a contract creation for signed transactions.
	EthSendRawTransaction(*EthSendRawTransactionParams) (*EthSendRawTransactionResult, error)
	// Returns an array of all logs matching a given filter object.
	EthSubmitHashrate(*EthSubmitHashrateParams) (*EthSubmitHashrateResult, error)
	// Used for submitting a proof-of-work solution.
	EthSubmitWork(*EthSubmitWorkParams) (*EthSubmitWorkResult, error)
	// Returns an object with data about the sync status or false.
	EthSyncing() (*EthSyncingResult, error)
	// Uninstalls a filter with given id. Should always be called when watch is no longer needed. Additionally Filters timeout when they aren't requested with eth_getFilterChanges for a period of time.
	EthUninstallFilter(*EthUninstallFilterParams) (*EthUninstallFilterResult, error)
}
type Web3ClientVersionResult struct {
	// client version
	ClientVersion string `json:"clientVersion"`
}
type Web3Sha3Params struct {
	// data to hash using the Keccak-256 algorithm
	Data string `json:"data"`
}
type Web3Sha3Result struct {
	// Hex representation of a Keccak 256 hash
	HashedData string `json:"hashedData"`
}
type NetListeningResult struct {
	// `true` if listening is active or `false` if listening is not active
	IsNetListening bool `json:"isNetListening"`
}
type NetPeerCountResult struct {
	// Hex representation of number of connected peers
	NumConnectedPeers string `json:"numConnectedPeers"`
}
type NetVersionResult struct {
	// chain ID associated with the current network
	ChainID string `json:"chainID"`
}
type BlockNumber struct {
	// The hex representation of the block's height
	BlockNumber string `json:"blockNumber"`
}
type EthBlockNumberResult struct {
	BlockNumber string `json:"blockNumber"`
}
type TransactionIndex struct {
	// Hex representation of the integer
	Integer string `json:"integer"`
}
type Transaction struct {
	// Integer of the transaction's index position in the block. null when its pending
	TransactionIndex string `json:"transactionIndex"`
	// Hash of the block where this transaction was in. null when its pending
	BlockHash string `json:"blockHash"`
	// Address of the sender
	From string `json:"from"`
	// Hex representation of a Keccak 256 hash
	Hash string `json:"hash"`
	// The data field sent with the transaction
	Data string `json:"data"`
	// A number only to be used once
	Nonce string `json:"nonce"`
	// The gas limit provided by the sender in Wei
	Gas string `json:"gas"`
	// Hex representation of a Keccak 256 hash
	Value string `json:"value"`
	// ECDSA recovery id
	V string `json:"v"`
	// ECDSA signature s
	S string `json:"s"`
	// The gas price willing to be paid by the sender in Wei
	GasPrice string `json:"gasPrice"`
	// address of the receiver. null when its a contract creation transaction
	To string `json:"to"`
	// Block number where this transaction was in. null when its pending
	BlockNumber string `json:"blockNumber"`
	// ECDSA signature r
	R string `json:"r"`
}
type BlockHash struct {
	// Hex representation of a Keccak 256 hash
	Keccak string `json:"keccak"`
}
type EthCallParams struct {
	Transaction

	BlockNumber string `json:"blockNumber"`
}
type EthCallResult struct {
	// Hex representation of a variable length byte array
	ReturnValue string `json:"returnValue"`
}
type EthChainIdResult struct {
	// hex format integer of the current chain id. Defaults are mainnet=61, morden=62.
	ChainId string `json:"chainId"`
}
type EthCoinbaseResult struct {
	// The address owned by the client that is used as default for things like the mining reward
	Address string `json:"address"`
}
type EthEstimateGasParams struct {
	Transaction
}
type EthEstimateGasResult struct {
	// Hex representation of the integer
	GasUsed string `json:"gasUsed"`
}
type EthGasPriceResult struct {
	// Hex representation of the integer
	GasPrice string `json:"gasPrice"`
}
type EthGetBalanceParams struct {
	// The address of the account or contract
	Address string `json:"address"`
	// The hex representation of the block's height
	BlockNumber string `json:"blockNumber"`
}
type GetBalanceResult struct {
	// Hex representation of the integer
	Integer string `json:"integer"`
}
type EthGetBalanceResult struct {
	GetBalanceResult string `json:"getBalanceResult"`
}
type EthGetBlockByHashParams struct {
	// The hex representation of the Keccak 256 of the RLP encoded block
	BlockHash string `json:"blockHash"`
	// If `true` it returns the full transaction objects, if `false` only the hashes of the transactions.
	IsTransactionsIncluded bool `json:"isTransactionsIncluded"`
}
type Block struct {
	// Hex representation of a Keccak 256 hash
	Sha3Uncles string `json:"sha3Uncles"`
	// Hex representation of a Keccak 256 hash
	TransactionsRoot string `json:"transactionsRoot"`
	// Hex representation of a Keccak 256 hash
	ParentHash string `json:"parentHash"`
	// The address of the beneficiary to whom the mining rewards were given or null when its the pending block
	Miner string `json:"miner"`
	// Integer of the difficulty for this block
	Difficulty string `json:"difficulty"`
	// The total used gas by all transactions in this block
	GasUsed string `json:"gasUsed"`
	// The unix timestamp for when the block was collated
	Timestamp string `json:"timestamp"`
	// Array of transaction objects, or 32 Bytes transaction hashes depending on the last given parameter
	Transactions []Transactions `json:"transactions"`
	// The block number or null when its the pending block
	Number string `json:"number"`
	// The block hash or null when its the pending block
	Hash string `json:"hash"`
	// Array of uncle hashes
	Uncles []string `json:"uncles"`
	// Hex representation of a Keccak 256 hash
	ReceiptsRoot string `json:"receiptsRoot"`
	// The 'extra data' field of this block
	ExtraData string `json:"extraData"`
	// Hex representation of a Keccak 256 hash
	StateRoot string `json:"stateRoot"`
	// Integer of the total difficulty of the chain until this block
	TotalDifficulty string `json:"totalDifficulty"`
	// Integer the size of this block in bytes
	Size string `json:"size"`
	// The maximum gas allowed in this block
	GasLimit string `json:"gasLimit"`
	// Randomly selected number to satisfy the proof-of-work or null when its the pending block
	Nonce string `json:"nonce"`
	// The bloom filter for the logs of the block or null when its the pending block
	LogsBloom string `json:"logsBloom"`
}
type Miner struct {
	Address string `json:"address"`
}
type Transactions struct {
	Transaction
}
type Number struct {
	// Hex representation of the integer
	Integer string `json:"integer"`
}
type Hash struct {
	// Hex representation of a Keccak 256 hash
	Keccak string `json:"keccak"`
}
type Uncles struct {
	// Hex representation of a Keccak 256 hash
	Keccak string `json:"keccak"`
}
type TotalDifficulty struct {
	// Hex representation of the integer
	Integer string `json:"integer"`
}
type Nonce struct {
	// Hex representation of the integer
	Integer string `json:"integer"`
}
type GetBlockByHashResult struct {
	Block
}
type EthGetBlockByHashResult struct {
	GetBlockByHashResult Block `json:"getBlockByHashResult"`
}
type EthGetBlockByNumberParams struct {
	BlockNumber string `json:"blockNumber"`
	// If `true` it returns the full transaction objects, if `false` only the hashes of the transactions.
	IsTransactionsIncluded bool `json:"isTransactionsIncluded"`
}
type GetBlockByNumberResult struct {
	Block
}
type EthGetBlockByNumberResult struct {
	GetBlockByNumberResult Block `json:"getBlockByNumberResult"`
}
type EthGetBlockTransactionCountByHashParams struct {
	// The hex representation of the Keccak 256 of the RLP encoded block
	BlockHash string `json:"blockHash"`
}
type BlockTransactionCountByHash struct {
	// Hex representation of the integer
	Integer string `json:"integer"`
}
type EthGetBlockTransactionCountByHashResult struct {
	// The Number of total transactions in the given block
	BlockTransactionCountByHash string `json:"blockTransactionCountByHash"`
}
type EthGetBlockTransactionCountByNumberParams struct {
	BlockNumber string `json:"blockNumber"`
}
type EthGetBlockTransactionCountByNumberResult struct {
	// The Number of total transactions in the given block
	BlockTransactionCountByHash string `json:"blockTransactionCountByHash"`
}
type EthGetCodeParams struct {
	// The address of the contract
	Address string `json:"address"`
	// The hex representation of the block's height
	BlockNumber string `json:"blockNumber"`
}
type EthGetCodeResult struct {
	// Hex representation of a variable length byte array
	Bytes string `json:"bytes"`
}
type EthGetFilterChangesParams struct {
	// An identifier used to reference the filter.
	FilterId string `json:"filterId"`
}
type Log struct {
	Topics []Topics `json:"topics"`
	// Hex representation of a Keccak 256 hash
	TransactionHash string `json:"transactionHash"`
	// Sender of the transaction
	Address string `json:"address"`
	// The hex representation of the Keccak 256 of the RLP encoded block
	BlockHash string `json:"blockHash"`
	// The hex representation of the block's height
	BlockNumber string `json:"blockNumber"`
	// Hex representation of a variable length byte array
	Data string `json:"data"`
	// Hex representation of the integer
	LogIndex string `json:"logIndex"`
	// Hex representation of the integer
	TransactionIndex string `json:"transactionIndex"`
}
type LogResult struct {
	// An indexed event generated during a transaction
	Log

	Topics []Topics `json:"topics"`
	// Hex representation of a Keccak 256 hash
	TransactionHash string `json:"transactionHash"`
	// Sender of the transaction
	Address string `json:"address"`
	// The hex representation of the Keccak 256 of the RLP encoded block
	BlockHash string `json:"blockHash"`
	// The hex representation of the block's height
	BlockNumber string `json:"blockNumber"`
	// Hex representation of a variable length byte array
	Data string `json:"data"`
	// Hex representation of the integer
	LogIndex string `json:"logIndex"`
	// Hex representation of the integer
	TransactionIndex string `json:"transactionIndex"`
}
type EthGetFilterChangesResult struct {
	LogResult []LogResult `json:"logResult"`
}
type EthGetFilterLogsParams struct {
	// An identifier used to reference the filter.
	FilterId string `json:"filterId"`
}
type Logs struct {
	// An indexed event generated during a transaction
	Log
	// Hex representation of the integer
	LogIndex string `json:"logIndex"`
	// Hex representation of the integer
	TransactionIndex string `json:"transactionIndex"`
	// Hex representation of a Keccak 256 hash
	TransactionHash string `json:"transactionHash"`
	// Sender of the transaction
	Address string `json:"address"`
	// The hex representation of the Keccak 256 of the RLP encoded block
	BlockHash string `json:"blockHash"`
	// The hex representation of the block's height
	BlockNumber string `json:"blockNumber"`
	// Hex representation of a variable length byte array
	Data string `json:"data"`

	Topics []Topics `json:"topics"`
}
type EthGetFilterLogsResult struct {
	Logs []Logs `json:"logs"`
}
type EthGetRawTransactionByHashParams struct {
	// Hex representation of a Keccak 256 hash
	TransactionHash string `json:"transactionHash"`
}
type EthGetRawTransactionByHashResult struct {
	// Hex representation of a variable length byte array
	RawTransactionByHash string `json:"rawTransactionByHash"`
}
type EthGetRawTransactionByBlockHashAndIndexParams struct {
	// The hex representation of the Keccak 256 of the RLP encoded block
	BlockHash string `json:"blockHash"`
	// Hex representation of the integer
	Index string `json:"index"`
}
type EthGetRawTransactionByBlockHashAndIndexResult struct {
	// Hex representation of a variable length byte array
	RawTransaction string `json:"rawTransaction"`
}
type EthGetRawTransactionByBlockNumberAndIndexParams struct {
	BlockNumber string `json:"blockNumber"`
	// Hex representation of the integer
	Index string `json:"index"`
}
type EthGetRawTransactionByBlockNumberAndIndexResult struct {
	// Hex representation of a variable length byte array
	RawTransaction string `json:"rawTransaction"`
}
type Filter struct {
	// The hex representation of the block's height
	FromBlock string `json:"fromBlock"`
	// The hex representation of the block's height
	ToBlock string `json:"toBlock"`

	Address string `json:"address"`
	// Array of 32 Bytes DATA topics. Topics are order-dependent. Each topic can also be an array of DATA with 'or' options
	Topics []string `json:"topics"`
}
type Address struct {
	// Address of the contract from which to monitor events
	Address string `json:"address"`
}
type Topics struct {
	// Hex representation of a 256 bit unit of data
	DataWord string `json:"dataWord"`
}
type EthGetLogsParams struct {
	// A filter used to monitor the blockchain for log/events
	Filter
}
type EthGetLogsResult struct {
	Logs []Logs `json:"logs"`
}
type EthGetStorageAtParams struct {
	Address string `json:"address"`
	// Hex representation of the storage slot where the variable exists
	Position string `json:"position"`

	BlockNumber string `json:"blockNumber"`
}
type EthGetStorageAtResult struct {
	// Hex representation of a 256 bit unit of data
	DataWord string `json:"dataWord"`
}
type EthGetTransactionByBlockHashAndIndexParams struct {
	// The hex representation of the Keccak 256 of the RLP encoded block
	BlockHash string `json:"blockHash"`
	// Hex representation of the integer
	Index string `json:"index"`
}
type TransactionResult struct {
	Transaction
}
type EthGetTransactionByBlockHashAndIndexResult struct {
	TransactionResult Transaction `json:"transactionResult"`
}
type EthGetTransactionByBlockNumberAndIndexParams struct {
	BlockNumber string `json:"blockNumber"`
	// Hex representation of the integer
	Index string `json:"index"`
}
type EthGetTransactionByBlockNumberAndIndexResult struct {
	TransactionResult Transaction `json:"transactionResult"`
}
type EthGetTransactionByHashParams struct {
	// Hex representation of a Keccak 256 hash
	TransactionHash string `json:"transactionHash"`
}
type EthGetTransactionByHashResult struct {
	Transaction
}
type EthGetTransactionCountParams struct {
	Address string `json:"address"`

	BlockNumber string `json:"blockNumber"`
}
type NonceOrNull struct {
	// A number only to be used once
	Nonce string `json:"nonce"`
}
type EthGetTransactionCountResult struct {
	NonceOrNull string `json:"nonceOrNull"`
}
type EthGetTransactionReceiptParams struct {
	// Hex representation of a Keccak 256 hash
	TransactionHash string `json:"transactionHash"`
}
type Receipt struct {
	// The hex representation of the block's height
	BlockNumber string `json:"blockNumber"`
	// Hex representation of the integer
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	// Hex representation of the integer
	GasUsed string `json:"gasUsed"`
	// An array of all the logs triggered during the transaction
	Logs []Logs `json:"logs"`
	// A 2048 bit bloom filter from the logs of the transaction. Each log sets 3 bits though taking the low-order 11 bits of each of the first three pairs of bytes in a Keccak 256 hash of the log's byte series
	TransactionIndex string `json:"transactionIndex"`
	// Whether or not the transaction threw an error.
	Status string `json:"status"`
	// The hex representation of the Keccak 256 of the RLP encoded block
	BlockHash string `json:"blockHash"`
	// The contract address created, if the transaction was a contract creation, otherwise null
	ContractAddress string `json:"contractAddress"`
	// The sender of the transaction
	From string `json:"from"`
	// A 2048 bit bloom filter from the logs of the transaction. Each log sets 3 bits though taking the low-order 11 bits of each of the first three pairs of bytes in a Keccak 256 hash of the log's byte series
	LogsBloom string `json:"logsBloom"`
	// Destination address of the transaction
	To string `json:"to"`
	// Hex representation of a Keccak 256 hash
	TransactionHash string `json:"transactionHash"`
}
type EthGetTransactionReceiptResult struct {
	// returns either a receipt or null
	Receipt
}
type EthGetUncleByBlockHashAndIndexParams struct {
	// The hex representation of the Keccak 256 of the RLP encoded block
	BlockHash string `json:"blockHash"`
	// Hex representation of the integer
	Index string `json:"index"`
}
type Uncle struct {
	// Randomly selected number to satisfy the proof-of-work or null when its the pending block
	Nonce string `json:"nonce"`
	// Hex representation of a Keccak 256 hash
	TransactionsRoot string `json:"transactionsRoot"`
	// Integer of the total difficulty of the chain until this block
	TotalDifficulty string `json:"totalDifficulty"`
	// Integer the size of this block in bytes
	Size string `json:"size"`
	// The total used gas by all transactions in this block
	GasUsed string `json:"gasUsed"`
	// Array of uncle hashes
	Uncles []string `json:"uncles"`
	// The block number or null when its the pending block
	Number string `json:"number"`
	// The block hash or null when its the pending block
	Hash string `json:"hash"`
	// Hex representation of a Keccak 256 hash
	Sha3Uncles string `json:"sha3Uncles"`
	// Hex representation of a Keccak 256 hash
	StateRoot string `json:"stateRoot"`
	// The 'extra data' field of this block
	ExtraData string `json:"extraData"`
	// The unix timestamp for when the block was collated
	Timestamp string `json:"timestamp"`
	// Hex representation of a Keccak 256 hash
	ReceiptsRoot string `json:"receiptsRoot"`
	// The address of the beneficiary to whom the mining rewards were given or null when its the pending block
	Miner string `json:"miner"`
	// The maximum gas allowed in this block
	GasLimit string `json:"gasLimit"`
	// Hex representation of a Keccak 256 hash
	ParentHash string `json:"parentHash"`
	// The bloom filter for the logs of the block or null when its the pending block
	LogsBloom string `json:"logsBloom"`
	// Integer of the difficulty for this block
	Difficulty string `json:"difficulty"`
}
type UncleOrNull struct {
	// Orphaned blocks that can be included in the chain but at a lower block reward. NOTE: An uncle doesn’t contain individual transactions.
	Uncle
}
type EthGetUncleByBlockHashAndIndexResult struct {
	UncleOrNull Uncle `json:"uncleOrNull"`
}
type EthGetUncleByBlockNumberAndIndexParams struct {
	// The hex representation of the block's height
	UncleBlockNumber string `json:"uncleBlockNumber"`
	// Hex representation of the integer
	Index string `json:"index"`
}
type UncleResult struct {
	// Orphaned blocks that can be included in the chain but at a lower block reward. NOTE: An uncle doesn’t contain individual transactions.
	Uncle
}
type EthGetUncleByBlockNumberAndIndexResult struct {
	// returns an uncle or null
	UncleResult Uncle `json:"uncleResult"`
}
type EthGetUncleCountByBlockHashParams struct {
	// The hex representation of the Keccak 256 of the RLP encoded block
	BlockHash string `json:"blockHash"`
}
type UncleCountOrNull struct {
	// Hex representation of the integer
	Integer string `json:"integer"`
}
type EthGetUncleCountByBlockHashResult struct {
	UncleCountOrNull string `json:"uncleCountOrNull"`
}
type EthGetUncleCountByBlockNumberParams struct {
	BlockNumber string `json:"blockNumber"`
}
type EthGetUncleCountByBlockNumberResult struct {
	UncleCountOrNull string `json:"uncleCountOrNull"`
}
type EthGetProofParams struct {
	// The address of the account or contract
	Address string `json:"address"`
	// The storage keys of all the storage slots being requested
	StorageKeys []string `json:"storageKeys"`

	BlockNumber string `json:"blockNumber"`
}
type StorageKeys struct {
	// Hex representation of the integer
	Integer string `json:"integer"`
}
type ProofAccount struct {
	// Hex representation of the integer
	Balance string `json:"balance"`
	// Hex representation of a Keccak 256 hash
	CodeHash string `json:"codeHash"`
	// A number only to be used once
	Nonce string `json:"nonce"`
	// Hex representation of a Keccak 256 hash
	StorageHash string `json:"storageHash"`
	// Current block header PoW hash.
	StorageProof []StorageProof `json:"storageProof"`
	// The address of the account or contract of the request
	Address string `json:"address"`
	// The set of node values needed to traverse a patricia merkle tree (from root to leaf) to retrieve a value
	AccountProof []string `json:"accountProof"`
}
type Proof struct {
	// Hex representation of a variable length byte array
	ProofNode string `json:"proofNode"`
}
type StorageProof struct {
	// The set of node values needed to traverse a patricia merkle tree (from root to leaf) to retrieve a value
	Proof []string `json:"proof"`
	// Hex representation of the integer
	Key string `json:"key"`
	// Hex representation of the integer
	Value string `json:"value"`
}
type AccountProof struct {
	// Hex representation of a variable length byte array
	ProofNode string `json:"proofNode"`
}
type ProofAccountOrNull struct {
	// The merkle proofs of the specified account connecting them to the blockhash of the block specified
	ProofAccount
}
type EthGetProofResult struct {
	ProofAccountOrNull ProofAccount `json:"proofAccountOrNull"`
}
type EthGetWorkResult struct {
	Work []string `json:"work"`
}
type EthHashrateResult struct {
	// Hex representation of the integer
	HashesPerSecond string `json:"hashesPerSecond"`
}
type EthMiningResult struct {
	// Whether of not the client is mining
	Mining bool `json:"mining"`
}
type EthNewBlockFilterResult struct {
	// Hex representation of the integer
	FilterId string `json:"filterId"`
}
type EthNewFilterParams struct {
	// A filter used to monitor the blockchain for log/events
	Filter
}
type EthNewFilterResult struct {
	// Hex representation of the integer
	FilterId string `json:"filterId"`
}
type EthNewPendingTransactionFilterResult struct {
	// Hex representation of the integer
	FilterId string `json:"filterId"`
}
type PendingTransactions struct {
	Transaction
	// Integer of the transaction's index position in the block. null when its pending
	TransactionIndex string `json:"transactionIndex"`
	// Hash of the block where this transaction was in. null when its pending
	BlockHash string `json:"blockHash"`
	// Address of the sender
	From string `json:"from"`
	// Hex representation of a Keccak 256 hash
	Hash string `json:"hash"`
	// The data field sent with the transaction
	Data string `json:"data"`
	// A number only to be used once
	Nonce string `json:"nonce"`
	// The gas limit provided by the sender in Wei
	Gas string `json:"gas"`
	// Hex representation of a Keccak 256 hash
	Value string `json:"value"`
	// ECDSA recovery id
	V string `json:"v"`
	// ECDSA signature s
	S string `json:"s"`
	// The gas price willing to be paid by the sender in Wei
	GasPrice string `json:"gasPrice"`
	// address of the receiver. null when its a contract creation transaction
	To string `json:"to"`
	// Block number where this transaction was in. null when its pending
	BlockNumber string `json:"blockNumber"`
	// ECDSA signature r
	R string `json:"r"`
}
type EthPendingTransactionsResult struct {
	PendingTransactions []PendingTransactions `json:"pendingTransactions"`
}
type EthProtocolVersionResult struct {
	// Hex representation of the integer
	ProtocolVersion string `json:"protocolVersion"`
}
type EthSignParams struct {
	Address string `json:"address"`
	// Hex representation of a variable length byte array
	Bytes string `json:"bytes"`
}
type EthSignResult struct {
	// Hex representation of a variable length byte array
	Signature string `json:"signature"`
}
type Addresses struct {
	Address string `json:"address"`
}
type EthAccountsResult struct {
	// addresses owned by the client
	Addresses []string `json:"addresses"`
}
type EthSendTransactionParams struct {
	Transaction
}
type EthSendTransactionResult struct {
	// Hex representation of a Keccak 256 hash
	TransactionHash string `json:"transactionHash"`
}
type EthSendRawTransactionParams struct {
	// Hex representation of a variable length byte array
	SignedTransactionData string `json:"signedTransactionData"`
}
type EthSendRawTransactionResult struct {
	// Hex representation of a Keccak 256 hash
	TransactionHash string `json:"transactionHash"`
}
type EthSubmitHashrateParams struct {
	// Hex representation of a 256 bit unit of data
	HashRate string `json:"hashRate"`
	// Hex representation of a 256 bit unit of data
	Id string `json:"id"`
}
type EthSubmitHashrateResult struct {
	// whether of not submitting went through successfully
	SubmitHashRateSuccess bool `json:"submitHashRateSuccess"`
}
type EthSubmitWorkParams struct {
	// A number only to be used once
	Nonce string `json:"nonce"`
	// Hex representation of a 256 bit unit of data
	PowHash string `json:"powHash"`
	// Hex representation of a 256 bit unit of data
	MixHash string `json:"mixHash"`
}
type EthSubmitWorkResult struct {
	// Whether or not the provided solution is valid
	SolutionValid bool `json:"solutionValid"`
}
type SyncStatus struct {
	// Hex representation of the integer
	StartingBlock string `json:"startingBlock"`
	// Hex representation of the integer
	CurrentBlock string `json:"currentBlock"`
	// Hex representation of the integer
	HighestBlock string `json:"highestBlock"`
	// Hex representation of the integer
	KnownStates string `json:"knownStates"`
	// Hex representation of the integer
	PulledStates string `json:"pulledStates"`
}
type Syncing struct {
	// An object with sync status data
	SyncStatus
}
type EthSyncingResult struct {
	Syncing SyncStatus `json:"syncing"`
}
type EthUninstallFilterParams struct {
	// An identifier used to reference the filter.
	FilterId string `json:"filterId"`
}
type EthUninstallFilterResult struct {
	// Whether of not the filter was successfully uninstalled
	FilterUninstalledSuccess bool `json:"filterUninstalledSuccess"`
}
