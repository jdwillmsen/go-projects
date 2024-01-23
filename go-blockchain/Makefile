clear:
	clear
run:
	go run cmd/main.go
bc-1: clear
	go run blockchain_server/main.go blockchain_server/blockchain_server.go -port 5000
bc-2: clear
	go run blockchain_server/main.go blockchain_server/blockchain_server.go -port 5001
bc-3: clear
	go run blockchain_server/main.go blockchain_server/blockchain_server.go -port 5002
wallet-1: clear
	go run wallet_server/main.go wallet_server/wallet_server.go -port 8080 -gateway http://127.0.0.1:5000
wallet-2: clear
	go run wallet_server/main.go wallet_server/wallet_server.go -port 8081 -gateway http://127.0.0.1:5001
wallet-3: clear
	go run wallet_server/main.go wallet_server/wallet_server.go -port 8082 -gateway http://127.0.0.1:5002
blockchains: bc-1 bc-2 bc-3
wallets: wallet-1 wallet-2 wallet-3