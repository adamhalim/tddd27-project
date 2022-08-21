package middleware

import (
	"errors"
	"net/http"

	"github.com/auth0/go-jwt-middleware/v2/validator"
	"gitlab.liu.se/adaab301/tddd27_2022_project/backend/chunk"
	"gitlab.liu.se/adaab301/tddd27_2022_project/lib/postgres"
)

// If the first chunk is uploaded (id == 0), a directory & session is created
func HandleChunkUpload(r *http.Request, token *validator.ValidatedClaims) error {
	id := r.URL.Query().Get("id")
	if id == "0" {
		fileName := r.URL.Query().Get("fileName")
		chunkName := r.URL.Query().Get("chunkName")
		if fileName == "" {
			return errors.New("no filename provided")
		}
		if chunkName == "" {
			return errors.New("no chunkName provided")
		}
		if token.RegisteredClaims.Subject == "" {
			return errors.New("invalid subject")
		}

		if err := chunk.CreateDirectory(chunkName); err != nil {
			return err
		}
		if err := chunk.NewSession(chunkName, fileName, token.RegisteredClaims.Subject); err != nil {
			return err
		}
		userExists := postgres.UserExists(token.RegisteredClaims.Subject)
		if !userExists {
			err := postgres.AddUser(postgres.User{
				Uid:      token.RegisteredClaims.Subject,
				Username: token.RegisteredClaims.Subject,
			})
			if err != nil {
				return err
			}
		}
	}
	return nil
}
