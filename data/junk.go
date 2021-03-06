package data

type Gizmo struct {
	Name        string
	ThreadCount int
}

// Get all threads in the database and returns it
func Gizmos() (gizmos []Gizmo, err error) {

	var sqlStr string = `
		select   u.name,
				 count(*) threadcount
		from     users u
		   inner join threads t
			  on t.user_id = u.id
		group by u.name
		order by threadcount desc,
			  	 u.name
	`

	rows, err := Db.Query(sqlStr)

	// rows, err := Db.Query("SELECT name, 1 threadcount FROM  users")

	if err != nil {
		return
	}

	for rows.Next() {

		gizmo := Gizmo{}

		if err = rows.Scan(&gizmo.Name, &gizmo.ThreadCount); err != nil {
			return
		}

		gizmos = append(gizmos, gizmo)
	}

	rows.Close()
	return
}
