angular.module('hackapp').service('Posts', function Posts($http, API_URL) {

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
})
