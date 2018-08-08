.PHONY: dep

dep:
	dep ensure -update -v
	rm -rf vendor
