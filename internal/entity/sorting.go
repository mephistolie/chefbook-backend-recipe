package entity

const (
	SortingCreationTimestamp = "creation_timestamp"
	SortingUpdateTimestamp   = "update_timestamp"
	SortingRating            = "rating"
	SortingVotes             = "votes"
	SortingTime              = "cooking_time"
	SortingCalories          = "calories"
)

var AvailableSortings = []string{
	SortingCreationTimestamp,
	SortingUpdateTimestamp,
	SortingRating,
	SortingVotes,
	SortingTime,
	SortingCalories,
}
