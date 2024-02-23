package main_domains

type PlanMetadata struct {
	value             float64
	durationDays      int16
	numberOfUsers     int16
	numberOfProjects  int16
	numberOfCompanies int16
}

func NewPlanMetadata(
	value float64,
	durationDays int16,
	numberOfUsers int16,
	numberOfProjects int16,
	numberOfCompanies int16,
) *PlanMetadata {
	return &PlanMetadata{
		value:             value,
		durationDays:      durationDays,
		numberOfUsers:     numberOfUsers,
		numberOfProjects:  numberOfProjects,
		numberOfCompanies: numberOfCompanies,
	}
}

func (this PlanMetadata) GetValue() float64 {
	return this.value
}

func (this PlanMetadata) GetDurationDays() int16 {
	return this.durationDays
}

func (this PlanMetadata) GetNumberOfUsers() int16 {
	return this.numberOfUsers
}

func (this PlanMetadata) GetNumberOfProjects() int16 {
	return this.numberOfProjects
}

func (this PlanMetadata) GetNumberOfCompanies() int16 {
	return this.numberOfCompanies
}
