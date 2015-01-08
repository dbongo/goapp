angular.module('hackapp').controller('LoginCtrl', function LoginCtrl($state, Auth) {
    var vm = this
    vm.alerts = []
    vm.user = {}

    function login(form) {
        if (form.$valid) {
            Auth.login({
                email: vm.user.email,
                password: vm.user.password
            })
            .then(function() {
                vm.alerts = []
                $state.go('posts')
            })
            .catch(function(err) {
                vm.alerts.push({
                    type: "danger",
                    msg: err.message
                })
                Auth.logout()
            })
        }
    }
    vm.login = login

    function closeAlert(index) {
        vm.alerts.splice(index, 1)
    }
    vm.closeAlert = closeAlert
})
