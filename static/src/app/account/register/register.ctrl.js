(function() {
	angular.module('app')
		.controller('RegisterCtrl', function RegisterCtrl($state, Auth) {
			var vm = this

			vm.alerts = []
			vm.user = {username: "", password: "", password2: ""}
			vm.register = register
			vm.closeAlert = closeAlert

			function register(form) {
				if (form.$valid) {
					Auth.createUser({
						username: vm.user.username,
						password: vm.user.password2
					})
                    .then(function() {
						vm.alerts = []
						$state.go('posts')
					})
                    .catch(function() {
						vm.alerts.push({type: "danger", msg: "The specified username is already in use"})
					})
				}
			}

			function closeAlert(index) {
				vm.alerts.splice(index, 1)
			}
		})
})();
