package entity

const (
	VisibilityPrivate = "private"
	VisibilityLink    = "link"
	VisibilityPublic  = "public"
)

var AvailableVisibilities = []string{
	VisibilityPrivate,
	VisibilityLink,
	VisibilityPublic,
}
