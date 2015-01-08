angular.module('hackapp').config(function($httpProvider, $stateProvider, $urlRouterProvider, $locationProvider) {

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

}).run(function($rootScope, $location, $state, $window, Auth) {

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
})
