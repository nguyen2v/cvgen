package cv

// CV is representation of data for document
type CV struct {
	Name string
	Contact struct {
		Phone string
		Email string
	}
	Address struct {
		City string
		Postcode string
		Street string
	}
	Header string

	Experience []ExperienceItem
	Education []EducationItem
	Skills []SkillItem
	Languages []LanguageItem
	Certificates []CertificateItem

	Hobby []string `yaml:",flow"`
}
