package context

import (
	"appengine"
	"appengine/blobstore"
	"net/http"
	"io"
	"errors"
)

type FileStore interface {
	GetUploadedFileKey(r *http.Request, name string) (err error, key string)
	WriteFile(w http.ResponseWriter, key string)
	CreateUploadUrl(url string) string
	RemoveFiles(fileKeys []string)
	GetReader(key string) io.Reader
}

type blobstoreFileStore struct {
	context appengine.Context
}

func NewFileStore(context appengine.Context) FileStore {
	return &blobstoreFileStore {
		context,
	}
}

func (fileStore *blobstoreFileStore) GetUploadedFileKey(r *http.Request, name string) (err error, key string) {

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

func (fileStore *blobstoreFileStore) WriteFile(w http.ResponseWriter, key string) {
	blobstore.Send(w, appengine.BlobKey(key))
}

func (fileStore *blobstoreFileStore) CreateUploadUrl(url string) string {
	uploadUrl, _ := blobstore.UploadURL(fileStore.context, url, nil)
	return uploadUrl.String()
}

func (fileStore *blobstoreFileStore) RemoveFiles(fileKeys []string) {
	keys := make([]appengine.BlobKey, len(fileKeys))
	for i := 0; i < len(fileKeys); i++ {
		keys[i] = appengine.BlobKey(fileKeys[i])
	}
	
	blobstore.DeleteMulti(fileStore.context, keys)
}

func (fileStore *blobstoreFileStore) GetReader(key string) io.Reader {
	reader := blobstore.NewReader(fileStore.context, appengine.BlobKey(key))
	return reader
}