package crypto

import (
	"fmt"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	_ = function.Register(&fnEncrypt{})
	_ = function.Register(&fnEncryptRsa{})
}

type fnEncrypt struct {
}
type fnEncryptRsa struct {
}

func (fnEncrypt) Name() string {
	return "encrypt"
}
func (fnEncryptRsa) Name() string {
	return "encryptrsa"
}

func (fnEncrypt) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeBytes, data.TypeBytes}, false
}
func (fnEncryptRsa) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeBytes, data.TypeBytes}, false
}

func (fnEncrypt) Eval(params ...interface{}) (interface{}, error) {
	key, err := coerce.ToBytes(params[0])
	if err != nil {
		return nil, fmt.Errorf("encrypt function first parameter (key) [%+v] must be bytes", params[0])
	}

	plaintext, err := coerce.ToBytes(params[1])
	if err != nil {
		return nil, fmt.Errorf("encrypt function second parameter (plaintext) [%+v] must be byte", params[1])
	}

	return Encrypt(key, plaintext)
}

func (fnEncryptRsa) Eval(params ...interface{}) (interface{}, error) {
	key, err := coerce.ToBytes(params[0])
	if err != nil {
		return nil, fmt.Errorf("encrypt function first parameter (key) [%+v] must be bytes", params[0])
	}

	plaintext, err := coerce.ToBytes(params[1])
	if err != nil {
		return nil, fmt.Errorf("encrypt function second parameter (publickey) [%+v] must be byte", params[1])
	}

	return EncryptRsa(key, plaintext)
}
