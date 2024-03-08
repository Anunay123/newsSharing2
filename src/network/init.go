package network

var (
	CommonRestService RestInterface
)

func DoInit() {
	CommonRestService = GetRestService()
}
