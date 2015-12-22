package main

import (
	"fmt"

	"github.com/quebec/req"
)

func main() {
	c := req.NewClient(nil)

	company, _ := c.GetNEQ("1143920115")
	fmt.Println(company.SectionInformationsGenerales.SousSecIdentification.NomEntreprise)
}
