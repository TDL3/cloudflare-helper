NAME := cloudflare-helper
PACKAGE_NAME := github.com/tdl3/cloudflare-helper
# VERSION := `git describe --dirty`
# COMMIT := `git rev-parse HEAD`

PLATFORM := linux
BUILD_DIR := build
# VAR_SETTING := -X $(PACKAGE_NAME)/constant.Version=$(VERSION) -X $(PACKAGE_NAME)/constant.Commit=$(COMMIT)
GOBUILD = env CGO_ENABLED=0 $(GO_DIR)go build -tags "full" -trimpath -ldflags="-s -w -buildid= $(VAR_SETTING)" -o $(BUILD_DIR)

.PHONY: cloudflare-helper release test
normal: clean cloudflare-helper
# all: darwin-amd64 darwin-arm64 linux-386 linux-amd64 linux-arm linux-armv8 linux-armv7 linux-armv6 windows-amd64 windows-386 windows-arm

clean:
	rm -rf $(BUILD_DIR)
	rm -f *.zip
	rm -f *.dat

cloudflare-helper:
	mkdir -p $(BUILD_DIR)
	$(GOBUILD)

install:
	mkdir -p /etc/$(NAME)
	mkdir -p /usr/share/$(NAME)
	cp example/*.yml /etc/$(NAME)
	cp $(BUILD_DIR)/$(NAME) /usr/bin/$(NAME)
	cp example/$(NAME).service /usr/lib/systemd/system/
	cp example/$(NAME)@.service /usr/lib/systemd/system/
	cp example/$(NAME).timer /usr/lib/systemd/system/
	systemctl daemon-reload

uninstall:
	rm /usr/lib/systemd/system/$(NAME).service
	rm /usr/lib/systemd/system/$(NAME)@.service
		rm /usr/lib/systemd/system/$(NAME).timer
	systemctl daemon-reload
	rm /usr/bin/$(NAME)
	rm -rd /etc/$(NAME)
	rm -rd /usr/share/$(NAME)

%.zip:
	@zip -du $(NAME)-$@ -j $(BUILD_DIR)/$</*
	@zip -du $(NAME)-$@ example/*
	@echo "<<< ---- $(NAME)-$@"

release: darwin-amd64.zip darwin-arm64.zip linux-386.zip linux-amd64.zip \
	linux-arm.zip linux-armv5.zip linux-armv6.zip linux-armv7.zip linux-armv8.zip \
	linux-mips-softfloat.zip linux-mips-hardfloat.zip linux-mipsle-softfloat.zip linux-mipsle-hardfloat.zip \
	linux-mips64.zip linux-mips64le.zip freebsd-386.zip freebsd-amd64.zip \
	windows-386.zip windows-amd64.zip windows-arm.zip windows-armv6.zip windows-armv7.zip

darwin-amd64:
	mkdir -p $(BUILD_DIR)/$@
	GOARCH=amd64 GOOS=darwin $(GOBUILD)/$@

darwin-arm64:
	mkdir -p $(BUILD_DIR)/$@
	GOARCH=arm64 GOOS=darwin $(GOBUILD)/$@

linux-386:
	mkdir -p $(BUILD_DIR)/$@
	GOARCH=386 GOOS=linux $(GOBUILD)/$@

linux-amd64:
	mkdir -p $(BUILD_DIR)/$@
	GOARCH=amd64 GOOS=linux $(GOBUILD)/$@

linux-arm:
	mkdir -p $(BUILD_DIR)/$@
	GOARCH=arm GOOS=linux $(GOBUILD)/$@

linux-armv5:
	mkdir -p $(BUILD_DIR)/$@
	GOARCH=arm GOOS=linux GOARM=5 $(GOBUILD)/$@

linux-armv6:
	mkdir -p $(BUILD_DIR)/$@
	GOARCH=arm GOOS=linux GOARM=6 $(GOBUILD)/$@

linux-armv7:
	mkdir -p $(BUILD_DIR)/$@
	GOARCH=arm GOOS=linux GOARM=7 $(GOBUILD)/$@

linux-armv8:
	mkdir -p $(BUILD_DIR)/$@
	GOARCH=arm64 GOOS=linux $(GOBUILD)/$@

linux-mips-softfloat:
	mkdir -p $(BUILD_DIR)/$@
	GOARCH=mips GOMIPS=softfloat GOOS=linux $(GOBUILD)/$@

linux-mips-hardfloat:
	mkdir -p $(BUILD_DIR)/$@
	GOARCH=mips GOMIPS=hardfloat GOOS=linux $(GOBUILD)/$@

linux-mipsle-softfloat:
	mkdir -p $(BUILD_DIR)/$@
	GOARCH=mipsle GOMIPS=softfloat GOOS=linux $(GOBUILD)/$@

linux-mipsle-hardfloat:
	mkdir -p $(BUILD_DIR)/$@
	GOARCH=mipsle GOMIPS=hardfloat GOOS=linux $(GOBUILD)/$@

linux-mips64:
	mkdir -p $(BUILD_DIR)/$@
	GOARCH=mips64 GOOS=linux $(GOBUILD)/$@

linux-mips64le:
	mkdir -p $(BUILD_DIR)/$@
	GOARCH=mips64le GOOS=linux $(GOBUILD)/$@

freebsd-386:
	mkdir -p $(BUILD_DIR)/$@
	GOARCH=386 GOOS=freebsd $(GOBUILD)/$@

freebsd-amd64:
	mkdir -p $(BUILD_DIR)/$@
	GOARCH=amd64 GOOS=freebsd $(GOBUILD)/$@

windows-386:
	mkdir -p $(BUILD_DIR)/$@
	GOARCH=386 GOOS=windows $(GOBUILD)/$@

windows-amd64:
	mkdir -p $(BUILD_DIR)/$@
	GOARCH=amd64 GOOS=windows $(GOBUILD)/$@

windows-arm:
	mkdir -p $(BUILD_DIR)/$@
	GOARCH=arm GOOS=windows $(GOBUILD)/$@

windows-armv6:
	mkdir -p $(BUILD_DIR)/$@
	GOARCH=arm GOOS=windows GOARM=6 $(GOBUILD)/$@

windows-armv7:
	mkdir -p $(BUILD_DIR)/$@
	GOARCH=arm GOOS=windows GOARM=7 $(GOBUILD)/$@