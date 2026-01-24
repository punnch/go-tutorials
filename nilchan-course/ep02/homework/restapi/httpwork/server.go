package httpwork

type HTTPServer struct {
	httpHandlers *HTTPHandlers
}

func NewHttpServer(httpHandlers *HTTPHandlers) *HTTPServer {
	return &HTTPServer{
		httpHandlers: httpHandlers,
	}
}

func (s *HTTPServer) Start() {

}
