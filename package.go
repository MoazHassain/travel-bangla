package main

// Package struct to hold package data
type Package struct {
	ID        int
	Title     string
	Location  string
	Days      int
	Persons   int
	Price     string
	Thumbnail string
	VideoURL  string
}

// Function to fetch packages from the database
func fetchPackages() ([]Package, error) {
	rows, err := db.Query("SELECT pk_id, pk_title, pk_location, pk_days, pk_persons, pk_price, pk_thumbnail FROM package")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var packages []Package

	for rows.Next() {
		var pkg Package
		err := rows.Scan(&pkg.ID, &pkg.Title, &pkg.Location, &pkg.Days, &pkg.Persons, &pkg.Price, &pkg.Thumbnail)
		if err != nil {
			return nil, err
		}
		packages = append(packages, pkg)
	}

	return packages, rows.Err()
}
