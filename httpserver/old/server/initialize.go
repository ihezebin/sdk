package server

type Initiator func()

func (s *Server) Init(i Initiator) *Server {
	i()
	return s
}
