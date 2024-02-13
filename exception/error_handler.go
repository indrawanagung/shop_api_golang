package exception

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/indrawanagung/shop_api_golang/model/web"
)

var ErrNotFound = errors.New("Resource was not found")

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	_, ok := err.(validator.ValidationErrors)
	if ok {
		return ctx.Status(400).JSON(web.WebResponse{
			Header: web.Header{
				Message: err.Error(),
				Error:   true,
			},
			Data: nil,
		})
	}

	_, ok = err.(NotFoundError)
	if ok {
		return ctx.Status(404).JSON(web.WebResponse{
			Header: web.Header{
				Message: err.Error(),
				Error:   true,
			},
			Data: nil,
		})
	}

	_, ok = err.(BadRequestError)
	if ok {
		return ctx.Status(400).JSON(web.WebResponse{
			Header: web.Header{
				Message: err.Error(),
				Error:   true,
			},
			Data: nil,
		})
	}

	_, ok = err.(UnauthorizedError)
	if ok {
		return ctx.Status(401).JSON(web.WebResponse{
			Header: web.Header{
				Message: err.Error(),
				Error:   true,
			},
			Data: nil,
		})
	}

	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	// Set Content-Type: text/plain; charset=utf-8
	ctx.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

	// Return status code with error message
	return ctx.Status(code).JSON(web.WebResponse{
		Header: web.Header{
			Message: err.Error(),
			Error:   true,
		},
		Data: nil,
	})

	//return ctx.Status(500).JSON(mobile.MobileResponse{
	//	Code:  500,
	//	Data:  nil,
	//	Error: err.Error(),
	//})

	//if err != nil {
	//	return nil
	//}
	//errResponse := notFoundError(ctx, err)
	//if errResponse != nil {
	//	mobileResponse := mobile.MobileResponse{
	//		Code:  http.StatusNotFound,
	//		Data:  nil,
	//		Error: errResponse.Error(),
	//	}
	//	return ctx.Status(mobileResponse.Code).JSON(mobileResponse)
	//}
	//errResponse = badRequestError(ctx, err)
	//if errResponse != nil {
	//	mobileResponse := mobile.MobileResponse{
	//		Code:  http.StatusBadRequest,
	//		Data:  nil,
	//		Error: errResponse.Error(),
	//	}
	//	return ctx.Status(mobileResponse.Code).JSON(mobileResponse)
	//}
	//
	//errResponse = unauthorizedError(ctx, err)
	//if errResponse != nil {
	//	mobileResponse := mobile.MobileResponse{
	//		Code:  http.StatusUnauthorized,
	//		Data:  nil,
	//		Error: errResponse.Error,
	//	}
	//	return ctx.Status(mobileResponse.Code).JSON(mobileResponse)
	//}
	//errResponse = validationErrors(ctx, err)
	//if errResponse != nil {
	//	mobileResponse := mobile.MobileResponse{
	//		Code:  http.StatusBadRequest,
	//		Data:  nil,
	//		Error: errResponse.Error,
	//	}
	//	return ctx.Status(mobileResponse.Code).JSON(mobileResponse)
	//}
	//
	//mobileResponse := mobile.MobileResponse{
	//	Code:  http.StatusInternalServerError,
	//	Data:  nil,
	//	Error: errResponse.Error,
	//}
	//return ctx.Status(mobileResponse.Code).JSON(mobileResponse)

	//if notFoundError(ctx, err) == nil{
	//	errNotFound := notFoundError(ctx)
	//	mobileResponse := mobile.MobileResponse{
	//		Code:  http.StatusNotFound,
	//		Data:  nil,
	//		Error: exception.Error,
	//	}
	//	return ctx.Status(mobileResponse.Code).JSON(mobileResponse)
	//}
	//unauthorizedError(ctx, err)
	//validationErrors(ctx, err)
	//badRequestError(ctx, err)
	//internalServerError(ctx, err)
	//return nil
}

//func validationErrors(ctx *fiber.Ctx, err error) error {
//	exception, ok := err.(validator.ValidationErrors)
//	if ok {
//		return exception
//	} else {
//		return nil
//	}
//}
//
//func notFoundError(ctx *fiber.Ctx, err error) error {
//	exception, ok := err.(NotFoundError)
//	if ok {
//		return exception
//	} else {
//		return nil
//	}
//}
//func badRequestError(ctx *fiber.Ctx, err error) error {
//	exception, ok := err.(BadRequestError)
//	if ok {
//		return exception
//	} else {
//		return nil
//	}
//}
//
//func unauthorizedError(ctx *fiber.Ctx, err error) error {
//	exception, ok := err.(UnauthorizedError)
//	if ok {
//		return exception
//	} else {
//		return nil
//	}
//}

//func internalServerError(ctx *fiber.Ctx, err interface{}) error {
//	mobileResponse := mobile.MobileResponse{
//		Code:  http.StatusInternalServerError,
//		Data:  nil,
//		Error: err.(error).Error(),
//	}
//
//	return ctx.Status(mobileResponse.Code).JSON(mobileResponse)
//
//}
