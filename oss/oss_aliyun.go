package oss

import (
	"bytes"
	"io"
	"io/ioutil"
	"path"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

//AliyunOss struct
type AliyunOss struct {
	client  *oss.Client
	bucket  *oss.Bucket
	rootDir string
}

//Init func
func (aliyunOss *AliyunOss) init(ossConfig Config) (err error) {
	// 创建OSSClient实例。
	if aliyunOss.client, err = oss.New(ossConfig.Endpoint, ossConfig.AccessKeyID, ossConfig.AccessKeySecret); err != nil {
		return
	}
	if aliyunOss.bucket, err = aliyunOss.client.Bucket(ossConfig.BucketName); err != nil {
		return
	}
	aliyunOss.rootDir = ossConfig.RootDir + "/"
	return
}

//PutObject func
func (aliyunOss *AliyunOss) PutObject(objectName string, object []byte) (err error) {
	objectFullName := path.Join(aliyunOss.rootDir, objectName)
	err = aliyunOss.bucket.PutObject(objectFullName, bytes.NewReader(object))
	return
}

//GetObject func
func (aliyunOss *AliyunOss) GetObject(objectName string) (bytes []byte, err error) {
	objectFullName := path.Join(aliyunOss.rootDir, objectName)
	var body io.ReadCloser
	body, err = aliyunOss.bucket.GetObject(objectFullName)
	if err != nil {
		return
	}
	defer body.Close()
	bytes, err = ioutil.ReadAll(body)
	return
}

//IsObjectExist func
func (aliyunOss *AliyunOss) IsObjectExist(objectName string) (isExist bool, err error) {
	objectFullName := path.Join(aliyunOss.rootDir, objectName)
	isExist, err = aliyunOss.bucket.IsObjectExist(objectFullName)
	return
}

//ListObjects func
func (aliyunOss *AliyunOss) ListObjects(subDir string) (objectNames []string, err error) {
	value := aliyunOss.rootDir
	if len(subDir) > 0 {
		value = path.Join(aliyunOss.rootDir, subDir)
	}

	prefix := oss.Prefix(value)
	marker := oss.Marker("")
	for {
		var lsRes oss.ListObjectsResult
		lsRes, err = aliyunOss.bucket.ListObjects(oss.MaxKeys(80), marker, prefix)
		if err != nil {
			return
		}
		prefix = oss.Prefix(lsRes.Prefix)
		marker = oss.Marker(lsRes.NextMarker)
		for _, objectProps := range lsRes.Objects {
			if strings.HasSuffix(objectProps.Key, "/") {
				continue
			}
			objectNames = append(objectNames, objectProps.Key)
		}
		if !lsRes.IsTruncated {
			break
		}
	}
	return
}
