.PHONY: all clean

vpath fanyi% cli

all: fanyi

fanyi: fanyi.go
	go build -ldflags -w -o $@ $<

clean:
	-$(RM) fanyi
