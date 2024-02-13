package util

import "github.com/indrawanagung/shop_api_golang/model/web"

func HeaderResponseSuccessfully() web.Header {
	return web.Header{
		Message: "Your request has been processed successfully",
		Error:   false,
	}
}
