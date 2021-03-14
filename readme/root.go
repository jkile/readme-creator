package readme

type Readme struct {
	imageUrl string
	license string
	projectName string
	description string
	installation string
	usage string
	contributing string
}

func (r *Readme) ImageUrl() string {
	return r.imageUrl
}

func (r *Readme) License() string {
	return r.license
}

func (r *Readme) ProjectName() string {
	return r.projectName
}

func (r *Readme) Description() string {
	return r.description
}

func (r *Readme) Installation() string {
	return r.installation
}

func (r *Readme) Usage() string {
	return r.usage
}

func (r *Readme) Contributing() string {
	return r.contributing
}

func (r *Readme) SetImageUrl(newUrl string) {
	r.imageUrl = newUrl
}

func (r *Readme) SetLicense(newLicense string) {
	r.license = newLicense
}

func (r *Readme) SetProjectName(newProjectName string) {
	r.projectName = newProjectName
}

func (r *Readme) SetDescription(newDescription string) {
	r.description = newDescription
}

func (r *Readme) SetInstallation(newInstallation string) {
	r.installation = newInstallation
}

func (r *Readme) SetUsage(newUsage string) {
	r.installation = newUsage
}

func (r *Readme) SetContributing(newContributing string) {
	r.contributing = newContributing
}