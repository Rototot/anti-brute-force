// Code generated by go-enum
// DO NOT EDIT!

package constants

import (
	"fmt"
)

const (
	// BucketTypeLogin is a BucketType of type Login
	BucketTypeLogin BucketType = iota
	// BucketTypeIp is a BucketType of type Ip
	BucketTypeIp
	// BucketTypePassword is a BucketType of type Password
	BucketTypePassword
)

const _BucketTypeName = "LoginIpPassword"

var _BucketTypeMap = map[BucketType]string{
	0: _BucketTypeName[0:5],
	1: _BucketTypeName[5:7],
	2: _BucketTypeName[7:15],
}

// String implements the Stringer interface.
func (x BucketType) String() string {
	if str, ok := _BucketTypeMap[x]; ok {
		return str
	}
	return fmt.Sprintf("BucketType(%d)", x)
}

var _BucketTypeValue = map[string]BucketType{
	_BucketTypeName[0:5]:  0,
	_BucketTypeName[5:7]:  1,
	_BucketTypeName[7:15]: 2,
}

// ParseBucketType attempts to convert a string to a BucketType
func ParseBucketType(name string) (BucketType, error) {
	if x, ok := _BucketTypeValue[name]; ok {
		return x, nil
	}
	return BucketType(0), fmt.Errorf("%s is not a valid BucketType", name)
}
