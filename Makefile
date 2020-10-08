GOPATH = $(PWD)/gopath
WD = $(PWD)

tboped: tboped.go Makefile $(GOPATH)/.ready
	go build

$(GOPATH)/.ready:
	mkdir -p $(GOPATH)
	go get -d
	cd $(GOPATH)/src/github.com/rivo/tview && \
		git checkout fe953220389fed743d3d5ee872695abab0ae68e0 && \
		patch -p1 -N -r- -l < $(WD)/tview.patch
	cd $(GOPATH)/src/github.com/gdamore/tcell && git checkout cb1e5d6fa606f3468b7af6ed4ae00a7fee63ce19
	touch $@

values.yaml: Makefile
	curl -O https://raw.githubusercontent.com/turbonomic/t8c-install/master/operator/helm-charts/xl/values.yaml