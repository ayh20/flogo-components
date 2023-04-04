# Crypto functions

This package adds function for encrypting and decrypting using AES GCM or RSA public/private keys pairs.

## Installation

```CLI
flogo install github.com/ayh20/flogo-components/functions/crypto
```

Link for TCI:

```
github.com/ayh20/flogo-components/functions/crypto
```

# AES GCM passphrase based encryption

## encrypt()
Encrypt given byte array with provide 128bit or 256 bit key.

Example usage
```
coerce.toString(crypto.encrypt(coerce.toBytes("AES256Key-32Characters1234567890"), coerce.toBytes("encrypt-me-text")))
```


### Input Args

| Arg         | Type        | Description                      |
|-------------|-------------|----------------------------------|
| key         | byte array  | 128bit or 256 bit key            |
| plaintext   | byte array  | Plaintext that will be encrypted |

### Output

| Arg         | Type       | Description           |
|-------------|------------|-----------------------|
| returnType  | byte array | Encrypted `plaintext` | 


## decrypt()
Decrypt given encrypted byte array with provide 128bit or 256 bit key.

Example usage
```
coerce.toString(crypto.decrypt(coerce.toBytes("AES256Key-32Characters1234567890"), coerce.toBytes($activity[Encrypt].output.name)))
```


### Input Args

| Arg        | Type       | Description                              |
|------------|------------|------------------------------------------|
| key        | byte array | 128bit or 256 bit key                    |
| ciphertext | byte array | Encrypted payload that will be decrypted |

### Output

| Arg        | Type       | Description            |
|------------|------------|------------------------|
| returnType | byte array | Decrypted `ciphertext` | 

## RSA Public/Private Key encryption

## encryptrsa()
Encrypt given byte array with provide 128bit or 256 bit key.

Example usage
```
coerce.toString(crypto.encryptrsa(coerce.toBytes("`-----BEGIN PUBLIC KEY-----MIIBIj.......7wIDAQAB-----END PUBLIC KEY-----`"), coerce.toBytes("encrypt-me-text")))
```


### Input Args

| Arg         | Type        | Description                      |
|-------------|-------------|----------------------------------|
| publickey   | byte array  | RSA public key string (PEM)      |
| plaintext   | byte array  | Plaintext that will be encrypted |

### Output

| Arg         | Type       | Description           |
|-------------|------------|-----------------------|
| returnType  | byte array | Encrypted `plaintext` | 


## decryptrsa()
Decrypt given encrypted byte array with provided private key.

Example usage
```
coerce.toString(crypto.decrypt(coerce.toBytes("-----BEGIN RSA PRIVATE KEY-----MIIEow.... jpdbL83Fbq-----END RSA PRIVATE KEY-----"), coerce.toBytes($activity[Encrypt].output.name)))
```


### Input Args

| Arg        | Type       | Description                              |
|------------|------------|------------------------------------------|
| privatekey | byte array | RSA private key string (PEM)             |
| ciphertext | byte array | Encrypted payload that will be decrypted |

### Output

| Arg        | Type       | Description            |
|------------|------------|------------------------|
| returnType | byte array | Decrypted `ciphertext` | 
