package iota_test

import (
	"context"
	"fmt"
	. "github.com/iotaledger/iota.go"
	. "github.com/iotaledger/iota.go/pow"
	"testing"
	"time"
)

const (
	workers     = 5
	targetScore = 4000.
)

var testWorker = New(workers)

func TestWorker_Small_Message_Mine(t *testing.T) {
	for i := 0; i < 100; i++ {
		msg := CreateNInputsMessage(1)
		start := time.Now()
		testWorker.Mine(context.Background(), msg[:len(msg)-NonceBytes], targetScore)
		duration := time.Since(start)
		fmt.Printf("%v", duration.Milliseconds())
	}
}

func TestWorker_Large_Message_Mine(t *testing.T) {
	for i := 0; i < 100; i++ {
		msg := CreateNInputsMessage(100)
		start := time.Now()
		testWorker.Mine(context.Background(), msg[:len(msg)-NonceBytes], targetScore)
		duration := time.Since(start)
		fmt.Printf("%v", duration.Milliseconds())
	}
}

func CreateNInputsMessage(n int) []byte {
	inputs := createInputs(n)
	outputs := createOutputs(1)
	unlock_blocks := createUnlockBlocks(1)

	txPayload := &Transaction{
		Essence: &TransactionEssence{
			Inputs:  inputs,
			Outputs: outputs,
			Payload: nil,
		},
		UnlockBlocks: unlock_blocks,
	}

	message := Message{
		NetworkID: 1,
		Parent1:   [MessageIDLength]byte{},
		Parent2:   [MessageIDLength]byte{},
		Payload:   txPayload,
		Nonce:     0,
	}

	bytes, err := message.Serialize(DeSeriModeNoValidation)
	must(err)
	return bytes
}

func createUnlockBlocks(n int) Serializables {
	var unlockBlocks Serializables
	for i := 0; i < n; i++ {
		block := SignatureUnlockBlock{
			Signature: func() Serializable {
				edSig, _ := randEd25519Signature()
				return edSig
			}(),
		}
		unlockBlocks = append(unlockBlocks, &block)
	}
	return unlockBlocks
}

func createOutputs(n int) Serializables {
	var outputs Serializables
	for i := 0; i < n; i++ {
		output := SigLockedSingleOutput{
			Address: func() Serializable {
				edAddr, _ := randEd25519Addr()
				return edAddr
			}(),
			Amount: uint64(i),
		}
		outputs = append(outputs, &output)
	}
	fmt.Println()
	return outputs
}

func createInputs(n int) Serializables {
	var inputs Serializables
	for i := 0; i < n; i++ {
		input := UTXOInput{
			TransactionID:          rand32ByteHash(),
			TransactionOutputIndex: uint16(i),
		}
		inputs = append(inputs, &input)
	}
	return inputs
}
