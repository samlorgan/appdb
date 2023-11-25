package types

type SalesItem struct {
	SalesCode            string
	SalesPrice           uint
	SalesName            string
	CountryOfManufacture Country
}

type Country struct {
	CountryCode string
	CountryName string
	Free        bool
}
