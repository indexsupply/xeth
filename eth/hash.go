package eth

import "golang.org/x/crypto/sha3"

func Keccak32(d []byte) [32]byte {
	return *(*[32]byte)(Keccak(d))
}

func Keccak(d []byte) []byte {
	k := sha3.NewLegacyKeccak256()
	k.Write(d)
	return k.Sum(nil)
}
