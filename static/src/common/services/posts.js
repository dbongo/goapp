(function() {
  angular.module('app')
    .service('Posts', function Posts($http, API_URL) {
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
    })
})();
