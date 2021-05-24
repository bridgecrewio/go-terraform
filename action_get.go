package goterraform

type TerraformGetParams struct {
	Directory string
}

func (t TerraformGetParams) Opts() map[string][]string {
	opts := make(map[string][]string)

	if t.Directory != "" {
		opts["directory"] = []string{t.Directory}
	}
	return opts
}

func (t TerraformGetParams) OptsString() string {
	return extractOptsString(t)
}

func (t TerraformGetParams) OptsStringSlice() []string {
	return extractOptsStringSlice(t)
}
