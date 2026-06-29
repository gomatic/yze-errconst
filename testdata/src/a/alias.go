package a

import e "errors"

// errAliased is created via an aliased errors import; the callee still resolves
// to errors.New through its fully-qualified name, so it must be flagged.
var errAliased = e.New("aliased") // want `use a sentinel error constant \(errs\.Const\) instead of errors\.New`
