package api

import (
	"testing"
	"time"

	"github.com/praelatus/backend/models"
	"github.com/praelatus/backend/store"
)

func TestMain(m *testing.M) {

}

type MockStore struct{}

func (ms MockStore) Users() store.UserStore {
	return MockUsersStore{}
}

func (ms MockStore) Teams() store.TeamStore {
	return MockTeamStore{}
}

func (ms MockStore) Labels() store.LabelStore {
	return MockLabelStore{}
}

func (ms MockStore) Fields() store.FieldStore {
	return MockFieldStore{}
}

func (ms MockStore) Tickets() store.TicketStore {
	return MockTicketStore{}
}

func (ms MockStore) Projects() store.ProjectStore {
	return MockProjectStore{}
}

func (ms MockStore) Statuses() store.StatusStore {
	return MockStatusStore{}
}

func (ms MockStore) Workflows() store.WorkflowStore {
	return MockWorkflowStore{}
}

//A Mock UsersStore struct
type MockUsersStore struct{}

func (ms MockUsersStore) Get(u *models.User) error {
	u.ID = 1
	u.Username = "foouser"
	u.Password = "foopass"
	u.Email = "foo@foo.com"
	u.FullName = "Foo McFooserson"
	u.IsActive = true
	return nil
}

func (ms MockUsersStore) GetAll() ([]models.User, error) {
	return []models.User{
		models.User{
			1,
			"foouser",
			"foopass",
			"foo@foo.com",
			"Foo McFooserson",
			"",
			"",
			false,
			true,
			models.Settings{},
		},
		models.User{
			2,
			"foouser",
			"foopass",
			"foo@foo.com",
			"Foo McFooserson",
			"",
			"",
			false,
			true,
			models.Settings{},
		},
	}, nil
}

func (ms MockUsersStore) New(u *models.User) error {
	u.ID = 1
	return nil
}

func (ms MockUsersStore) Save(u models.User) error {
	return nil
}

func (ms MockUsersStore) Remove(u models.User) error {
	return nil
}

//A Mock TeamStore struct
type MockTeamStore struct{}

func (ms MockTeamStore) Get(t *models.Team) error {
	t.ID = 1
	t.Name = "A"
	t.Lead = models.User{
		1,
		"foouser",
		"foopass",
		"foo@foo.com",
		"Foo McFooserson",
		"",
		"",
		false,
		true,
		models.Settings{},
	}
	t.Members = []models.User{
		models.User{
			1,
			"foouser",
			"foopass",
			"foo@foo.com",
			"Foo McFooserson",
			"",
			"",
			false,
			true,
			models.Settings{},
		},
		models.User{
			2,
			"foouser",
			"foopass",
			"foo@foo.com",
			"Foo McFooserson",
			"",
			"",
			false,
			true,
			models.Settings{},
		},
	}
	return nil
}

func (ms MockTeamStore) GetAll() ([]models.Team, error) {
	return []models.Team{
			models.Team{
				ID:   1,
				Name: "A",
				Lead: models.User{
					1,
					"foouser",
					"foopass",
					"foo@foo.com",
					"Foo McFooserson",
					"",
					"",
					false,
					true,
					models.Settings{},
				},
				Members: []models.User{
					models.User{
						1,
						"foouser",
						"foopass",
						"foo@foo.com",
						"Foo McFooserson",
						"",
						"",
						false,
						true,
						models.Settings{},
					},
					models.User{
						2,
						"foouser",
						"foopass",
						"foo@foo.com",
						"Foo McFooserson",
						"",
						"",
						false,
						true,
						models.Settings{},
					},
				},
			},
			models.Team{
				ID:   1,
				Name: "A",
				Lead: models.User{
					2,
					"foouser3",
					"foopass",
					"foo@foo3.com",
					"Foo McFooserson3",
					"",
					"",
					false,
					true,
					models.Settings{},
				},
				Members: []models.User{
					models.User{
						3,
						"foouser3",
						"foopass",
						"foo@foo3.com",
						"Foo McFooserson3",
						"",
						"",
						false,
						true,
						models.Settings{},
					},
					models.User{
						4,
						"foouser4",
						"foopass",
						"foo@foo4.com",
						"Foo McFooserson4",
						"",
						"",
						false,
						true,
						models.Settings{},
					},
				},
			},
		},
		nil
}

