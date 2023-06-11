package clientRepository

import (
	"e-commerce/db"
	"e-commerce/model"

	"github.com/valyala/fasthttp"
)


func Register(ctx *fasthttp.RequestCtx, client *model.Client)(err error) {
	_, err = db.Conn.Collections["client"].InsertOne(ctx, client)
	return err
}