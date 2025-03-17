package jumpserver

func (j *JumpServer) HostsFile() string {
	var hosts string
	for _, host := range j.Hostnames {
		hosts += host.Hostname + "=" + host.Base64() + "\n"
	}
	return hosts
}
