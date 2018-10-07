package box_test

import (
	"box"
	"github.com/Hexilee/gotten"
	"github.com/Hexilee/gotten/unmarshalers"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestUpload(t *testing.T) {
	resp, err := box.GetService().Upload(&box.UploadParam{File: gotten.FilePath("testAssets/avatar.jpg")})
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	var result box.UploadResult
	assert.Nil(t, resp.Unmarshal(&result))
	assert.Equal(t, "avatar.jpg", result.Data.Filename)
}

func TestDownload(t *testing.T) {
	resp, err := box.GetService().Upload(&box.UploadParam{File: gotten.FilePath("testAssets/avatar.jpg")})
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	var result box.UploadResult
	assert.Nil(t, resp.Unmarshal(&result))
	assert.Equal(t, "avatar.jpg", result.Data.Filename)

	resp, err = box.GetService().Download(&box.TokenParam{Token: result.Data.Token})
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	var info unmarshalers.FileInfo
	assert.Nil(t, resp.Unmarshal(&info))
	assert.Equal(t, "avatar.jpg", info.Filename)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func TestChange(t *testing.T) {
	resp, err := box.GetService().Upload(&box.UploadParam{File: gotten.FilePath("testAssets/avatar.jpg")})
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	var result box.UploadResult
	assert.Nil(t, resp.Unmarshal(&result))
	assert.Equal(t, "avatar.jpg", result.Data.Filename)

	newToken := RandStringRunes(10)
	resp, err = box.GetService().Change(&box.ChangeParam{
		NewToken:   newToken,
		OldToken:   result.Data.Token,
		SecureId:   result.Data.SecureId,
		Expiration: 86400,
	})
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	var changeResult box.ChangeResult
	assert.Nil(t, resp.Unmarshal(&changeResult))
	assert.Equal(t, 0, changeResult.Status)
}