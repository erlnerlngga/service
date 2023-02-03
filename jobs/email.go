package jobs

import (
	"context"
	"log"

	"github.com/maragudk/errors"

	"github.com/maragudk/service/model"
)

func SendEmail(r registry, log *log.Logger) {
	r.Register("send-email", func(ctx context.Context, m model.Map) error {
		switch m["type"] {
		case "signup":
			log.Println("TODO: Send signup email with token:", m["token"])
		default:
			return errors.Newf(`unknown email type "%v"`, m["type"])
		}
		return nil
	})
}
