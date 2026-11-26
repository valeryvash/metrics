.PHONY: statictest sbuild sclean 

statictest:
	STATIC_TEST=$$(command -v statictest); \
	echo "Using STATIC_TEST: $$STATIC_TEST"; \
	go vet -v -vettool=$$STATIC_TEST ./...

sbuild:
	go build -o ./cmd/server/server ./cmd/server/main.go

sclean:
	rm ./cmd/server/server 