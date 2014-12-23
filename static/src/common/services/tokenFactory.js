(function() {
  angular.module('app')
    .factory('TokenFactory', function TokenFactory($window) {
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
    })
})();
