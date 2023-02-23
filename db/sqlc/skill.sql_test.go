package db

import (
	"context"
	"database/sql"
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

func TestQueries_GetAllSkills(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomSkill(t)
	}

	skills, err := testStore.GetAllSkills(context.Background(), GetAllSkillsParams{
		Limit:  5,
		Offset: 5,
	})

	require.NoError(t, err)
	require.NotEmpty(t, skills)
	require.Len(t, skills, 5)

	for _, skill := range skills {
		require.NotEmpty(t, skill)
	}
}

func TestQueries_DeleteSkillByID(t *testing.T) {
	skill := createRandomSkill(t)
	err := testStore.DeleteSkillByID(context.Background(), skill.ID)
	require.NoError(t, err)
	skill, err = testStore.GetSkillByID(context.Background(), skill.ID)
	require.Error(t, err)
	require.Empty(t, skill)
	require.ErrorIs(t, err, sql.ErrNoRows)
}

func createSkillByLearner(t *testing.T, learner string) Skill {
	params := CreateSkillParams{
		Learner: learner,
		Name:    utils.RandomName(5),
	}
	skill, err := testStore.CreateSkill(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, skill)

	require.Equal(t, skill.Learner, params.Learner)
	require.Equal(t, skill.Name, params.Name)
	return skill

}

func TestQueries_GetSkillsByLearner(t *testing.T) {
	firstSkill := createRandomSkill(t)
	skills, err := testStore.GetSkillsByLearner(context.Background(), GetSkillsByLearnerParams{
		Learner: firstSkill.Learner,
		Limit:   5,
		Offset:  0,
	})
	require.NoError(t, err)
	require.Len(t, skills, 1)

	skills, err = testStore.GetSkillsByLearner(context.Background(), GetSkillsByLearnerParams{
		Learner: "done",
		Limit:   5,
		Offset:  0,
	})

	require.NoError(t, err)
	require.Len(t, skills, 0)

	for i := 0; i < 10; i++ {
		createSkillByLearner(t, "random")
	}

	skills, err = testStore.GetSkillsByLearner(context.Background(), GetSkillsByLearnerParams{
		Learner: "random",
		Limit:   8,
		Offset:  0,
	})
	require.NoError(t, err)
	require.Len(t, skills, 8)

	for _, skill := range skills {
		require.NotEmpty(t, skill)
	}
}
