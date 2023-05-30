package vendorRepository

import (
	"e-commerce/db"
	"e-commerce/model"
	"github.com/valyala/fasthttp"
	"log"
)

func Create(vendorStruct *model.Vendor, context *fasthttp.RequestCtx) (err error) {
	_, err = db.Conn.Collections["vendor"].InsertOne(context, vendorStruct)
	if err != nil {
		log.Println("An error ocurred during insertion")
		return err
	}
	return nil
}
