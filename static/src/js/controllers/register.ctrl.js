angular.module('app').controller('RegisterCtrl', function RegisterCtrl($state, Auth) {
	var vm = this

	vm.alerts = []
	vm.user = {}
	vm.register = register
	vm.closeAlert = closeAlert

	function register(form) {
		if (form.$valid) {
			Auth.register({email: vm.user.email, username: vm.user.username, password: vm.user.password2})
			.then(function() {
				vm.alerts = []
				$state.go('posts')
			})
			.catch(function(err) {
				vm.alerts.push({type: "danger", msg: err.message})
				Auth.logout()
			})
		}
	}

	function closeAlert(index) {
		vm.alerts.splice(index, 1)
	}
})
