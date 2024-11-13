package middlware

type Middleware struct{}

type Config struct {
}

func New(conf Config) *Middleware {
	return &Middleware{}
}

type logMeta struct {
	Methods  string `json:"methods"`
	Path     string `json:"path"`
	Latency  string `json:"latency"`
	Status   int    `json:"status"`
	Response string `json:"response"`
}
