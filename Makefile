.PHONY: plugin rpc-plugin

submodules:
	git submodule update --init --recursive
drafter:
	$(MAKE) -C ./adapter/ext/drafter drafter
clean:
	$(MAKE) -C ./adapter/ext/drafter distclean
plugin:
	go build -buildmode=plugin -o drafter.so ./plugin
rpc-plugin:
	go build -o drafter-rpc ./rpc-plugin