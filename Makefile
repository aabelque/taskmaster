build:
	go build -o taskmasterctl ./cmd/client 
	go build -o taskmasterd ./cmd/server 

clean:
	rm -f taskmasterctl taskmasterd
