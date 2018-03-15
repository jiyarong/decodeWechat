package decodeWechat

import(
	"encoding/base64"
	"crypto/aes"
	"crypto/cipher"
)

type DecodeWechatEncrypted struct {
	AppId string
	SessionKey string
}

func (decoder *DecodeWechatEncrypted) DecryptData(encryptedData string, iv string) string {
	sessionKey, _ := base64.StdEncoding.DecodeString(decoder.SessionKey)
	_encryptedData, _ := base64.StdEncoding.DecodeString(encryptedData)
	_iv, _ := base64.StdEncoding.DecodeString(iv)

	block, _ := aes.NewCipher(sessionKey)
	mode := cipher.NewCBCDecrypter(block, _iv)
	mode.CryptBlocks(_encryptedData, _encryptedData)
	return string(_encryptedData)
}
