
## gomodimports

Like `goimports` but for keeping a tidy go.mod file


```
$ gomodimports -w -f go.mod
```

Takes an ugly go mod file

```
module myrepo.com/package

go 1.13

require vcs.com/pkg/testing v1.5.1
replace vcs.com/pkg/testing => ../testing
require (
	vcs.com/other-packages v0.0.0
)
```

and turns it inoto a go.mod file with a single `require` and `replace` block

```
module myrepo.com/package

go 1.13

require (
        vcs.com/pkg/testing v1.5.1
        vcs.com/other-packages v0.0.0
)

replace (
        vcs.com/pkg/testing => ../testing
)
```