package internal

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
	"time"
)

func TestServerPass(t *testing.T) {
	zerolog.SetGlobalLevel(zerolog.Disabled)

	// Create server
	srv := &Server{
		Addr:       ":8088",
		Difficulty: 2,
	}

	// Run server
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Err(err).Msg("server error")
		}
	}()

	// Ensure that server is started
	time.Sleep(time.Millisecond * 10)

	// Call to server
	_, err := Call(":8088", 2)
	if err != nil {
		println(err.Error())
	}

	assert.NoError(t, err)
}

func TestServerFail(t *testing.T) {
	zerolog.SetGlobalLevel(zerolog.Disabled)

	srv := &Server{
		Addr:       ":8089",
		Difficulty: 4,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Err(err).Msg("server error")
		}
	}()

	// Ensure that server is started
	time.Sleep(time.Millisecond * 10)

	conn, err := net.Dial("tcp", ":8089")
	assert.NoError(t, err)
	defer conn.Close()

	buffer := make([]byte, 1024)
	_, err = conn.Read(buffer)
	assert.NoError(t, err)
	_, err = conn.Write([]byte("1"))
	assert.NoError(t, err)

	_, err = conn.Read(buffer)
	assert.Error(t, err, "EOF")
}

func TestVerifyHashcash(t *testing.T) {
	t.Run("valid hash", func(t *testing.T) {
		assert.True(t, verifyHashcash("example_question", 4, 10636))
	})

	t.Run("invalid hash", func(t *testing.T) {
		assert.False(t, verifyHashcash("challenge", 4, 1))
	})
}
