'use strict';

angular.module('app', ['ngResource', 'ui.router', 'ui.bootstrap'])

angular.module('app').config(['$httpProvider', '$stateProvider', '$urlRouterProvider', '$locationProvider', function($httpProvider, $stateProvider, $urlRouterProvider, $locationProvider) {

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

}]).run(['$rootScope', '$location', '$state', '$window', 'Auth', function($rootScope, $location, $state, $window, Auth) {

    $rootScope.$on('$stateChangeStart', function(event, next) {
        if (next.authenticate) $state.go('login')
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

angular.module('app').constant('API_URL', 'http://localhost:3000/api')

angular.module('app').controller('AppCtrl', ['Auth', function AppCtrl(Auth) {
	var vm = this

	vm.getCurrentUser = Auth.getCurrentUser
	vm.logout = Auth.logout
}])

angular.module('app').controller('LoginCtrl', ['$state', 'Auth', function LoginCtrl($state, Auth) {
	var vm = this

	vm.alerts = []
	vm.user = {}
	vm.login = login
	vm.closeAlert = closeAlert

	function login(form) {
		if (form.$valid) {
			Auth.login({email: vm.user.email, password: vm.user.password})
			.then(function() {
				vm.alerts = []
				$state.go('posts')
			})
			.catch(function(err) {
				vm.alerts.push({type: "danger", msg: err.message})
				Auth.logout()
			})
		}
	}

	function closeAlert(index) {
		vm.alerts.splice(index, 1)
	}
}])

angular.module('app').controller('PostsCtrl', ['Posts', 'Auth', function PostsCtrl(Posts, Auth) {
	var vm = this

	vm.posts = []
	vm.addPost = addPost

	Posts.fetch().then(function(posts) {
		vm.posts = posts
	})

	function addPost() {
		if (vm.postBody) {
			Posts.create({username: Auth.getCurrentUser().username, body: vm.postBody})
			.then(function(res) {
				vm.postBody = null
				vm.posts.push(res.data)
			})
		}
	}
}])

angular.module('app').controller('RegisterCtrl', ['$state', 'Auth', function RegisterCtrl($state, Auth) {
	var vm = this

	vm.alerts = []
	vm.user = {}
	vm.register = register
	vm.closeAlert = closeAlert

	function register(form) {
		if (form.$valid) {
			Auth.register({email: vm.user.email, username: vm.user.username, password: vm.user.password2})
			.then(function() {
				vm.alerts = []
				$state.go('posts')
			})
			.catch(function(err) {
				vm.alerts.push({type: "danger", msg: err.message})
				Auth.logout()
			})
		}
	}

	function closeAlert(index) {
		vm.alerts.splice(index, 1)
	}
}])

angular.module('app').directive('match', function match() {
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
})

angular.module('app').service('Auth', ['$http', 'TokenFactory', 'API_URL', function Auth($http, TokenFactory, API_URL) {
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
}])

angular.module('app').factory('AuthInterceptor', ['$q', '$rootScope', 'TokenFactory', function AuthInterceptor($q, $rootScope, TokenFactory) {

    function request(config) {
        config.headers = config.headers || {}
        var token = TokenFactory.get()
        if (token) config.headers.Authorization = 'Bearer ' + token
        return config
    }

    function responseError(rejection) {
        if ((rejection.status === 401) || (rejection.status === 403)) {
            $rootScope.$broadcast('Auth:Required')
        } else if (rejection.status === 419) {
            $rootScope.$broadcast('Auth:Forbidden')
        }
        return $q.reject(rejection)
    }

    return {
        request: request,
        responseError: responseError
    }
}])

angular.module('app').service('Posts', ['$http', 'API_URL', function Posts($http, API_URL) {

	function fetch() {
		return $http.get(API_URL + '/api/posts').then(function(res) {
			return res.data
		})
	}

	function create(post) {
		return $http.post(API_URL + '/api/posts', post)
	}

	var service = {
		fetch: fetch,
		create: create
	}
	return service
}])

angular.module('app').factory('TokenFactory', ['$window', function TokenFactory($window) {
	var store = $window.localStorage
	var key = 'access_token'

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

	return {
		get: getToken,
		set: setToken
	}
}])

// angular.module('app').factory('User', function User($resource, API_URL) {
//     return $resource(API_URL + '/api/users/:id/:controller', {id: '@_id'}, {
// 		get: {
// 			method: 'GET',
// 			params: {
// 				id: 'me'
// 			}
// 		}
// 	})
// })