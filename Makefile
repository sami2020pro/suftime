SUFTIME_INSTALLATION_PATH = /usr/bin
DATA_INSTALLATION_PATH = /etc/suftime

help:
	@echo -e "First: make compile\nSecond: sudo make install"

compile:
	@go build ./src/suftime.go
	@mv ./suftime ./src

install:
	@if [ -d $(DATA_INSTALLATION_PATH) ]; then \
		rm -r $(DATA_INSTALLATION_PATH); \
	fi
	@cp ./src/suftime $(SUFTIME_INSTALLATION_PATH)
	@mkdir $(DATA_INSTALLATION_PATH)
	@cp ./assets/help.txt $(DATA_INSTALLATION_PATH)
	@echo -e "\x1b[37msuftime\x1b[0m installed successfuly in '$(SUFTIME_INSTALLATION_PATH)'.\nRun suftime"
