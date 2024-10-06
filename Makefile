build: install
	go build -o make/zscnetworksupport-api
install:
	go install

clean:
	rm -rf make/zscnetworksupport-api

update:

serve: build
	mkdir opt/zscww/api

	cp make/zscnetworksupport-api opt/zscww/api/zscnetorksupport-api
	cp config.json opt/zscww/api/config.json

	cp -r static /opt/zscww/api/static
	echo "构建即将完成，请输入 /opt/zscww/api/zscnetworksupport-api ${配置文件路径} 来启动API服务"
	chmod +x /opt/zscww/api/zscnetworksupport-api

help:

default: build

