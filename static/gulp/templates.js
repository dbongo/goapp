var gulp = require('gulp')
var ngHtml2Js = require('gulp-ng-html2js')
var concat = require('gulp-concat')
//var minifyHtml = require('gulp-minify-html')
//var uglify = require('gulp-uglify')


gulp.task('templates', function () {
  return gulp.src(['app/**/*.html'])
  //.pipe(minifyHtml({empty: true, spare: true, quotes: true}))
  //.pipe(ngHtml2Js({moduleName: 'templates', prefix: ''}))
  .pipe(ngHtml2Js({moduleName: 'app', prefix: ''}))
  .pipe(concat('templates.js'))
  //.pipe(uglify())
  .pipe(gulp.dest('dist'));
})

// gulp.task('watch:templates', ['templates'], function () {
//   gulp.watch('templates/**/*.html', ['templates'])
// })
