package pg

import (
	"database/sql"
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/praelatus/backend/models"
)

// TicketStore contains methods for storing and retrieving Tickets from
// Postgres DB
type TicketStore struct {
	db *sql.DB
}

func getOpts(db *sql.DB, fid int64, fo *models.FieldOption) error {
	rows, err := db.Query(`SELECT option FROM field_options 
						   WHERE field_id = $1`, fid)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var opt string

		err = rows.Scan(&opt)
		if err != nil {
			return err
		}

		fo.Options = append(fo.Options, opt)
	}

	return nil
}

func populateFields(db *sql.DB, t *models.Ticket) error {
	rows, err := db.Query(`
		SELECT fv.id, f.name, f.data_type, 
			   fv.int_value, fv.flt_value, fv.str_value, 
			   fv.opt_value, fv.dte_value, f.id
		FROM field_values AS fv
		JOIN fields AS f ON f.id = fv.field_id
		WHERE fv.ticket_id = $1`, t.ID)
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		// We need to be able to scan in all the values then determine which
		// actually goes into the model.
		var fv *models.FieldValue
		var i int
		var f float64
		var s, o string
		var d time.Time
		var fID int64

		err = rows.Scan(fv.ID, fv.Name, fv.DataType, &i, &f, &s, &o, &d, &fID)
		if err != nil {
			return err
		}

		// By Odin's Beard I can't think of a better way to wrangle this mess.
		switch fv.DataType {
		case "FLOAT":
			fv.Value = f
		case "INT":
			fv.Value = i
		case "STRING":
			fv.Value = s
		case "DATE":
			fv.Value = d
		case "OPT":
			fo := models.FieldOption{}
			fo.Selected = o

			// Fill out the options and defaults.
			e := getOpts(db, fID, &fo)
			if e != nil {
				return e
			}

			fv.Value = fo
		default:
			fv.Value = nil
		}

		t.Fields = append(t.Fields, *fv)
	}

	return nil
}

func intoTicket(row rowScanner, db *sql.DB, t *models.Ticket) error {
	var ajson, rjson, sjson, tjson json.RawMessage

	err := row.Scan(&t.ID, &t.Key, &t.CreatedDate, &t.UpdatedDate, &t.Summary,
		&t.Description, &ajson, &rjson, &sjson, &tjson)
	if err != nil {
		return handlePqErr(err)
	}

	dberr := make(chan error)
	done := make(chan struct{})
	defer close(done)

	go func() {
		defer close(dberr)
		select {
		case dberr <- populateFields(db, t):
			return
		case <-done:
			return
		}
	}()

	err = json.Unmarshal(ajson, &t.Assignee)
	if err != nil {
		done <- struct{}{}
		return err
	}

	err = json.Unmarshal(rjson, &t.Reporter)
	if err != nil {
		done <- struct{}{}
		return err
	}

	err = json.Unmarshal(sjson, &t.Status)
	if err != nil {
		done <- struct{}{}
		return err
	}

	err = json.Unmarshal(tjson, &t.Type)
	if err != nil {
		done <- struct{}{}
		return err
	}

	err = <-dberr
	if err != nil {
		log.Println("Errored while getting fields.")
	}
	return handlePqErr(err)
}

// Get gets a Ticket from a postgres DB by it's ID
func (ts *TicketStore) Get(t *models.Ticket) error {
	row := ts.db.QueryRow(`SELECT t.id, t.key, t.created_date, 
									 t.updated_date, t.summary, t.description, 
									 row_to_json(a.*) AS assignee, 
									 row_to_json(r.*) AS reporter, 
									 row_to_json(s.*) AS status, 
									 row_to_json(tt.*) AS ticket_type 
						   FROM tickets AS t 
						   JOIN users AS a ON a.id = t.assignee_id
						   JOIN users AS r ON r.id = t.reporter_id
						   JOIN statuses AS s ON s.id = t.status_id
						   JOIN ticket_types AS tt ON tt.id = t.ticket_type_id
						   JOIN projects AS p ON p.id = t.project_id
						   WHERE t.id = $1 
						   OR t.key = $2`, t.ID, t.Key)

	err := intoTicket(row, ts.db, t)
	return handlePqErr(err)
}

