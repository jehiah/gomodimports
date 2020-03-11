module myrepo.com/package

go 1.13

require (
	vcs.com/pkg/testing v1.5.1
	vcs.com/other-packages v0.0.0
)

replace (
	vcs.com/pkg/testing => ../testing
)
