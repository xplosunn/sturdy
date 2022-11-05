test:
	@# -count=1 is to stop go test using a cache.
	go test ./... -count=1