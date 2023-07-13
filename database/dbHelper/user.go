package dbHelper

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"math"
	"rms/models"
)

func AddRestaurant(db *sqlx.DB, restaurant models.Restaurants, userId int) error {
	SQL := `INSERT INTO restaurants
    		(restaurant_name,longitude,latitude,address,opening_time,closing_time,created_by) 
			VALUES ($1,$2,$3,$4,$5,$6,$7)`
	_, err := db.Exec(SQL, restaurant.Name, restaurant.Longitude, restaurant.Latitude, restaurant.Address, restaurant.OpeningTime, restaurant.ClosingTime, userId)
	if err != nil {
		return err
	}
	return nil
}
func AddDish(db *sqlx.DB, dishes models.Dishes, userId int) error {
	SQL := `insert into dishes
    		(dish_name, dish_type, restaurant_id, created_by)
			values ($1,$2,$3,$4)`
	_, err := db.Exec(SQL, dishes.Name, dishes.Type, dishes.RestaurantId, userId)
	if err != nil {
		return err
	}
	return nil
}
func AllRestaurants(db sqlx.Ext) ([]models.Restaurants, error) {
	var allRestaurants []models.Restaurants
	SQL := `SELECT id,restaurant_name,longitude,latitude,opening_time,closing_time 
			FROM restaurants`
	rows, err := db.Query(SQL)
	if err != nil {
		return allRestaurants, err
	}
	for rows.Next() {
		var Restaurant models.Restaurants
		err = rows.Scan(&Restaurant.Id, &Restaurant.Name, &Restaurant.Longitude, &Restaurant.Latitude, &Restaurant.OpeningTime, &Restaurant.ClosingTime)
		if err != nil {

			return allRestaurants, err
		}
		allRestaurants = append(allRestaurants, Restaurant)
	}
	return allRestaurants, nil
}
func AllSubAdmin(db *sqlx.DB) ([]models.Users, error) {
	var allSudAdmin []models.Users
	SQL := `SELECT u.user_name, u.user_email, u.created_by
			FROM users u
			JOIN user_role r ON u.id = r.user_id
			WHERE r.user_type = 'SubAdmin'`
	rows, err := db.Query(SQL)
	if err != nil {
		return allSudAdmin, err
	}
	for rows.Next() {
		var SubAdmin models.Users
		err = rows.Scan(&SubAdmin.Name, &SubAdmin.Email, &SubAdmin.CreatedBy)
		if err != nil {

			return allSudAdmin, err
		}
		allSudAdmin = append(allSudAdmin, SubAdmin)
	}
	return allSudAdmin, nil
}
func AllUsers(db *sqlx.DB) ([]models.AllInfo, error) {
	var userInfo []models.AllInfo
	SQL := `SELECT r.restaurant_name, r.address, r.opening_time, r.closing_time,
 	        d.dish_name, d.dish_type,
    	    u.user_name, u.user_email
			from users u
			JOIN dishes d ON d.created_by = u.id
			JOIN restaurants r ON u.id = r.created_by;`
	//todo: use select or get
	rows, err := db.Query(SQL)
	if err != nil {
		return userInfo, err
	}
	for rows.Next() {
		var info models.AllInfo
		err = rows.Scan(&info.RestaurantName, &info.Address, &info.OpeningTime, &info.ClosingTime, &info.DishName, &info.Type, &info.UserName, &info.UserEmail)
		if err != nil {
			return userInfo, err
		}
		userInfo = append(userInfo, info)
	}
	return userInfo, nil
}
func AllCustomers(db *sqlx.DB, userId int) ([]models.AllInfo, error) {
	var customerInfo []models.AllInfo
	SQL := `SELECT r.restaurant_name, r.address, r.opening_time, r.closing_time,
       		u.user_name, u.user_email
			from restaurants r
			JOIN users u ON r.created_by = u.created_by AND u.id = r.created_by
			where r.created_by = $1;`
	rows, err := db.Query(SQL, userId)
	if err != nil {
		return customerInfo, err
	}
	for rows.Next() {
		var info models.AllInfo
		err = rows.Scan(&info.RestaurantName, &info.Address, &info.OpeningTime, &info.ClosingTime, &info.DishName, &info.Type, &info.StartServe, &info.EndServe, &info.UserName, &info.UserEmail)
		if err != nil {
			return customerInfo, err
		}
		customerInfo = append(customerInfo, info)
	}
	return customerInfo, nil
}
func RestaurantsByDish(db *sqlx.DB) ([]models.RestaurantByDishes, error) {
	var restaurantsByDishes []models.RestaurantByDishes
	SQL := `SELECT r.restaurant_name, r.address, r.opening_time, r.closing_time,
       		d.dish_name, d.dish_type
			from restaurants r
			join dishes d ON r.id = d.restaurant_id;
`
	rows, err := db.Query(SQL)
	if err != nil {
		return restaurantsByDishes, err
	}
	for rows.Next() {
		var info models.RestaurantByDishes
		err = rows.Scan(&info.RestaurantName, &info.Address, &info.OpeningTime, &info.ClosingTime, &info.DishName, &info.Type)
		if err != nil {
			return restaurantsByDishes, err
		}
		restaurantsByDishes = append(restaurantsByDishes, info)
	}
	return restaurantsByDishes, nil
}
func ValidateUser(db *sqlx.DB, userId int) (bool, error) {
	SQL := `select user_type from user_role where user_id=$1 and user_type='Admin' OR user_type='SubAdmin'`
	var userType string
	err := db.Get(&userType, SQL, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
func Distance(db *sqlx.DB, userId int, name string) (float64, error) {
	var distance float64
	var lon1 float64
	var lon2 float64
	var lat1 float64
	var lat2 float64
	SQL := `select r.longitude, r.latitude, a.longitude, a.latitude
			from restaurants r
			join address a ON r.restaurant_name = $1 AND a.id = $2;
`
	err := db.QueryRowx(SQL, name, userId).Scan(&lon1, &lat1, &lon2, &lat2)
	if err != nil {
		if err == sql.ErrNoRows {
			return distance, nil
		}
		return distance, err
	}
	lat1Rad := lat1 * (math.Pi / 180)
	lon1Rad := lon1 * (math.Pi / 180)
	lat2Rad := lat2 * (math.Pi / 180)
	lon2Rad := lon2 * (math.Pi / 180)

	// Haversine formula
	dLat := lat2Rad - lat1Rad
	dLon := lon2Rad - lon1Rad
	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*
			math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	distance = 6371 * c
	return distance, nil

}
