package app

func StartApplication() {
	urlMappers()
	//todo log

	router.Run(":8181")

}
