.PHONY: build-binary
build:
	go build -o bin/qp_flutter

.PHONY: clean-binary
clean-build:
	rm -f bin/qp_flutter

.PHONY: move-binary
move:
	@echo "This target requires sudo privileges"
	sudo mv bin/qp_flutter /usr/local/bin/