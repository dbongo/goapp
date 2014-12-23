(function () {
  // this module is the entry point into our angular app
  angular.module('app', ['ngResource','ui.router','ui.bootstrap'])
    .config(['$httpProvider', '$urlRouterProvider', '$locationProvider', function ($httpProvider, $urlRouterProvider, $locationProvider) {
      $urlRouterProvider.otherwise('/')
      $locationProvider.html5Mode(true)
      $httpProvider.interceptors.push('AuthInterceptor')
    }])
    .factory('AuthInterceptor', ['$q', '$rootScope', 'TokenFactory', function AuthInterceptor($q, $rootScope, TokenFactory) {
      return {
        request: function(config) {
          config.headers = config.headers || {}
          var token = TokenFactory.get()
          if (token) config.headers.Authorization = 'Bearer ' + token
          return config
        },
        responseError: function(rejection) {
          if ((rejection.status === 401) || (rejection.status === 403))
            $rootScope.$broadcast('Auth:Required')
          else if (rejection.status === 419)
            $rootScope.$broadcast('Auth:Forbidden')
          return $q.reject(rejection)
        }
      }
    }])
    .run(['$rootScope', '$location', '$state', '$window', 'Auth', function($rootScope, $location, $state, $window, Auth) {
      $rootScope.$on('$stateChangeStart', function(event, next) {
        Auth.isLoggedInAsync(function(loggedIn) {
          if (next.authenticate && !loggedIn) $state.go('login')
        })
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
})();

(function () {
  angular.module('app')
    .controller('AppCtrl', ['Auth', function AppCtrl(Auth) {
      var vm = this
      vm.getCurrentUser = Auth.getCurrentUser
      vm.isLoggedIn = Auth.isLoggedIn
      vm.isAdmin = Auth.isAdmin
      vm.logout = Auth.logout
    }])
})();

(function () {
  angular.module('app')
    .config(['$stateProvider', function ($stateProvider) {
      $stateProvider
        .state('login', {
          url: '/login',
          templateUrl: 'account/login/login.html',
          controller: 'LoginCtrl',
          controllerAs: 'vm'
        })
        .state('register', {
          url: '/register',
          templateUrl: 'account/register/register.html',
          controller: 'RegisterCtrl',
          controllerAs: 'vm'
        })
    }])
})();

(function () {
  angular.module('app')
    .controller('LoginCtrl', ['$state', 'Auth', function LoginCtrl($state, Auth) {
      var vm = this
      vm.alerts = []
      vm.user = {username: "", password: ""}
      vm.login = login
      vm.closeAlert = closeAlert

      function login(form) {
        if(form.$valid) {
          Auth.login(vm.user).then(function() {
            vm.alerts = []
            $state.go('posts')
          }).catch(function(err) {
            vm.alerts.push({type: "danger", msg: err.message})
          })
        }
      }

      function closeAlert(index) {
		    vm.alerts.splice(index, 1)
	    }
    }])
})();

(function() {
  angular.module('app')
    .controller('RegisterCtrl', ['$state', 'Auth', function RegisterCtrl($state, Auth) {
      var vm = this
      vm.alerts = []
      vm.user = {username: "", password: "", password2: ""}
      vm.register = register
      vm.closeAlert = closeAlert

      function register(form) {
        if(form.$valid) {
          Auth.createUser({username: vm.user.username, password: vm.user.password2}).then(function() {
            vm.alerts = []
            $state.go('posts')
          }).catch(function() {
            vm.alerts.push({type: "danger", msg: "The specified username is already in use"})
          })
        }
      }

      function closeAlert(index) {
        vm.alerts.splice(index, 1)
      }
    }])
})();

(function () {
  angular.module('app')
    .config(['$stateProvider', function ($stateProvider) {
      $stateProvider
        .state('home', {
          url: '/',
          templateUrl: 'home/home.html'
        })
    }])
})();

(function () {
  angular.module('app')
    .config(['$stateProvider', function ($stateProvider) {
      $stateProvider
        .state('posts', {
          url: '/posts',
          templateUrl: 'user/posts/posts.html',
          controller: 'PostsCtrl',
          controllerAs: 'vm',
          authenticate: true
        })
    }])
})();

(function() {
	angular.module('app')
		.controller('PostsCtrl', ['Posts', 'Auth', function PostsCtrl(Posts, Auth) {
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
		}])
})();

(function () {
  angular.module('app')
    .constant('API_URL', 'http://localhost:3000')
})();

(function() {
  angular.module('app')
    .directive('match', function match() {
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
})();

(function() {
	angular.module('app')
		.service('Auth', ['$http', 'TokenFactory', 'API_URL', 'User', '$q', function Auth($http, TokenFactory, API_URL, User, $q) {
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
		}])
})();

(function() {
  angular.module('app')
    .service('Posts', ['$http', 'API_URL', function Posts($http, API_URL) {
      var service = {
        fetch: fetch,
        create: create
      }
      return service

      function fetch() {
        return $http.get(API_URL + '/api/posts').then(function (res) {
          return res.data
        })
      }

      function create(post) {
        return $http.post(API_URL + '/api/posts', post)
      }
    }])
})();

(function() {
  angular.module('app')
    .factory('TokenFactory', ['$window', function TokenFactory($window) {
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
        if (token) store.setItem(key, token)
        else store.removeItem(key)
      }
    }])
})();

(function() {
	angular.module('app')
		.factory('User', ['$resource', 'API_URL', function User($resource, API_URL) {
			return $resource(API_URL + '/api/users/:id/:controller', {id: '@_id'}, {
				get: {
					method: 'GET',
					params: {
						id: 'me'
					}
				}
			})
		}])
})();

(function(module) {
try {
  module = angular.module('app');
} catch (e) {
  module = angular.module('app', []);
}
module.run(['$templateCache', function($templateCache) {
  $templateCache.put('home/home.tpl.html',
    '<div class="container">\n' +
    '  <div class="jumbotron">\n' +
    '    <h1>Evant Admin Portal</h1>\n' +
    '  </div>\n' +
    '</div>\n' +
    '');
}]);
})();

(function(module) {
try {
  module = angular.module('app');
} catch (e) {
  module = angular.module('app', []);
}
module.run(['$templateCache', function($templateCache) {
  $templateCache.put('account/login/login.tpl.html',
    '<div class="container">\n' +
    '  <alert ng-repeat="alert in vm.alerts" type="{{alert.type}}" close="vm.closeAlert($index)">{{alert.msg}}</alert>\n' +
    '  <form name="form" role="form" class="form-horizontal" ng-submit="vm.login(form)">\n' +
    '    <fieldset>\n' +
    '      <legend>Sign in</legend>\n' +
    '      <div class="form-group" ng-class="form.username.$invalid?\'has-error\':\'has-success\'">\n' +
    '        <label for="username" class="col-sm-2 control-label">Username</label>\n' +
    '        <div class="col-sm-10">\n' +
    '          <div class="input-group">\n' +
    '            <span class="input-group-addon"><i class="fa fa-envelope"></i></span>\n' +
    '            <input id="username" name="username" type="text" class="form-control" placeholder="Username" ng-model="vm.user.username" required>\n' +
    '          </div>\n' +
    '          <span class="help-block" ng-show="form.username.$error.username">Please enter a valid username</span>\n' +
    '        </div>\n' +
    '      </div>\n' +
    '      <div class="form-group" ng-class="form.password.$invalid?\'has-error\':\'has-success\'">\n' +
    '        <label for="password" class="col-sm-2 control-label">Password</label>\n' +
    '        <div class="col-sm-10">\n' +
    '          <div class="input-group">\n' +
    '            <span class="input-group-addon"><i class="fa fa-key"></i></span>\n' +
    '            <input id="password" name="password" type="password" class="form-control" placeholder="Password" ng-model="vm.user.password" required>\n' +
    '          </div>\n' +
    '        </div>\n' +
    '      </div>\n' +
    '      <div class="form-group">\n' +
    '        <div class="col-sm-offset-2 col-sm-10">\n' +
    '          <button type="submit" class="btn btn-success">Login</button>\n' +
    '          <a class="btn btn-primary" ui-sref="register">Sign up</a>\n' +
    '          <a class="btn btn-link" href="#">Forgot your password?</a>\n' +
    '        </div>\n' +
    '      </div>\n' +
    '    </fieldset>\n' +
    '  </form>\n' +
    '</div>\n' +
    '');
}]);
})();

(function(module) {
try {
  module = angular.module('app');
} catch (e) {
  module = angular.module('app', []);
}
module.run(['$templateCache', function($templateCache) {
  $templateCache.put('account/register/register.tpl.html',
    '<div class="container">\n' +
    '  <alert ng-repeat="alert in vm.alerts" type="{{alert.type}}" close="vm.closeAlert($index)">{{alert.msg}}</alert>\n' +
    '  <form name="form" role="form" class="form-horizontal" ng-submit="vm.register(form)">\n' +
    '    <fieldset>\n' +
    '      <legend>Sign up</legend>\n' +
    '      <div class="form-group" ng-class="form.username.$invalid?\'has-error\':\'has-success\'">\n' +
    '        <label for="username" class="col-sm-2 control-label">Username</label>\n' +
    '        <div class="col-sm-10">\n' +
    '          <div class="input-group">\n' +
    '            <span class="input-group-addon"><i class="fa fa-envelope"></i></span>\n' +
    '            <input id="username" name="username" type="text" class="form-control" placeholder="Username" ng-model="vm.user.username" required>\n' +
    '          </div>\n' +
    '          <span class="help-block" ng-show="form.username.$error.username">Please enter a valid username</span>\n' +
    '        </div>\n' +
    '      </div>\n' +
    '      <div class="form-group" ng-class="form.password.$invalid?\'has-error\':\'has-success\'">\n' +
    '        <label for="password" class="col-sm-2 control-label">Password</label>\n' +
    '        <div class="col-sm-10">\n' +
    '          <div class="input-group">\n' +
    '            <span class="input-group-addon"><i class="fa fa-key"></i></span>\n' +
    '            <input id="password" name="password" type="password" class="form-control" placeholder="Password" ng-model="vm.user.password" required>\n' +
    '          </div>\n' +
    '        </div>\n' +
    '      </div>\n' +
    '      <div ng-show=\'form.password.$valid\' class="form-group" ng-class="form.password2.$error.match?\'has-error\':\'has-success\'">\n' +
    '        <label for="password2" class="col-sm-2 control-label">Repeat Password</label>\n' +
    '        <div class="col-sm-10">\n' +
    '          <input id="password2" name="password2" type="password" class="form-control" placeholder="Repeat Password" ng-model="vm.user.password2" match="vm.user.password">\n' +
    '          <span class="help-block" ng-show="form.password2.$error.match">Both passwords must be the same</span>\n' +
    '        </div>\n' +
    '      </div>\n' +
    '      <div class="form-group">\n' +
    '        <div class="col-sm-offset-2 col-sm-10">\n' +
    '          <button type="submit" class="btn btn-success">Create</button>\n' +
    '          <a class="btn btn-link" ui-sref="login">Already have an account?</a>\n' +
    '        </div>\n' +
    '      </div>\n' +
    '    </fieldset>\n' +
    '  </form>\n' +
    '</div>\n' +
    '');
}]);
})();

(function(module) {
try {
  module = angular.module('app');
} catch (e) {
  module = angular.module('app', []);
}
module.run(['$templateCache', function($templateCache) {
  $templateCache.put('user/posts/posts.tpl.html',
    '<div class=\'container\'>\n' +
    '  <h1>Recent Posts</h1>\n' +
    '  <form role=\'form\'>\n' +
    '    <div class=\'form-group\'>\n' +
    '      <div class=\'input-group\'>\n' +
    '        <input ng-model=\'vm.postBody\' class=\'form-control\'>\n' +
    '        <span class=\'input-group-btn\'>\n' +
    '          <button ng-click=\'vm.addPost()\' class=\'btn btn-default\'>Add Post</button>\n' +
    '        </span>\n' +
    '      </div>\n' +
    '    </div>\n' +
    '  </form>\n' +
    '  <ul class=\'posts list-group\'>\n' +
    '    <li ng-repeat="post in vm.posts | orderBy:\'-created\'" class=\'list-group-item\'>\n' +
    '      <strong>@{{ post.username }}</strong>\n' +
    '      <span>{{ post.body }}</span>\n' +
    '    </li>\n' +
    '  </ul>\n' +
    '</div>\n' +
    '');
}]);
})();
