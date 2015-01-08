angular.module('hackapp').controller('AppCtrl', function AppCtrl(Auth) {
  var vm = this

  vm.getCurrentUser = Auth.getCurrentUser
  vm.logout = Auth.logout
})
