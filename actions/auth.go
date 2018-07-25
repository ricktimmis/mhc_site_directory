package actions

import (
	"fmt"
	"os"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/gplus"
	"github.com/markbates/goth/providers/instagram"
	"github.com/markbates/goth/providers/linkedin"
	"github.com/markbates/goth/providers/meetup"
	"github.com/markbates/goth/providers/paypal"
	"github.com/markbates/goth/providers/twitter"
	"github.com/pkg/errors"
	"github.com/ricktimmis/mhclub/models"
)

func init() {
	gothic.Store = App().SessionStore

	goth.UseProviders(
		twitter.New(os.Getenv("TWITTER_KEY"), os.Getenv("TWITTER_SECRET"), fmt.Sprintf("%s%s", App().Host, "/auth/twitter/callback")),
		facebook.New(os.Getenv("FACEBOOK_KEY"), os.Getenv("FACEBOOK_SECRET"), fmt.Sprintf("%s%s", App().Host, "/auth/facebook/callback")),
		linkedin.New(os.Getenv("LINKEDIN_KEY"), os.Getenv("LINKEDIN_SECRET"), fmt.Sprintf("%s%s", App().Host, "/auth/linkedin/callback")),
		gplus.New(os.Getenv("GPLUS_KEY"), os.Getenv("GPLUS_SECRET"), fmt.Sprintf("%s%s", App().Host, "/auth/gplus/callback")),
		instagram.New(os.Getenv("INSTAGRAM_KEY"), os.Getenv("INSTAGRAM_SECRET"), fmt.Sprintf("%s%s", App().Host, "/auth/instagram/callback")),
		meetup.New(os.Getenv("MEETUP_KEY"), os.Getenv("MEETUP_SECRET"), fmt.Sprintf("%s%s", App().Host, "/auth/meetup/callback")),
		paypal.New(os.Getenv("PAYPAL_KEY"), os.Getenv("PAYPAL_SECRET"), fmt.Sprintf("%s%s", App().Host, "/auth/paypal/callback")),
	)
}

func AuthCallback(c buffalo.Context) error {
	user, err := gothic.CompleteUserAuth(c.Response(), c.Request())
	if err != nil {
		return c.Error(401, err)
	}

	/* Create a New User, or Do something with the user, maybe register them/sign them in

	   ToDo - Grab Avatar image, and use as profile pic
	   ToDo - Check for existing username, and
	*/
	tx := c.Value("tx").(*pop.Connection)
	q := tx.Where("loginprovider = ? and loginprovider_id = ?", user.Provider, user.UserID)
	exists, err := q.Exists(&models.User{})
	if err != nil {
		return errors.WithStack(err)
	}
	u := &models.User{}
	if exists {
		err = q.First(u)
		if err != nil {
			return errors.WithStack(err)
		}
	} else {

		u.Firstname = user.LastName
		u.Lastname = user.LastName
		u.Username = user.Name
		u.Email = user.Email
		u.LoginProvider = user.Provider
		u.LoginProviderID = user.UserID
		u.Membertype = "Guest" // All new users start as guests, and return to guest if Subscription expires
		err = tx.Save(u)
		if err != nil {
			return errors.WithStack(err)
		}
	}
	c.Session().Set("current_user_id", u.ID)
	err = c.Session().Save()
	if err != nil {
		return errors.WithStack(err)
	}
	return c.Redirect(302, "/")

	// Render raw data to browser for debugging purposes.
	return c.Render(200, r.JSON(user))
}

func SetCurrentUser(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if uid := c.Session().Get("current_user_id"); uid != nil {
			u := &models.User{}
			tx := c.Value("tx").(*pop.Connection)
			err := tx.Find(u, uid)
			if err != nil {
				return errors.WithStack(err)
			}
			c.Set("current_user", u)
		}
		return next(c)
	}
}

/*
Authorise access only if User is of a certain Membership level
see SPECIFICATION.md for Membership Benefits
FIXME Ideally this should be a single function that can determine the handler being called, and figure out if access is allowed
*/
func AuthorizeGuest(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if uid := c.Session().Get("current_user_id"); uid != nil {
			u := &models.User{}
			tx := c.Value("tx").(*pop.Connection)
			err := tx.Find(u, uid)
			if err != nil {
				return errors.WithStack(err)
			}
			if u.Membertype != "Guest" {
				c.Flash().Add("danger", "Your membership level does not grant you access to that feature")
				return c.Redirect(302, "/")
			}
		} else {
			c.Flash().Add("danger", "Your membership level does not grant you access to that feature")
			return c.Redirect(302, "/")
		}
		return next(c)
	}
}
