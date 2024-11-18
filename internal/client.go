package internal

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/rs/zerolog/log"
	"net"
	"strconv"
	"strings"
)

// Call calls to server with given address and difficulty and returns the quote.
func Call(addr string, difficulty int) (string, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		return "", fmt.Errorf("can't read challenge: %w", err)
	}

	challenge := strings.TrimSpace(string(buffer[:n]))
	log.Trace().Msgf("challenge: %s", challenge)

	nonce := generateHashcash(challenge, difficulty)
	log.Trace().Msgf("nonce: %d", nonce)
	_, err = conn.Write([]byte(strconv.Itoa(nonce)))
	if err != nil {
		return "", fmt.Errorf("can't write nonce: %w", err)
	}

	buffer = buffer[:1024]
	n, err = conn.Read(buffer)
	if err != nil {
		return "", fmt.Errorf("can't read quote: %w", err)
	}
	return string(buffer[:n]), nil
}

// generateHashcash generates a hashcash and returns the nonce.
func generateHashcash(challenge string, difficulty int) int {
	var nonce int
	var hash string
	prefix := strings.Repeat("0", difficulty)

	for {
		data := fmt.Sprintf("%s:%d", challenge, nonce)

		hashBytes := sha256.Sum256([]byte(data))
		hash = hex.EncodeToString(hashBytes[:])

		if strings.HasPrefix(hash, prefix) {
			break
		}
		nonce++
	}

	return nonce
}
