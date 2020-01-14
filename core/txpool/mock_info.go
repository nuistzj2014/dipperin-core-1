// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dipperin/dipperin-core/core/txpool (interfaces: BlockChain)

// Package tx_pool is a generated GoMock package.
package txpool

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/dipperin/dipperin-core/common"
	"github.com/dipperin/dipperin-core/core/chain-config"
	"github.com/dipperin/dipperin-core/core/chain/state-processor"
	"github.com/dipperin/dipperin-core/core/economy-model"
	"github.com/dipperin/dipperin-core/core/model"
	"github.com/dipperin/dipperin-core/third-party/crypto"
	"github.com/dipperin/dipperin-core/third-party/crypto/cs-crypto"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/golang/mock/gomock"
	"golang.org/x/crypto/sha3"
	"hash"
	"math/big"
	"os"
	"reflect"
	"strconv"
	"sync"
)

var testPriv1 = "289c2857d4598e37fb9647507e47a309d6133539bf21a8b9cb6df88fd5232031"
var testPriv2 = "289c2857d4598e37fb9647507e47a309d6133539bf21a8b9cb6df88fd5232032"
var testPriv3 = "289c2857d4598e37fb9647507e47a309d6133539bf21a8b9cb6df88fd5232033"
var tx1hash = common.HexToHash("0x528131488f97c6314b2fa0dff404f1037067e787b65cb244d79c7ecea007c0d5")
var tx2hash = common.HexToHash("0x0aedd7a6779339cc44fe1e51cdf42b4bf3a557d52e646390e6d6bf6d489a5de3")
var path = "./transaction.out"

type transactions []model.AbstractTransaction

var testTxFee = economy_model.GetMinimumTxFee(200)
var threshold = new(big.Int).Div(new(big.Int).Mul(testTxFee, big.NewInt(100+int64(DefaultTxPoolConfig.FeeBump))), big.NewInt(100))
var testRoot = "0x54bbe8ffddc42dd501ab37438c2496d1d3be51d9c562531d56b48ea3bea66708"
var testTxPoolConfig TxPoolConfig
var ms = model.NewSigner(big.NewInt(1))

// MockBlockChain is a mock of BlockChain interface
type MockBlockChain struct {
	ctrl     *gomock.Controller
	recorder *MockBlockChainMockRecorder
}

// MockBlockChainMockRecorder is the mock recorder for MockBlockChain
type MockBlockChainMockRecorder struct {
	mock *MockBlockChain
}

// NewMockBlockChain creates a new mock instance
func NewMockBlockChain(ctrl *gomock.Controller) *MockBlockChain {
	mock := &MockBlockChain{ctrl: ctrl}
	mock.recorder = &MockBlockChainMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBlockChain) EXPECT() *MockBlockChainMockRecorder {
	return m.recorder
}

// CurrentBlock mocks base method
func (m *MockBlockChain) CurrentBlock() model.AbstractBlock {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentBlock")
	ret0, _ := ret[0].(model.AbstractBlock)
	return ret0
}

// CurrentBlock indicates an expected call of CurrentBlock
func (mr *MockBlockChainMockRecorder) CurrentBlock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentBlock", reflect.TypeOf((*MockBlockChain)(nil).CurrentBlock))
}

// GetBlockByNumber mocks base method
func (m *MockBlockChain) GetBlockByNumber(arg0 uint64) model.AbstractBlock {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockByNumber", arg0)
	ret0, _ := ret[0].(model.AbstractBlock)
	return ret0
}

// GetBlockByNumber indicates an expected call of GetBlockByNumber
func (mr *MockBlockChainMockRecorder) GetBlockByNumber(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockByNumber", reflect.TypeOf((*MockBlockChain)(nil).GetBlockByNumber), arg0)
}