func (ms MockTeamStore) GetForUser(m models.User) ([]models.Team, error) {
	return []models.Team{
		models.Team{
			ID:   1,
			Name: "A",
			Lead: models.User{
				1,
				"foouser",
				"foopass",
				"foo@foo.com",
				"Foo McFooserson",
				"",
				"",
				false,
				true,
				models.Settings{},
			},
			Members: []models.User{
				models.User{
					1,
					"foouser",
					"foopass",
					"foo@foo.com",
					"Foo McFooserson",
					"",
					"",
					false,
					true,
					models.Settings{},
				},
				models.User{
					2,
					"foouser",
					"foopass",
					"foo@foo.com",
					"Foo McFooserson",
					"",
					"",
					false,
					true,
					models.Settings{},
				},
			},
		},
		models.Team{
			ID:   1,
			Name: "A",
			Lead: models.User{
				2,
				"foouser3",
				"foopass",
				"foo@foo3.com",
				"Foo McFooserson3",
				"",
				"",
				false,
				true,
				models.Settings{},
			},
			Members: []models.User{
				models.User{
					3,
					"foouser3",
					"foopass",
					"foo@foo3.com",
					"Foo McFooserson3",
					"",
					"",
					false,
					true,
					models.Settings{},
				},
				models.User{
					4,
					"foouser4",
					"foopass",
					"foo@foo4.com",
					"Foo McFooserson4",
					"",
					"",
					false,
					true,
					models.Settings{},
				},
			},
		},
	}, nil
}

func (ms MockTeamStore) AddMembers(m models.Team, u ...models.User) error {
	return nil
}

func (ms MockTeamStore) New(t *models.Team) error {
	t.ID = 1
	return nil
}

func (ms MockTeamStore) Save(t models.Team) error {
	return nil
}

func (ms MockTeamStore) Remove(t models.Team) error {
	return nil
}

//A Mock LabelStore struct
type MockLabelStore struct{}

func (ms MockLabelStore) Get(l *models.Label) error {
	l.ID = 1
	l.Name = "mock"
	return nil
}

func (ms MockLabelStore) GetAll() ([]models.Label, error) {
	return []models.Label{
		models.Label{
			ID:   1,
			Name: "mock",
		},
		models.Label{
			ID:   2,
			Name: "fake",
		},
	}, nil
}

func (ms MockLabelStore) New(l *models.Label) error {
	l.ID = 1
	return nil
}

func (ms MockLabelStore) Save(l models.Label) error {
	return nil
}

func (ms MockLabelStore) Remove(l models.Label) error {
	return nil
}

//A Mock FieldStore struct
type MockFieldStore struct{}

func (MockFieldStore) Get(f *models.Field) error {
	f.ID = 1
	f.Name = "String Field"
	f.DataType = "STRING"
	return nil
}

func (MockFieldStore) GetAll() ([]models.Field, error) {
	return []models.Field{
		models.Field{
			ID:       1,
			Name:     "String Field",
			DataType: "STRING",
		},
		models.Field{
			ID:       2,
			Name:     "Int Field",
			DataType: "INT",
		},
	}, nil
}

func (MockFieldStore) GetByProject(p models.Project) ([]models.Field, error) {
	return []models.Field{
		models.Field{
			ID:       1,
			Name:     "String Field",
			DataType: "STRING",
		},
		models.Field{
			ID:       2,
			Name:     "Int Field",
			DataType: "INT",
		},
	}, nil
}

func (MockFieldStore) AddToProject(p models.Project, f *models.Field, t ...models.TicketType) error {
	return nil
}

