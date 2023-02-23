package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"goskill/db/utils"
	"testing"
	"time"
)

func createRandomSkill(t *testing.T) Skill {
	params := CreateSkillParams{
		Learner: utils.RandomName(6),
		Name:    utils.RandomName(5),
	}
	skill, err := testStore.CreateSkill(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, skill)

	require.Equal(t, skill.Learner, params.Learner)
	require.Equal(t, skill.Name, params.Name)
	return skill
}

func TestQueries_CreateSkill(t *testing.T) {
	createRandomSkill(t)
}

func TestQueries_GetSkillByID(t *testing.T) {
	skill := createRandomSkill(t)
	gSkill, err := testStore.GetSkillByID(context.Background(), skill.ID)
	require.NotEmpty(t, gSkill)
	require.NoError(t, err)
	require.Equal(t, skill.ID, gSkill.ID)
	require.Equal(t, skill.Name, gSkill.Name)
	require.Equal(t, skill.Learner, gSkill.Learner)
	require.Equal(t, skill.Score, gSkill.Score)
	require.WithinDuration(t, skill.CreatedAt, gSkill.CreatedAt, time.Second)
}
