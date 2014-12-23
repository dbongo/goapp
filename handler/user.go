package handler

// // CurrentUser ...
// func CurrentUser(c web.C, w http.ResponseWriter, r *http.Request) {
// 	var user = ToUser(c)
// 	if user == nil {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}
// 	// return private data for the currently authenticated user,
// 	// specifically, their auth token.
// 	data := struct {
// 		*model.User
// 		Token string `json:"token"`
// 	}{user, user.Token}
// 	json.NewEncoder(w).Encode(&data)
// }
