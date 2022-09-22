build:
	go build -o farcaster-export main.go

abigen:
	abigen --out ./contracts/farcaster_registry/farcaster_registry.go --pkg farcaster_registry --type FarcasterRegistry --abi ./contracts/farcaster_registry/abi.json
