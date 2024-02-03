all: compile

compile:
	@echo "Compiling"
	env GOOS=linux GOARCH=arm GOARM=5 go build
	scp ./family-dashboard pi@epaper01:~/family-dashboard/