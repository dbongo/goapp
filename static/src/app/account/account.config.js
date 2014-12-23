(function () {
  angular.module('app')
    .config(function ($stateProvider) {
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
    })
})();
