package util

import (
	"service_api/src/util/consts"

	pb "bitbucket.org/klaraeng/package_proto/service_api/client"
)

func AppendArrayCategory[R []*pb.GetCatalogCategoryReply_Type](
	reply R,
	types int,
	category string,
	key string,
	value string) []*pb.GetCatalogCategoryReply_Type {
	switch types {
	case consts.TypeCastValue:
		for _, r := range reply {
			if r.TypeName == category {
				for _, k := range r.Key {
					if k.KeyName == key {
						isAppend := false
						for _, p := range k.Value {
							if p == value {
								isAppend = true
							}
						}
						if false == isAppend {
							k.Value = append(k.Value, value)
						}
					}
				}
			}
		}
		break
	case consts.TypeCastKey:
		for _, r := range reply {
			if r.TypeName == category {
				isAppend := false
				for _, k := range r.Key {
					if k.KeyName == key {
						isAppend = true
					}
				}
				if false == isAppend {
					r.Key = append(r.Key, &pb.GetCatalogCategoryReply_Type_Keys{
						KeyName:          key,
						KeyTranslateName: value,
						Value:            []string{},
					})
				}
			}
		}
		break
	case consts.TypeCastType:
		isAppend := false
		for _, r := range reply {
			if r.TypeName == category {
				isAppend = true
			}
		}

		if false == isAppend {
			reply = append(reply, &pb.GetCatalogCategoryReply_Type{
				TypeName: category,
				Key:      []*pb.GetCatalogCategoryReply_Type_Keys{},
			})
		}
		break
	}

	return reply
}
