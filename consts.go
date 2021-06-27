package shellcligen

const (
	scriptFileName               = "script.sh"
	scriptConfigFileName         = "script.conf"
	templateWithConflictChecking = `hello`
	safeFlagsTemplateTag         = `@safe_flags@`
	safeFlagsTemplate            = `
set -o errexit
set -o nounset 
set -o pipefail	
`
)
