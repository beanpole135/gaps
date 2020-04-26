build:
	@go get
	@go build

cleanbuild:
	@go clean

clean:
	@go clean

package:
	@go clean
	GOARCH=amd64 GOOS=linux go build -o bin/${GOOS}-${GOARCH}/gaps
	GOARCH=amd64 GOOS=windows go build -o bin/${GOOS}-${GOARCH}/gaps
	GOARCH=amd64 GOOS=darwin go build -o bin/mac-${GOARCH}/gaps
	GOARCH=amd64 GOOS=freebsd go build -o bin/${GOOS}-${GOARCH}/gaps
	GOARCH=amd64 GOOS=openbsd go build -o bin/${GOOS}-${GOARCH}/gaps
	GOARCH=amd64 GOOS=solaris go build -o bin/${GOOS}-${GOARCH}/gaps
	GOARCH=amd64 GOOS=dragonfly go build -o bin/${GOOS}-${GOARCH}/gaps
	GOARCH=amd64 GOOS=plan9 go build -o bin/${GOOS}-${GOARCH}/gaps
