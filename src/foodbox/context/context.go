package context

import (
	"appengine"
	"appengine/user"
	"net/http"
)

type FoodboxContext struct {
	requestURL string
	requestContext appengine.Context
	CurrentUser *user.User
}

func NewContext(r *http.Request) *FoodboxContext {
	c := appengine.NewContext(r)
	u := user.Current(c)
	return &FoodboxContext {
		requestURL: r.URL.String(),
		requestContext: c,
		CurrentUser: u,
	}
}

func (context *FoodboxContext) RedirectIfNotLoggedIn(w http.ResponseWriter) bool {
    if context.CurrentUser != nil {
    	return false
    }
    
    url := context.LoginURL()
    
    w.Header().Set("Location", url)
    w.WriteHeader(http.StatusFound)
    return true
}

func (context *FoodboxContext) LoginURL() string {
	url, _ := user.LoginURL(context.requestContext, context.requestURL)
	return url
}

func (context *FoodboxContext) LogoutURL() string {
	url, _ := user.LogoutURL(context.requestContext, context.requestURL)
	return url
}

func (context *FoodboxContext) GetAppengineContext() appengine.Context {
	return context.requestContext
}