// GetAll gets all the Tickets from the database
func (ts *TicketStore) GetAll() ([]models.Ticket, error) {
	var tickets []models.Ticket

	rows, err := ts.db.Query(`SELECT t.id, t.key, t.created_date, 
									  t.updated_date, t.summary, t.description, 
									  row_to_json(a.*) AS assignee, 
									  row_to_json(r.*) AS reporter, 
									  row_to_json(s.*) AS status, 
									  row_to_json(tt.*) AS ticket_type 
							  FROM tickets AS t 
							  JOIN users AS a ON a.id = t.assignee_id
							  JOIN users AS r ON r.id = t.reporter_id
							  JOIN statuses AS s ON s.id = t.status_id
							  JOIN ticket_types AS tt ON tt.id = t.ticket_type_id`)
	if err != nil {
		return tickets, handlePqErr(err)
	}

	for rows.Next() {
		var t models.Ticket

		err = intoTicket(rows, ts.db, &t)
		if err != nil {
			log.Println("Error getting tickets")
			return tickets, handlePqErr(err)
		}

		tickets = append(tickets, t)
	}

	return tickets, nil
}

// GetAllByProject gets all the Tickets from the database based on the given
// project
func (ts *TicketStore) GetAllByProject(p models.Project) ([]models.Ticket, error) {
	var tickets []models.Ticket

	rows, err := ts.db.Query(`SELECT t.id, t.key, t.created_date, 
									  t.updated_date, t.summary, t.description, 
									  row_to_json(a.*) AS assignee, 
									  row_to_json(r.*) AS reporter, 
									  row_to_json(s.*) AS status, 
									  row_to_json(tt.*) AS ticket_type 
							  FROM tickets AS t 
							  JOIN users AS a ON a.id = t.assignee_id
							  JOIN users AS r ON r.id = t.reporter_id
							  JOIN projects AS p ON p.id = t.project_id
							  JOIN statuses AS s ON s.id = t.status_id
							  JOIN ticket_types AS tt ON tt.id = t.ticket_type_id
							  WHERE p.id = $1
							  OR p.key = $2`, p.ID, p.Key)
	if err != nil {
		return tickets, handlePqErr(err)
	}

	for rows.Next() {
		var t models.Ticket

		err = intoTicket(rows, ts.db, &t)
		if err != nil {
			return tickets, handlePqErr(err)
		}

		tickets = append(tickets, t)
	}

	return tickets, nil
}

// Save will update an existing ticket in the postgres DB
func (ts *TicketStore) Save(ticket models.Ticket) error {
	_, err := ts.db.Exec(`UPDATE tickets SET 
						  (summary, description, updated_date) = ($1, $2, $3) 
						  WHERE id = $4`,
		ticket.Summary, ticket.Description, time.Now(), ticket.ID)

	for _, fv := range ticket.Fields {
		jsn, err := json.Marshal(fv.Value)
		if err != nil {
			return err
		}

		_, err = ts.db.Exec(`UPDATE field_values 
							 SET (name, data_type, value) = ($1, $2, $3)
							 WHERE id = $4`, fv.Name, fv.DataType, jsn, fv.ID)
		if err != nil {
			return handlePqErr(err)
		}
	}

	return handlePqErr(err)
}

// Remove will update an existing ticket in the postgres DB
func (ts *TicketStore) Remove(ticket models.Ticket) error {
	tx, err := ts.db.Begin()
	if err != nil {
		return handlePqErr(tx.Rollback())
	}

	_, err = tx.Exec(`DELETE FROM field_values WHERE ticket_id = $1;`, ticket.ID)
	if err != nil {
		return handlePqErr(tx.Rollback())
	}

	_, err = tx.Exec(`DELETE FROM tickets_labels WHERE ticket_id = $1;`, ticket.ID)
	if err != nil {
		return handlePqErr(tx.Rollback())
	}

	_, err = tx.Exec(`DELETE FROM tickets WHERE id = $1;`, ticket.ID)
	if err != nil {
		return handlePqErr(tx.Rollback())
	}

	return handlePqErr(tx.Commit())
}

