package sql_test

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/maragudk/is"

	"github.com/maragudk/service/model"
	"github.com/maragudk/service/sqltest"
)

func TestDatabase_Signup(t *testing.T) {
	t.Run("signs up an account, group, and user, and returns a token", func(t *testing.T) {
		db := sqltest.CreateDatabase(t)

		u := model.User{
			Name:  "Me",
			Email: "Me@example.com",
		}
		token, err := db.Signup(context.Background(), &u)
		is.NotError(t, err)
		is.Equal(t, 34, len(token))
		is.True(t, strings.HasPrefix(token, "t_"))

		is.Equal(t, 34, len(u.ID))
		is.True(t, strings.HasPrefix(u.ID.String(), "u_"))
		is.True(t, time.Since(u.Created.T) < time.Second)
		is.True(t, time.Since(u.Updated.T) < time.Second)
		is.Equal(t, "Me", u.Name)
		is.Equal(t, "me@example.com", u.Email.String())
		is.True(t, !u.Confirmed)
		is.True(t, u.Active)

		var a model.Account
		err = db.DB.Get(&a, `select * from accounts where id = ?`, u.AccountID)
		is.NotError(t, err)
		is.Equal(t, 34, len(a.ID))
		is.True(t, strings.HasPrefix(a.ID.String(), "a_"))
		is.True(t, time.Since(a.Created.T) < time.Second)
		is.True(t, time.Since(a.Updated.T) < time.Second)
		is.Equal(t, "Me", a.Name)

		var g model.Group
		err = db.DB.Get(&g, `select * from groups where accountID = ?`, u.AccountID)
		is.NotError(t, err)
		is.Equal(t, 34, len(g.ID))
		is.True(t, strings.HasPrefix(g.ID.String(), "g_"))
		is.True(t, time.Since(g.Created.T) < time.Second)
		is.True(t, time.Since(g.Updated.T) < time.Second)
		is.Equal(t, "Me", g.Name)

		var exists bool
		err = db.DB.Get(&exists, `select exists (select * from group_membership where userID = ? and groupID = ?)`,
			u.ID, g.ID)
		is.NotError(t, err)
		is.True(t, exists)
	})

	t.Run("errors on duplicate email", func(t *testing.T) {
		db := sqltest.CreateDatabase(t)

		u := model.User{
			Name:  "Me",
			Email: "Me@example.com",
		}
		_, err := db.Signup(context.Background(), &u)
		is.NotError(t, err)

		_, err = db.Signup(context.Background(), &u)
		is.Error(t, model.ErrorEmailConflict, err)
	})
}
