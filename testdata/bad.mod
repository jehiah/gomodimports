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

