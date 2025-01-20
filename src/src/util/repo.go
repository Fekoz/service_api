package util

import "gorm.io/gorm"

func IsErrQb(qb *gorm.DB) bool {
	if qb.Error != nil || qb == nil{
		return true
	}
	return false
}
