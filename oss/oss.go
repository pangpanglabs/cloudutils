package oss

import (
	"fmt"
	"strings"
)

//Config struct
type Config struct {
	Endpoint        string
	BucketName      string
	RootDir         string
	AccessKeyID     string
	AccessKeySecret string
}

//Client interface
type Client interface {
	init(ossConfig Config) (err error)

	PutObject(objectName string, bytes []byte) (err error)
	GetObject(objectName string) (bytes []byte, err error)
	IsObjectExist(objectName string) (isExist bool, err error)
	ListObjects(subDir string) (objectNames []string, err error)
}

//New func
func New(ossConfig Config) (ossClient Client, err error) {
	if strings.Contains(ossConfig.Endpoint, "aliyuncs.com") {
		ossClient = &AliyunOss{}
		if err = ossClient.init(ossConfig); err != nil {
			return
		}
		return
	} else if strings.Contains(ossConfig.Endpoint, "myhuaweicloud.com") {
		ossClient = &HuaweiOss{}
		if err = ossClient.init(ossConfig); err != nil {
			return
		}
		return
	}

	err = fmt.Errorf("Unsupport endpoint(%s)", ossConfig.Endpoint)
	return
}
