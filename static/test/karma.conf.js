module.exports = function(config) {
  config.set({
    basePath: './',
    frameworks: ['mocha', 'chai', 'sinon-chai'],
    files: [
      'dist/angular/angular.js',
      'dist/angular-ui-router/release/angular-ui-router.js',
      'dist/angular-mocks/angular-mocks.js',
      'dist/templates.js',
      'app/module.js',
      'app/*.js',
      'test/*.spec.js'
    ],
    reporters: ['progress'],
    port: 9876,
    colors: true,
    logLevel: config.LOG_INFO,
    autoWatch: true,
    browsers: ['PhantomJS'],
    singleRun: false
  })
}
