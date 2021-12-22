module myrepo.com/package

go 1.13

require (
	vcs.com/other-packages v0.0.0
	vcs.com/pkg/testing v1.5.1
)

require vcs.com/pkg/testing/b v1.5.1 // indirect

// comment about pkg/testing
replace vcs.com/pkg/testing => ../testing
