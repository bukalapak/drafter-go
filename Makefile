.PHONY: plugin

submodules:
	git submodule update --init --recursive
drafter:
	$(MAKE) -C ./ext/drafter drafter
clean:
	$(MAKE) -C ./ext/drafter distclean
plugin:
	go build -buildmode=plugin -o drafter.so ./plugin
rpc-plugin:
	go build -o drafter-rpc ./rpc-plugin