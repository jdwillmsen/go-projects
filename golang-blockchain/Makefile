go-download:
	go install golang.org/dl/go1.18.10@latest && \
    go1.18.10 download
tidy:
	go1.18.10 mod tidy

create-tmp:
	mkdir tmp -p

# CLI Commands
print:
	go1.18.10 run main.go printchain
bal:
	go1.18.10 run main.go getbalance -address $(address)
bc:
	go1.18.10 run main.go createblockchain -address $(address)
send:
	go1.18.10 run main.go send -amount $(amount) -from $(from) -to $(to)
wallet: create-tmp
	go1.18.10 run main.go createwallet
list:
	go1.18.10 run main.go listaddresses
reindex:
	go1.18.10 run main.go reindexutxo

