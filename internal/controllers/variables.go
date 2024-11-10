package controllers

const (
	PathFeed            = "/feed"
	PathLogin           = "/login"
	PathSignup          = "/signup"
	PathLogout          = "/logout"
	PathPublishArticles = "/publish"
	PathDefault         = "/"

	PathAdmin       = "/admin"
	PathAddCategory = PathAdmin + "/AddCategory"

	PathUserPage string = "/users" + "/:" + PathParamUserValue

	PathParamUserValue = "username"
)
