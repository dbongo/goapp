var gulp   = require('gulp')
var concat = require('gulp-concat')
//var uglify = require('gulp-uglify')
var ngAnnotate = require('gulp-ng-annotate')
//var sourcemaps = require('gulp-sourcemaps')

gulp.task('js', function () {
  return gulp.src(['app/module.js','app/**/*.js','dist/templates.js'])
  //.pipe(ngAnnotate())
  //.pipe(sourcemaps.init())
  .pipe(ngAnnotate())
  .pipe(concat('app.js'))
  //.pipe(ngAnnotate())
  //.pipe(uglify())
  //.pipe(sourcemaps.write())
  .pipe(gulp.dest('dist'))
})

// gulp.task('watch:js', ['js'], function () {
//   gulp.watch('app/**/*.js', ['js'])
// })