func (MockFieldStore) New(f *models.Field) error {
	f.ID = 1
	return nil
}

func (MockFieldStore) Save(f models.Field) error {
	return nil
}

func (MockFieldStore) Remove(f models.Field) error {
	return nil
}

// //A Mock TicketStore struct
type MockTicketStore struct{}

func (MockTicketStore) Get(t *models.Ticket) error {
	t.ID = 1

	t.CreatedDate = time.Now()
	t.UpdatedDate = time.Now()

	t.Key = "TEST-1"

	t.Summary = "A mock issue"
	t.Description = "This issue is a fake."

	t.Fields = []models.FieldValue{
		models.FieldValue{
			ID:       1,
			Name:     "String Field",
			DataType: "STRING",
			Value:    "This is a string",
		},
		models.FieldValue{
			ID:       2,
			Name:     "Int Field",
			DataType: "INT",
			Value:    3,
		},
	}

	t.Labels = []models.Label{
		models.Label{
			ID:   1,
			Name: "mock",
		},
	}

	t.Type = models.TicketType{1, "Bug"}

	t.Reporter = models.User{
		1,
		"foouser",
		"foopass",
		"foo@foo.com",
		"Foo McFooserson",
		"",
		"",
		false,
		true,
		models.Settings{},
	}

	t.Assignee = models.User{
		2,
		"baruser",
		"barpass",
		"bar@bar.com",
		"Bar McBarserson",
		"",
		"",
		true,
		true,
		models.Settings{},
	}

	t.Status = models.Status{
		ID:   1,
		Name: "In Progress",
	}

	return nil
}

