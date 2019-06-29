package cv

type ExperienceItem struct {
	Name string
	StartDate string `yaml:"start_date"`
	EndDate string `yaml:"end_date"`
	Position string
	Description string
}

// EducationItem is representation of education history
type EducationItem struct {
	Name string
	StartDate string `yaml:"start_date"`
	EndDate string `yaml:"end_date"`
	Specialization string
}

// SkillItem store information about skills
type SkillItem struct {
	Name string
	Description string
}

// LanguageItem store information about language and level
type LanguageItem struct {
	Name string
	Level string
}

// CertificateItem store information about achievements
type CertificateItem struct {
	Name string
	When string
}