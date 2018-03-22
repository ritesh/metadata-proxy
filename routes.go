package main

import "net/http"

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []route

var routes = Routes{
	route{
		"latest",
		"GET",
		"/latest/iam/security-credentials/{role-name}",
		iamSecurityCredentials,
	},
}
