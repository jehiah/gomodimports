#/bin/bash
set -e
go build
./gomodimports -f testdata/bad.mod > testdata/output.mod
cmp testdata/good.mod testdata/output.mod
GOT=$(./gomodimports -f testdata/bad.mod -l)
if [ -n "${GOT}" ]; then
    echo "expected no output from gomodimports; got ${GOT}"
    exit 1
fi
echo "" >> testdata/output.mod
GOT=$(./gomodimports -f testdata/bad.mod -l)
if [ "${GOT}" != "testdata/bad.mod" ]; then
    echo "expected 'testdata/bad.mod' from gomodimports; got ${GOT}"
    exit 1
fi
