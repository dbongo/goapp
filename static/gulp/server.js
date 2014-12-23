var gulp = require('gulp')
var connect = require('gulp-connect')

gulp.task('server', function() {
  connect.server({
    root: ['dist'],
    port: 8080,
    livereload: true
  })
})
