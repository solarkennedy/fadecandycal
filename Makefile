go:
	go build .
	bash -c 'FADECANDYCAL_DATE=`date  "+%B %d"` ./fadecandycal'

stop:
	ssh root@fadecandycal -- /etc/init.d/fadecandycal stop

restart:
	ssh root@fadecandycal -- /etc/init.d/fadecandycal restart

deploy: fadecandycal.mips
	scp fadecandycal.mips root@fadecandycal:/tmp/

fadecandycal.mips: fadecandycal.go colors/colors.go
	GOOS=linux GOARCH=mips GOMIPS=softfloat go build -o fadecandycal.mips .

test:
	go test -v .
	cd colors && go test -v

fmt:
	go fmt ...

clean:
	rm fadecandycal.mips fadecandycal
