export DEBUG_I2P=true

build:
	go build -o jumpserver ./jumpd

run:
	./jumpserver