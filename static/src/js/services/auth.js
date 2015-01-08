angular.module('app').service('Auth', function Auth($http, TokenFactory, API_URL) {
	var currentUser = {}

	function login(user) {
		return $http.post(API_URL + '/auth/login', user).then(function(res) {
			currentUser = res.data
			TokenFactory.set(currentUser.token)
			return res.data
		})
	}

	function register(user) {
		return $http.post(API_URL + '/auth/register', user).then(function(res) {
			currentUser = res.data
			TokenFactory.set(currentUser.token)
			return res.data
		})
	}

	function logout() {
		TokenFactory.set()
		currentUser = {}
	}

	function getCurrentUser() {
		return currentUser
	}

	var service = {
		login: login,
		register: register,
		logout: logout,
		getCurrentUser: getCurrentUser
	}
	return service
})
