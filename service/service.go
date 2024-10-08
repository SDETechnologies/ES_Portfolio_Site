package service

import (
    "context"
    "main/database"
)

type IService interface {
    GetAllJobs(ctx context.Context) ([]database.Job, error)
    GetExperienceItemsByID(ctx context.Context, jobID int) ([]database.ExperienceItem, error)
    GetAllSchools(ctx context.Context) ([]*database.School, error)
    GetSchoolClassesFromSchoolID(ctx context.Context, schoolID int) ([]database.Class, error)
    GetAllProjects(ctx context.Context, typeID int) ([]*database.Project, error)
    GetProjectImagesFromID(ctx context.Context, projectID int) ([]*database.ProjectImage, error)
    GetProjectTypeFromID(ctx context.Context, typeID int) (*database.ProjectType, error)
}

type Service struct {
    Database database.IDatabase
}

func NewService(database database.IDatabase) *Service {
    return &Service{
        Database: database,
    }
}
