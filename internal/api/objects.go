package api

type Exercise struct {
	ExerciseID  string `json:"exerciseId"`
	Name        string `json:"name"`
	Repetitions int    `json:"repetitions"`
	Sets        int    `json:"sets"`
	//Loading Loading TODO: Define how weights can be loaded (e.g. left hand, right hand, waist, etc?)
}
