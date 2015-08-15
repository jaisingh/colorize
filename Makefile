SRC=colorize.go
OUT=colorize
INST=$(GOPATH)/bin/$(OUT)

BUILDCMD=go build -o $(OUT) $(SRC)
TESTCMD=go test

build: $(OUT)

$(OUT): $(SRC)
	$(BUILDCMD)

$(INST): $(OUT)
	cp $(OUT) $(INST)

install: $(INST)

clean:
	rm $(OUT)

superclean: clean
	rm $(INST)

test:
	$(TESTCMD)
