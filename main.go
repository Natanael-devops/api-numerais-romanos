package main

import (
	"github.com/Natanael-devops/api-numerais-romanos/database"
	"github.com/Natanael-devops/api-numerais-romanos/routes"
)

func main() {
	database.ConectaComBancoDeDados()

	routes.CarregaRotas()

}
