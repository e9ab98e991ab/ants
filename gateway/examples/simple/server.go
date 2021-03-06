package main

import (
	tp "github.com/henrylee2cn/teleport"
	micro "github.com/henrylee2cn/tp-micro"
	"github.com/henrylee2cn/tp-micro/discovery"
	"github.com/henrylee2cn/tp-micro/discovery/etcd"
	"github.com/xiaoenai/ants/helper"
	html "github.com/xiaoenai/ants/helper/mod-html"
)

func init() {
	// html.ParseGlob("*.tpl")

	html.Parse("home", `<!DOCTYPE html>
<html>
	<head>
	    <title>Home</title>
	</head>
	<body>
		<br/>
		<h2><center>{{.}}</center></h2>
	</body>
</html>`)
}

// Home HTML home page
func Home(ctx tp.PullCtx, args *struct{}) ([]byte, *tp.Rerror) {
	return html.Render(ctx, "home", "Home Page Test!")
}

// Home2 HTML home page
func Home2(ctx tp.PullCtx, args *struct{}) ([]byte, *tp.Rerror) {
	return nil, helper.Redirect(ctx, 302, "http://localhost:5000/home")
}

// Args args
type Args struct {
	A int
	B int `param:"<range:1:>"`
}

// Math handler
type Math struct {
	tp.PullCtx
}

// Divide divide API
func (m *Math) Divide(args *Args) (int, *tp.Rerror) {
	return args.A / args.B, nil
}

func main() {
	cfg := micro.SrvConfig{
		ListenAddress:   ":9090",
		EnableHeartbeat: true,
		PrintDetail:     true,
	}
	srv := micro.NewServer(cfg, discovery.ServicePlugin(
		cfg.InnerIpPort(),
		etcd.EasyConfig{
			Endpoints: []string{"http://127.0.0.1:2379"},
		},
	))
	srv.RoutePullFunc(Home)
	srv.RoutePullFunc(Home2)
	srv.RoutePull(new(Math))
	srv.ListenAndServe()
}
