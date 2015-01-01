function Auth($http, TokenFactory, API_URL) {
	var currentUser = {}

	var service = {
		login: login,
		register: register,
		logout: logout,
		getCurrentUser: getCurrentUser
	}
	return service

	function login(u) {
		$http.post(API_URL + '/login', {
			email: u.email,
			password: u.password
		}).success(function(data) {
			currentUser = data
			TokenFactory.set(data.token)
			return currentUser
		}).error(function(err) {
			this.logout()
			return err
		})

	}

	function register(u) {
		$http.post(API_URL + '/register', {
			email: u.email,
			username: u.username,
			password: u.password
		}).success(function(data) {
			currentUser = data
			TokenFactory.set(data.token)
			return currentUser
		}).error(function(err) {
			this.logout()
			return err
		})
	}

	function logout() {
		TokenFactory.set()
		currentUser = {}
	}

	function getCurrentUser() {
		return currentUser
	}
}

angular.module('app')
.service('Auth', Auth)
