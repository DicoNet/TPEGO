package model

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

type Result struct {
	Type   string
	Length int
	Value  string
}

func VerificarCaracteres(Caracteres string) (Result, error) {

	Caracteres = strings.TrimSpace(Caracteres)
	TipoCaracteres := string(Caracteres[0]) + string(Caracteres[1])
	LargoCaracteres := VerificarLargo(string(Caracteres[2]))
	LargoCaracteres += string(Caracteres[3])
	ValorCaracteres, CantidadValorCaracteres := ValoresFaltantes(Caracteres)

	errType := VerificarTipo(TipoCaracteres, ValorCaracteres)
	LargoCaracteresInteger, errLength := strconv.Atoi(LargoCaracteres)

	if errType == nil {
		if errLength != nil {
			return NuevaEntidad(TipoCaracteres, LargoCaracteresInteger, ValorCaracteres), errors.New("Deben ser números")
		}

		if LargoCaracteresInteger == CantidadValorCaracteres {
			return NuevaEntidad(TipoCaracteres, LargoCaracteresInteger, ValorCaracteres), nil
		}

		return NuevaEntidad(TipoCaracteres, LargoCaracteresInteger, ValorCaracteres), errors.New("no coinciden las longitudes")
	}
	return NuevaEntidad(TipoCaracteres, LargoCaracteresInteger, ValorCaracteres), errType
}

// Verificamos que type sea TX o NN y si contienen error en los caracteres

func VerificarTipo(TipoCaracteres string, ValorCaracteres string) error {
	if TipoCaracteres == "TX" {
		return ContieneNumero(ValorCaracteres)
	}

	if TipoCaracteres == "NN" {
		return ContieneCaracter(ValorCaracteres)
	}
	return errors.New("no es TX ni NN")
}

//Verificamos que el primer valor no sea 0, sino se elimina

func VerificarLargo(LargoCaracteres string) string {
	if string(LargoCaracteres) == "0" {
		LargoCaracteres = strings.Replace(string(LargoCaracteres), "0", "", 1)
	}
	return LargoCaracteres
}

func ContieneCaracter(Caracteres string) error {
	_, err := strconv.Atoi(Caracteres)
	if err != nil {
		return errors.New("de tipo NN y el valor contiene como minimo un caracter")
	}
	return nil
}

func ContieneNumero(ValorCaracteres string) error {
	for _, c := range ValorCaracteres {
		if unicode.IsDigit(c) {
			return errors.New("de tipo TX y contiene como minimo un número")
		}
	}
	return nil
}
func NuevaEntidad(t string, l int, v string) Result {
	return Result{t, l, v}
}

func ValoresFaltantes(Caracteres string) (string, int) {
	return Caracteres[4:], len(Caracteres[4:])
}
