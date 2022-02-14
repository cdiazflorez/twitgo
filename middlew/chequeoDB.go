package middlew

import (
	"net/http"

	"github.com/cdiazflorez/twitgo/bd"
)

func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == 0 {
			http.Error(rw, "Conexión perdida con la Base de Datos", 500)
			return
		}
		next.ServeHTTP(rw, r)
	}
}
