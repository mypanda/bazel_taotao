all:
	bazel build //...
	cat bazel-out/stable-status.txt
test:
	bazel test //...
gofmt:
	gofmt -w -s pkg/ cmd/
gazelle:
	bezel run //:gazelle
clean:
	bazel clean
