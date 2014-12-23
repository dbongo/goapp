(function() {
	angular.module('app')
		.factory('User', function User($resource, API_URL) {
			return $resource(API_URL + '/api/users/:id/:controller', {id: '@_id'}, {
				get: {
					method: 'GET',
					params: {
						id: 'me'
					}
				}
			})
		})
})();
