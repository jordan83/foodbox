package context

import (
	"appengine"
	"appengine/blobstore"
	"net/http"
	"errors"
)

type FileUploadUrlProvider interface {
	CreateUploadUrl(url string) string
}

type blobstoreFileUploadUrlProvider struct {
	context appengine.Context
}

func NewFileUploadUrlProvider(context appengine.Context) FileUploadUrlProvider {
	return &blobstoreFileUploadUrlProvider {
		context,
	}
}

func (urlProvider *blobstoreFileUploadUrlProvider) CreateUploadUrl(url string) string {
	uploadUrl, _ := blobstore.UploadURL(urlProvider.context, url, nil)
	return uploadUrl.String()
}

type ImageStore interface {
	GetUploadedImageKey(r *http.Request, name string) (err error, key string)
	WriteImage(w http.ResponseWriter, key string)
}

type blobstoreImageStore struct {
	context appengine.Context
}

func NewImageStore(context appengine.Context) ImageStore {
	return &blobstoreImageStore {
		context,
	}
}

func (imageStore *blobstoreImageStore) GetUploadedImageKey(r *http.Request, name string) (err error, key string) {

	var blobs map[string][]*blobstore.BlobInfo
	blobs, _, err = blobstore.ParseUpload(r)
	if err != nil {
		key = ""
        return
	}
	
	file := blobs[name]

	if len(file) == 0 {
		err = errors.New("No file uploaded")
        key = ""
        return
	}
	
	key = string(file[0].BlobKey)
	err = nil
	return 
}

func (imageStore *blobstoreImageStore) WriteImage(w http.ResponseWriter, key string) {
	blobstore.Send(w, appengine.BlobKey(key))
}