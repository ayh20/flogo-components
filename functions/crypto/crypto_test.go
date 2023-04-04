package crypto

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var flogoEncrypt = &fnEncrypt{}

var flogoDecrypt = &fnDecrypt{}

var flogoEncryptRsa = &fnEncryptRsa{}

var flogoDecryptRsa = &fnDecryptRsa{}

func TestEncryptedDecryptedTextValue(t *testing.T) {
	key := []byte("AES256Key-32Characters1234567890")
	plaintext := []byte("example plain text")

	ciphertextInterface, err := flogoEncrypt.Eval(key, plaintext)

	var encryptedText []byte = ciphertextInterface.([]byte)

	assert.Nil(t, err)
	assert.NotNil(t, ciphertextInterface)
	assert.NotNil(t, encryptedText)

	// Decrypt same text
	plaintextInterface, err := flogoDecrypt.Eval(key, encryptedText)
	var decryptedText []byte = plaintextInterface.([]byte)

	assert.Nil(t, err)
	assert.NotNil(t, ciphertextInterface)
	assert.NotNil(t, decryptedText)
	assert.Equal(t, plaintext, decryptedText)
}

func TestFailingEncryptedDecryptedTextValue(t *testing.T) {
	key := []byte("AES256Key-32Characters1234567890")
	key2 := []byte("AES256Key2-32Characters123456789")
	plaintext := []byte("example plain text")

	ciphertextInterface, err := flogoEncrypt.Eval(key, plaintext)

	var encryptedText []byte = ciphertextInterface.([]byte)

	assert.Nil(t, err)
	assert.NotNil(t, ciphertextInterface)
	assert.NotNil(t, encryptedText)

	// Decrypt same text
	plaintextInterface, err := flogoDecrypt.Eval(key2, encryptedText)
	var decryptedText []byte = plaintextInterface.([]byte)

	assert.NotNil(t, err)
	assert.NotNil(t, ciphertextInterface)
	assert.Nil(t, decryptedText)
	assert.NotEqual(t, plaintext, decryptedText)
}

