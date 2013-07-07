package context

import (
	"appengine"
	"appengine/blobstore"
	"net/http"
	"errors"
)

type ImageStore interface {
	GetUploadedImageKey(r *http.Request, name string) (err error, key string)
	WriteImage(w http.ResponseWriter, key string)
	CreateUploadUrl(url string) string
	RemoveImages(imageKeys []string)
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

func (imageStore *blobstoreImageStore) CreateUploadUrl(url string) string {
	uploadUrl, _ := blobstore.UploadURL(imageStore.context, url, nil)
	return uploadUrl.String()
}

func (imageStore *blobstoreImageStore) RemoveImages(imageKeys []string) {
	keys := make([]appengine.BlobKey, len(imageKeys))
	for i := 0; i < len(imageKeys); i++ {
		keys[i] = appengine.BlobKey(imageKeys[i])
	}
	
	blobstore.DeleteMulti(imageStore.context, keys)
}