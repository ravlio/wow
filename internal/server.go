package internal

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"math/rand"
	"net"
	"strconv"
	"strings"
)

// Server is a server with hashcash DDOS protection.
type Server struct {
	Addr       string
	Difficulty int
}

// ListenAndServe listens on the TCP network address and serves incoming connections.
func (s *Server) ListenAndServe() error {
	listener, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}

		go s.handleConnection(conn)
	}
}

// handleConnection handles incoming connection.
func (s *Server) handleConnection(c net.Conn) {
	defer c.Close()

	// generate challenge
	challenge := uuid.New().String()
	log.Trace().Msgf("challenge: %s", challenge)

	// write challenge
	_, err := c.Write([]byte(challenge))
	if err != nil {
		log.Err(err).Msg("can't write challenge")
		return
	}

	// read nonce
	buffer := make([]byte, 64)
	n, err := c.Read(buffer)
	if err != nil {
		log.Err(err).Msg("can't receive nonce")
		return
	}
	nonceStr := string(buffer[:n])
	log.Trace().Msgf("nonce: %s", nonceStr)
	nonce, err := strconv.Atoi(nonceStr)
	if err != nil {
		log.Trace().Msgf("can't parce nonce %q: %v", nonceStr, err.Error())
		return
	}

	// verify nonce
	if !verifyHashcash(challenge, s.Difficulty, nonce) {
		log.Trace().Msgf("invalid nonce %d", nonce)
		return
	}

	quote := wowQuotes[rand.Intn(len(wowQuotes))]
	_, err = c.Write([]byte(quote))
	if err != nil {
		log.Err(err).Msg("can't write quote")
	}

	return
}

// verifyHashcash verifies hashcash nonce.
func verifyHashcash(challenge string, difficulty int, nonce int) bool {
	prefix := strings.Repeat("0", difficulty)
	data := fmt.Sprintf("%s:%d", challenge, nonce)

	hashBytes := sha256.Sum256([]byte(data))
	hash := hex.EncodeToString(hashBytes[:])

	return strings.HasPrefix(hash, prefix)
}
