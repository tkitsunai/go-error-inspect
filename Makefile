CMD_PATH=/usr/local/go/bin/go

vtest:
	${CMD_PATH} test -v ./... -count=1
