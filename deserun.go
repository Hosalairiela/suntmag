var daysOfWeek = [...]string{
    "Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday",
}

func parseWeekday(v string) (time.Weekday, error) {
    for i := range daysOfWeek {
        if daysOfWeek[i] == v {
            return time.Weekday(i), nil
        }
    }
    return time.Sunday, fmt.Errorf("invalid weekday '%s'", v)
}
