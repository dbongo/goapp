var gulp = require('gulp')

gulp.task('index', function () {
  return gulp.src('./index.html')
  .pipe(gulp.dest('dist'))
})

gulp.task('favicon', function () {
  return gulp.src('./favicon.ico')
  .pipe(gulp.dest('dist'))
})
