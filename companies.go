package req

// GetNEQ queries the Registre des entreprises du Québec to get company
// information using the NEQ number
func (c *Client) GetNEQ(id string) (*Company, error) {

	data := &APIRequest{Critere: APICriteriaRequest{Id: id}}
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

// An APIRequest is used as a wrapper to request the API
type APIRequest struct {
	Critere APICriteriaRequest `json:"Critere"`
}

// An APICriteriaRequest contains criteria for a request to the API
type APICriteriaRequest struct {
	Id string `json:"Id,omitempty"`
}

// A CompanyResponse is used to wrap responses from the API
type CompanyResponse struct {
	D Company `json:"d"`
}

// A Company represents information about a company from
// the registry
type Company struct {
	Consigne              string `json:"Consigne"`
	EnDateDu              string `json:"EnDateDu"`
	Message               string `json:"Message"`
	SectionAdministration struct {
		IDSection           int         `json:"IdSection"`
		MessageAucuneValeur interface{} `json:"MessageAucuneValeur"`
		SousSecActionnaires struct {
			IDSousSection     int `json:"IdSousSection"`
			ListeActionnaires []struct {
				Adresse            string      `json:"Adresse"`
				DateDebut          interface{} `json:"DateDebut"`
				DateFin            interface{} `json:"DateFin"`
				Fonction           interface{} `json:"Fonction"`
				IDPersonne         int         `json:"IdPersonne"`
				MessageActionnaire interface{} `json:"MessageActionnaire"`
				NomPersMoral       string      `json:"NomPersMoral"`
				NomPersPhysique    interface{} `json:"NomPersPhysique"`
				PrenomPersPhysique interface{} `json:"PrenomPersPhysique"`
				Statut             interface{} `json:"Statut"`
			} `json:"ListeActionnaires"`
			MessageAucuneValeur    interface{} `json:"MessageAucuneValeur"`
			TexteConventionUnanime interface{} `json:"TexteConventionUnanime"`
		} `json:"SousSecActionnaires"`
		SousSecAdminBienAutrui struct {
			IDSousSection             int           `json:"IdSousSection"`
			ListeAdminBienAutrui      []interface{} `json:"ListeAdminBienAutrui"`
			ListeAdminBienAutruiHisto interface{}   `json:"ListeAdminBienAutruiHisto"`
			MessageAucuneValeur       string        `json:"MessageAucuneValeur"`
		} `json:"SousSecAdminBienAutrui"`
		SousSecAdministrateurs struct {
			IDSousSection        int `json:"IdSousSection"`
			ListeAdministrateurs []struct {
				Adresse              string `json:"Adresse"`
				DateDebut            string `json:"DateDebut"`
				DateFin              string `json:"DateFin"`
				Fonction             string `json:"Fonction"`
				NomAdministrateur    string `json:"NomAdministrateur"`
				PrenomAdministrateur string `json:"PrenomAdministrateur"`
				IDAdministrateur     int    `json:"idAdministrateur"`
			} `json:"ListeAdministrateurs"`
			ListeAdministrateursHisto []struct {
				Adresse              string `json:"Adresse"`
				DateDebut            string `json:"DateDebut"`
				DateFin              string `json:"DateFin"`
				Fonction             string `json:"Fonction"`
				NomAdministrateur    string `json:"NomAdministrateur"`
				PrenomAdministrateur string `json:"PrenomAdministrateur"`
				IDAdministrateur     int    `json:"idAdministrateur"`
			} `json:"ListeAdministrateursHisto"`
			MessageAucuneValeur interface{} `json:"MessageAucuneValeur"`
		} `json:"SousSecAdministrateurs"`
		SousSecAssocies   interface{} `json:"SousSecAssocies"`
		SousSecConvention struct {
			IDSousSection              int         `json:"IdSousSection"`
			ListeActionnaireTiers      interface{} `json:"ListeActionnaireTiers"`
			ListeActionnaireTiersHisto interface{} `json:"ListeActionnaireTiersHisto"`
			MessageAucuneValeur        string      `json:"MessageAucuneValeur"`
			TexteConventionRetire      interface{} `json:"TexteConventionRetire"`
			TexteExistConvention       interface{} `json:"TexteExistConvention"`
		} `json:"SousSecConvention"`
		SousSecDirigeants struct {
			IDSousSection   int `json:"IdSousSection"`
			ListeDirigeants []struct {
				Adresse            string      `json:"Adresse"`
				DateDebut          interface{} `json:"DateDebut"`
				DateFin            interface{} `json:"DateFin"`
				Fonction           string      `json:"Fonction"`
				IDPersonne         int         `json:"IdPersonne"`
				NomPersMoral       interface{} `json:"NomPersMoral"`
				NomPersPhysique    string      `json:"NomPersPhysique"`
				PrenomPersPhysique string      `json:"PrenomPersPhysique"`
				Statut             interface{} `json:"Statut"`
				TexteAucun         interface{} `json:"TexteAucun"`
			} `json:"ListeDirigeants"`
			MessageAucuneValeur interface{} `json:"MessageAucuneValeur"`
		} `json:"SousSecDirigeants"`
		SousSecFCE           interface{} `json:"SousSecFCE"`
		SousSecFondesPouvoir struct {
			IDSousSection       int           `json:"IdSousSection"`
			ListeFondesPouvoir  []interface{} `json:"ListeFondesPouvoir"`
			MessageAucuneValeur string        `json:"MessageAucuneValeur"`
		} `json:"SousSecFondesPouvoir"`
	} `json:"SectionAdministration"`
	SectionDocument struct {
		IDSection                 int         `json:"IdSection"`
		MessageAucuneValeur       interface{} `json:"MessageAucuneValeur"`
		SousSecDocumentsConserves struct {
			GrilleDocuments struct {
				EntetesColonnes []struct {
					ID      int    `json:"ID"`
					Libelle string `json:"Libelle"`
				} `json:"EntetesColonnes"`
				LibelleGrille string `json:"LibelleGrille"`
				Lignes        []struct {
					ID            int `json:"ID"`
					ListeCellules []struct {
						ID     int    `json:"ID"`
						Valeur string `json:"Valeur"`
					} `json:"ListeCellules"`
				} `json:"Lignes"`
			} `json:"GrilleDocuments"`
			IDSousSection       int         `json:"IdSousSection"`
			MessageAucuneValeur interface{} `json:"MessageAucuneValeur"`
		} `json:"SousSecDocumentsConserves"`
		SousSecDocumentsEnTraitement struct {
			GrilleDocuments     interface{} `json:"GrilleDocuments"`
			IDSousSection       int         `json:"IdSousSection"`
			MessageAucuneValeur string      `json:"MessageAucuneValeur"`
		} `json:"SousSecDocumentsEnTraitement"`
	} `json:"SectionDocument"`
	SectionEtablissement struct {
		IDSection                 int         `json:"IdSection"`
		MessageAucuneValeur       interface{} `json:"MessageAucuneValeur"`
		SousSecEtablissementAutre struct {
			IDSousSection      int `json:"IdSousSection"`
			ListeEtablissement []struct {
				Adresse                string      `json:"Adresse"`
				DateDebutEtablissement interface{} `json:"DateDebutEtablissement"`
				DateDebutNom           interface{} `json:"DateDebutNom"`
				DateFinEtablissement   interface{} `json:"DateFinEtablissement"`
				DateFinNom             interface{} `json:"DateFinNom"`
				EtablissementPrincipal interface{} `json:"EtablissementPrincipal"`
				IDEtablissement        int         `json:"IdEtablissement"`
				ListeActivites         []struct {
					CAE         string      `json:"CAE"`
					Description string      `json:"Description"`
					IDActivite  int         `json:"IdActivite"`
					Precision   interface{} `json:"Precision"`
					Titre       string      `json:"Titre"`
				} `json:"ListeActivites"`
				NoReference  string `json:"NoReference"`
				Nom          string `json:"Nom"`
				TypeActivite string `json:"TypeActivite"`
			} `json:"ListeEtablissement"`
			MessageAucuneValeur interface{} `json:"MessageAucuneValeur"`
		} `json:"SousSecEtablissementAutre"`
		SousSecEtablissementPrincipal struct {
			IDSousSection      int `json:"IdSousSection"`
			ListeEtablissement []struct {
				Adresse                string      `json:"Adresse"`
				DateDebutEtablissement interface{} `json:"DateDebutEtablissement"`
				DateDebutNom           interface{} `json:"DateDebutNom"`
				DateFinEtablissement   interface{} `json:"DateFinEtablissement"`
				DateFinNom             interface{} `json:"DateFinNom"`
				EtablissementPrincipal interface{} `json:"EtablissementPrincipal"`
				IDEtablissement        int         `json:"IdEtablissement"`
				ListeActivites         []struct {
					CAE         string      `json:"CAE"`
					Description string      `json:"Description"`
					IDActivite  int         `json:"IdActivite"`
					Precision   interface{} `json:"Precision"`
					Titre       string      `json:"Titre"`
				} `json:"ListeActivites"`
				NoReference  string `json:"NoReference"`
				Nom          string `json:"Nom"`
				TypeActivite string `json:"TypeActivite"`
			} `json:"ListeEtablissement"`
			MessageAucuneValeur interface{} `json:"MessageAucuneValeur"`
		} `json:"SousSecEtablissementPrincipal"`
	} `json:"SectionEtablissement"`
	SectionInformationsGenerales struct {
		IDSection                 int         `json:"IdSection"`
		MessageAucuneValeur       interface{} `json:"MessageAucuneValeur"`
		SousSecActiviteFCE        interface{} `json:"SousSecActiviteFCE"`
		SousSecActiviteNbrEmploye struct {
			IDSousSection         int `json:"IdSousSection"`
			ListeSecteursActivite []struct {
				CAE                 string      `json:"CAE"`
				Description         string      `json:"Description"`
				IDSecteur           int         `json:"IdSecteur"`
				MessageAucuneValeur interface{} `json:"MessageAucuneValeur"`
				Precision           string      `json:"Precision"`
			} `json:"ListeSecteursActivite"`
			MessageAucuneValeur interface{} `json:"MessageAucuneValeur"`
			NombreEmployes      string      `json:"NombreEmployes"`
		} `json:"SousSecActiviteNbrEmploye"`
		SousSecAdresseCor struct {
			CodePostal             interface{} `json:"CodePostal"`
			IDSousSection          int         `json:"IdSousSection"`
			LigneAdresse           string      `json:"LigneAdresse"`
			Localite               interface{} `json:"Localite"`
			MessageAucuneValeur    interface{} `json:"MessageAucuneValeur"`
			NomEntreprise          interface{} `json:"NomEntreprise"`
			NomPersonnePhysique    string      `json:"NomPersonnePhysique"`
			PrenomPersonnePhysique string      `json:"PrenomPersonnePhysique"`
			Statut                 interface{} `json:"Statut"`
		} `json:"SousSecAdresseCor"`
		SousSecAdresseEntrpr struct {
			CodePostal          interface{} `json:"CodePostal"`
			IDSousSection       int         `json:"IdSousSection"`
			LigneAdresse        string      `json:"LigneAdresse"`
			Localite            interface{} `json:"Localite"`
			MessageAucuneValeur interface{} `json:"MessageAucuneValeur"`
			Statut              interface{} `json:"Statut"`
		} `json:"SousSecAdresseEntrpr"`
		SousSecAutresInformations interface{} `json:"SousSecAutresInformations"`
		SousSecContTran           struct {
			IDSousSection          int `json:"IdSousSection"`
			ListeBlocsContinuation []struct {
				DateContin  string      `json:"DateContin"`
				Lieu        interface{} `json:"Lieu"`
				Loi         interface{} `json:"Loi"`
				MessageCont string      `json:"MessageCont"`
			} `json:"ListeBlocsContinuation"`
			MessageAucuneValeur interface{} `json:"MessageAucuneValeur"`
		} `json:"SousSecContTran"`
		SousSecContinuationFCE interface{} `json:"SousSecContinuationFCE"`
		SousSecDateMiseAJour   struct {
			DateDernierDeclaration string      `json:"DateDernierDeclaration"`
			DateFinPerAnneeCour    string      `json:"DateFinPerAnneeCour"`
			DateFinPerAnneePrec    string      `json:"DateFinPerAnneePrec"`
			DateMajEtat            string      `json:"DateMajEtat"`
			IDSousSection          int         `json:"IdSousSection"`
			MessageAucuneValeur    interface{} `json:"MessageAucuneValeur"`
		} `json:"SousSecDateMiseAJour"`
		SousSecFaillite struct {
			IDSousSection       int         `json:"IdSousSection"`
			MessageAucuneValeur interface{} `json:"MessageAucuneValeur"`
			TexteFaillite       string      `json:"TexteFaillite"`
		} `json:"SousSecFaillite"`
		SousSecFormeJuridique struct {
			AssujetiVolontaire  interface{} `json:"AssujetiVolontaire"`
			DateDebutRespLimite interface{} `json:"DateDebutRespLimite"`
			DateEtat            interface{} `json:"DateEtat"`
			DateFinRespLimite   interface{} `json:"DateFinRespLimite"`
			DateFormation       string      `json:"DateFormation"`
			EtatJuridique       interface{} `json:"EtatJuridique"`
			IDSousSection       int         `json:"IdSousSection"`
			LieuConstitution    interface{} `json:"LieuConstitution"`
			MessageAucuneValeur interface{} `json:"MessageAucuneValeur"`
			PrecisFormeJuri     interface{} `json:"PrecisFormeJuri"`
			RegimeConstitutif   string      `json:"RegimeConstitutif"`
			RegimeCourant       string      `json:"RegimeCourant"`
			TexteRespLimite     interface{} `json:"TexteRespLimite"`
			Type                string      `json:"Type"`
		} `json:"SousSecFormeJuridique"`
		SousSecFusion struct {
			GrilleFusion  interface{} `json:"GrilleFusion"`
			GrilleFusions []struct {
				EntetesColonnes []struct {
					ID      int    `json:"ID"`
					Libelle string `json:"Libelle"`
				} `json:"EntetesColonnes"`
				LibelleGrille string `json:"LibelleGrille"`
				Lignes        []struct {
					ID            int `json:"ID"`
					ListeCellules []struct {
						ID     int    `json:"ID"`
						Valeur string `json:"Valeur"`
					} `json:"ListeCellules"`
				} `json:"Lignes"`
			} `json:"GrilleFusions"`
			IDSousSection       int         `json:"IdSousSection"`
			MessageAucuneValeur interface{} `json:"MessageAucuneValeur"`
			Resultante          interface{} `json:"Resultante"`
		} `json:"SousSecFusion"`
		SousSecIdentification struct {
			DateEntreeEnVigueur  interface{} `json:"DateEntreeEnVigueur"`
			DateFinUtilisation   interface{} `json:"DateFinUtilisation"`
			IDSousSection        int         `json:"IdSousSection"`
			MessageAucuneValeur  interface{} `json:"MessageAucuneValeur"`
			NEQ                  string      `json:"NEQ"`
			NomEntreprise        string      `json:"NomEntreprise"`
			NomFamille           interface{} `json:"NomFamille"`
			Prenom               interface{} `json:"Prenom"`
			Situation            interface{} `json:"Situation"`
			Statut               interface{} `json:"Statut"`
			TableauAutresLangues interface{} `json:"TableauAutresLangues"`
		} `json:"SousSecIdentification"`
		SousSecImmatriculation struct {
			DateCessationPrevue   string      `json:"DateCessationPrevue"`
			DateImmatriculation   string      `json:"DateImmatriculation"`
			DateStatut            string      `json:"DateStatut"`
			IDSousSection         int         `json:"IdSousSection"`
			MessageAucuneValeur   interface{} `json:"MessageAucuneValeur"`
			StatutImmatriculation string      `json:"StatutImmatriculation"`
		} `json:"SousSecImmatriculation"`
		SousSecLiquidationDissolution struct {
			Cessation           interface{} `json:"Cessation"`
			IDSousSection       int         `json:"IdSousSection"`
			MessageAucuneValeur interface{} `json:"MessageAucuneValeur"`
			Texte               string      `json:"Texte"`
		} `json:"SousSecLiquidationDissolution"`
		SousSecObjectifs         interface{} `json:"SousSecObjectifs"`
		SousSecRecours           interface{} `json:"SousSecRecours"`
		SousSectionRapportAnnuel interface{} `json:"SousSectionRapportAnnuel"`
	} `json:"SectionInformationsGenerales"`
	SectionNomsAutreNoms struct {
		Consigne            string      `json:"Consigne"`
		IDSection           int         `json:"IdSection"`
		MessageAnalyse      interface{} `json:"MessageAnalyse"`
		MessageAucuneValeur interface{} `json:"MessageAucuneValeur"`
		SousSecAutresNoms   struct {
			GrilleNoms struct {
				EntetesColonnes []struct {
					ID      int    `json:"ID"`
					Libelle string `json:"Libelle"`
				} `json:"EntetesColonnes"`
				LibelleGrille string `json:"LibelleGrille"`
				Lignes        []struct {
					ID            int `json:"ID"`
					ListeCellules []struct {
						ID     int    `json:"ID"`
						Valeur string `json:"Valeur"`
					} `json:"ListeCellules"`
				} `json:"Lignes"`
			} `json:"GrilleNoms"`
			IDSousSection       int         `json:"IdSousSection"`
			MessageAucuneValeur interface{} `json:"MessageAucuneValeur"`
		} `json:"SousSecAutresNoms"`
		SousSecNoms struct {
			GrilleNoms struct {
				EntetesColonnes []struct {
					ID      int    `json:"ID"`
					Libelle string `json:"Libelle"`
				} `json:"EntetesColonnes"`
				LibelleGrille string `json:"LibelleGrille"`
				Lignes        []struct {
					ID            int `json:"ID"`
					ListeCellules []struct {
						ID     int    `json:"ID"`
						Valeur string `json:"Valeur"`
					} `json:"ListeCellules"`
				} `json:"Lignes"`
			} `json:"GrilleNoms"`
			IDSousSection       int         `json:"IdSousSection"`
			MessageAucuneValeur interface{} `json:"MessageAucuneValeur"`
		} `json:"SousSecNoms"`
	} `json:"SectionNomsAutreNoms"`
	TitreEtat      string `json:"TitreEtat"`
	TypeResultat   string `json:"TypeResultat"`
	ZoneImportante string `json:"ZoneImportante"`
	_Type          string `json:"__type"`
}
