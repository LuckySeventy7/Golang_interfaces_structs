package main


import ("fmt"
		"./multimedia"
		)

func main(){
	
	cws:=  multimedia.ContenidoWeb{}//slice 
	//mmi:= []multimediaInterface{}
	
	var op int
	
	for op != 5{
	
	
	
	fmt.Println("\n---------------------------------")
	fmt.Println("1. Capturar Imagen \n2. Capturar Audio \n3. Capturar Video \n4. Mostrar \n5. Salir")
	fmt.Printf("- Opcion: ")
	fmt.Scan(&op)
	switch op{
		case 1:{
			img := multimedia.Imagen{}
			fmt.Printf("\nImagen, titulo: ")
			fmt.Scan(&img.Titulo)
			fmt.Printf("Imagen, formato: ")
			fmt.Scan(&img.Formato)
			fmt.Printf("Imagen, canales: ")
			fmt.Scan(&img.Canales)

			//mmi = append(mmi, &img)
			//mms = multimediaSlice{ multimed: []multimediaInterface{
			//	&aud,
			//}, }
			
			cws.CwSlice = append(cws.CwSlice,&img) 

		}
		case 2:{
			aud := multimedia.Audio{}
			fmt.Printf("\nAudio, titulo: ")
			fmt.Scan(&aud.Titulo)
			fmt.Printf("Audio, formato: ")
			fmt.Scan(&aud.Formato)
			fmt.Printf("Audio, duracion: ")
			fmt.Scan(&aud.Duracion)

			cws.CwSlice = append(cws.CwSlice,&aud) 
			
			
		}
		case 3:{
			vid := multimedia.Video{}
			fmt.Printf("\nVideo, titulo: ")
			fmt.Scan(&vid.Titulo)
			fmt.Printf("Video, formato: ")
			fmt.Scan(&vid.Formato)
			fmt.Printf("Video, frames: ")
			fmt.Scan(&vid.Frames)

			cws.CwSlice = append(cws.CwSlice,&vid)
			
		}
		case 4:
			cws.Mostrar()

	}
}

		
}