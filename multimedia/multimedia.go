package multimedia

import "fmt"

type MultimediaInterface interface{
	Mostrar()// podra usar el mostrar de cualquier de las estructuras multimedia
}
type ContenidoWeb struct {
	CwSlice []MultimediaInterface
}

func (cws *ContenidoWeb) Mostrar(){

	for _,i:= range cws.CwSlice{
		fmt.Println(i)

	}
}
type Imagen struct {
	 Titulo string
	 Formato string
	 Canales string
}
type Audio struct {
	Titulo string
	Formato string
	Duracion string
}
type Video struct {
	Titulo string
	Formato string
	Frames string
}

func (im Imagen)Mostrar() {
	fmt.Println("Imagen, titulo: ", im.Titulo)
	fmt.Println(im.Formato)
	fmt.Println(im.Canales)
}
func (au Audio)Mostrar() {
	fmt.Println(au.Titulo)
	fmt.Println(au.Formato)
	fmt.Println(au.Duracion)
}
func (vi Video) Mostrar() {
	fmt.Println(vi.Titulo)
	fmt.Println(vi.Formato)
	fmt.Println(vi.Frames)
}
