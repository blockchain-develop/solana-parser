package solanaparser

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"os"
	"testing"
)

func TestTransaction_WhirlPool_ParseSwapV2(t *testing.T) {
	solClient := rpc.New(rpc.MainNetBeta_RPC)
	result, err := solClient.GetTransaction(
		context.Background(),
		solana.MustSignatureFromBase58("254kJ7VTZJ4tuU7YGrpD1quntbAxQuSR2kmUUgSmEKE7fNv3PgdT3jDZUkcUSe8vr4Fb5wum5RxQaUo1ZdVsW75T"),
		&rpc.GetTransactionOpts{
			Commitment:                     rpc.CommitmentConfirmed,
			MaxSupportedTransactionVersion: &rpc.MaxSupportedTransactionVersion1,
		})
	if err != nil {
		panic(err)
	}
	transaction, _ := result.Transaction.GetTransaction()
	transactionParsed := &rpc.TransactionParsed{
		Transaction: transaction,
		Meta:        result.Meta,
	}
	txRawJson, _ := json.MarshalIndent(transactionParsed, "", "    ")
	os.WriteFile(fmt.Sprintf("tx_raw.json"), txRawJson, 0644)
	tx := ParseTransaction(0, transaction, result.Meta)
	if tx == nil {
		panic("invalid transaction")
	}
	txJson, _ := json.MarshalIndent(tx, "", "    ")
	os.WriteFile(fmt.Sprintf("tx.json"), txJson, 0644)
}

func TestTransaction_WhirlPool_ParseSwapV2_2(t *testing.T) {
	solClient := rpc.New(rpc.MainNetBeta_RPC)
	result, err := solClient.GetTransaction(
		context.Background(),
		solana.MustSignatureFromBase58("5NmBwGfqWdQ95kpKNMCe5koGMMChUccP4V9ZK972uN23Mf7ts47xUmWTXVbFEjUN9qMM6iPH23wjdPqcGD6Xa8wf"),
		&rpc.GetTransactionOpts{
			Commitment:                     rpc.CommitmentConfirmed,
			MaxSupportedTransactionVersion: &rpc.MaxSupportedTransactionVersion1,
		})
	if err != nil {
		panic(err)
	}
	transaction, _ := result.Transaction.GetTransaction()
	transactionParsed := &rpc.TransactionParsed{
		Transaction: transaction,
		Meta:        result.Meta,
	}
	txRawJson, _ := json.MarshalIndent(transactionParsed, "", "    ")
	os.WriteFile(fmt.Sprintf("tx_raw.json"), txRawJson, 0644)
	tx := ParseTransaction(0, transaction, result.Meta)
	if tx == nil {
		panic("invalid transaction")
	}
	txJson, _ := json.MarshalIndent(tx, "", "    ")
	os.WriteFile(fmt.Sprintf("tx.json"), txJson, 0644)
}

func TestTransaction_WhirlPool_ParseSwap(t *testing.T) {
	solClient := rpc.New(rpc.MainNetBeta_RPC)
	result, err := solClient.GetTransaction(
		context.Background(),
		solana.MustSignatureFromBase58("wtWnUqobtQxVJCFFHC8gcsWnDZAeymHvb8HXmkHFGWtcWFiopsrgYs7NwxLpkki4829jKoou8nGRYQJGdJ1kFAh"),
		&rpc.GetTransactionOpts{
			Commitment:                     rpc.CommitmentConfirmed,
			MaxSupportedTransactionVersion: &rpc.MaxSupportedTransactionVersion1,
		})
	if err != nil {
		panic(err)
	}
	transaction, _ := result.Transaction.GetTransaction()
	transactionParsed := &rpc.TransactionParsed{
		Transaction: transaction,
		Meta:        result.Meta,
	}
	txRawJson, _ := json.MarshalIndent(transactionParsed, "", "    ")
	os.WriteFile(fmt.Sprintf("tx_raw.json"), txRawJson, 0644)
	tx := ParseTransaction(0, transaction, result.Meta)
	if tx == nil {
		panic("invalid transaction")
	}
	txJson, _ := json.MarshalIndent(tx, "", "    ")
	os.WriteFile(fmt.Sprintf("tx.json"), txJson, 0644)
}

