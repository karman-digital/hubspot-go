package shared

type ObjectNames struct {
	Contact string
	Deal    string
	Company string
}

var DefaultObjectNames = ObjectNames{
	Contact: "contact",
	Deal:    "deal",
	Company: "company",
}
