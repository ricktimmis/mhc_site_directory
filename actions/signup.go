package actions

import "github.com/gobuffalo/buffalo"

// SignupSignupHandler default implementation.
func SignupSignupHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("signup/signup_handler.html"))
}
