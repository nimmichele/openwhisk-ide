package main

func dockerVersion() string {
	return Sys("@docker version --format {{.Server.Version}}")
}
