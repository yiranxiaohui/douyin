// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	feed "douyin/biz/router/feed"
	publish "douyin/biz/router/publish"
	user "douyin/biz/router/user"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	user.Register(r)

	feed.Register(r)

	publish.Register(r)
}
