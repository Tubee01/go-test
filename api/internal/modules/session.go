package modules

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/session"
)

var Session = session.New()

func SetSession(ctx fiber.Ctx, key string, value interface{}) (string, error) {
	sess, err := Session.Get(ctx)
	if err != nil {
		return "", err
	}

	sess.Set(key, value)
	id := sess.ID()

	sess.Save()

	return id, nil
}

func GetSessionData(ctx fiber.Ctx, key string) (interface{}, error) {
	sess, err := Session.Get(ctx)
	if err != nil {
		return nil, err
	}

	return sess.Get(key), nil
}

func DestroySession(ctx fiber.Ctx) error {
	sess, err := Session.Get(ctx)
	if err != nil {
		return err
	}

	err = sess.Destroy()
	if err != nil {
		return err
	}

	return nil
}
