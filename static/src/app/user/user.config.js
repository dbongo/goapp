(function () {
  angular.module('app')
    .config(function ($stateProvider) {
      $stateProvider
        .state('posts', {
          url: '/posts',
          templateUrl: 'user/posts/posts.html',
          controller: 'PostsCtrl',
          controllerAs: 'vm',
          authenticate: true
        })
    })
})();
