package tokenGenerator

import (
	"crypto/rand"
	"log"
	"math/big"
)

const (
	uppercase   = "ABCDEFGHIJKLMOPQRSTUVWXYZ"
	lowercase   = "abcdefghijklmopqrstuvwxyz"
	numbers     = "0123456789"
	symbols     = ".,;:!?./-\"'#{([-|\\@)]=}*+"
	maxTokenLen = 512 // Max token size
)

// Fonction pour générer un token
func generateToken(length int, noSymbols bool) (string, error) {
	var charset string
	if noSymbols {
		charset = uppercase + lowercase + numbers
	} else {
		charset = uppercase + lowercase + numbers + symbols
	}

	token := make([]byte, length)

	for i := 0; i < length; i++ {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		token[i] = charset[randomIndex.Int64()]
	}

	return string(token), nil
}

func GetToken(length int, noSymbols bool) string {
	// Lire la taille minimum du token
	if length < 2 {
		log.Fatal("Veuillez fournir une taille de token")
	}

	// Vérifier que la taille du token ne dépasse pas maxTokenLen
	if length > maxTokenLen {
		log.Fatalf("La taille maximale du token est de %d caractères.", maxTokenLen)
	}

	// Générer le token
	token, err := generateToken(length, noSymbols)
	if err != nil {
		log.Fatal("Erreur lors de la génération du token:", err)
	}

	return token
}
