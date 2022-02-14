package routers

import (
	"encoding/json"
	"net/http"

	"github.com/cdiazflorez/twitgo/bd"
	"github.com/cdiazflorez/twitgo/models"
)

func Registro(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El Email de usuario es requerido "+err.Error(), 400)
	}

	if len(t.Password) < 6 {
		http.Error(w, "El password tiene menos de 6 caracteres"+err.Error(), 400)
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)

	if encontrado == true {

		http.Error(w, "Ya existe un usuario registrado con ese email "+err.Error(), 400)

	}

	_, status, err := bd.InsertoRegistro(t)

	if err != nil {
		http.Error(w, "Ocurrio un error al crear el usuario"+err.Error(), 400)
	}

	if status == false {
		http.Error(w, "Ocurrio un error al intentar guardar el registro del usuario", 400)
	}

	w.WriteHeader(http.StatusCreated)

}
