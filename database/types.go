package database

type Job struct {
    ID int
    Title string
    CompanyName string
    Location string
    StartDate string
    EndDate string
    ExperienceItems []ExperienceItem
}

type ExperienceItem struct {
    ID int
    JobID int
    Text string
}

type School struct {
    ID int
    Name string
    Major *string
    Graduated int
    Degree *string
    SchoolName *string
    SchoolTypeID int
    GPA float64
    Location string
    TranscriptPath string
    LogoPath string
    Classes []Class
}

type Class struct {
    ID int
    SchoolID int
    Abbreviation string
    Name string
    Grade string
    Term string
    Credits int
    ClassFilesURL string
}

type Project struct {
    ID int
    Name string
    URL *string
    GithubRepoURL *string
    Description *string
    ProjectImages []*ProjectImage
    ProjectType ProjectType
    Disabled int
    ListOrder int
}

type ProjectType struct {
    ID int
    Name string
    Description *string
}

type ProjectImage struct {
    ID int
    ProjectID int
    ImageLink string
    Caption string
    AltLink string
    ImageIsVideo int
}
