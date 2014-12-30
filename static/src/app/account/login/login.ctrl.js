(function() {
	angular.module('app')
		.controller('LoginCtrl', function LoginCtrl($state, Auth) {
			var vm = this
            
			vm.alerts = []
			vm.user = {username: "", password: ""}
			vm.login = login
			vm.closeAlert = closeAlert

			function login(form) {
				if (form.$valid) {
					Auth.login(vm.user)
                    .then(function() {
						vm.alerts = []
						$state.go('posts')
					})
                    .catch(function(err) {
						vm.alerts.push({type: "danger", msg: err.message})
					})
				}
			}

			function closeAlert(index) {
				vm.alerts.splice(index, 1)
			}
		})
})();
