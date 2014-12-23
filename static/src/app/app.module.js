(function () {
  // this module is the entry point into our angular app
  angular.module('app', ['ngResource','ui.router','ui.bootstrap'])
    .config(function ($httpProvider, $urlRouterProvider, $locationProvider) {
      $urlRouterProvider.otherwise('/')
      $locationProvider.html5Mode(true)
      $httpProvider.interceptors.push('AuthInterceptor')
    })
    .factory('AuthInterceptor', function AuthInterceptor($q, $rootScope, TokenFactory) {
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
    })
    .run(function($rootScope, $location, $state, $window, Auth) {
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
    })
})();
