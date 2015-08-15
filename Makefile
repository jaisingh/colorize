OUT="colorize"
INST="$(GOPATH)/bin/$(OUT)"

build: $(OUT)

$(OUT):
	go build $(OUT).go

$(INST): $(OUT)
	cp $(OUT) $(INST)

install: $(INST)

clean:
	rm $(OUT)
