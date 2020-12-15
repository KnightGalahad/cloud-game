package shared

import (
	"github.com/giongto35/cloud-game/v2/pkg/environment"
	"github.com/spf13/pflag"
)

type Environment environment.Env

type Server struct {
	Port       int
	HttpsPort  int
	HttpsKey   string
	HttpsChain string
}

func (s *Server) WithFlags(fs *pflag.FlagSet) {
	fs.IntVar(&s.Port, "port", 8000, "HTTP server port")
	fs.IntVar(&s.HttpsPort, "httpsPort", 443, "HTTPS server port (just why?)")
	fs.StringVar(&s.HttpsKey, "httpsKey", "", "HTTPS key")
	fs.StringVar(&s.HttpsChain, "httpsChain", "", "HTTPS chain")
}

func (env *Environment) Get() environment.Env {
	return (environment.Env)(*env)
}

func (env *Environment) WithFlags(fs *pflag.FlagSet) {
	val := string(*env)
	fs.StringVar(&val, "env", "dev", "Specify environment type: [dev, staging, prod]")
}