package main

import (
	"fmt"

	"github.com/quebec/req"
)

func main() {
	c := req.NewClient(nil)

	result, err := c.Search("mrc", nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, company := range result.ListeEntreprises {
		fmt.Println(" (" + company.NumeroDossier + ") " + company.Nom)
	}
}
