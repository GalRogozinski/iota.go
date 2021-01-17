package iota_test

import (
	"fmt"
	. "github.com/iotaledger/iota.go"
	"github.com/stretchr/testify/assert"
	"testing"
)

//1. create simple message and pow it
//2. create message with 100 inputs and PoW

func TestSimpleMessage(t *testing.T) {
	signedTxPayload := oneInputOutputTransaction()
	
	message := Message{
		NetworkID: 1,
		Parent1:   [MessageIDLength]byte{},
		Parent2:   [MessageIDLength]byte{},
		Payload: signedTxPayload,
		Nonce:     0,
	}

	bytes, err := message.Serialize(DeSeriModeNoValidation)
	must(err)
	fmt.Println(bytes)
	assert.NotEmpty(t, bytes)
}

func Test100Message(t *testing.T) {
	inputs := createInputs(1)
	outputs := createOutputs(100)
	unlock_blocks := createUnlockBlocks(100)

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
		Payload: txPayload,
		Nonce:     0,
	}

	bytes, err := message.Serialize(DeSeriModeNoValidation)
	must(err)
	fmt.Printf("%x", bytes)
	assert.NotEmpty(t, message)
}

func createUnlockBlocks(n int) Serializables {
	var unlockBlocks Serializables
	for i := 0; n < 100; i++ {
		block := SignatureUnlockBlock {
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
		output := SigLockedSingleOutput {
			Address: func() Serializable {
				edAddr, _ := randEd25519Addr()
				return edAddr
			}(),
			Amount: uint64(i),
		}
		outputs = append(outputs, output)
		fmt.Println(outputs)
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