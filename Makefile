
build:
	go build -o cfsync-proxy

install:
	install -t /usr/bin cfsync-proxy

clean:
	rm -f /usr/bin/cfsync-proxy
	rm -f cfsync-proxy

run:
	/usr/bin/cfsync-proxy

dpkg:
	mkdir -p debian
