(function() {
	angular.module('app')
		.config(function($stateProvider) {
			$stateProvider
			.state('home', {
				url: '/',
				templateUrl: 'home/home.tpl.html',
				resolve: {
					template: function($templateCache) {
						return $templateCache.get('home/home.tpl.html')
					}
				}
			})
		})
})();
