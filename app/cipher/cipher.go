package cipher

import (
	"github.com/ProtonMail/gopenpgp/v2/crypto"
	"os"
)

type ICipher interface {
	Encrypt(filename string) error
	Decrypt(filename string) error
}

type Cipher struct {
	PathReq        string
	PathRes        string
	PathPublicKey  string
	PathPrivateKey string
	PathPassword   string
	ExtText        string
	ExtPGP         string
}

func NewCipher(c Cipher) *Cipher {
	return &c
}

func (c *Cipher) Encrypt(filename string) error {
	plainData, err := os.ReadFile(c.PathReq + filename)
	if err != nil {
		return err
	}

	publicKeyData, err := os.ReadFile(c.PathPublicKey)
	if err != nil {
		return err
	}

	publicKey, err := crypto.NewKeyFromArmored(string(publicKeyData))
	if err != nil {
		return err
	}

	publicRing, err := crypto.NewKeyRing(publicKey)
	if err != nil {
		return err
	}

	message := crypto.NewPlainMessage(plainData)
	encMessage, err := publicRing.Encrypt(message, nil)
	if err != nil {
		return err
	}

	data, err := encMessage.GetArmored()
	if err != nil {
		return err
	}

	newFileName := c.PathRes + getNameWithoutExt(filename) + c.ExtText + c.ExtPGP
	err = os.WriteFile(newFileName, []byte(data), 0644)
	if err != nil {
		return err
	}

	return nil
}

func (c *Cipher) Decrypt(filename string) error {
	privateKeyData, err := os.ReadFile(c.PathPrivateKey)
	if err != nil {
		return err
	}

	key, err := crypto.NewKeyFromArmored(string(privateKeyData))
	if err != nil {
		return err
	}

	password, err := os.ReadFile(c.PathPassword)
	if err != nil {
		return err
	}

	unlockedKey, err := key.Unlock(password)
	if err != nil {
		return err
	}

	keyRing, err := crypto.NewKeyRing(unlockedKey)
	if err != nil {
		return err
	}

	encData, err := os.ReadFile(c.PathReq + filename)
	if err != nil {
		return err
	}

	encMessage, err := crypto.NewPGPMessageFromArmored(string(encData))
	if err != nil {
		return err
	}

	decrypted, err := keyRing.Decrypt(encMessage, nil, 0)
	if err != nil {
		return err
	}

	newFileName := c.PathRes + getNameWithoutExt(filename) + c.ExtText
	err = os.WriteFile(newFileName, decrypted.GetBinary(), 0644)
	if err != nil {
		return err
	}

	return nil
}
