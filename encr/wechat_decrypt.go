package encr

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

// CBC 模式
//解密
/**
* rawData 原始加密数据
* key  密钥
* iv  向量
 */
// WxDecrypt 微信解密
func WxDecrypt(rawData, sessionKey, iv string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(rawData)
	if err != nil {
		return "", err
	}
	key, err := base64.StdEncoding.DecodeString(sessionKey)
	if err != nil {
		return "", err
	}
	i, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return "", err
	}
	dnData, err := aesCBCDecrypt(data, key, i)
	if err != nil {
		return "", err
	}
	return string(dnData), nil
}

// 解密
func aesCBCDecrypt(encryptData, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	blockSize := block.BlockSize()
	if len(encryptData) < blockSize {
		return nil, errors.New("cipherText too short")
	}
	if len(encryptData)%blockSize != 0 {
		return nil, errors.New("cipherText is not a multiple of the block size")
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(encryptData, encryptData)
	if len(encryptData) == 0 {
		return nil, errors.New("解密结果为空")
	}
	// 解填充
	encryptData = pKCS7UnPadding(encryptData)
	return encryptData, nil
}

//去除填充
func pKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}
