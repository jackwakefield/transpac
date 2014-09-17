build:
	go build

install_deps:
	go get gopkg.in/alecthomas/kingpin.v1
	go get github.com/jackwakefield/gopac
	go get gopkg.in/elazarl/goproxy.v1
	go get github.com/spf13/cobra
	go get github.com/spf13/viper
