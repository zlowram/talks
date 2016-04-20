// Any service must implement the following interface

type Service interface {
	Name() string
	GetBanner(ip string, port string) Banner
}

g := gsd.NewGsd([]string{"127.0.0.1","192.168.1.2"}, []string{"80","443"})
services := []gsd.Service{
	gsd.NewHttpsService(),
	gsd.NewHttpService(),
	gsd.NewTCPService(),
	gsd.NewTCPTLSService(),
}
g.AddServices(services)
