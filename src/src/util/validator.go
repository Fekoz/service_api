package util

import (
	"regexp"
	"strings"

	pb "bitbucket.org/klaraeng/package_proto/service_api/client"
)

func ValidateClient(req *pb.Personal) bool {
	if req == nil {
		return false
	}
	if len(req.Name) < 1 || len(req.Phone) < 1 {
		return false
	}

	cEmail := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if cEmail.MatchString(req.Email) == false {
		return false
	}

	newPhone := req.Phone
	newPhone = strings.Replace(newPhone, "-", "", -1)
	newPhone = strings.Replace(newPhone, "(", "", -1)
	newPhone = strings.Replace(newPhone, ")", "", -1)
	newPhone = strings.Replace(newPhone, " ", "", -1)
	newPhone = strings.Replace(newPhone, "+", "", -1)
	req.Phone = newPhone
	cPhone := regexp.MustCompile(`^[0-9]{10,12}$`)
	if cPhone.MatchString(req.Phone) == false {
		return false
	}

	return true
}
