// Package aes implements a codec to encrypt/decrypt content with aes/pkcs7padding algorithm.
// The password length decide the block size of aes algorithm, it could be 128/192/256 bits
// The length of iv should always be 128 bit, otherwise it will panic. It is recommended that the iv should be different for each encryption/decryption pair

package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

type AESCodec struct {
	Password []byte
	block    cipher.Block
}

func NewAESCodec(password []byte) (*AESCodec, error) {

	block, err := aes.NewCipher(password)
	if err != nil {
		return nil, err
	}

	codec := new(AESCodec)
	codec.Password = password
	codec.block = block

	return codec, nil
}

func (codec *AESCodec) ChangePassword(password string) error {
	pd := []byte(password)

	block, err := aes.NewCipher(pd)
	if err != nil {
		return err
	}

	codec.Password = pd
	codec.block = block

	return nil
}

func (codec *AESCodec) Encrypt(src, iv []byte) []byte {

	encrypter := cipher.NewCBCEncrypter(codec.block, iv)
	src = PKCS7Padding(src, codec.block.BlockSize())

	dst := make([]byte, len(src))
	encrypter.CryptBlocks(dst, src)

	return dst
}

func (codec *AESCodec) Decrypt(src, iv []byte) []byte {
	decrypter := cipher.NewCBCDecrypter(codec.block, iv)

	dst := make([]byte, len(src))
	decrypter.CryptBlocks(dst, src)
	dst = PKCS7UnPadding(dst, codec.block.BlockSize())

	return dst
}

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(plantText []byte, blockSize int) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	return plantText[:(length - unpadding)]
}
