.PHONY: all
all: slide22 slide27 slide30

%: %.go
	go build -o $@ $<

.PHONY: clean
clean:
	rm -f slide22 slide27 slide30