func TestEncryptedDecryptedTextValueRsa(t *testing.T) {
	privkey := []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEA4f5wg5l2hKsTeNem/V41fGnJm6gOdrj8ym3rFkEU/wT8RDtnSgFEZOQpHEgQ7JL38xUfU0Y3g6aYw9QT0hJ7mCpz9Er5qLaMXJwZxzHzAahlfA0icqabvJOMvQtzD6uQv6wPEyZtDTWiQi9AXwBpHssPnpYGIn20ZZuNlX2BrClciHhCPUIIZOQn/MmqTD31jSyjoQoV7MhhMTATKJx2XrHhR+1DcKJzQBSTAGnpYVaqpsARap+nwRipr3nUTuxyGohBTSmjJ2usSeQXHI3bODIRe1AuTyHceAbewn8b462yEWKARdpd9AjQW5SIVPfdsz5B6GlYQ5LdYKtznTuy7wIDAQABAoIBAQCwia1k7+2oZ2d3n6agCAbqIE1QXfCmh41ZqJHbOY3oRQG3X1wpcGH4Gk+O+zDVTV2JszdcOt7E5dAyMaomETAhRxB7hlIOnEN7WKm+dGNrKRvV0wDU5ReFMRHg31/Lnu8c+5BvGjZX+ky9POIhFFYJqwCRlopGSUIxmVj5rSgtzk3iWOQXr+ah1bjEXvlxDOWkHN6YfpV5ThdEKdBIPGEVqa63r9n2h+qazKrtiRqJqGnOrHzOECYbRFYhexsNFz7YT02xdfSHn7gMIvabDDP/Qp0PjE1jdouiMaFHYnLBbgvlnZW9yuVf/rpXTUq/njxIXMmvmEyyvSDnFcFikB8pAoGBAPF77hK4m3/rdGT7X8a/gwvZ2R121aBcdPwEaUhvj/36dx596zvYmEOjrWfZhF083/nYWE2kVquj2wjs+otCLfifEEgXcVPTnEOPO9Zg3uNSL0nNQghjFuD3iGLTUBCtM66oTe0jLSslHe8gLGEQqyMzHOzYxNqibxcOZIe8Qt0NAoGBAO+UI5+XWjWEgDmvyC3TrOSf/KCGjtu0TSv30ipv27bDLMrpvPmD/5lpptTFwcxvVhCs2b+chCjlghFSWFbBULBrfci2FtliClOVMYrlNBdUSJhf3aYSG2Doe6Bgt1n2CpNn/iu37Y3NfemZBJA7hNl4dYe+f+uzM87cdQ214+jrAoGAXA0XxX8ll2+ToOLJsaNTOvNB9h9Uc5qK5X5w+7G7O998BN2PC/MWp8H+2fVqpXgNENpNXttkRm1hk1dych86EunfdPuqsX+as44oCyJGFHVBnWpm33eWQw9YqANRI+pCJzP08I5WK3osnPiwshd+hR54yjgfYhBFNI7B95PmEQkCgYBzFSz7h1+s34Ycr8SvxsOBWxymG5zaCsUbPsL04aCgLScCHb9J+E86aVbbVFdglYa5Id7DPTL61ixhl7WZjujspeXZGSbmq0KcnckbmDgqkLECiOJW2NHP/j0McAkDLL4tysF8TLDO8gvuvzNC+WQ6drO2ThrypLVZQ+ryeBIPmwKBgEZxhqa0gVvHQG/7Od69KWj4eJP28kq13RhKay8JOoN0vPmspXJo1HY3CKuHRG+AP579dncdUnOMvfXOtkdM4vk0+hWASBQzM9xzVcztCa+koAugjVaLS9A+9uQoqEeVNTckxx0S2bYevRy7hGQmUJTyQm3j1zEUR5jpdbL83Fbq
-----END RSA PRIVATE KEY-----`)
	// 	pubkey2 := []byte(`-----BEGIN PUBLIC KEY-----
	// MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA4f5wg5l2hKsTeNem/V41
	// fGnJm6gOdrj8ym3rFkEU/wT8RDtnSgFEZOQpHEgQ7JL38xUfU0Y3g6aYw9QT0hJ7
	// mCpz9Er5qLaMXJwZxzHzAahlfA0icqabvJOMvQtzD6uQv6wPEyZtDTWiQi9AXwBp
	// HssPnpYGIn20ZZuNlX2BrClciHhCPUIIZOQn/MmqTD31jSyjoQoV7MhhMTATKJx2
	// XrHhR+1DcKJzQBSTAGnpYVaqpsARap+nwRipr3nUTuxyGohBTSmjJ2usSeQXHI3b
	// ODIRe1AuTyHceAbewn8b462yEWKARdpd9AjQW5SIVPfdsz5B6GlYQ5LdYKtznTuy
	// 7wIDAQAB
	// -----END PUBLIC KEY-----`)
	pubkey := []byte(`-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA4f5wg5l2hKsTeNem/V41fGnJm6gOdrj8ym3rFkEU/wT8RDtnSgFEZOQpHEgQ7JL38xUfU0Y3g6aYw9QT0hJ7mCpz9Er5qLaMXJwZxzHzAahlfA0icqabvJOMvQtzD6uQv6wPEyZtDTWiQi9AXwBpHssPnpYGIn20ZZuNlX2BrClciHhCPUIIZOQn/MmqTD31jSyjoQoV7MhhMTATKJx2XrHhR+1DcKJzQBSTAGnpYVaqpsARap+nwRipr3nUTuxyGohBTSmjJ2usSeQXHI3bODIRe1AuTyHceAbewn8b462yEWKARdpd9AjQW5SIVPfdsz5B6GlYQ5LdYKtznTuy7wIDAQAB
-----END PUBLIC KEY-----`)
	plaintext := []byte("example plain text")

	ciphertextInterface, err := flogoEncryptRsa.Eval(pubkey, plaintext)

	var encryptedText []byte = ciphertextInterface.([]byte)

	assert.Nil(t, err)
	assert.NotNil(t, ciphertextInterface)
	assert.NotNil(t, encryptedText)
	fmt.Printf("%v \n", encryptedText)

	// Decrypt same text
	plaintextInterface, err := flogoDecryptRsa.Eval(privkey, encryptedText)
	var decryptedText []byte = plaintextInterface.([]byte)

	assert.Nil(t, err)
	assert.NotNil(t, ciphertextInterface)
	assert.NotNil(t, decryptedText)
	assert.Equal(t, plaintext, decryptedText)
	fmt.Printf("%v \n", decryptedText)
}
