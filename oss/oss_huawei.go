package oss

import (
	"obs"
)

//HuaweiOss struct
type HuaweiOss struct {
	obsClient *obs.ObsClient
	bucket    *obs.Bucket
	rootDir   string
}

func (huawaiOss *HuaweiOss) init(ossConfig Config) (err error) {
	if huawaiOss.obsClient, err = obs.New(ossConfig.AccessKeyID, ossConfig.AccessKeySecret, ossConfig.Endpoint); err != nil {
		return
	}

	return err
}

//PutObject func
func (huawaiOss *HuaweiOss) PutObject(objectName string, bytes []byte) (err error) { return }

//GetObject func
func (huawaiOss *HuaweiOss) GetObject(objectName string) (bytes []byte, err error) { return }

//IsObjectExist func
func (huawaiOss *HuaweiOss) IsObjectExist(objectName string) (isExist bool, err error) { return }

//ListObjects func
func (huawaiOss *HuaweiOss) ListObjects(subDir string) (objectNames []string, err error) { return }
