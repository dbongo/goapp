angular.module('hackapp').factory('AuthInterceptor', function AuthInterceptor($q, $rootScope, TokenFactory) {

  function request(config) {
    config.headers = config.headers || {}
    var token = TokenFactory.get()
    if (token) {
      config.headers.Authorization = 'Bearer ' + token
    }
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
})