func (ms MockTicketStore) GetAll() ([]models.Ticket, error) {
	return []models.Ticket{
		models.Ticket{
			ID:          1,
			CreatedDate: time.Now(),
			UpdatedDate: time.Now(),

			Key: "TEST-1",

			Summary:     "A mock issue",
			Description: "This issue is a fake.",

			Fields: []models.FieldValue{
				models.FieldValue{
					ID:       1,
					Name:     "String Field",
					DataType: "STRING",
					Value:    "This is a string",
				},
				models.FieldValue{
					ID:       2,
					Name:     "Int Field",
					DataType: "INT",
					Value:    3,
				},
			},

			Labels: []models.Label{
				models.Label{
					ID:   1,
					Name: "mock",
				},
			},

			Type: models.TicketType{1, "Bug"},

			Reporter: models.User{
				1,
				"foouser",
				"foopass",
				"foo@foo.com",
				"Foo McFooserson",
				"",
				"",
				false,
				true,
				models.Settings{},
			},

			Assignee: models.User{
				2,
				"baruser",
				"barpass",
				"bar@bar.com",
				"Bar McBarserson",
				"",
				"",
				true,
				true,
				models.Settings{},
			},

			Status: models.Status{
				ID:   1,
				Name: "In Progress",
			},
		},

		models.Ticket{
			ID:          2,
			CreatedDate: time.Now(),
			UpdatedDate: time.Now(),

			Key: "TEST-2",

			Summary:     "A mock issue",
			Description: "This issue is a fake.",

			Fields: []models.FieldValue{
				models.FieldValue{
					ID:       1,
					Name:     "String Field",
					DataType: "STRING",
					Value:    "This is a string",
				},
				models.FieldValue{
					ID:       2,
					Name:     "Int Field",
					DataType: "INT",
					Value:    3,
				},
			},

			Labels: []models.Label{
				models.Label{
					ID:   1,
					Name: "mock",
				},
			},

			Type: models.TicketType{1, "Bug"},

			Reporter: models.User{
				1,
				"foouser",
				"foopass",
				"foo@foo.com",
				"Foo McFooserson",
				"",
				"",
				false,
				true,
				models.Settings{},
			},

			Assignee: models.User{
				2,
				"baruser",
				"barpass",
				"bar@bar.com",
				"Bar McBarserson",
				"",
				"",
				true,
				true,
				models.Settings{},
			},

			Status: models.Status{
				ID:   1,
				Name: "In Progress",
			},
		},
	}, nil
}
func (ms MockTicketStore) GetAllByProject(p models.Project) ([]models.Ticket, error) {
	return []models.Ticket{
		models.Ticket{
			ID:          1,
			CreatedDate: time.Now(),
			UpdatedDate: time.Now(),

			Key: "TEST-1",

			Summary:     "A mock issue",
			Description: "This issue is a fake.",

			Fields: []models.FieldValue{
				models.FieldValue{
					ID:       1,
					Name:     "String Field",
					DataType: "STRING",
					Value:    "This is a string",
				},
				models.FieldValue{
					ID:       2,
					Name:     "Int Field",
					DataType: "INT",
					Value:    3,
				},
			},

			Labels: []models.Label{
				models.Label{
					ID:   1,
					Name: "mock",
				},
			},

			Type: models.TicketType{1, "Bug"},

			Reporter: models.User{
				1,
				"foouser",
				"foopass",
				"foo@foo.com",
				"Foo McFooserson",
				"",
				"",
				false,
				true,
				models.Settings{},
			},

			Assignee: models.User{
				2,
				"baruser",
				"barpass",
				"bar@bar.com",
				"Bar McBarserson",
				"",
				"",
				true,
				true,
				models.Settings{},
			},

			Status: models.Status{
				ID:   1,
				Name: "In Progress",
			},
		},

		models.Ticket{
			ID:          2,
			CreatedDate: time.Now(),
			UpdatedDate: time.Now(),

			Key: "TEST-2",

			Summary:     "A mock issue",
			Description: "This issue is a fake.",

			Fields: []models.FieldValue{
				models.FieldValue{
					ID:       1,
					Name:     "String Field",
					DataType: "STRING",
					Value:    "This is a string",
				},
				models.FieldValue{
					ID:       2,
					Name:     "Int Field",
					DataType: "INT",
					Value:    3,
				},
			},

			Labels: []models.Label{
				models.Label{
					ID:   1,
					Name: "mock",
				},
			},

			Type: models.TicketType{1, "Bug"},

			Reporter: models.User{
				1,
				"foouser",
				"foopass",
				"foo@foo.com",
				"Foo McFooserson",
				"",
				"",
				false,
				true,
				models.Settings{},
			},

			Assignee: models.User{
				2,
				"baruser",
				"barpass",
				"bar@bar.com",
				"Bar McBarserson",
				"",
				"",
				true,
				true,
				models.Settings{},
			},

			Status: models.Status{
				ID:   1,
				Name: "In Progress",
			},
		},
	}, nil
}

func (ms MockTicketStore) GetComments(t models.Ticket) ([]models.Comment, error) {
	return []models.Comment{
		models.Comment{
			1,
			time.Now(),
			time.Now(),
			"This is a fake comment",
			models.User{
				2,
				"baruser",
				"barpass",
				"bar@bar.com",
				"Bar McBarserson",
				"",
				"",
				true,
				true,
				models.Settings{},
			},
		},
	}, nil
}

func (ms MockTicketStore) NewComment(t models.Ticket, c *models.Comment) error {
	c.ID = 1
	return nil
}

func (ms MockTicketStore) SaveComment(c models.Comment) error {
	return nil
}

func (ms MockTicketStore) RemoveComment(c models.Comment) error {
	return nil
}

func (ms MockTicketStore) NextTicketKey(p models.Project) string {
	return "TEST-2"
}

func (ms MockTicketStore) New(p models.Project, t *models.Ticket) error {
	t.ID = 1
	return nil
}

func (ms MockTicketStore) Save(t models.Ticket) error {
	return nil
}

func (ms MockTicketStore) Remove(t models.Ticket) error {
	return nil
}

// A Mock TypeStore struct
type MockTypeStore struct{}

