package main

// ServerRoute : 設定Route
func ServerRoute() map[string]func(*App) {
	route := make(map[string]func(*App))

	/**
	  |----------------
	  | 設定路由清單
	  |----------------
	*/

	// ====== BEGIN ======
	route["/"] = indexHandler
	route["/chat"] = socketHandler
	route["/login"] = loginHandler
	route["/price"] = priceHandler
	route["/sql"] = sqlHandler
	// ======  END  ======

	return route
}
