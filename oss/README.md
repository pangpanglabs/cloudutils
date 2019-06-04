# cloudutils/oss

Install `github.com/pangpanglabs/cloudutils/oss` package.
```golang
go get -u github.com/pangpanglabs/cloudutils/oss
```

## Getting Started

Import `github.com/pangpanglabs/cloudutils/oss` package.
```golang
import "github.com/pangpanglabs/cloudutils/oss"
```

Use `oss.Client` interface type
```golang
var ossClient oss.Client
```

Create new oss clinet
```golang
ossConfig = oss.Config{
    Endpoint:        "Your endpoint",    // ex. oss-cn-shanghai.aliyuncs.com
    BucketName:      "Your bucket name", // ex. web-static-files
    RootDir:         "Your root dir",    // ex. fos
    AccessKeyID:     "Your AK",          // AccessKeyID
    AccessKeySecret: "Your SK",          // AccessKeySecret
}
ossClient, err := oss.New(ossConfig)     // Create new oss clinet
if err != nil {
    fmt.Println(err)
    return
}
```

Put a new object
```golang
objectName := fmt.Sprintf("test/%s.txt", time.Now().Format("2006-01-02_15-04-05"))
fmt.Println("PubObject:", objectName)
if err := ossClient.PutObject(objectName, []byte("yourObjectValueByteArrary")); err != nil {
    fmt.Println(err)
    return
} else {
    fmt.Println("PutObject Success")
}
```

Get a object
```golang
fmt.Println("GetObject:")
if content, err := ossClient.GetObject(objectName); err != nil {
    fmt.Println(err)
    return
} else {
    fmt.Println(string(content))
}
```

Browse objects in root dir: /bucket_name/root_dir/*
```golang
fmt.Println("ListObjects:")
if files, err := ossClient.ListObjects(""); err != nil {
    fmt.Println(err)
    return
} else {
    for _, file := range files {
        fmt.Println(file)
    }
}
```

Browse objects in specific dir: /bucket_name/root_dir/test/*
```golang
fmt.Println("ListObjects:")
if files, err := ossClient.ListObjects("test"); err != nil {
    fmt.Println(err)
    return
} else {
    for _, file := range files {
        fmt.Println(file)
    }
}
```

Check the object exists
```golang
fmt.Println("IsObjectExist:")
if isExist, err := ossClient.IsObjectExist(objectName); err != nil {
    fmt.Println(err)
    return
} else {
    fmt.Println(isExist)
}
```
