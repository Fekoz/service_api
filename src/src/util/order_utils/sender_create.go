package order_utils

import (
	"service_api/src/dtos"
	"service_api/src/entitys"
	"service_api/src/util"
)

func SenderCarpetEmailCreate(
	params map[string]string,
	tmp *entitys.MailTemplate,
	client *entitys.Client,
	product *entitys.Product,
	price *entitys.Price,
	images *entitys.Images,
) []byte {
	if len(tmp.Title) < 1 || len(tmp.Message) < 1 {
		return nil
	}

	title := util.ReplaceSubstrings(tmp.Title, params)
	title = util.ReplaceString("Client", title, client)

	message := util.ReplaceSubstrings(tmp.Message, params)
	message = util.ReplaceString("Client", message, client)

	if len(product.Article) > 1 {
		message = util.ReplaceString("Product", message, product)
		title = util.ReplaceString("Product", title, product)
	}

	if len(price.Mid) > 1 {
		message = util.ReplaceString("Price", message, price)
		title = util.ReplaceString("Price", title, price)
	}

	if len(images.Filename) > 1 && len(images.Dir) > 1 {
		message = util.ReplaceString("Image", message, images)
		title = util.ReplaceString("Image", title, images)
	}

	messageEmail := &dtos.SenderEmailDto{
		To:      client.Email,
		Title:   title,
		Message: message,
	}

	data, err := util.EncodeValue[*dtos.SenderEmailDto](&messageEmail)
	if err != nil {
		return nil
	}

	return data
}

func SenderCarpetTelegramCreate(message string) []byte {
	messageEmail := &dtos.SenderTelegramDto{
		Message: message,
	}

	data, err := util.EncodeValue[*dtos.SenderTelegramDto](&messageEmail)
	if err != nil {
		return nil
	}

	return data
}
