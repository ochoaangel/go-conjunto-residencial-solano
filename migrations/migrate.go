package migrations

import "crs/initializers"

func init() {
	initializers.LoadEnvVariables()
	initializers.InitDB()
}
