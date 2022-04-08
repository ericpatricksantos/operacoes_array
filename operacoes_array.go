package operacoes_array

func UneDuasListasSemElementosRepetidos(lista1, lista2 []string) (listaResultante []string, mapa map[string]string, tamanho int) {

	// Preenchendo o Mapa com a lista1
	mapa = PreenchendoMapa(lista1)
	// Preenchendo o Mapa com a lista 2
	PreenchendoMapaComLista(lista2, mapa)

	for _, valor := range mapa {
		listaResultante = append(listaResultante, valor)
	}
	return listaResultante, mapa, len(mapa)
}

func PreenchendoMapa(lista []string) map[string]string {
	mapa := map[string]string{}
	for _, str := range lista {
		_, existe := mapa[str]
		if existe {
			continue
		} else {
			mapa[str] = str
		}
	}

	return mapa
}

func UneVariasMapas(ListaMapa ...map[string]string) map[string]string {
	mapaResultante := map[string]string{}
	for _, mapa := range ListaMapa {
		for _, valor := range mapa {
			_, existe := mapaResultante[valor]
			if existe {
				continue
			} else {
				mapaResultante[valor] = valor
			}
		}
	}

	return mapaResultante
}

func PreenchendoMapaComLista(lista []string, mapa map[string]string) {
	for _, valor := range lista {
		_, ok := mapa[valor]
		if ok {
			continue
		} else {
			mapa[valor] = valor
		}
	}
}
