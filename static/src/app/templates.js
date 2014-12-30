(function(module) {
try {
  module = angular.module('app');
} catch (e) {
  module = angular.module('app', []);
}
module.run(['$templateCache', function($templateCache) {
  $templateCache.put('home/home.tpl.html',
    '<div class="container">\n' +
    '    <div class="jumbotron">\n' +
    '        <h1>Evant Admin Portal</h1>\n' +
    '    </div>\n' +
    '</div>\n' +
    '');
}]);
})();

(function(module) {
try {
  module = angular.module('app');
} catch (e) {
  module = angular.module('app', []);
}
module.run(['$templateCache', function($templateCache) {
  $templateCache.put('account/login/login.tpl.html',
    '<div class="container">\n' +
    '    <alert ng-repeat="alert in vm.alerts" type="{{alert.type}}" close="vm.closeAlert($index)">{{alert.msg}}</alert>\n' +
    '    <form name="form" role="form" class="form-horizontal" ng-submit="vm.login(form)">\n' +
    '        <fieldset>\n' +
    '            <legend>Sign in</legend>\n' +
    '            <div class="form-group" ng-class="form.username.$invalid?\'has-error\':\'has-success\'">\n' +
    '                <label for="username" class="col-sm-2 control-label">Username</label>\n' +
    '                <div class="col-sm-10">\n' +
    '                    <div class="input-group">\n' +
    '                        <span class="input-group-addon"><i class="fa fa-envelope"></i></span>\n' +
    '                        <input id="username" name="username" type="text" class="form-control" placeholder="Username" ng-model="vm.user.username" required>\n' +
    '                    </div>\n' +
    '                    <span class="help-block" ng-show="form.username.$error.username">Please enter a valid username</span>\n' +
    '                </div>\n' +
    '            </div>\n' +
    '            <div class="form-group" ng-class="form.password.$invalid?\'has-error\':\'has-success\'">\n' +
    '                <label for="password" class="col-sm-2 control-label">Password</label>\n' +
    '                <div class="col-sm-10">\n' +
    '                    <div class="input-group">\n' +
    '                        <span class="input-group-addon"><i class="fa fa-key"></i></span>\n' +
    '                        <input id="password" name="password" type="password" class="form-control" placeholder="Password" ng-model="vm.user.password" required>\n' +
    '                    </div>\n' +
    '                </div>\n' +
    '            </div>\n' +
    '            <div class="form-group">\n' +
    '                <div class="col-sm-offset-2 col-sm-10">\n' +
    '                    <button type="submit" class="btn btn-success">Login</button>\n' +
    '                    <a class="btn btn-primary" ui-sref="register">Sign up</a>\n' +
    '                    <a class="btn btn-link" href="#">Forgot your password?</a>\n' +
    '                </div>\n' +
    '            </div>\n' +
    '        </fieldset>\n' +
    '    </form>\n' +
    '</div>\n' +
    '');
}]);
})();

(function(module) {
try {
  module = angular.module('app');
} catch (e) {
  module = angular.module('app', []);
}
module.run(['$templateCache', function($templateCache) {
  $templateCache.put('account/register/register.tpl.html',
    '<div class="container">\n' +
    '    <alert ng-repeat="alert in vm.alerts" type="{{alert.type}}" close="vm.closeAlert($index)">{{alert.msg}}</alert>\n' +
    '    <form name="form" role="form" class="form-horizontal" ng-submit="vm.register(form)">\n' +
    '        <fieldset>\n' +
    '            <legend>Sign up</legend>\n' +
    '            <div class="form-group" ng-class="form.username.$invalid?\'has-error\':\'has-success\'">\n' +
    '                <div class="col-sm-10">\n' +
    '                    <div class="input-group">\n' +
    '                        <span class="input-group-addon"><i class="fa fa-envelope"></i></span>\n' +
    '                        <input id="username" name="username" type="text" class="form-control" placeholder="Username" ng-model="vm.user.username" required>\n' +
    '                    </div>\n' +
    '                    <span class="help-block" ng-show="form.username.$error.username">Please enter a valid username</span>\n' +
    '                </div>\n' +
    '            </div>\n' +
    '            <div class="form-group" ng-class="form.password.$invalid?\'has-error\':\'has-success\'">\n' +
    '                <label for="password" class="col-sm-2 control-label">Password</label>\n' +
    '                <div class="col-sm-10">\n' +
    '                    <div class="input-group">\n' +
    '                        <span class="input-group-addon"><i class="fa fa-key"></i></span>\n' +
    '                        <input id="password" name="password" type="password" class="form-control" placeholder="Password" ng-model="vm.user.password" required>\n' +
    '                    </div>\n' +
    '                </div>\n' +
    '            </div>\n' +
    '            <div ng-show=\'form.password.$valid\' class="form-group" ng-class="form.password2.$error.match?\'has-error\':\'has-success\'">\n' +
    '                <label for="password2" class="col-sm-2 control-label">Repeat Password</label>\n' +
    '                <div class="col-sm-10">\n' +
    '                    <input id="password2" name="password2" type="password" class="form-control" placeholder="Repeat Password" ng-model="vm.user.password2" match="vm.user.password">\n' +
    '                    <span class="help-block" ng-show="form.password2.$error.match">Both passwords must be the same</span>\n' +
    '                </div>\n' +
    '            </div>\n' +
    '            <div class="form-group">\n' +
    '                <div class="col-sm-offset-2 col-sm-10">\n' +
    '                    <button type="submit" class="btn btn-success">Create</button>\n' +
    '                    <a class="btn btn-link" ui-sref="login">Already have an account?</a>\n' +
    '                </div>\n' +
    '            </div>\n' +
    '        </fieldset>\n' +
    '    </form>\n' +
    '</div>\n' +
    '');
}]);
})();

(function(module) {
try {
  module = angular.module('app');
} catch (e) {
  module = angular.module('app', []);
}
module.run(['$templateCache', function($templateCache) {
  $templateCache.put('user/posts/posts.tpl.html',
    '<div class=\'container\'>\n' +
    '    <h1>Recent Posts</h1>\n' +
    '    <form role=\'form\'>\n' +
    '        <div class=\'form-group\'>\n' +
    '            <div class=\'input-group\'>\n' +
    '                <input ng-model=\'vm.postBody\' class=\'form-control\'>\n' +
    '                <span class=\'input-group-btn\'>\n' +
    '                    <button ng-click=\'vm.addPost()\' class=\'btn btn-default\'>Add Post</button>\n' +
    '                </span>\n' +
    '            </div>\n' +
    '        </div>\n' +
    '    </form>\n' +
    '    <ul class=\'posts list-group\'>\n' +
    '        <li ng-repeat="post in vm.posts | orderBy:\'-created\'" class=\'list-group-item\'>\n' +
    '            <strong>@{{ post.username }}</strong>\n' +
    '            <span>{{ post.body }}</span>\n' +
    '        </li>\n' +
    '    </ul>\n' +
    '</div>\n' +
    '');
}]);
})();
