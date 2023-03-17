package timeonly

type TimeOnlyRange [2]TimeOnly

// new TimeOnlyRange object
func NewTimeOnlyRange(firstT TimeOnly, secondT TimeOnly) TimeOnlyRange {
	return TimeOnlyRange{firstT, secondT}
}

func (t TimeOnlyRange) IsSecondGreaterThanFirst() bool {
	return t[1].Sub(t[0]) > 0
}
