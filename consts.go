package shellcligen

const (
	templateWithConflictChecking = `hello`
	safeFlagsTemplateTag         = `@safe_flags@`
	safeFlagsTemplate            = `
set -o errexit
set -o nounset 
set -o pipefail	
`
)
