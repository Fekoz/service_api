package util

import (
	"service_api/src/entitys"

	pb "bitbucket.org/klaraeng/package_proto/service_api/client"
	"github.com/go-kratos/kratos/v2/errors"
)

func FormatToProduct(item *entitys.Product) (*pb.GetCarpetOnlyProductReply, error) {
	if item == nil {
		return nil, errors.New(500, "CARPET_PRODUCT_ITEM_HAS_NULL", "Указан не верный параметр ARTICLE, UUID, ID")
	}
	return &pb.GetCarpetOnlyProductReply{
		Success: true,
		Product: &pb.Product{
			Id:       item.ID,
			Uuid:     item.Uuid,
			Article:  item.Article,
			Name:     item.Name,
			IsActive: item.IsActive,
		},
	}, nil
}

func FormatToSpecification(item []entitys.Specification) (*pb.GetCarpetOnlySpecificationReply, error) {
	if item == nil || len(item) == 0 {
		return nil, errors.New(500, "CARPET_SPECIFICATION_ITEM_HAS_NULL", "Указан не верный параметр ARTICLE, UUID, ID")
	}
	var element []*pb.Specification
	for _, el := range item {
		element = append(element, &pb.Specification{
			Name:  el.Name,
			Value: el.Value,
		})
	}
	return &pb.GetCarpetOnlySpecificationReply{
		Success:       true,
		Specification: element,
	}, nil
}

func FormatToImages(item []entitys.Images) (*pb.GetCarpetOnlyImagesReply, error) {
	if item == nil || len(item) == 0 {
		return nil, errors.New(500, "CARPET_IMAGES_ITEM_HAS_NULL", "Указан не верный параметр ARTICLE, UUID, ID")
	}
	var element []*pb.Images
	for _, el := range item {
		element = append(element, &pb.Images{
			Name: el.Filename,
			Dir:  el.Dir,
			Type: el.Type,
		})
	}
	return &pb.GetCarpetOnlyImagesReply{
		Success: true,
		Images:  element,
	}, nil
}

func FormatToAttribute(item []entitys.Attribute) (*pb.GetCarpetOnlyAttributeReply, error) {
	if item == nil || len(item) == 0 {
		return nil, errors.New(500, "CARPET_ATTRIBUTE_ITEM_HAS_NULL", "Указан не верный параметр ARTICLE, UUID, ID")
	}
	var element []*pb.Attribute
	for _, el := range item {
		element = append(element, &pb.Attribute{
			Name:  el.Name,
			Value: el.Value,
		})
	}
	return &pb.GetCarpetOnlyAttributeReply{
		Success:   true,
		Attribute: element,
	}, nil
}
