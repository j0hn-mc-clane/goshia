all: install_golang create_dirs copy_config build_goshia install_goshia init_service
local: create_dirs copy_config build_goshia install_goshia init_service
install_golang:
	rm -rf /root/.go
	curl -L https://git.io/vQhTU | bash
	rm -f /usr/bin/go
	ln -s /root/.go/bin/go /usr/bin/go

create_dirs:
	echo "Creating required directories..."
	mkdir -p /home/goshia
	mkdir -p /var/log/goshia

copy_config:
	echo "Creating configuration..."
	cp configuration.yml /home/goshia/configuration.yml

build_goshia:
	echo "Building binary"
	go build

install_goshia:
	export GIN_MODE=release
	echo "Stopping existing goshia service..."
	systemctl stop goshia &>/dev/null || echo "goshia not found..."
	echo "Copying executable to /usr/bin/goshia"
	cp goshia /usr/bin/goshia

init_service:
	echo "Init goshia as service..."
	cp ./goshia.service /etc/systemd/system/goshia.service
	systemctl daemon-reload
	systemctl enable goshia
	systemctl start goshia