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

func TestTransaction_StableSwap_ParseSwap(t *testing.T) {
	solClient := rpc.New(rpc.MainNetBeta_RPC)
	result, err := solClient.GetTransaction(
		context.Background(),
		solana.MustSignatureFromBase58("oatTsPD1GHEs9gWXdRbmaZf6EJK44qY7awqptR2WdrjH3qxBKHJYG11Yv9RX3bhciG73CYhGrFGoFmboRbxDo18"),
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

func TestTransaction_StableSwap_ParseSwap_2(t *testing.T) {
	solClient := rpc.New(rpc.MainNetBeta_RPC)
	result, err := solClient.GetTransaction(
		context.Background(),
		solana.MustSignatureFromBase58("4Q5SjR3c3Xbpf9xaGDPGr3zJCsHgsNLtFgaq2G4Xkij3p1LzjfF1y9CjdnjPm24wUhHuymfkKwfpmNMig4KZMPQa"),
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

func TestTransaction_StableSwap_ParseSwap_3(t *testing.T) {
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

func TestTransaction_StableSwap_ParseSwapV2(t *testing.T) {
	solClient := rpc.New(rpc.MainNetBeta_RPC)
	result, err := solClient.GetTransaction(
		context.Background(),
		solana.MustSignatureFromBase58("ESeR3usVhdFoj36RKSbt2rYHDJLruCAEJQaAhTqQhor5pHWGihGJytQVWx8BN14VqCnCnLvQAzUfUQQrZ2whGPp"),
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

func TestTransaction_StableSwap_ParseWithdraw(t *testing.T) {
	solClient := rpc.New(rpc.MainNetBeta_RPC)
	result, err := solClient.GetTransaction(
		context.Background(),
		solana.MustSignatureFromBase58("bRfDDa4iL8KvdxZDXyM8Jx8CGbaMFdknu7dYbvHirLwA17NKfwu2AcA3ssa1dmr5R4bNxNYotTdEpTUhGRPwTX9"),
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

func TestTransaction_StableSwap_ParseDeposit(t *testing.T) {
	solClient := rpc.New(rpc.MainNetBeta_RPC)
	result, err := solClient.GetTransaction(
		context.Background(),
		solana.MustSignatureFromBase58("2yMqYtjBEz5qxjcfT53zgh1KjHx8hJ6v6R78kGgKi8P9DefpxytoKoMEEnga83jiVcy8nYKYetor9DZd3yZMk7vt"),
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
