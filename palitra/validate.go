package main

const (
	maxWidth         = 700
	maxLimit         = 5
	errValidateWidth = "The width field must be between 1 and 700"
	errValidateURL   = "The url field is required"
	errValidateLimit = "The limit field must be between 1 and 5"

	fieldWidth = "width"
	fieldURL   = "url"
	fieldLimit = "limit"
)

func (p *PalitraParams) validate() (PalitraErrors, bool) {
	errs := make(PalitraErrors)

	if p.Width > maxWidth || p.Width == 0 {
		errs[fieldWidth] = errValidateWidth
	}

	if p.URL == "" {
		errs[fieldURL] = errValidateURL
	}

	if p.Limit > maxLimit {
		errs[fieldLimit] = errValidateLimit
	}

	hasErrors := len(errs) != 0
	return errs, hasErrors
}
