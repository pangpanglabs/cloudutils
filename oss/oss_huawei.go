package oss

import (
	lib_bytes "bytes"
	"io/ioutil"
	"obs"
	"path"
	"strings"
)

//HuaweiOss struct
type HuaweiOss struct {
	obsClient  *obs.ObsClient
	bucketName string
	rootDir    string
}

func (huawaiOss *HuaweiOss) init(ossConfig Config) (err error) {
	if huawaiOss.obsClient, err = obs.New(ossConfig.AccessKeyID, ossConfig.AccessKeySecret, ossConfig.Endpoint); err != nil {
		return
	}
	huawaiOss.bucketName = ossConfig.BucketName
	huawaiOss.rootDir = ossConfig.RootDir + "/"
	return
}

//PutObject func
func (huawaiOss *HuaweiOss) PutObject(objectName string, bytes []byte) (err error) {
	objectFullName := path.Join(huawaiOss.rootDir, objectName)
	input := &obs.PutObjectInput{}
	input.Bucket = huawaiOss.bucketName
	input.Key = objectFullName
	input.Body = lib_bytes.NewReader(bytes)
	_, err = huawaiOss.obsClient.PutObject(input)
	return
}

//GetObject func
func (huawaiOss *HuaweiOss) GetObject(objectName string) (bytes []byte, err error) {
	objectFullName := path.Join(huawaiOss.rootDir, objectName)
	input := &obs.GetObjectInput{}
	input.Bucket = huawaiOss.bucketName
	input.Key = objectFullName
	output, err := huawaiOss.obsClient.GetObject(input)
	if err != nil {
		return
	}
	defer output.Body.Close()
	bytes, err = ioutil.ReadAll(output.Body)
	return
}

//IsObjectExist func
func (huawaiOss *HuaweiOss) IsObjectExist(objectName string) (isExist bool, err error) {
	objectFullName := path.Join(huawaiOss.rootDir, objectName)
	input := &obs.GetObjectMetadataInput{}
	input.Bucket = huawaiOss.bucketName
	input.Key = objectFullName
	_, err = huawaiOss.obsClient.GetObjectMetadata(input)
	if err == nil {
		return true, nil
	} else if obsError, ok := err.(obs.ObsError); ok && obsError.StatusCode == 404 {
		return false, nil
	}
	return false, err
}

//ListObjects func
func (huawaiOss *HuaweiOss) ListObjects(subDir string) (objectNames []string, err error) {
	value := huawaiOss.rootDir
	if len(subDir) > 0 {
		value = path.Join(huawaiOss.rootDir, subDir)
	}

	input := &obs.ListObjectsInput{}
	input.Bucket = huawaiOss.bucketName
	input.MaxKeys = 100
	input.Prefix = value
	for {
		var output *obs.ListObjectsOutput
		output, err = huawaiOss.obsClient.ListObjects(input)
		if err != nil {
			return
		}
		for _, val := range output.Contents {
			if strings.HasSuffix(val.Key, "/") {
				continue
			}
			objectNames = append(objectNames, val.Key)
		}
		if output.IsTruncated {
			input.Marker = output.NextMarker
		} else {
			break
		}
	}
	return
}
