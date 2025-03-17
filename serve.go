package jumpserver

import "net/http"

func (j *JumpServer) Serve() error {
	l, err := j.Garlic.Listen()
	if err != nil {
		return err
	}
	defer l.Close()
	return http.Serve(l, j)
}
