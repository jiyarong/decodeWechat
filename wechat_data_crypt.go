package decodeWechat

import(
	"encoding/base64"
	"crypto/aes"
	"crypto/cipher"
)

type DecodeWechatEncrypted struct {
	appId string
	sessionKey string
}

func (decoder *DecodeWechatEncrypted) decryptData(encryptedData string, iv string) string {
	sessionKey, _ := base64.StdEncoding.DecodeString(decoder.sessionKey)
	_encryptedData, _ := base64.StdEncoding.DecodeString(encryptedData)
	_iv, _ := base64.StdEncoding.DecodeString(iv)

	block, _ := aes.NewCipher(sessionKey)
	mode := cipher.NewCBCDecrypter(block, _iv)
	mode.CryptBlocks(_encryptedData, _encryptedData)
	return string(_encryptedData)
}
