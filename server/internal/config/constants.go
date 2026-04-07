package config

const ROOT_FOLDER_VAR = "PROJECT_ROOT_FOLDER"
const PPE = "p5aNLsDH3r$YDQQ5"
const PANIER_OUVERT = 4
const PANIER_FERMER = 5
const COMMANDE_OUVERT = 1
const COMMANDE_FERMER = 3
const COMMANDE_EN_COURS = 2
const STATUS_ASSIGN = 0

// routes constant
const (
	//swager
	SwaggerPath = "/swagger/*any"

	//adminPath
	AdminPath = "/admin"

	//dashboardPaath
	DashPath = "/dash"

	// authentification
	AuthPath = "/auth"
	Login    = "/login"
	Logout   = "/logout"
	Refresh  = "/refresh"

	// user
	UserPath = "/user"
	Creat    = "/creatUser"
	Delete   = "/:idUser/delete"
	Update   = "/updateUser"

	//profil
	ProfilPath   = "/profil"
	GetProfil    = "/list"
	AssignProfil = "/assignProfil"

	//article
	ArticlePath               = "/article"
	ListArticle               = "/list"
	AddArticle                = "/add"
	DeleteArticle             = "/:id/delete"
	FilterArticleByCommercant = "/filterCommerc/:commercant"
	FilterArticleByName       = "/filterArticle/:article"
	FilterArticleByCategorie  = "/filterCategorie/:categorie"
	FilterArticleAll          = "/filter-articles"
	ListeArticleByCommercant  = "/listArticleByCommerçant/:idCommercant"
	FilterArticles            = "/dash/article/commercant"
	UpdateArticle             = "/dash/article/update"
	GetArticleDetail          = "/articles/:id"
	Checkout                  = "/orders/create"

	WalletPath = "/wallet"
	Montant    = "/montant"

	//categorie
	CategoriePath  = "/categorie"
	ListCategorie  = "/list"
	ListCategories = "/categories/list"

	//panier
	PanierPath   = "/panier"
	AjoutPanier  = "/add"
	DetailPanier = "/:id_client/detail"
	DeletePanier = "/:id_panier/delete"

	//commande
	CommandePath      = "/commande"
	AjoutCommande     = "/add"
	AssignCommande    = "/:id_commande/assign/:id_livreur"
	ListCommandeCreer = "/commandeOuvert"
	CommandeAssign    = "/commandeAssign/:user_id"

	//livreur
	LivreurPath       = "/livreur"
	AjoutLivreur      = "/add"
	ListLivreur       = "/listLivreur"
	ListLivreurFilter = "/listLivreurFilter"

	//commercant
	CommercantPath    = "/commercant"
	ChercheCommercant = "/ChercheCommercant"

	//Delivery
	Delivery     = "/api/delivery"
	Tickets      = "/tickets"
	UpdateTicket = "/tickets/:id"
	AssignTicket = "/tickets/:id/assign"
	ListOrders   = "/orders/list"

	// reservation
	ReservationPath = "/reservations"
	ReservationID   = "/:id"
)
