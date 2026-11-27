package routes

import (
	"net/http"

)

// MainRouter crée le routeur principal et y ajoute toutes les routes du site.
func MainRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// Fichiers statiques : /static/css/style.css → assets/css/style.css
	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("GET /assets/", http.StripPrefix("/assets/", fs))



	filamentRoutes(mux)

	
	return mux
}
