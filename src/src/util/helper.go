package util

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"service_api/src/dtos"
	"service_api/src/util/consts"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/google/uuid"
)

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func getAddrIsVal(md metadata.Metadata, addr string) (string, error) {
	ip := md.Get(addr)
	if len(ip) > 0 {
		return ip, nil
	}

	return "", errors.New(500, "IP", "Не указан IP")
}

func GetClientIp(ctx context.Context) (string, error) {
	var ip string
	var err error
	if md, ok := metadata.FromServerContext(ctx); ok {
		ip, err = getAddrIsVal(md, consts.MetadataIP)
		if err == nil {
			return ip, nil
		}

		ip, err = getAddrIsVal(md, consts.MetadataIPProxy)
		if err == nil {
			return ip, nil
		}
	}

	return "", errors.New(500, "IP", "Не указан IP. Не доступно для использования")
}

func GenSession(data *dtos.AdminTransfer) string {
	out, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	body := fmt.Sprintf("%s::%s", out, uuid.NewString())
	hash := md5.New()
	hash.Write([]byte(body))
	return hex.EncodeToString(hash.Sum(nil))
}

func IsInstanceOf(objectPtr, typePtr interface{}) bool {
	return reflect.TypeOf(objectPtr) == reflect.TypeOf(typePtr)
}

func GenKey[T any](prefix string, req *T) (string, error) {
	out, err := json.Marshal(req)
	if err != nil {
		return "", err
	}
	hash := md5.New()
	hash.Write(out)
	return fmt.Sprintf(consts.MaskCacheKey, prefix, hex.EncodeToString(hash.Sum(nil))), nil
}

func EncodeValue[T any](resp *T) ([]byte, error) {
	data, err := json.Marshal(resp)
	if err != nil {
		return []byte(""), err
	}
	return data, nil
}

func DecodeValue[T any](data []byte) (T, error) {
	var resp T
	err := json.Unmarshal(data, &resp)
	return resp, err
}