func (ms MockTypeStore) Get(t models.TicketType) error {
	t.ID = 1
	t.Name = "Mock Type"
	return nil
}

func (ms MockTypeStore) GetAll() ([]models.TicketType, error) {
	return []models.TicketType{
		models.TicketType{
			ID:   1,
			Name: "Mock Type",
		},
		models.TicketType{
			ID:   2,
			Name: "Fake Type",
		},
	}, nil
}

func (ms MockTypeStore) New(t *models.TicketType) error {
	t.ID = 1
	return nil
}

func (ms MockTypeStore) Save(t models.TicketType) error {
	return nil
}

func (ms MockTypeStore) Remove(t models.TicketType) error {
	return nil
}

// A Mock ProjectStore struct
type MockProjectStore struct{}

func (ms MockProjectStore) Get(p *models.Project) error {
	p.ID = 1
	p.Name = "Test Project"
	p.Key = "TEST"
	p.CreatedDate = time.Now()
	p.Lead = models.User{
		2,
		"baruser",
		"barpass",
		"bar@bar.com",
		"Bar McBarserson",
		"",
		"",
		true,
		true,
		models.Settings{},
	}
	return nil
}

func (ms MockProjectStore) GetAll() ([]models.Project, error) {
	return []models.Project{
		models.Project{
			ID:          1,
			CreatedDate: time.Now(),
			Name:        "Test Project",
			Key:         "TEST",
			Lead: models.User{
				2,
				"baruser",
				"barpass",
				"bar@bar.com",
				"Bar McBarserson",
				"",
				"",
				true,
				true,
				models.Settings{},
			},
		},
		models.Project{
			ID:          2,
			Name:        "Mock Project",
			Key:         "MOCK",
			CreatedDate: time.Now(),
			Lead: models.User{
				1,
				"foouser",
				"foopass",
				"foo@foo.com",
				"Foo McFooserson",
				"",
				"",
				false,
				true,
				models.Settings{},
			},
		},
	}, nil
}

func (ms MockProjectStore) New(p *models.Project) error {
	p.ID = 1
	return nil
}

func (ms MockProjectStore) Save(p models.Project) error {
	return nil
}

func (ms MockProjectStore) Remove(p models.Project) error {
	return nil
}

// A Mock StatusStore struct
type MockStatusStore struct{}

func (ms MockStatusStore) Get(s *models.Status) error {
	s.ID = 1
	s.Name = "Mock Status"
	return nil
}

func (ms MockStatusStore) GetAll() ([]models.Status, error) {
	return []models.Status{
		models.Status{
			1,
			"Mock Status",
		},
		models.Status{
			2,
			"Fake Status",
		},
	}, nil
}

func (ms MockStatusStore) New(s *models.Status) error {
	s.ID = 1
	return nil
}

func (ms MockStatusStore) Save(p models.Status) error {
	return nil
}

func (ms MockStatusStore) Remove(p models.Status) error {
	return nil
}

// A Mock Workflow Store
type MockWorkflowStore struct{}

func (ms MockWorkflowStore) Get(w *models.Workflow) error {
	w = &models.Workflow{
		Name: "Simple Workflow",
		Transitions: map[string][]models.Transition{
			"Backlog": []models.Transition{
				models.Transition{
					Name:     "In Progress",
					ToStatus: models.Status{ID: 2},
					Hooks:    []models.Hook{},
				},
			},
			"In Progress": []models.Transition{
				models.Transition{
					Name:     "Done",
					ToStatus: models.Status{ID: 3},
					Hooks:    []models.Hook{},
				},
				models.Transition{
					Name:     "Backlog",
					ToStatus: models.Status{ID: 1},
					Hooks:    []models.Hook{},
				},
			},
			"Done": []models.Transition{
				models.Transition{
					Name:     "ReOpen",
					ToStatus: models.Status{ID: 1},
					Hooks:    []models.Hook{},
				},
			},
		},
	}

	return nil
}

