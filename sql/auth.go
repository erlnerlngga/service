package sql

import (
	"context"
	"crypto/rand"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/maragudk/errors"

	"github.com/maragudk/service/model"
)

// Signup creates an account, a personal group, an unconfirmed user, and a token.
// Also creates a job to send an email with a token.
func (d *Database) Signup(ctx context.Context, u *model.User) error {
	return d.inTransaction(ctx, func(tx *sqlx.Tx) error {
		token, err := createToken()
		if err != nil {
			return err
		}

		var a model.Account
		if err := tx.GetContext(ctx, &a, `insert into accounts (name) values (?) returning *`, u.Name); err != nil {
			return errors.Wrap(err, "error creating account")
		}

		var g model.Group
		query := `insert into groups (accountID, name) values (?, ?) returning *`
		if err := tx.GetContext(ctx, &g, query, a.ID, u.Name); err != nil {
			return errors.Wrap(err, "error creating group")
		}

		var exists bool
		query = `select exists (select * from users where email = ?)`
		if err := tx.GetContext(ctx, &exists, query, u.Email.ToLower()); err != nil {
			return errors.Wrap(err, "error getting user by email")
		}
		if exists {
			return model.ErrorEmailConflict
		}

		query = `insert into users (accountID, name, email) values (?, ?, ?) returning *`
		if err := tx.GetContext(ctx, u, query, a.ID, u.Name, u.Email.ToLower()); err != nil {
			return errors.Wrap(err, "error creating user")
		}

		query = `insert into group_membership (groupID, userID) values (?, ?)`
		if _, err := tx.ExecContext(ctx, query, g.ID, u.ID); err != nil {
			return err
		}

		query = `insert into tokens (value, userID) values (?, ?)`
		if _, err := tx.ExecContext(ctx, query, token, u.ID); err != nil {
			return errors.Wrap(err, "error creating token")
		}

		m := model.Map{
			"type":  "signup",
			"token": token,
		}
		if err := d.createJobInTx(ctx, tx, "send-email", m, 10*time.Second); err != nil {
			return err
		}

		return nil
	})
}

func createToken() (string, error) {
	secret := make([]byte, 16)
	if _, err := rand.Read(secret); err != nil {
		return "", err
	}
	return fmt.Sprintf("t_%x", secret), nil
}
