package mercury

import (
	"testing"

	"github.com/rs/zerolog"
	"github.com/smartcontractkit/chainlink/integration-tests/client"
	"github.com/stretchr/testify/require"
)

var (
	adminId = "02185d5a-f1ee-40d1-a52a-bf39871b614c"
	// DATABASE_ENCRYPTION_KEY=key go run cmd/server/main.go encrypt "key"
	adminSecret = "key"
)

func TestMercuryServerHMAC(t *testing.T) {
	l := zerolog.New(zerolog.NewTestWriter(t))

	mercuryserver := client.NewMercuryServer("http://localhost:3000")

	user, _, err := mercuryserver.GetUsers(adminId, adminSecret)

	// Create new user
	// newUserSecret := "key"
	// newUserRole := "user"
	// newUserDisabled := false
	// user, _, err := mercuryserver.AddUser(adminId, adminSecret, newUserSecret, newUserRole, newUserDisabled)
	require.NoError(t, err)
	_ = user

	// Get report

	l.Log().Msgf("asdsa")
}
