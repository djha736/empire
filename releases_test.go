package empire

import (
	"testing"
)

func TestReleasesServiceCreate(t *testing.T) {
	var scheduled bool

	app := &App{}
	config := &Config{}
	slug := &Slug{}

	p := &mockProcessesRepository{}
	r := &mockReleasesRepository{}
	m := &mockManager{
		ScheduleReleaseFunc: func(release *Release, config *Config, slug *Slug, formation Formation) error {
			scheduled = true
			return nil
		},
	}
	s := &releasesService{
		ReleasesRepository:  r,
		ProcessesRepository: p,
		Manager:             m,
	}

	if _, err := s.Create(app, config, slug, ""); err != nil {
		t.Fatal(err)
	}

	if got, want := scheduled, true; got != want {
		t.Fatal("Expected a release to be created")
	}
}

type mockReleasesRepository struct {
	ReleasesRepository // Just to satisfy the interface.

	HeadFunc                    func(AppName) (*Release, error)
	CreateFunc                  func(*Release) (*Release, error)
	FindByAppNameAndVersionFunc func(AppName, ReleaseVersion) (*Release, error)
}

func (s *mockReleasesRepository) Head(name AppName) (*Release, error) {
	if s.HeadFunc != nil {
		return s.HeadFunc(name)
	}

	return nil, nil
}

func (s *mockReleasesRepository) FindByAppNameAndVersion(a AppName, v ReleaseVersion) (*Release, error) {
	if s.FindByAppNameAndVersionFunc != nil {
		return s.FindByAppNameAndVersionFunc(a, v)
	}

	return nil, nil
}

func (s *mockReleasesRepository) Create(release *Release) (*Release, error) {
	if s.CreateFunc != nil {
		return s.CreateFunc(release)
	}

	return release, nil
}

type mockReleasesService struct {
	ReleasesService // Just to satisfy the interface.

	CreateFunc func(*App, *Config, *Slug, string) (*Release, error)
}

func (s *mockReleasesService) Create(app *App, config *Config, slug *Slug, desc string) (*Release, error) {
	if s.CreateFunc != nil {
		return s.CreateFunc(app, config, slug, desc)
	}

	return nil, nil
}
