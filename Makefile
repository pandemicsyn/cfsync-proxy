DESTDIR =
PREFIX  = /usr
BINDIR 	= $(PREFIX)/bin

build:
		go build -o cfsync-proxy

install:
		install -d $(DESTDIR)$(BINDIR)
		install cfsync-proxy $(DESTDIR)$(BINDIR)

clean:
		rm -f cfsync-proxy

run:
		/usr/bin/cfsync-proxy

dpkg:
		mkdir -p debian
