(function() {
	angular.module('app')
		.service('Auth', function Auth($http, TokenFactory, API_URL, User, $q) {
			var currentUser = {}

			if (TokenFactory.get()) {
				currentUser = User.get()
			}

			var service = {
				login: login,
				createUser: createUser,
				logout: logout,
				isLoggedIn: isLoggedIn,
				isLoggedInAsync: isLoggedInAsync,
				isAdmin: isAdmin,
				getCurrentUser: getCurrentUser
			}
			return service

			function login(user, callback) {
				var cb = callback || angular.noop
				var deferred = $q.defer()
				$http.post(API_URL + '/auth/local', {username: user.username, password: user.password})
					.success(function(data) {
						TokenFactory.set(data.token)
						currentUser = User.get()
						deferred.resolve(data)
						return cb()
					})
					.error(function(err) {
						this.logout()
						deferred.reject(err)
						return cb(err)
					}.bind(this))
				return deferred.promise
			}

			function createUser(user, callback) {
				var cb = callback || angular.noop
				TokenFactory.set()
				return User.save(user, function(data) {
					TokenFactory.set(data.token)
					currentUser = User.get()
					return cb(user)
				}, function(err) {
					this.logout()
					return cb(err)
				}.bind(this)).$promise
			}

			function logout() {
				TokenFactory.set()
				currentUser = {}
			}

			function isLoggedIn() {
				return currentUser.hasOwnProperty('role')
			}

			function isLoggedInAsync(cb) {
				if (currentUser.hasOwnProperty('$promise')) {
					currentUser.$promise.then(function() {
						cb(true)
					}).catch(function() {
						cb(false)
					})
				} else if (isLoggedIn()) cb(true)
				else cb(false)
			}

			function isAdmin() {
				return currentUser.role === 'admin'
			}

			function getCurrentUser() {
				return currentUser
			}
		})
})();
