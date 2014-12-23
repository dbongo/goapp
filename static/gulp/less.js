var gulp = require('gulp')
var concat = require('gulp-concat')
var less = require('gulp-less')

gulp.task('less', function () {
  return gulp.src('less/app.less')
  .pipe(concat('app.css'))
  .pipe(less())
  .pipe(gulp.dest('dist'))
})

gulp.task('watch:less', ['less'], function () {
  gulp.watch('less/**/*.less', ['less'])
})
