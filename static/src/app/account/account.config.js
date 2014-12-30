(function() {
	angular.module('app')
		.config(function($stateProvider) {
			$stateProvider
			.state('login', {
				url: '/login',
				templateUrl: 'account/login/login.tpl.html',
				controller: 'LoginCtrl',
				controllerAs: 'vm',
				resolve: {
					template: function($templateCache) {
						return $templateCache.get('account/login/login.tpl.html')
					}
				}
			})
			.state('register', {
				url: '/register',
				templateUrl: 'account/register/register.tpl.html',
				controller: 'RegisterCtrl',
				controllerAs: 'vm',
				resolve: {
					template: function($templateCache) {
						return $templateCache.get('account/register/register.tpl.html')
					}
				}
			})
		})
})();
