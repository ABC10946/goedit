goedit:
	go build .

clean:
	rm goedit

install:
	mv goedit $(HOME)/.local/bin/goedit

uninstall:
	rm $(HOME)/.local/bin/goedit
