package user

type Bar struct {
	Current int `json:"current"`
	Maximum int `json:"maximum"`
}

type Chain struct {
	ID       int     `json:"id"`
	Current  int     `json:"current"`
	Max      int     `json:"max"`
	Timeout  int     `json:"timeout"`
	Modifier float64 `json:"modifier"`
	Cooldown int     `json:"cooldown"`
	Start    int     `json:"start"`
	End      int     `json:"end"`
}

type Bars struct {
	Energy Bar    `json:"energy"`
	Nerve  Bar    `json:"nerve"`
	Happy  Bar    `json:"happy"`
	Life   Bar    `json:"life"`
	Chain  *Chain `json:"chain"`
}

type Status struct {
	Description string `json:"description"`
	Details     string `json:"details"`
	State       string `json:"state"`
	Color       string `json:"color"`
	Until       *int   `json:"until"`
}

type LastAction struct {
	Status    string `json:"status"`
	Timestamp int    `json:"timestamp"`
	Relative  string `json:"relative"`
}

type Basic struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Level  int    `json:"level"`
	Gender string `json:"gender"`
	Status Status `json:"status"`
}

type Property struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Spouse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Status      string `json:"status"`
	DaysMarried int    `json:"days_married"`
}

type Profile struct {
	ID            int        `json:"id"`
	Name          string     `json:"name"`
	Level         int        `json:"level"`
	Rank          string     `json:"rank"`
	Title         string     `json:"title"`
	DonatorStatus *string    `json:"donator_status"`
	Age           int        `json:"age"`
	SignedUp      int        `json:"signed_up"`
	FactionID     *int       `json:"faction_id"`
	HonorID       int        `json:"honor_id"`
	Property      Property   `json:"property"`
	Image         *string    `json:"image"`
	Gender        string     `json:"gender"`
	Revivable     bool       `json:"revivable"`
	Role          string     `json:"role"`
	Status        Status     `json:"status"`
	Spouse        *Spouse    `json:"spouse"`
	Awards        int        `json:"awards"`
	Friends       int        `json:"friends"`
	Enemies       int        `json:"enemies"`
	ForumPosts    int        `json:"forum_posts"`
	Karma         int        `json:"karma"`
	LastAction    LastAction `json:"last_action"`
	Life          Bar        `json:"life"`
}

type BattleStatModifier struct {
	Effect string `json:"effect"`
	Value  int    `json:"value"`
	Type   string `json:"type"`
}

type BattleStat struct {
	Value     int64                `json:"value"`
	Modifier  int                  `json:"modifier"`
	Modifiers []BattleStatModifier `json:"modifiers"`
}

type BattleStats struct {
	Strength  BattleStat `json:"strength"`
	Defense   BattleStat `json:"defense"`
	Speed     BattleStat `json:"speed"`
	Dexterity BattleStat `json:"dexterity"`
	Total     int64      `json:"total"`
}

type barsResponse struct {
	Bars Bars `json:"bars"`
}

type basicResponse struct {
	Profile Basic `json:"profile"`
}

type profileResponse struct {
	Profile Profile `json:"profile"`
}

type battleStatsResponse struct {
	BattleStats BattleStats `json:"battlestats"`
}
