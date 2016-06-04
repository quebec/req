package req

const (
	DomaineREQ                           = 1
	DomaineFichierCentraleDesEntreprises = 2
	DomaineAutoritesPubliques            = 3

	TypeNom            = 1
	TypeMots           = 2
	TypeMotsApparentes = 3

	EtendueSocietesDePersonnes      = 1
	EtendueEntreprisesIndividuelles = 2
	EtenduePersonnesMorales         = 3
	EtendueTous                     = 4
)

// GetNEQ queries the Registre des entreprises du Québec to get company
// information using the NEQ number
func (c *Client) GetNEQ(id string) (*Company, error) {

	data := &APIRequest{Critere: APICriteriaRequest{Id: id, UtilisateurAccepteConditionsUtilisation: true}}
	endpoint := "/RQAnonymeGR/GR/GR03/GR03A2_20A_PIU_RechEntMob_PC/ServiceCommunicationInterne.asmx/ObtenirEtatsRensEntreprise"

	req, err := c.NewRequest("POST", endpoint, data)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Referer", "https://www.registreentreprises.gouv.qc.ca/RQAnonymeGR/GR/GR03/GR03A2_20A_PIU_RechEntMob_PC/index.html")

	var company CompanyResponse
	err = c.Do(req, &company)

	if err != nil {
		return nil, err
	}

	return &company.D, nil
}

// Search is used to search for a list of companies
func (c *Client) Search(keywords string, opt *SearchOptions) (*SearchResults, error) {
	data := &APIRequest{
		Critere: APICriteriaRequest{
			Domaine:                                 DomaineREQ,
			Type:                                    TypeMots,
			Etendue:                                 EtendueTous,
			Texte:                                   keywords,
			PageCourante:                            0,
			UtilisateurAccepteConditionsUtilisation: true,
		},
	}

	endpoint := "/RQAnonymeGR/GR/GR03/GR03A2_20A_PIU_RechEntMob_PC/ServiceCommunicationInterne.asmx/ObtenirListeEntreprises"

	req, err := c.NewRequest("POST", endpoint, data)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Referer", "https://www.registreentreprises.gouv.qc.ca/RQAnonymeGR/GR/GR03/GR03A2_20A_PIU_RechEntMob_PC/index.html")

	var results SearchResultsResponse
	err = c.Do(req, &results)

	return &results.D, err
}

type SearchOptions struct {
	Domaine int
	Etendue int
	Page    int
	Type    int
}

// An APIRequest is used as a wrapper to request the API
type APIRequest struct {
	Critere APICriteriaRequest `json:"critere"`
}

// An APICriteriaRequest contains criteria for a request to the API
type APICriteriaRequest struct {
	Id                                      string `json:"Id,omitempty"`
	Domaine                                 int    `json:"Domaine,omitempty"`
	Etendue                                 int    `json:"Etendue,omitempty"`
	PageCourante                            int    `json:"PageCourante,omitempty"`
	Texte                                   string `json:"Texte,omitempty"`
	Type                                    int    `json:"Type,omitempty"`
	UtilisateurAccepteConditionsUtilisation bool   `json:"UtilisateurAccepteConditionsUtilisation,omitempty"`
}

type SearchResultsResponse struct {
	D SearchResults `json:"d"`
}

type SearchResults struct {
	CleSession       string `json:"CleSession"`
	ListeEntreprises []struct {
		AdressePrimaire    string `json:"AdressePrimaire"`
		DateChangementEtat string `json:"DateChangementEtat"`
		DateFinale         string `json:"DateFinale"`
		DateInitiale       string `json:"DateInitiale"`
		ID                 string `json:"ID"`
		Nom                string `json:"Nom"`
		NumeroDossier      string `json:"NumeroDossier"`
		Statut             string `json:"Statut"`
		StatutDuNom        string `json:"StatutDuNom"`
	} `json:"ListeEntreprises"`
	Message              string `json:"Message"`
	NombrePages          int    `json:"NombrePages"`
	PageCourante         int    `json:"PageCourante"`
	TotalEnregistrements int    `json:"TotalEnregistrements"`
	TypeResultat         string `json:"TypeResultat"`
}

// A CompanyResponse is used to wrap responses from the API
type CompanyResponse struct {
	D Company `json:"d"`
}

