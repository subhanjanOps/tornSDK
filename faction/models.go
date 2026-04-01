package faction

type Basic struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Tag        string `json:"tag"`
	TagImage   string `json:"tag_image"`
	LeaderID   int    `json:"leader_id"`
	CoLeaderID int    `json:"co_leader_id"`
	Respect    int    `json:"respect"`
	DaysOld    int    `json:"days_old"`
	Capacity   int    `json:"capacity"`
	Members    int    `json:"members"`
	IsEnlisted *bool  `json:"is_enlisted"`
	Rank       Rank   `json:"rank"`
	BestChain  int    `json:"best_chain"`
	Note       string `json:"note"`
}

type Rank struct {
	Level    int    `json:"level"`
	Name     string `json:"name"`
	Division int    `json:"division"`
	Position int    `json:"position"`
	Wins     int    `json:"wins"`
}

type basicResponse struct {
	Basic Basic `json:"basic"`
}
