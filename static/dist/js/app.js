angular.module('app', ['ngResource', 'ui.router', 'ui.bootstrap'])

angular.module('app')
.config(['$httpProvider', '$stateProvider', '$urlRouterProvider', '$locationProvider', function($httpProvider, $stateProvider, $urlRouterProvider, $locationProvider) {
    $urlRouterProvider.otherwise('/')
    $stateProvider
    .state('home', {
        url: '/',
        templateUrl: 'templates/home.tpl.html'
    })
    .state('login', {
        url: '/login',
        templateUrl: 'templates/login.tpl.html',
        controller: 'LoginCtrl',
        controllerAs: 'vm'
    })
    .state('register', {
        url: '/register',
        templateUrl: 'templates/register.tpl.html',
        controller: 'RegisterCtrl',
        controllerAs: 'vm'
    })
    .state('posts', {
        url: '/posts',
        templateUrl: 'templates/posts.tpl.html',
        controller: 'PostsCtrl',
        controllerAs: 'vm',
        authenticate: true
    })
    $locationProvider.html5Mode(true)
    $httpProvider.interceptors.push('AuthInterceptor')
}])
.run(['$rootScope', '$location', '$state', '$window', 'Auth', function($rootScope, $location, $state, $window, Auth) {
    $rootScope.$on('$stateChangeStart', function(event, next) {
        if (next.authenticate) {
            $state.go('login')
        }
    })
    $rootScope.$on('Auth:Required', function() {
        Auth.logout()
        $state.go('login')
    })
    $rootScope.$on('Auth:Forbidden', function() {
        Auth.logout()
        $state.go('login')
    })
}])

angular.module('app')
.constant('API_URL', 'http://localhost:3000')

function AppCtrl(Auth) {
	var vm = this

	vm.getCurrentUser = Auth.getCurrentUser
	//vm.isLoggedIn = Auth.isLoggedIn
	//vm.isAdmin = Auth.isAdmin
	vm.logout = Auth.logout
}
AppCtrl.$inject = ['Auth'];

angular.module('app')
.controller('AppCtrl', AppCtrl)

function LoginCtrl($state, Auth) {
	var vm = this

	vm.alerts = []
	vm.user = {email: "", password: ""}
	vm.login = login
	vm.closeAlert = closeAlert

	function login(form) {
		if (form.$valid) {
			Auth.login(vm.user).then(function() {
				vm.alerts = []
				$state.go('posts')
			}).catch(function(err) {
				vm.alerts.push({
					type: "danger",
					msg: err.message
				})
			})
		}
	}

	function closeAlert(index) {
		vm.alerts.splice(index, 1)
	}
}
LoginCtrl.$inject = ['$state', 'Auth'];

angular.module('app')
.controller('LoginCtrl', LoginCtrl)

function PostsCtrl(Posts, Auth) {
	var vm = this

	vm.posts = []
	vm.addPost = addPost

	Posts.fetch().then(function(posts) {
		vm.posts = posts
	})

	function addPost() {
		if (vm.postBody) {
			Posts.create({
				username: Auth.getCurrentUser().username,
				body: vm.postBody
			}).then(function(res) {
				vm.postBody = null
				vm.posts.push(res.data)
			})
		}
	}
}
PostsCtrl.$inject = ['Posts', 'Auth'];

angular.module('app')
.controller('PostsCtrl', PostsCtrl)

function RegisterCtrl($state, Auth) {
	var vm = this

	vm.alerts = []
	vm.user = {username: "", password: "", password2: ""}
	vm.register = register
	vm.closeAlert = closeAlert

	function register(form) {
		if (form.$valid) {
			Auth.createUser({
				username: vm.user.username,
				password: vm.user.password2
			}).then(function() {
				vm.alerts = []
				$state.go('posts')
			}).catch(function() {
				vm.alerts.push({
					type: "danger",
					msg: "The specified username is already in use"
				})
			})
		}
	}

	function closeAlert(index) {
		vm.alerts.splice(index, 1)
	}
}
RegisterCtrl.$inject = ['$state', 'Auth'];

angular.module('app')
.controller('RegisterCtrl', RegisterCtrl)

function match() {
	return {
		require: 'ngModel',
		restrict: 'A',
		scope: {
			match: '='
		},
		link: function(scope, elem, attrs, ctrl) {
			scope.$watch(function() {
				return (ctrl.$pristine && angular.isUndefined(ctrl.$modelValue)) || scope.match === ctrl.$modelValue
			}, function(currentValue) {
				ctrl.$setValidity('match', currentValue)
			})
		}
	}
}

angular.module('app')
.directive('match', match)

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
Auth.$inject = ['$http', 'TokenFactory', 'API_URL'];

angular.module('app')
.service('Auth', Auth)

function AuthInterceptor($q, $rootScope, TokenFactory) {
    return {
        request: function(config) {
            config.headers = config.headers || {}
            var token = TokenFactory.get()
            if (token) {
                config.headers.Authorization = 'Bearer ' + token
            }
            return config
        },
        responseError: function(rejection) {
            if ((rejection.status === 401) || (rejection.status === 403)) {
                $rootScope.$broadcast('Auth:Required')
            } else if (rejection.status === 419) {
                $rootScope.$broadcast('Auth:Forbidden')
            }
            return $q.reject(rejection)
        }
    }
}
AuthInterceptor.$inject = ['$q', '$rootScope', 'TokenFactory'];

angular.module('app')
.factory('AuthInterceptor', AuthInterceptor)

function Posts($http, API_URL) {
	var service = {
		fetch: fetch,
		create: create
	}
	return service

	function fetch() {
		return $http.get(API_URL + '/api/posts').then(function(res) {
			return res.data
		})
	}

	function create(post) {
		return $http.post(API_URL + '/api/posts', post)
	}
}
Posts.$inject = ['$http', 'API_URL'];


angular.module('app')
.service('Posts', Posts)

function TokenFactory($window) {
	var store = $window.localStorage
	var key = 'access_token'

	return {
		get: getToken,
		set: setToken
	}

	function getToken() {
		return store.getItem(key)
	}

	function setToken(token) {
		if (token) {
            store.setItem(key, token)
        } else {
            store.removeItem(key)
        }
	}
}
TokenFactory.$inject = ['$window'];

angular.module('app')
.factory('TokenFactory', TokenFactory)

// angular.module('app').factory('User', User)
// function User($resource, API_URL) {
//      return $resource(API_URL + '/api/users/:id/:controller', {id: '@_id'}, {
// 			get: {
// 				method: 'GET',
// 				params: {
// 					id: 'me'
// 				}
// 			}
// 		})
// 	})