func TestTransaction_WhirlPool_ParseIncreaseLiquidityV2(t *testing.T) {
	solClient := rpc.New(rpc.MainNetBeta_RPC)
	result, err := solClient.GetTransaction(
		context.Background(),
		solana.MustSignatureFromBase58("51TV1rEVJGgrJG7yNrdUjfAhQYgJzFqKYbmGXmpqtDdLE8GgApd6KWXJcrmKKL81aE7qfay3NwXYJBRqQCYzJKxm"),
		&rpc.GetTransactionOpts{
			Commitment:                     rpc.CommitmentConfirmed,
			MaxSupportedTransactionVersion: &rpc.MaxSupportedTransactionVersion1,
		})
	if err != nil {
		panic(err)
	}
	transaction, _ := result.Transaction.GetTransaction()
	transactionParsed := &rpc.TransactionParsed{
		Transaction: transaction,
		Meta:        result.Meta,
	}
	txRawJson, _ := json.MarshalIndent(transactionParsed, "", "    ")
	os.WriteFile(fmt.Sprintf("tx_raw.json"), txRawJson, 0644)
	tx := ParseTransaction(0, transaction, result.Meta)
	if tx == nil {
		panic("invalid transaction")
	}
	txJson, _ := json.MarshalIndent(tx, "", "    ")
	os.WriteFile(fmt.Sprintf("tx.json"), txJson, 0644)
}

func TestTransaction_WhirlPool_ParseTwoHopSwapV2(t *testing.T) {
	solClient := rpc.New(rpc.MainNetBeta_RPC)
	result, err := solClient.GetTransaction(
		context.Background(),
		solana.MustSignatureFromBase58("42357Cdt3ASZnVFvbt7a6PsVdsSwyiUGnye57h3f2RLBGboeJbustrNr47MD1krCvLm3Ea1nwYUCfEvSS7KszFDF"),
		&rpc.GetTransactionOpts{
			Commitment:                     rpc.CommitmentConfirmed,
			MaxSupportedTransactionVersion: &rpc.MaxSupportedTransactionVersion1,
		})
	if err != nil {
		panic(err)
	}
	transaction, _ := result.Transaction.GetTransaction()
	transactionParsed := &rpc.TransactionParsed{
		Transaction: transaction,
		Meta:        result.Meta,
	}
	txRawJson, _ := json.MarshalIndent(transactionParsed, "", "    ")
	os.WriteFile(fmt.Sprintf("tx_raw.json"), txRawJson, 0644)
	tx := ParseTransaction(0, transaction, result.Meta)
	if tx == nil {
		panic("invalid transaction")
	}
	txJson, _ := json.MarshalIndent(tx, "", "    ")
	os.WriteFile(fmt.Sprintf("tx.json"), txJson, 0644)
}

func TestTransaction_WhirlPool_ParseDecreaseLiquidityV2(t *testing.T) {
	solClient := rpc.New(rpc.MainNetBeta_RPC)
	result, err := solClient.GetTransaction(
		context.Background(),
		solana.MustSignatureFromBase58("5TH5HaYRDRacVS8YRqvTZjPwvfn19AiN6qKdyEBnroCP8yX1QHNwxDXoaTcQwG4Pfu2Wm4L3kHXPiLtMCSHPe8Yd"),
		&rpc.GetTransactionOpts{
			Commitment:                     rpc.CommitmentConfirmed,
			MaxSupportedTransactionVersion: &rpc.MaxSupportedTransactionVersion1,
		})
	if err != nil {
		panic(err)
	}
	transaction, _ := result.Transaction.GetTransaction()
	transactionParsed := &rpc.TransactionParsed{
		Transaction: transaction,
		Meta:        result.Meta,
	}
	txRawJson, _ := json.MarshalIndent(transactionParsed, "", "    ")
	os.WriteFile(fmt.Sprintf("tx_raw.json"), txRawJson, 0644)
	tx := ParseTransaction(0, transaction, result.Meta)
	if tx == nil {
		panic("invalid transaction")
	}
	txJson, _ := json.MarshalIndent(tx, "", "    ")
	os.WriteFile(fmt.Sprintf("tx.json"), txJson, 0644)
}

func TestTransaction_WhirlPool_ParseInitializePoolV2(t *testing.T) {
	solClient := rpc.New(rpc.MainNetBeta_RPC)
	result, err := solClient.GetTransaction(
		context.Background(),
		solana.MustSignatureFromBase58("43TbXXDGHvHSTr68dz9Yx79bV9s2tSZxSndTJmfCVc2jus7LL7MFctdqoUkHPnEz2qRojbmLfD8Z7WVfFv7CjmM9"),
		&rpc.GetTransactionOpts{
			Commitment:                     rpc.CommitmentConfirmed,
			MaxSupportedTransactionVersion: &rpc.MaxSupportedTransactionVersion1,
		})
	if err != nil {
		panic(err)
	}
	transaction, _ := result.Transaction.GetTransaction()
	transactionParsed := &rpc.TransactionParsed{
		Transaction: transaction,
		Meta:        result.Meta,
	}
	txRawJson, _ := json.MarshalIndent(transactionParsed, "", "    ")
	os.WriteFile(fmt.Sprintf("tx_raw.json"), txRawJson, 0644)
	tx := ParseTransaction(0, transaction, result.Meta)
	if tx == nil {
		panic("invalid transaction")
	}
	txJson, _ := json.MarshalIndent(tx, "", "    ")
	os.WriteFile(fmt.Sprintf("tx.json"), txJson, 0644)
}
