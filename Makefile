HELPERS = internal/cli/add_account.go \
		  internal/cli/add.go \
		  internal/cli/init.go \
		  internal/cli/list.go \
		  internal/cli/remove.go \
		  internal/cli/edit.go \
		  internal/util/decrypt_store.go \
		  internal/util/encrypt.go \
		  internal/util/encrypt_store.go \
		  internal/util/get_credentials.go \
		  internal/util/get_db.go \
		  internal/util/pad.go \
		  internal/util/verify.go \
		  internal/server/start.go

.PHONY: all client server clean

all: client server

client: bin/pipasso

bin/pipasso: cmd/client/main.go $(HELPERS)
	go build -gcflags '-N -l' -o $@ $<

server: bin/pipasso-server

bin/pipasso-server: cmd/server/main.go $(HELPERS)
	go build -gcflags '-N -l' -o $@ $<

clean:
	rm -r bin
