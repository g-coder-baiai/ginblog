package model

import (
	"context"
	"ginblog/utils"
	"ginblog/utils/errmsg"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"log"
	"mime/multipart"
)


var AccessKey = utils.AccessKey
var SecretKey = utils.SecretKey
var Bucket = utils.Bucket
var ImgUrl = utils.QiniuServer


func UpLoadFile(file multipart.File,fileSize int64)(string,int){
	putPolicy:= storage.PutPolicy{
		Scope: Bucket,
	}
	mac:=qbox.NewMac(AccessKey,SecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg:=storage.Config{
		Zone:&storage.ZoneHuabei,
		UseCdnDomains: false,  //要钱的
		UseHTTPS: false,       //要钱的
	}

	putExtra := storage.PutExtra{}

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	err:=formUploader.PutWithoutKey(context.Background(),&ret,upToken,file,fileSize,&putExtra)
	if err!=nil{
		log.Println(err)
		return "",errmsg.ERROR
	}

	url := ImgUrl + ret.Key
	// url := fmt.Sprintf("%s%s",ImgUrl,ret.Key)

	return url,errmsg.SUCCSE
}