VERSION := 1.0.0

build:
	go build

install_deps:
	go get gopkg.in/alecthomas/kingpin.v1
	go get github.com/jackwakefield/gopac
	go get github.com/elazarl/goproxy
	go get github.com/spf13/cobra
	go get github.com/spf13/viper

dist: BASE := transpac-$(VERSION)
dist:
	rm -f $(BASE).tar $(BASE).tar.bz2
	git archive --format=tar.gz --prefix $(BASE)/ -o $(BASE).tar.gz HEAD
