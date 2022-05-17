package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Natanael-devops/api-numerais-romanos/database"
	"github.com/Natanael-devops/api-numerais-romanos/models"
	"github.com/gin-gonic/gin"
)

//EXPLICAÇÃO
/*
A função CriaPalavra recebe o body request e chama a
função VerificaPalavra para selecionar os algarismos romanos,
depois é chamada a função FazSlice que transforma os algarismos romanos em um slice
de algarismos arábicos.
Depois a função CalculaMaior vai calcular qual é o maior algarismo da slice.
Por fim, a função ToRoman apenas reescreve o mesmo valor final para numerais romanos
e assim a função CriaPalavra termina retornando o body request reescrito.*/

var numero = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

var numInv = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}
var maxTable = []int{
	1000,
	900,
	500,
	400,
	100,
	90,
	50,
	40,
	10,
	9,
	5,
	4,
	1,
}

func ApresentaNumeros(c *gin.Context) {
	var numeros []models.Numero
	database.DB.Find(&numeros)
	c.JSON(200, numeros)
}

type Romano struct{}

func NovoRomano() *Romano {
	return &Romano{}
}

func CriaPalavra(c *gin.Context) {
	//função para criar a palavra no banco de dados
	var palavranova models.Numero
	if err := c.BindJSON(&palavranova); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	fmt.Println(palavranova.Text)
	pfiltrada := VerificaPalavra(palavranova.Text)
	q := FazSlice(pfiltrada)

	r := CalculaMaior(q)
	s := NovoRomano().ToRoman(r)

	palavranova.Number = s
	palavranova.Value = r
	database.DB.Create(&palavranova)
	c.JSON(200, palavranova)
}

func VerificaPalavra(p string) []string {
	var novoSlice = []string{}
	palavra := strings.ToUpper(p)
	numerais := "IVXLCDM"
	for i := 0; i < len(palavra); i++ {

		if strings.Contains(numerais, string(palavra[i])) {
			if i > 0 && !strings.Contains(numerais, string(palavra[i-1])) {
				novoSlice = append(novoSlice, string(palavra[i]))
			}

			//for 2
			for n := i + 1; n < len(palavra); n++ {
				if strings.Contains(numerais, string(palavra[n])) && numero[string(palavra[i])]+numero[string(palavra[n])] > numero[string(palavra[i])] {
					if i > 0 && !strings.Contains(numerais, string(palavra[i-1])) {
						novoSlice = append(novoSlice, string(palavra[i])+string(palavra[n]))
					}

					//for 3
					for o := n + 1; o < len(palavra); o++ {
						if strings.Contains(numerais, string(palavra[o])) && numero[string(palavra[i])]+numero[string(palavra[n])]+numero[string(palavra[o])] > numero[string(palavra[i])]+numero[string(palavra[n])] {

							if i > 0 && !strings.Contains(numerais, string(palavra[i-1])) {
								novoSlice = append(novoSlice, string(palavra[i])+string(palavra[n])+string(palavra[o]))
							}

							for p := o + 1; p < len(palavra); p++ {
								if strings.Contains(numerais, string(palavra[p])) && numero[string(palavra[i])]+numero[string(palavra[n])]+numero[string(palavra[o])]+numero[string(palavra[p])] > numero[string(palavra[i])]+numero[string(palavra[n])]+numero[string(palavra[o])] {

									if i > 0 && !strings.Contains(numerais, string(palavra[i-1])) {
										novoSlice = append(novoSlice, string(palavra[i])+string(palavra[n])+string(palavra[o])+string(palavra[p]))
									}
								} else {
									break
								}
							}

						} else {
							break
						}
						if strings.Contains(numerais, string(palavra[o])) && numero[string(palavra[i])]+numero[string(palavra[n])]+numero[string(palavra[o])] < numero[string(palavra[i])]+numero[string(palavra[n])] {
							novoSlice = append(novoSlice, string(palavra[i])+string(palavra[n]))
						}
					}
				} else {

					break
				}
				if strings.Contains(numerais, string(palavra[n])) && numero[string(palavra[i])]+numero[string(palavra[n])] < numero[string(palavra[i])] {
					novoSlice = append(novoSlice, string(palavra[i]))
				}
			}
		}
	}
	fmt.Println(novoSlice)
	return novoSlice
}

func FazSlice(n []string) []int {
	var sliceNumeral = []int{}
	for i := 0; i < len(n); i++ {
		r := NovoRomano()
		v := r.Arabe(n[i])
		sliceNumeral = append(sliceNumeral, v)
	}
	if len(n) > 0 {
		fmt.Println(sliceNumeral)
	}
	return sliceNumeral
}

func CalculaMaior(c []int) int {
	if len(c) == 0 {
		c = append(c, 0)
	}
	resultado := c[0]
	for _, valor := range c {

		if valor > resultado {
			resultado = valor
		}
	}
	return resultado
}

func (r *Romano) Arabe(n string) int {
	saida := 0
	tamanho := len(n)
	for i := 0; i < tamanho; i++ {
		c := string(n[i])
		cc := numero[c]
		if i < tamanho-1 {
			cproximo := string(n[i+1])
			ccproximo := numero[cproximo]
			if cc < ccproximo {
				saida += ccproximo - cc
				i++
			} else {
				saida += cc
			}

		} else {
			saida += cc
		}
	}
	return saida
}

func (r *Romano) ToRoman(n int) string {
	out := ""
	for n > 0 {
		v := highestDecimal(n)
		out += numInv[v]
		n -= v
	}
	return out
}

func highestDecimal(n int) int {
	for _, v := range maxTable {
		if v <= n {
			return v
		}
	}
	return 1
}
