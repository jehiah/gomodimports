
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
replace (
    // comment about pkg/testing
    vcs.com/pkg/testing => ../testing
) 
require (
	vcs.com/other-packages v0.0.0
)
```

and turns it inoto a go.mod file with a single `require` and `replace` block preserving comments.

```
module myrepo.com/package

go 1.13

require (
	vcs.com/other-packages v0.0.0
	vcs.com/pkg/testing v1.5.1
)

// comment about pkg/testing
replace vcs.com/pkg/testing => ../testing
```