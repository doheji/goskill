package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Store struct {
	db *sql.Conn
	*Queries
}

func NewStore(db *sql.Conn) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (s *Store) execTransaction(ctx context.Context, f func(queries *Queries) error) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	query := New(tx)
	if err = f(query); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("the transaction error was %s and the rollback failed with %s", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

type CompleteSessionResult struct {
	Skill   Skill   `json:"skill"`
	Session Session `json:"session"`
}

func (s *Store) CompleteSession(ctx context.Context, sessionID int64) (CompleteSessionResult, error) {
	var result CompleteSessionResult
	err := s.execTransaction(ctx,
		func(q *Queries) error {
			// Update the session attendance
			ses, err := q.UpdateSession(ctx, UpdateSessionParams{
				Toupdatetaskid:      false,
				Updatedtaskid:       0,
				Toupdatename:        false,
				Updatedname:         "",
				Toupdatedescription: false,
				Updateddescription:  sql.NullString{},
				Toupdategoal:        false,
				Updatedgoal:         sql.NullString{},
				Toupdatelocation:    false,
				Updatedlocation:     sql.NullString{},
				Toupdateattended:    true,
				Updatedattended:     true,
				Toupdatestarttime:   false,
				Updatedstarttime:    time.Time{},
				Toupdateduration:    false,
				Updatedduration:     0,
				Sessionid:           sessionID,
			})

			if err != nil {
				return err
			}

			result.Session = ses

			// Get the skill
			skill, err := q.GetSkillBySession(ctx, sessionID)
			if err != nil {
				return err
			}

			// Increase it
			skill, err = q.IncreaseSkillScore(ctx, IncreaseSkillScoreParams{
				Score: 1,
				ID:    skill.ID,
			})
			if err != nil {
				return err
			}
			result.Skill = skill
			return nil
		},
	)
	if err != nil {
		return result, err
	}
	return result, nil
}
