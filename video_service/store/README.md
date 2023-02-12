# 文件上传功能

这里也是面向接口编程，因为文件上传可能会有很多总方式：
1. 上传本地
2. 上传阿里云
3. 上传腾讯云
4. ...

所以面向接口编程是一个不错的选择

```go
// Uploader 文件上传接口
type Uploader interface {
	// Upload 上传文件到云端
	Upload(*UploadParam) (*UploadResult, error)
}
```

当然，目前只有一个上传的接口，按理来说至少还需要一个删除废物文件的接口.
