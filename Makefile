# SPDX-FileCopyrightText: © 2021-2022 Nadim Kobeissi <nadim@symbolic.software>
# SPDX-License-Identifier: GPL-3.0-only

all:
	@go install
	@make -s windows
	@make -s linux
	@make -s macos
	@make -s freebsd

windows:
	@GOOS="windows" GOARCH="amd64" go build -trimpath -gcflags="-e" -ldflags="-s -w" -o build/windows .

linux:
	@GOOS="linux" GOARCH="amd64" go build -trimpath -gcflags="-e" -ldflags="-s -w" -o build/linux .

macos:
	@GOOS="darwin" GOARCH="arm64" go build -trimpath -gcflags="-e" -ldflags="-s -w" -o build/macos .
	@mv build/macos/cpr build/macos/cpr_applesilicon
	@GOOS="darwin" GOARCH="amd64" go build -trimpath -gcflags="-e" -ldflags="-s -w" -o build/macos .

freebsd:
	@GOOS="freebsd" GOARCH="amd64" go build -trimpath -gcflags="-e" -ldflags="-s -w" -o build/freebsd .

lint:
	@/bin/echo "[Verifpal] Running golangci-lint..."
	@golangci-lint run

clean:
	@$(RM) build/*/cp*

.PHONY: all windows linux macos freebsd lint clean build
