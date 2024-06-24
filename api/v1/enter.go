package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/bbs"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	videoApi "github.com/flipped-aurora/gin-vue-admin/server/api/v1/video"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
	BbsApiGroup     bbsapi.ApiGroup
	VideoApiGroup   videoApi.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