func (ms MockWorkflowStore) GetByProject(p models.Project) ([]models.Workflow, error) {
	return []models.Workflow{
		models.Workflow{
			ID:   1,
			Name: "Simple Workflow",
			Transitions: map[string][]models.Transition{
				"Backlog": []models.Transition{
					models.Transition{
						Name:     "In Progress",
						ToStatus: models.Status{ID: 2},
						Hooks:    []models.Hook{},
					},
				},
				"In Progress": []models.Transition{
					models.Transition{
						Name:     "Done",
						ToStatus: models.Status{ID: 3},
						Hooks:    []models.Hook{},
					},
					models.Transition{
						Name:     "Backlog",
						ToStatus: models.Status{ID: 1},
						Hooks:    []models.Hook{},
					},
				},
				"Done": []models.Transition{
					models.Transition{
						Name:     "ReOpen",
						ToStatus: models.Status{ID: 1},
						Hooks:    []models.Hook{},
					},
				},
			},
		},

		models.Workflow{
			ID:   2,
			Name: "Another Simple Workflow",
			Transitions: map[string][]models.Transition{
				"Backlog": []models.Transition{
					models.Transition{
						Name:     "In Progress",
						ToStatus: models.Status{ID: 2},
						Hooks:    []models.Hook{},
					},
				},
				"In Progress": []models.Transition{
					models.Transition{
						Name:     "Done",
						ToStatus: models.Status{ID: 3},
						Hooks:    []models.Hook{},
					},
					models.Transition{
						Name:     "Backlog",
						ToStatus: models.Status{ID: 1},
						Hooks:    []models.Hook{},
					},
				},
				"Done": []models.Transition{
					models.Transition{
						Name:     "ReOpen",
						ToStatus: models.Status{ID: 1},
						Hooks:    []models.Hook{},
					},
				},
			},
		},
	}, nil
}

func (ms MockWorkflowStore) GetAll() ([]models.Workflow, error) {
	return []models.Workflow{
		models.Workflow{
			ID:   1,
			Name: "Simple Workflow",
			Transitions: map[string][]models.Transition{
				"Backlog": []models.Transition{
					models.Transition{
						Name:     "In Progress",
						ToStatus: models.Status{ID: 2},
						Hooks:    []models.Hook{},
					},
				},
				"In Progress": []models.Transition{
					models.Transition{
						Name:     "Done",
						ToStatus: models.Status{ID: 3},
						Hooks:    []models.Hook{},
					},
					models.Transition{
						Name:     "Backlog",
						ToStatus: models.Status{ID: 1},
						Hooks:    []models.Hook{},
					},
				},
				"Done": []models.Transition{
					models.Transition{
						Name:     "ReOpen",
						ToStatus: models.Status{ID: 1},
						Hooks:    []models.Hook{},
					},
				},
			},
		},

		models.Workflow{
			ID:   2,
			Name: "Another Simple Workflow",
			Transitions: map[string][]models.Transition{
				"Backlog": []models.Transition{
					models.Transition{
						Name:     "In Progress",
						ToStatus: models.Status{ID: 2},
						Hooks:    []models.Hook{},
					},
				},
				"In Progress": []models.Transition{
					models.Transition{
						Name:     "Done",
						ToStatus: models.Status{ID: 3},
						Hooks:    []models.Hook{},
					},
					models.Transition{
						Name:     "Backlog",
						ToStatus: models.Status{ID: 1},
						Hooks:    []models.Hook{},
					},
				},
				"Done": []models.Transition{
					models.Transition{
						Name:     "ReOpen",
						ToStatus: models.Status{ID: 1},
						Hooks:    []models.Hook{},
					},
				},
			},
		},
	}, nil
}

func (ms MockWorkflowStore) New(p models.Project, w *models.Workflow) error {
	w.ID = 1
	return nil
}

func (ms MockWorkflowStore) Save(p models.Workflow) error {
	return nil
}

func (ms MockWorkflowStore) Remove(p models.Workflow) error {
	return nil
}