// StateAtByStateRoot mocks base method
func (m *MockBlockChain) StateAtByStateRoot(arg0 common.Hash) (*state_processor.AccountStateDB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateAtByStateRoot", arg0)
	ret0, _ := ret[0].(*state_processor.AccountStateDB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateAtByStateRoot indicates an expected call of StateAtByStateRoot
func (mr *MockBlockChainMockRecorder) StateAtByStateRoot(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateAtByStateRoot", reflect.TypeOf((*MockBlockChain)(nil).StateAtByStateRoot), arg0)
}

func createKey() (*ecdsa.PrivateKey, *ecdsa.PrivateKey, *ecdsa.PrivateKey) {
	key1, err1 := crypto.HexToECDSA(testPriv1)
	key2, err2 := crypto.HexToECDSA(testPriv2)
	key3, err3 := crypto.HexToECDSA(testPriv3)
	if err1 != nil || err2 != nil || err3 != nil {
		return nil, nil, nil
	}
	return key1, key2, key3
}

func createKeyBatch(num int) (keys []*ecdsa.PrivateKey) {
	keyBase := []byte(testPriv1)
	baseLen := len(keyBase)

	pat := []byte{}
	patLen := len([]byte(strconv.Itoa(num)))
	for i := 0; i < patLen; i++ {
		pat = append(pat, '0')
	}

	keyBase = append(keyBase[:baseLen-patLen], pat...)

	for i := 0; i < num; i++ {
		s := []byte(strconv.Itoa(i))
		slice := append(keyBase[:baseLen-len(s)], s...)
		//fmt.Println(i, "=", string(slice))
		key, err := crypto.HexToECDSA(string(slice))
		if err != nil {
			return nil
		}
		keys = append(keys, key)
	}

	return
}

func createTxList(n int) []*model.Transaction {
	keyAlice, keyBob, _ := createKey()
	ms := model.NewSigner(big.NewInt(1))

	bob := cs_crypto.GetNormalAddress(keyBob.PublicKey)

	var res []*model.Transaction
	for i := 0; i < n; i++ {
		temptx := model.NewTransaction(uint64(i+1), bob, big.NewInt(int64(i)), big.NewInt(0).Mul(big.NewInt(int64(i)), model.TestGasPrice), model.TestGasLimit, []byte{})
		temptx.SignTx(keyAlice, ms)
		res = append(res, temptx)
	}
	return res
}

func createTxListWithFee(n int) []*model.Transaction {
	keyAlice, keyBob, _ := createKey()
	ms := model.NewSigner(big.NewInt(1))

	bob := cs_crypto.GetNormalAddress(keyBob.PublicKey)

	var res []*model.Transaction
	for i := 0; i < n; i++ {
		temptx := model.NewTransaction(uint64(i+1), bob, big.NewInt(int64(i)), model.TestGasPrice, model.TestGasLimit, []byte{})

		temptx.SignTx(keyAlice, ms)
		res = append(res, temptx)
	}
	return res
}

// generate n transactions list file on the specified path
func createTxListFile(n int, path string) {
	txs := createTxList(n)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Create(path)
	} else {
		os.Remove(path)
	}

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, tx := range txs {
		err := rlp.Encode(f, tx)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// convert tx list to map of address, todo: this should be optimized in the future
func txListToAddrMap(txs []model.AbstractTransaction) map[common.Address][]model.AbstractTransaction {
	_, keyBob, _ := createKey()
	bob := cs_crypto.GetNormalAddress(keyBob.PublicKey)
	m := make(map[common.Address][]model.AbstractTransaction)

	for _, tx := range txs {
		m[bob] = append(m[bob], tx)
	}

	return m
}

type fakePool struct {
	txs []model.AbstractTransaction
}

func newFakePool() *fakePool {
	return &fakePool{
		txs: make([]model.AbstractTransaction, 1),
	}
}

func (p *fakePool) Add(txs []model.AbstractTransaction) []error {
	p.txs = []model.AbstractTransaction{}

	res := make([]error, len(txs))

	for i, txs := range txs {
		p.txs = append(p.txs, txs)
		res[i] = nil
	}

	return res
}

type MockPool struct {
	all     *txLookup
	feeList *txFeeList
}

func init() {
	testTxPoolConfig = DefaultTxPoolConfig
	testTxPoolConfig.NoLocals = true
	testTxPoolConfig.GlobalSlots = 4096
	testTxPoolConfig.AccountSlots = 16
	testTxPoolConfig.GlobalQueue = 4096
	testTxPoolConfig.AccountQueue = 512
	testTxPoolConfig.Journal = "./locals.out"
}

type testBlockChain struct {
	statedb *state_processor.AccountStateDB
}

func (bc *testBlockChain) CurrentBlock() model.AbstractBlock {
	header := model.NewHeader(0, 0, common.Hash{}, common.Hash{}, common.Difficulty{}, big.NewInt(0), common.Address{}, common.BlockNonce{})
	return model.NewBlock(header, nil, nil)
}

func (bc *testBlockChain) GetBlockByNumber(number uint64) model.AbstractBlock {
	return bc.CurrentBlock()
}

func (bc *testBlockChain) StateAtByStateRoot(root common.Hash) (*state_processor.AccountStateDB, error) {
	return bc.statedb, nil
}

func transaction(nonce uint64, to common.Address, amount *big.Int, gasPrice *big.Int, gasLimit uint64, key *ecdsa.PrivateKey) model.AbstractTransaction {

	uTx := model.NewTransaction(nonce, to, amount, gasPrice, gasLimit, nil)
	tx, _ := uTx.SignTx(key, ms)
	return tx
}

func createTestStateDB() (ethdb.Database, common.Hash) {
	db := ethdb.NewMemDatabase()
	tdb := state_processor.NewStateStorageWithCache(db)
	teststatedb, _ := state_processor.NewAccountStateDB(common.Hash{}, tdb)

	key1, key2, key3 := createKey()
	aliceAddr := cs_crypto.GetNormalAddress(key1.PublicKey)
	bobAddr := cs_crypto.GetNormalAddress(key2.PublicKey)
	chalieAddr := cs_crypto.GetNormalAddress(key3.PublicKey)
	teststatedb.NewAccountState(aliceAddr)
	teststatedb.NewAccountState(bobAddr)
	teststatedb.NewAccountState(chalieAddr)
	teststatedb.SetNonce(aliceAddr, uint64(20))
	teststatedb.SetNonce(bobAddr, uint64(30))
	teststatedb.SetNonce(chalieAddr, uint64(30))
	teststatedb.SetBalance(aliceAddr, big.NewInt(8400003000))
	teststatedb.SetBalance(bobAddr, big.NewInt(8400003000))
	teststatedb.SetBalance(chalieAddr, big.NewInt(8400003000))
	root, _ := teststatedb.Commit()
	tdb.TrieDB().Commit(root, false)
	return db, root
}

func createTestAddrs(num int) ([]common.Address, []*ecdsa.PrivateKey) {
	keys := createKeyBatch(num)
	addrs := []common.Address{}
	for i := 0; i < num; i++ {
		addr := cs_crypto.GetNormalAddress(keys[i].PublicKey)
		addrs = append(addrs, addr)
	}
	return addrs, keys
}

func createTestStateDBWithBatch(num int) (ethdb.Database, common.Hash) {
	db := ethdb.NewMemDatabase()
	tdb := state_processor.NewStateStorageWithCache(db)
	teststatedb, _ := state_processor.NewAccountStateDB(common.Hash{}, tdb)

	addrs, _ := createTestAddrs(num)
	for i := 0; i < num; i++ {
		teststatedb.NewAccountState(addrs[i])
		teststatedb.SetNonce(addrs[i], uint64(20))
		teststatedb.SetBalance(addrs[i], big.NewInt(1000000))
	}
	root, _ := teststatedb.Commit()
	tdb.TrieDB().Commit(root, false)
	return db, root
}

func setupTxPoolBatch(num int) *TxPool {
	db, root := createTestStateDBWithBatch(num)
	teststatedb, _ := state_processor.NewAccountStateDB(root, state_processor.NewStateStorageWithCache(db))

	blockchain := &testBlockChain{statedb: teststatedb}

	pool := NewTxPool(testTxPoolConfig, chain_config.ChainConfig{ChainId: big.NewInt(1)}, blockchain)

	pool.signer = ms

	return pool
}

func setupTxPool() *TxPool {
	db, root := createTestStateDB()
	teststatedb, _ := state_processor.NewAccountStateDB(root, state_processor.NewStateStorageWithCache(db))
	//con := newFakeValidator()
	blockchain := &testBlockChain{statedb: teststatedb}

	pool := NewTxPool(testTxPoolConfig, chain_config.ChainConfig{ChainId: big.NewInt(1)}, blockchain)

	pool.signer = ms

	return pool
}

type txHelper struct {
	cacher *model.TxCacher
	wg     sync.WaitGroup
}

func (th *txHelper) help_TxRecover(txs []model.AbstractTransaction) {
	th.cacher.TxRecover(txs)
	th.wg.Done()
}

func rlpHash(x interface{}) (h common.Hash, err error) {
	hw := sha3.NewLegacyKeccak256()
	err = rlp.Encode(hw, x)
	if err != nil {
		return
	}
	hw.Sum(h[:0])
	return
}

func rlpHashNew(hw hash.Hash, data []byte) (h common.Hash, err error) {
	hw.Write(data)
	hw.Sum(h[:0])
	return
}