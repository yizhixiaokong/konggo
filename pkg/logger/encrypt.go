package logger

import (
	"bufio"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"io"
	"os"
)

var PwdKey = []byte("ZJG**#K$&DERSKDI")

// PKCS7 填充模式
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// 填充的反向操作，删除填充字符串
func PKCS7UnPadding(origData []byte) ([]byte, error) {
	length := len(origData)
	if length == 0 {
		return nil, errors.New("加密字符串错误！")
	} else {
		unpadding := int(origData[length-1])
		return origData[:(length - unpadding)], nil
	}
}

// 加密
func AesEcrypt(origData []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS7Padding(origData, blockSize)
	blocMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blocMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// 解密
func AesDeCrypt(cypted []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(cypted))
	blockMode.CryptBlocks(origData, cypted)
	origData, err = PKCS7UnPadding(origData)
	if err != nil {
		return nil, err
	}
	return origData, err
}

// 加密
func EnPwdCode(pwd []byte) (string, error) {
	//return string(pwd), nil
	result, err := AesEcrypt(pwd, PwdKey)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(result), err
}

// 解密
func DePwdCode(pwd string) ([]byte, error) {
	pwdByte, err := base64.StdEncoding.DecodeString(pwd)
	if err != nil {
		return nil, err
	}
	return AesDeCrypt(pwdByte, PwdKey)
}

// 解密日志文件
func DeLogFile(src string) error {
	fp, err := os.Open(src)
	defer fp.Close()
	if err != nil {
		return err
	}

	fw, err := os.Create("已解密" + src)
	defer fw.Close()
	if err != nil {
		return err
	}

	bufReader := bufio.NewReader(fp)
	bufWriter := bufio.NewWriter(fw)

	for {
		r, _, c := bufReader.ReadLine()
		rs := bytes.Split(r, []byte(" "))
		rx, _ := DePwdCode(string(rs[len(rs)-1]))

		if c == io.EOF {
			break
		}
		_, _ = bufWriter.Write(r[:len(r)-len(rs[len(rs)-1])])
		_, _ = bufWriter.Write(rx)
		_, _ = bufWriter.Write([]byte("\n"))
	}

	err = bufWriter.Flush()
	return err
}
