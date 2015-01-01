function AppCtrl(Auth) {
	var vm = this

	vm.getCurrentUser = Auth.getCurrentUser
	//vm.isLoggedIn = Auth.isLoggedIn
	//vm.isAdmin = Auth.isAdmin
	vm.logout = Auth.logout
}

angular.module('app')
.controller('AppCtrl', AppCtrl)
