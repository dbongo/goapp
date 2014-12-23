(function () {
  angular.module('app')
    .config(function ($stateProvider) {
      $stateProvider
        .state('home', {
          url: '/',
          templateUrl: 'home/home.html'
        })
    })
})();
