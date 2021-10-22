package schemas

import _ "embed"

//go:embed step.schema.json
var StepSchema string

//go:embed env_var.schema.json
var EnvVarSchema string
