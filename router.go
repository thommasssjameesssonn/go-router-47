package main
type Router struct { routes map[string]string }
func NewRouter() *Router { return &Router{make(map[string]string)} }