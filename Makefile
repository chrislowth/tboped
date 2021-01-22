OPERATOR_TAG=7.22			# Git tag for the version to be used
GOPATH = $(PWD)/gopath
WD = $(PWD)
GO = go

tboped: tboped.go Makefile $(GOPATH)/.ready t8cop/* t8cop/operator_type.go ui/*.go
	$(GO) build

clean:
	rm -fv tboped t8cop/operator_type.go t8cop/structs.go
	rm -rfv t8c-install gopath

$(GOPATH)/.ready:
	mkdir -p $(GOPATH)
	$(GO) get -d
	cd tools && $(GO) get -d
	cd $(GOPATH)/src/github.com/rivo/tview && \
		git checkout fe953220389fed743d3d5ee872695abab0ae68e0 && \
		patch -p1 -N -r- -l < $(WD)/tview.patch
	cd $(GOPATH)/src/github.com/gdamore/tcell && git checkout cb1e5d6fa606f3468b7af6ed4ae00a7fee63ce19
	touch $@

values.yaml: Makefile
	curl -O https://raw.githubusercontent.com/turbonomic/t8c-install/master/operator/helm-charts/xl/values.yaml

t8c-install/.ready:
	rm -rf t8c-install
	git clone https://github.com/turbonomic/t8c-install.git
	cd t8c-install && git checkout $(OPERATOR_TAG)
	touch $@

t8cop/operator_type.go: Makefile tools/mkopgo.go t8cop/utils.go \
		t8c-install/.ready \
		t8c-install/operator/deploy/crds/charts_v1alpha1_xl_crd.yaml \
		t8c-install/operator/helm-charts/xl/requirements.yaml
	rm -f $@ tools/structs.go t8cop/structs.go
	cd tools && go run -tags mkopgo mkopgo.go ../t8c-install/operator > /tmp/operator_type.go
	uniq tools/structs.go > t8cop/structs.go
	mv /tmp/operator_type.go $@

dlv:
	$(GO) get -u github.com/go-delve/delve/cmd/dlv
