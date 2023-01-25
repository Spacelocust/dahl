RUN = go run main.go

go-conf:
	${RUN} run go-test \
		--props '{ "label": { "valid": "avec conf" } }'

go-no-conf: 
	${RUN} template run \
		--from /golang/go-test.dahl \
		--to ./temp-test \
		-n test \
		-s Controller \
		-e .go \
		--props '{ "package": "test", "nameFunc": "isWorking", "label": { "valid": "sans conf" } }'

.PHONY: c
go-cmd:
	${RUN} run go-cmd -n ${c} --to ./cmd/${c}

go-dahl:
	${RUN} ${c}