// A Company represents information about a company from
// the registry
type Company struct {
	EnDateDu              string `json:"EnDateDu"`
	SectionAdministration struct {
		SousSecActionnaires struct {
			ListeActionnaires []struct {
				Adresse            string      `json:"Adresse"`
				DateDebut          interface{} `json:"DateDebut"`
				DateFin            interface{} `json:"DateFin"`
				Fonction           interface{} `json:"Fonction"`
				MessageActionnaire interface{} `json:"MessageActionnaire"`
				NomPersMoral       string      `json:"NomPersMoral"`
				NomPersPhysique    interface{} `json:"NomPersPhysique"`
				PrenomPersPhysique interface{} `json:"PrenomPersPhysique"`
				Statut             interface{} `json:"Statut"`
			} `json:"ListeActionnaires"`
			TexteConventionUnanime interface{} `json:"TexteConventionUnanime"`
		} `json:"SousSecActionnaires"`
		SousSecAdminBienAutrui struct {
			ListeAdminBienAutrui      []interface{} `json:"ListeAdminBienAutrui"`
			ListeAdminBienAutruiHisto interface{}   `json:"ListeAdminBienAutruiHisto"`
		} `json:"SousSecAdminBienAutrui"`
		SousSecAdministrateurs struct {
			ListeAdministrateurs []struct {
				Adresse              string `json:"Adresse"`
				DateDebut            string `json:"DateDebut"`
				DateFin              string `json:"DateFin"`
				Fonction             string `json:"Fonction"`
				NomAdministrateur    string `json:"NomAdministrateur"`
				PrenomAdministrateur string `json:"PrenomAdministrateur"`
			} `json:"ListeAdministrateurs"`
			ListeAdministrateursHisto []struct {
				Adresse              string `json:"Adresse"`
				DateDebut            string `json:"DateDebut"`
				DateFin              string `json:"DateFin"`
				Fonction             string `json:"Fonction"`
				NomAdministrateur    string `json:"NomAdministrateur"`
				PrenomAdministrateur string `json:"PrenomAdministrateur"`
			} `json:"ListeAdministrateursHisto"`
		} `json:"SousSecAdministrateurs"`
		SousSecAssocies   interface{} `json:"SousSecAssocies"`
		SousSecConvention struct {
			ListeActionnaireTiers      interface{} `json:"ListeActionnaireTiers"`
			ListeActionnaireTiersHisto interface{} `json:"ListeActionnaireTiersHisto"`
			TexteConventionRetire      interface{} `json:"TexteConventionRetire"`
			TexteExistConvention       interface{} `json:"TexteExistConvention"`
		} `json:"SousSecConvention"`
		SousSecDirigeants struct {
			ListeDirigeants []struct {
				Adresse            string      `json:"Adresse"`
				DateDebut          interface{} `json:"DateDebut"`
				DateFin            interface{} `json:"DateFin"`
				Fonction           string      `json:"Fonction"`
				NomPersMoral       interface{} `json:"NomPersMoral"`
				NomPersPhysique    string      `json:"NomPersPhysique"`
				PrenomPersPhysique string      `json:"PrenomPersPhysique"`
				Statut             interface{} `json:"Statut"`
				TexteAucun         interface{} `json:"TexteAucun"`
			} `json:"ListeDirigeants"`
		} `json:"SousSecDirigeants"`
		SousSecFCE           interface{} `json:"SousSecFCE"`
		SousSecFondesPouvoir struct {
			ListeFondesPouvoir []interface{} `json:"ListeFondesPouvoir"`
		} `json:"SousSecFondesPouvoir"`
	} `json:"SectionAdministration"`
	SectionDocument struct {
		SousSecDocumentsConserves struct {
			GrilleDocuments struct {
				LibelleGrille string `json:"LibelleGrille"`
				Lignes        []struct {
					ListeCellules []struct {
						Valeur string `json:"Valeur"`
					} `json:"ListeCellules"`
				} `json:"Lignes"`
			} `json:"GrilleDocuments"`
		} `json:"SousSecDocumentsConserves"`
		SousSecDocumentsEnTraitement struct {
			GrilleDocuments interface{} `json:"GrilleDocuments"`
		} `json:"SousSecDocumentsEnTraitement"`
	} `json:"SectionDocument"`
	SectionEtablissement struct {
		SousSecEtablissementAutre struct {
			ListeEtablissement []struct {
				Adresse                string      `json:"Adresse"`
				DateDebutEtablissement interface{} `json:"DateDebutEtablissement"`
				DateDebutNom           interface{} `json:"DateDebutNom"`
				DateFinEtablissement   interface{} `json:"DateFinEtablissement"`
				DateFinNom             interface{} `json:"DateFinNom"`
				EtablissementPrincipal interface{} `json:"EtablissementPrincipal"`
				ListeActivites         []struct {
					CAE         string      `json:"CAE"`
					Description string      `json:"Description"`
					Precision   interface{} `json:"Precision"`
					Titre       string      `json:"Titre"`
				} `json:"ListeActivites"`
				NoReference  string `json:"NoReference"`
				Nom          string `json:"Nom"`
				TypeActivite string `json:"TypeActivite"`
			} `json:"ListeEtablissement"`
		} `json:"SousSecEtablissementAutre"`
		SousSecEtablissementPrincipal struct {
			ListeEtablissement []struct {
				Adresse                string      `json:"Adresse"`
				DateDebutEtablissement interface{} `json:"DateDebutEtablissement"`
				DateDebutNom           interface{} `json:"DateDebutNom"`
				DateFinEtablissement   interface{} `json:"DateFinEtablissement"`
				DateFinNom             interface{} `json:"DateFinNom"`
				EtablissementPrincipal interface{} `json:"EtablissementPrincipal"`
				ListeActivites         []struct {
					CAE         string      `json:"CAE"`
					Description string      `json:"Description"`
					Precision   interface{} `json:"Precision"`
					Titre       string      `json:"Titre"`
				} `json:"ListeActivites"`
				NoReference  string `json:"NoReference"`
				Nom          string `json:"Nom"`
				TypeActivite string `json:"TypeActivite"`
			} `json:"ListeEtablissement"`
		} `json:"SousSecEtablissementPrincipal"`
	} `json:"SectionEtablissement"`
	SectionInformationsGenerales struct {
		SousSecActiviteFCE        interface{} `json:"SousSecActiviteFCE"`
		SousSecActiviteNbrEmploye struct {
			ListeSecteursActivite []struct {
				CAE         string `json:"CAE"`
				Description string `json:"Description"`
				Precision   string `json:"Precision"`
			} `json:"ListeSecteursActivite"`
			NombreEmployes string `json:"NombreEmployes"`
		} `json:"SousSecActiviteNbrEmploye"`
		SousSecAdresseCor struct {
			CodePostal             interface{} `json:"CodePostal"`
			LigneAdresse           string      `json:"LigneAdresse"`
			Localite               interface{} `json:"Localite"`
			NomEntreprise          interface{} `json:"NomEntreprise"`
			NomPersonnePhysique    string      `json:"NomPersonnePhysique"`
			PrenomPersonnePhysique string      `json:"PrenomPersonnePhysique"`
			Statut                 interface{} `json:"Statut"`
		} `json:"SousSecAdresseCor"`
		SousSecAdresseEntrpr struct {
			CodePostal   interface{} `json:"CodePostal"`
			LigneAdresse string      `json:"LigneAdresse"`
			Localite     interface{} `json:"Localite"`
			Statut       interface{} `json:"Statut"`
		} `json:"SousSecAdresseEntrpr"`
		SousSecAutresInformations interface{} `json:"SousSecAutresInformations"`
		SousSecContTran           struct {
			ListeBlocsContinuation []struct {
				DateContin  string      `json:"DateContin"`
				Lieu        interface{} `json:"Lieu"`
				Loi         interface{} `json:"Loi"`
				MessageCont string      `json:"MessageCont"`
			} `json:"ListeBlocsContinuation"`
		} `json:"SousSecContTran"`
		SousSecContinuationFCE interface{} `json:"SousSecContinuationFCE"`
		SousSecDateMiseAJour   struct {
			DateDernierDeclaration string `json:"DateDernierDeclaration"`
			DateFinPerAnneeCour    string `json:"DateFinPerAnneeCour"`
			DateFinPerAnneePrec    string `json:"DateFinPerAnneePrec"`
			DateMajEtat            string `json:"DateMajEtat"`
		} `json:"SousSecDateMiseAJour"`
		SousSecFaillite struct {
			TexteFaillite string `json:"TexteFaillite"`
		} `json:"SousSecFaillite"`
		SousSecFormeJuridique struct {
			AssujetiVolontaire  interface{} `json:"AssujetiVolontaire"`
			DateDebutRespLimite interface{} `json:"DateDebutRespLimite"`
			DateEtat            interface{} `json:"DateEtat"`
			DateFinRespLimite   interface{} `json:"DateFinRespLimite"`
			DateFormation       string      `json:"DateFormation"`
			EtatJuridique       interface{} `json:"EtatJuridique"`
			LieuConstitution    interface{} `json:"LieuConstitution"`
			PrecisFormeJuri     interface{} `json:"PrecisFormeJuri"`
			RegimeConstitutif   string      `json:"RegimeConstitutif"`
			RegimeCourant       string      `json:"RegimeCourant"`
			TexteRespLimite     interface{} `json:"TexteRespLimite"`
			Type                string      `json:"Type"`
		} `json:"SousSecFormeJuridique"`
		SousSecFusion struct {
			GrilleFusion  interface{} `json:"GrilleFusion"`
			GrilleFusions []struct {
				LibelleGrille string `json:"LibelleGrille"`
				Lignes        []struct {
					ListeCellules []struct {
						Valeur string `json:"Valeur"`
					} `json:"ListeCellules"`
				} `json:"Lignes"`
			} `json:"GrilleFusions"`
			Resultante interface{} `json:"Resultante"`
		} `json:"SousSecFusion"`
		SousSecIdentification struct {
			DateEntreeEnVigueur  interface{} `json:"DateEntreeEnVigueur"`
			DateFinUtilisation   interface{} `json:"DateFinUtilisation"`
			NEQ                  string      `json:"NEQ"`
			NomEntreprise        string      `json:"NomEntreprise"`
			NomFamille           interface{} `json:"NomFamille"`
			Prenom               interface{} `json:"Prenom"`
			Situation            interface{} `json:"Situation"`
			Statut               interface{} `json:"Statut"`
			TableauAutresLangues interface{} `json:"TableauAutresLangues"`
		} `json:"SousSecIdentification"`
		SousSecImmatriculation struct {
			DateCessationPrevue   string `json:"DateCessationPrevue"`
			DateImmatriculation   string `json:"DateImmatriculation"`
			DateStatut            string `json:"DateStatut"`
			StatutImmatriculation string `json:"StatutImmatriculation"`
		} `json:"SousSecImmatriculation"`
		SousSecLiquidationDissolution struct {
			Cessation interface{} `json:"Cessation"`
			Texte     string      `json:"Texte"`
		} `json:"SousSecLiquidationDissolution"`
		SousSecObjectifs         interface{} `json:"SousSecObjectifs"`
		SousSecRecours           interface{} `json:"SousSecRecours"`
		SousSectionRapportAnnuel interface{} `json:"SousSectionRapportAnnuel"`
	} `json:"SectionInformationsGenerales"`
	SectionNomsAutreNoms struct {
		Consigne          string      `json:"Consigne"`
		MessageAnalyse    interface{} `json:"MessageAnalyse"`
		SousSecAutresNoms struct {
			GrilleNoms struct {
				LibelleGrille string `json:"LibelleGrille"`
				Lignes        []struct {
					ListeCellules []struct {
						Valeur string `json:"Valeur"`
					} `json:"ListeCellules"`
				} `json:"Lignes"`
			} `json:"GrilleNoms"`
		} `json:"SousSecAutresNoms"`
		SousSecNoms struct {
			GrilleNoms struct {
				LibelleGrille string `json:"LibelleGrille"`
				Lignes        []struct {
					ListeCellules []struct {
						Valeur string `json:"Valeur"`
					} `json:"ListeCellules"`
				} `json:"Lignes"`
			} `json:"GrilleNoms"`
		} `json:"SousSecNoms"`
	} `json:"SectionNomsAutreNoms"`
	TypeResultat string `json:"TypeResultat"`
}
