package req

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNEQ(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/RQAnonymeGR/GR/GR03/GR03A2_20A_PIU_RechEntMob_PC/ServiceCommunicationInterne.asmx/ObtenirEtatsRensEntreprise", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"d":{"EnDateDu":"Renseignements en date du 2015-12-21 14:06:39","SectionInformationsGenerales":{"SousSecIdentification":{"NomEntreprise":"BOMBARDIER INC.","NEQ":"1143920115"}},"TypeResultat":"E","Message":""}}`)
	})

	entreprise, err := client.GetNEQ("1143920115")

	assert.NotNil(t, entreprise)
	assert.Nil(t, err)
	assert.Equal(t, "1143920115", entreprise.SectionInformationsGenerales.SousSecIdentification.NEQ)
	assert.Equal(t, "BOMBARDIER INC.", entreprise.SectionInformationsGenerales.SousSecIdentification.NomEntreprise)
}

func TestGetNEQError(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/RQAnonymeGR/GR/GR03/GR03A2_20A_PIU_RechEntMob_PC/ServiceCommunicationInterne.asmx/ObtenirEtatsRensEntreprise", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"ExceptionType": "", "Message": "An unexpected error occured", "StackTrace": ""}`)
	})

	entreprise, err := client.GetNEQ("1143920115")

	assert.NotNil(t, err)
	assert.Nil(t, entreprise)
	assert.Contains(t, err.Error(), "500 An unexpected error occured")
}
