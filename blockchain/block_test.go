package blockchain


import "bytes"
import "testing"
import "encoding/hex"

import "github.com/deroproject/derosuite/config"
import "github.com/deroproject/derosuite/crypto"

func Test_Genesis_block_serdes(t *testing.T) {

	mainnet_genesis_block_hex := "010000000000000000000000000000000000000000000000000000000000000000000010270000023c01ff0001ffffffffffff07020bf6522f9152fa26cd1fc5c022b1a9e13dab697f3acf4b4d0ca6950a867a194321011d92826d0656958865a035264725799f39f6988faa97d532f972895de849496d0000"

	mainnet_genesis_block, _ := hex.DecodeString(mainnet_genesis_block_hex)

	var bl Block
	err := bl.Deserialize(mainnet_genesis_block)

	if err != nil {
		t.Errorf("Deserialization test failed for NULL block err %s\n", err)
	}

	// test the block serializer and deserializer whether it gives the same
	serialized := bl.Serialize()

	if !bytes.Equal(serialized, mainnet_genesis_block) {
		t.Errorf("Serialization test failed for Genesis block %X\n", serialized)
	}

	// calculate POW hash
	powhash := bl.GetPoWHash()
	if powhash != crypto.Hash([32]byte{0xa7, 0x3b, 0xd3, 0x7a, 0xba, 0x34, 0x54, 0x77, 0x6b, 0x40, 0x73, 0x38, 0x54, 0xa8, 0x34, 0x9f, 0xe6, 0x35, 0x9e, 0xb2, 0xc9, 0x1d, 0x93, 0xbc, 0x72, 0x7c, 0x69, 0x43, 0x1c, 0x1d, 0x1f, 0x95}) {
		t.Errorf("genesis block POW failed %x\n", powhash[:])
	}

	// test block id
	if bl.GetHash() != config.Mainnet.Genesis_Block_Hash {
		t.Error("genesis block ID failed \n")
	}

}
