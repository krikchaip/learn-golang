package main

import "net/http"

func securityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// adding new header keys
		// NOTE: MUST BE executed before any call to WriteHeader() or Write()

		// restrict where the resources for your web page can be loaded from.
		// ref: https://developer.mozilla.org/en-US/docs/Web/HTTP/CSP
		// - only load 'stylesheets' from self (current origin), 'fonts.googleapis.com'
		// - only load 'fonts' from 'fonts.gstatic.com'
		// - for everything else, allows only from self (current origin)
		w.Header().Set("Content-Security-Policy",
			"default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com")

		// control what information to include in a 'Referrer' header
		// in this case, include the full URL (with path, querystring, ...)
		// for the same-origin requests and only include host for cross-origin requests
		// ref: https://developer.mozilla.org/en-US/docs/Web/Security/Same-origin_policy
		w.Header().Set("Referrer-Policy", "origin-when-cross-origin")

		// instructs browsers to not MIME-type sniff the content-type of the response
		// ref: https://security.stackexchange.com/questions/7506/using-file-extension-and-mime-type-as-output-by-file-i-b-combination-to-dete/7531#7531
		w.Header().Set("X-Content-Type-Options", "nosniff")

		// help prevent clickjacking attacks in older browsers that don't support CSP headers
		// ref: https://developer.mozilla.org/en-US/docs/Web/Security/Types_of_attacks#click-jacking
		w.Header().Set("X-Frame-Options", "deny")

		// should be disabled when using CSP headers
		// ref: https://owasp.org/www-project-secure-headers/#x-xss-protection
		w.Header().Set("X-XSS-Protection", "0")

		next.ServeHTTP(w, r)
	})
}
