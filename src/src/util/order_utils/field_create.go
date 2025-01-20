package order_utils

import (
	"service_api/src/entitys"
	"service_api/src/util/consts"
	"strconv"

	pb "bitbucket.org/klaraeng/package_proto/service_api/client"
)

func OrderFieldCreate(orderId uint, req *pb.CreateOrderLandingRequest) []entitys.OrderField {
	var field []entitys.OrderField

	field = append(field, entitys.OrderField{
		Key:     consts.CARPET_FORM,
		Value:   pb.CARPET_FORM_name[int32(req.Form)],
		OrderID: orderId,
	})

	field = append(field, entitys.OrderField{
		Key:     consts.CARPET_DESIGN,
		Value:   pb.CARPET_DESIGN_name[int32(req.Design)],
		OrderID: orderId,
	})

	field = append(field, entitys.OrderField{
		Key:     consts.CARPET_COLOR,
		Value:   pb.CARPET_COLOR_name[int32(req.Color)],
		OrderID: orderId,
	})

	if req.Size == nil {
		req.Size = &pb.CreateOrderLandingRequest_SizeCarpet{
			Width:  1,
			Height: 1,
		}
	}

	field = append(field, entitys.OrderField{
		Key:     consts.CARPET_SIZE_WIDTH,
		Value:   strconv.FormatInt(req.Size.Width, 10),
		OrderID: orderId,
	})

	field = append(field, entitys.OrderField{
		Key:     consts.CARPET_SIZE_HEIGHT,
		Value:   strconv.FormatInt(req.Size.Height, 10),
		OrderID: orderId,
	})

	field = append(field, entitys.OrderField{
		Key:     consts.CARPET_PAY_LIMIT,
		Value:   strconv.FormatInt(req.CarpetPayLimit, 10),
		OrderID: orderId,
	})

	return field
}