// New will add a new Ticket to the postgres DB
func (ts *TicketStore) New(project models.Project, ticket *models.Ticket) error {
	// TODO update fields?
	err := ts.db.QueryRow(`INSERT INTO tickets 
						   (summary, description, project_id, assignee_id, 
						   reporter_id, ticket_type_id, status_id, key) 
						   VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
						   RETURNING id;`,
		ticket.Summary, ticket.Description, project.ID,
		ticket.Assignee.ID, ticket.Reporter.ID, ticket.Type.ID,
		ticket.Status.ID, ticket.Key).
		Scan(&ticket.ID)

	for _, fv := range ticket.Fields {
		jsn, err := json.Marshal(fv)
		if err != nil {
			return err
		}

		err = ts.db.QueryRow(`INSERT INTO field_values 
							  VALUES (name, data_type, value) = ($1, $2, $3)
							  RETURNING id`, fv.Name, fv.DataType, jsn).
			Scan(&fv.ID)
	}

	return handlePqErr(err)
}

// GetComments will return all comments for a ticket based on it's ID
func (ts *TicketStore) GetComments(t models.Ticket) ([]models.Comment, error) {
	var comments []models.Comment

	rows, err := ts.db.Query(`SELECT c.id, c.created_date, c.updated_date, 
									 c.body, row_to_json(users.*) as author 
							  FROM comments AS c
							  JOIN tickets AS t ON t.id = c.ticket_id
							  JOIN users ON users.id = c.author_id
							  WHERE t.id = $1
							  OR t.key = $2`, t.ID, t.Key)

	if err != nil {
		return comments, handlePqErr(err)
	}

	for rows.Next() {
		var c models.Comment
		var ajson json.RawMessage

		err := rows.Scan(&c.ID, &c.CreatedDate, &c.UpdatedDate, &c.Body, &ajson)
		if err != nil {
			return comments, handlePqErr(err)
		}

		err = json.Unmarshal(ajson, &c.Author)
		if err != nil {
			return comments, handlePqErr(err)
		}

		comments = append(comments, c)
	}

	return comments, nil
}

// NewComment will add a new Comment to the postgres DB
func (ts *TicketStore) NewComment(t models.Ticket, c *models.Comment) error {
	_, err := ts.db.Exec(`UPDATE tickets SET (updated_date) = ($1) 
					      WHERE id = $2;`, time.Now(), t.ID)
	if err != nil {
		return handlePqErr(err)
	}

	err = ts.db.QueryRow(`INSERT INTO comments 
						  (body, ticket_id, author_id) VALUES ($1, $2, $3)
						  RETURNING id;`, c.Body, t.ID, c.Author.ID).
		Scan(&c.ID)

	return handlePqErr(err)
}

// SaveComment will add a new Comment to the postgres DB
func (ts *TicketStore) SaveComment(c models.Comment) error {
	_, err := ts.db.Exec(`UPDATE comments 
						  SET (body, updated_date, author_id) = ($1, $2, $3)
						  WHERE id = $4`,
		c.Body, time.Now(), c.Author.ID, c.ID)

	return handlePqErr(err)
}

// RemoveComment will add a new Comment to the postgres DB
func (ts *TicketStore) RemoveComment(c models.Comment) error {
	_, err := ts.db.Exec("DELETE FROM comments WHERE id = $1", c.ID)
	return handlePqErr(err)
}

// NextTicketKey will generate the appropriate number for a ticket key
func (ts *TicketStore) NextTicketKey(p models.Project) string {
	var count int

	err := ts.db.QueryRow(`SELECT COUNT(t.id) FROM tickets AS t
						   JOIN projects AS p ON p.id = t.project_id
						   WHERE p.id = $1 OR p.key = $2`, p.ID, p.Key).
		Scan(&count)
	if err != nil {
		handlePqErr(err)
		return p.Key + strconv.Itoa(1)
	}

	return p.Key + strconv.Itoa(count+1)
}
