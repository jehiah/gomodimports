#/bin/bash
set -e
go build
./gomodimports -f testdata/bad.mod > testdata/output.mod
cmp testdata/good.mod testdata/output.mod