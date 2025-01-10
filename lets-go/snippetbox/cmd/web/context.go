package main

// we have to define our custom key type to prevent collisions
type contextKey string

const isAuthenticatedContextKey = contextKey("isAuthenticated")
