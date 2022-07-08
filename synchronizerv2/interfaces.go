package synchronizerv2

import (
	"context"
	"math/big"

	etherman "github.com/0xPolygonHermez/zkevm-node/ethermanv2"
	state "github.com/0xPolygonHermez/zkevm-node/statev2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/jackc/pgx/v4"
)

// ethermanInterface contains the methods required to interact with ethereum.
type ethermanInterface interface {
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
	GetRollupInfoByBlockRange(ctx context.Context, fromBlock uint64, toBlock *uint64) ([]etherman.Block, map[common.Hash][]etherman.Order, error)
	EthBlockByNumber(ctx context.Context, blockNumber uint64) (*types.Block, error)
	GetLatestBatchNumber() (uint64, error)
}

// stateInterface gathers the methods required to interact with the state.
type stateInterface interface {
	GetLastBlock(ctx context.Context, dbTx pgx.Tx) (*state.Block, error)
	AddGlobalExitRoot(ctx context.Context, exitRoot *state.GlobalExitRoot, dbTx pgx.Tx) error
	AddForcedBatch(ctx context.Context, forcedBatch *state.ForcedBatch, dbTx pgx.Tx) error
	AddBlock(ctx context.Context, block *state.Block, dbTx pgx.Tx) error
	Reset(ctx context.Context, blockNumber uint64, dbTx pgx.Tx) error
	GetPreviousBlock(ctx context.Context, offset uint64, dbTx pgx.Tx) (*state.Block, error)
	GetLastBatchNumber(ctx context.Context, dbTx pgx.Tx) (uint64, error)
	GetBatchByNumber(ctx context.Context, batchNumber uint64, dbTx pgx.Tx) (*state.Batch, error)
	ResetTrustedState(ctx context.Context, batchNumber uint64, dbTx pgx.Tx) error
	AddVirtualBatch(ctx context.Context, virtualBatch *state.VirtualBatch, dbTx pgx.Tx) error
	// GetNextForcedBatches returns the next forcedBatches in FIFO order
	GetNextForcedBatches(ctx context.Context, nextForcedBatches int, dbTx pgx.Tx) ([]state.ForcedBatch, error)
	AddBatchNumberInForcedBatch(ctx context.Context, forceBatchNumber, batchNumber uint64, dbTx pgx.Tx) error
	AddVerifiedBatch(ctx context.Context, verifiedBatch *state.VerifiedBatch, dbTx pgx.Tx) error

	ProcessAndStoreClosedBatch(ctx context.Context, processingCtx state.ProcessingContext, encodedTxs []byte, dbTx pgx.Tx) error

	BeginStateTransaction(ctx context.Context) (pgx.Tx, error)
	RollbackStateTransaction(ctx context.Context, dbTx pgx.Tx) error
	CommitStateTransaction(ctx context.Context, dbTx pgx.Tx) error
}