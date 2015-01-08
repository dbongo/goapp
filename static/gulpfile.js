var gulp = require('gulp')
var usemin = require('gulp-usemin')
var connect = require('gulp-connect')
var minifyCss = require('gulp-minify-css')
var minifyJs = require('gulp-uglify')
var concat = require('gulp-concat')
var less = require('gulp-less')
var rename = require('gulp-rename')
var ngAnnotate = require('gulp-ng-annotate')
//var minifyHTML = require('gulp-minify-html')

var paths = {
  scripts: 'src/js/**/*.*',
  styles: 'src/less/hackapp.less',
  templates: 'src/templates/**/*.html',
  index: 'src/index.html',
  bower_fonts: 'src/vendor/**/*.{ttf,woff,eof,svg}',
}

/**
 * Handle bower components from index
 */
gulp.task('usemin', function() {
  return gulp.src(paths.index)
  .pipe(usemin({
    js: [minifyJs(), 'concat'],
    css: [minifyCss({
        keepSpecialComments: 0
      }), 'concat'],
  }))
  .pipe(gulp.dest('dist/'))
})

/**
 * Copy assets
 */
gulp.task('copy-bower_fonts', function() {
  return gulp.src(paths.bower_fonts)
  .pipe(rename({
    dirname: '/fonts'
  }))
  .pipe(gulp.dest('dist/lib'))
})

/**
 * Handle custom files
 */
gulp.task('custom-js', function() {
  return gulp.src(paths.scripts)
  //.pipe(minifyJs())
  .pipe(ngAnnotate({
    add: true,
    single_quotes: true
  }))
  .pipe(concat('hackapp.js'))
  .pipe(gulp.dest('dist/js'))
})

gulp.task('custom-less', function() {
  return gulp.src(paths.styles)
  .pipe(less())
  .pipe(gulp.dest('dist/css'))
})

gulp.task('custom-templates', function() {
  return gulp.src(paths.templates)
  //.pipe(minifyHTML())
  .pipe(gulp.dest('dist/templates'))
})

/**
 * Live reload server
 */
gulp.task('webserver', function() {
  connect.server({
    root: 'dist',
    livereload: true,
    port: 8888
  })
})

/**
 * Gulp tasks
 */
gulp.task('build-assets', ['copy-bower_fonts'])

gulp.task('build-custom', ['custom-js', 'custom-less', 'custom-templates'])

gulp.task('build', ['usemin', 'build-assets', 'build-custom'])

gulp.task('default', ['build', 'webserver'])
