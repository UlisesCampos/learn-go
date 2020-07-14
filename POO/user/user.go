package user

import "time"

type Usuario struct {
	Id        int
	Nombre    string
	FechaAlta time.Time
	status    bool
}

func (this *Usuario) AltaUsuario(id int, nombre string, fechaalta time.Time, Status bool) {
	this.Id = id
	this.Nombre = nombre
	this.FechaAlta = fechaalta
	this.status = Status
}
