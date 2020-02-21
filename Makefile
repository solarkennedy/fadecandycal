go:
	go build .
	bash -c 'FADECANDYCAL_DATE=`date  "+%B %d"` ./fadecandycal'

stop:
	ssh root@10.0.2.113 -- /etc/init.d/fadecandycal stop

restart:
	ssh root@10.0.2.113 -- /etc/init.d/fadecandycal restart

deploy: fadecandycal.mips
	scp fadecandycal.mips root@10.0.2.113:/tmp/

fadecandycal.mips: fadecandycal.go
	GOOS=linux GOARCH=mips GOMIPS=softfloat go build -o fadecandycal.mips .

test:
	go test -v .
	cd colors && go test -v

fmt:
	go fmt ...

clean:
	rm fadecandycal.mips fadecandycal
