package vm

import (
	"math/big"
	"github.com/dipperin/dipperin-core/third-party/rpc"
	"github.com/dipperin/dipperin-core/common"
	"strings"
	"github.com/dipperin/dipperin-core/third-party/log"
	"github.com/dipperin/dipperin-core/core/rpc-interface"
	"github.com/dipperin/dipperin-core/core/vm/model"
	"github.com/dipperin/dipperin-core/common/consts"
	"io/ioutil"
	"github.com/dipperin/dipperin-core/tests/node-cluster"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/ethereum/go-ethereum/rlp"
	"path/filepath"
	"github.com/dipperin/dipperin-core/common/util"
	"fmt"
)

var (
	AbiPath  = filepath.Join(util.HomeDir(), "go/src/github.com/dipperin/dipperin-core/core/vm/event/event.cpp.abi.json")
	WASMPath = filepath.Join(util.HomeDir(), "go/src/github.com/dipperin/dipperin-core/core/vm/event/event.wasm")
)

func LogTestPrint(function, msg string, ctx ...interface{}) {
	printMsg := "[~wjw~" + function + "]" + msg
	log.Info(printMsg, ctx...)
}

func GetRpcTXMethod(methodName string) string {
	return "dipperin_" + strings.ToLower(methodName[0:1]) + methodName[1:]
}

func SendTransaction(client *rpc.Client, from, to common.Address, value, fee *big.Int, data []byte) (common.Hash, error) {
	var resp common.Hash
	if err := client.Call(&resp, GetRpcTXMethod("SendTransaction"), from, to, value, fee, data, nil); err != nil {
		LogTestPrint("Test", "SendTransaction failed", "err", err)
		return common.Hash{}, err
	}
	LogTestPrint("Test", "SendTransaction Successful", "txId", resp.Hex())
	return resp, nil
}

func SendTransactionContract(client *rpc.Client, from, to common.Address, value, gasLimit, gasPrice *big.Int, data []byte) (common.Hash, error) {
	var resp common.Hash
	if err := client.Call(&resp, GetRpcTXMethod("SendTransactionContract"), from, to, value, gasLimit, gasPrice, data, nil); err != nil {
		LogTestPrint("Test", "SendContract failed", "err", err)
		return common.Hash{}, err
	}
	LogTestPrint("Test", "SendContract Successful", "txId", resp.Hex())
	return resp, nil
}

func Transaction(client *rpc.Client, hash common.Hash) (bool, uint64) {
	var resp *rpc_interface.TransactionResp
	if err := client.Call(&resp, GetRpcTXMethod("Transaction"), hash); err != nil {
		return false, 0
	}
	if resp.BlockNumber == 0 {
		return false, 0
	}
	return true, resp.BlockNumber
}

func GetReceiptByTxHash(client *rpc.Client, hash common.Hash) *model.Receipt {
	var resp *model.Receipt
	if err := client.Call(&resp, GetRpcTXMethod("GetReceiptByTxHash"), hash); err != nil {
		LogTestPrint("Test", "call GetReceiptByTxHash failed", "err", err)
		return nil
	}
	return resp
}

func GetReceiptsByBlockNum(client *rpc.Client, num uint64) model.Receipts {
	var resp model.Receipts
	if err := client.Call(&resp, GetRpcTXMethod("GetReceiptsByBlockNum"), num); err != nil {
		LogTestPrint("Test", "call GetReceiptsByBlockNum failed", "err", err)
		return nil
	}
	return resp
}

func GetContractAddressByTxHash(client *rpc.Client, hash common.Hash) common.Address {
	var resp common.Address
	if err := client.Call(&resp, GetRpcTXMethod("GetContractAddressByTxHash"), hash); err != nil {
		LogTestPrint("Test", "call GetContractAddressByTxHash failed", "err", err)
		return common.Address{}
	}
	return resp
}

func GetBlockByNumber(client *rpc.Client, num uint64) rpc_interface.BlockResp {
	var respBlock rpc_interface.BlockResp
	if err := client.Call(&respBlock, GetRpcTXMethod("GetBlockByNumber"), num); err != nil {
		LogTestPrint("Test", "call GetBlockByNumber failed", "err", err)
		return rpc_interface.BlockResp{}
	}
	return respBlock
}

func CreateContract(t *testing.T, cluster *node_cluster.NodeCluster, nodeName string, times int) []common.Hash {
	client := cluster.NodeClient[nodeName]
	from, err := cluster.GetNodeMainAddress(nodeName)
	LogTestPrint("Test", "From", "addr", from.Hex())
	assert.NoError(t, err)

	to := common.HexToAddress(common.AddressContractCreate)
	value := big.NewInt(100)
	gasLimit := big.NewInt(2 * consts.DIP)
	gasPrice := big.NewInt(1)
	//txFee := big.NewInt(0).Mul(gasLimit, gasPrice)

	abiBytes, err := ioutil.ReadFile(AbiPath)
	assert.NoError(t, err)
	WASMBytes, err := ioutil.ReadFile(WASMPath)
	assert.NoError(t, err)
	ExtraData, err := rlp.EncodeToBytes([]interface{}{WASMBytes, abiBytes})
	assert.NoError(t, err)

	var txHashList []common.Hash
	for i := 0; i < times; i++ {
		txHash, innerErr := SendTransactionContract(client, from, to, value, gasLimit, gasPrice, ExtraData)
		assert.NoError(t, innerErr)
		txHashList = append(txHashList, txHash)

		/*		txHash, innerErr = SendTransaction(client, from, factory.AliceAddrV, value, txFee, nil)
				assert.NoError(t, innerErr)
				txHashList = append(txHashList, txHash)*/
	}
	return txHashList
}

func CallContract(t *testing.T, cluster *node_cluster.NodeCluster, nodeName string, addrList []common.Address) []common.Hash {
	client := cluster.NodeClient[nodeName]
	from, err := cluster.GetNodeMainAddress(nodeName)
	LogTestPrint("Test", "From", "addr", from.Hex())
	assert.NoError(t, err)

	value := big.NewInt(100)
	gasLimit := big.NewInt(2 * consts.DIP)
	gasPrice := big.NewInt(1)

	var txHashList []common.Hash
	for i := 0; i < len(addrList); i++ {
		input := genInput(t, "hello", fmt.Sprintf("Event,%v", 100*i))
		txHash, innerErr := SendTransactionContract(client, from, addrList[i], value, gasLimit, gasPrice, input)
		assert.NoError(t, innerErr)
		txHashList = append(txHashList, txHash)
	}
	return txHashList
}

func genInput(t *testing.T, funcName, param string) []byte {
	input := []interface{}{
		funcName,
		param,
	}

	result, err := rlp.EncodeToBytes(input)
	assert.NoError(t, err)
	return result
}