function AuthInterceptor($q, $rootScope, TokenFactory) {
    return {
        request: function(config) {
            config.headers = config.headers || {}
            var token = TokenFactory.get()
            if (token) {
                config.headers.Authorization = 'Bearer ' + token
            }
            return config
        },
        responseError: function(rejection) {
            if ((rejection.status === 401) || (rejection.status === 403)) {
                $rootScope.$broadcast('Auth:Required')
            } else if (rejection.status === 419) {
                $rootScope.$broadcast('Auth:Forbidden')
            }
            return $q.reject(rejection)
        }
    }
}

angular.module('app')
.factory('AuthInterceptor', AuthInterceptor)
