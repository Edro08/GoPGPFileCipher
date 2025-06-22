package main

import (
	"GoPGPFileCipher/app/cipher"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	basePathReq        = "./files/$profile/request/"
	basePathRes        = "./files/$profile/response/"
	basePathPublicKey  = "./keys/$profile/public.key"
	basePathPrivateKey = "./keys/$profile/private.key"
	basePathPassword   = "./keys/$profile/passphrase.txt"

	extText        = ".txt"
	extPGP         = ".pgp"
	replaceProfile = "$profile"
)

func execute(profile string) {
	pathReq := strings.Replace(basePathReq, replaceProfile, profile, 1)
	pathRes := strings.Replace(basePathRes, replaceProfile, profile, 1)
	pathPublicKey := strings.Replace(basePathPublicKey, replaceProfile, profile, 1)
	pathPrivateKey := strings.Replace(basePathPrivateKey, replaceProfile, profile, 1)
	pathPassword := strings.Replace(basePathPassword, replaceProfile, profile, 1)

	c := cipher.NewCipher(cipher.Cipher{
		PathReq:        pathReq,
		PathRes:        pathRes,
		PathPublicKey:  pathPublicKey,
		PathPrivateKey: pathPrivateKey,
		PathPassword:   pathPassword,
		ExtText:        extText,
		ExtPGP:         extPGP,
	})

	entries, err := getDirEntries(pathReq)
	if err != nil {
		log.Printf("❌ Error leyendo archivos de %s: %v", pathReq, err)
		return
	}

	if len(entries) == 0 {
		log.Printf("📂 No se encontraron archivos en %s", pathReq)
		return
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		filename := entry.Name()
		fullPath := filepath.Join(pathReq, filename)

		var opErr error
		switch {
		case strings.HasSuffix(strings.ToLower(filename), extText):
			// Se pasa solo el nombre porque Cipher ya tiene los paths configurados
			opErr = c.Encrypt(filename)
		case strings.HasSuffix(strings.ToLower(filename), extPGP):
			// Se pasa solo el nombre porque Cipher ya tiene los paths configurados
			opErr = c.Decrypt(filename)
		default:
			continue
		}

		if opErr != nil {
			log.Printf("❌ Error procesando %s: %v", filename, opErr)
			continue
		}

		if err := os.Remove(fullPath); err != nil {
			log.Printf("⚠️ Error al eliminar %s: %v", filename, err)
		} else {
			log.Printf("✔️ Procesado y eliminado: %s", filename)
		}
	}
}

func getDirEntries(path string) ([]os.DirEntry, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	return entries, nil
}
