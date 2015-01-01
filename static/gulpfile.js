// var gulp = require('gulp')
// var ngHtml2Js = require('gulp-ng-html2js')
// var $ = require('gulp-load-plugins')()
var gulp = require('gulp')
var usemin = require('gulp-usemin')
//var wrap = require('gulp-wrap')
var connect = require('gulp-connect')
//var watch = require('gulp-watch')
var minifyCss = require('gulp-minify-css')
var minifyJs = require('gulp-uglify')
var concat = require('gulp-concat')
var less = require('gulp-less')
var rename = require('gulp-rename')
var ngAnnotate = require('gulp-ng-annotate')
//var minifyHTML = require('gulp-minify-html')

var paths = {
	scripts: 'src/js/**/*.*',
	styles: 'src/less/app.less',
	templates: 'src/templates/**/*.html',
	index: 'src/index.html',
	bower_fonts: 'src/vendor/**/*.{ttf,woff,eof,svg}',
}

// var paths = {
// 	styles: './src/less/app.less',
// 	index: './src/index.html',
// 	vendor: [
// 		'!./src/vendor/**/*.{js,css,less,scss,json,md,gzip,txt}',
// 		'!./src/vendor/angular-scenario',
// 		'!./src/vendor/angular-mocks',
// 		'./src/vendor/**/*.*'
// 	],
// 	js: [
// 		'./src/app/**/*.module.js',
// 		'./src/app/**/*.js',
// 		'./src/common/**/*.js'
// 	],
// 	dest: './dist'
// }

/**
* Handle bower components from index
*/
gulp.task('usemin', function() {
	return gulp.src(paths.index)
	.pipe(usemin({
		js: [minifyJs(), 'concat'],
		css: [minifyCss({keepSpecialComments: 0}), 'concat'],
	}))
	.pipe(gulp.dest('dist/'))
})

/**
* Copy assets
*/
gulp.task('build-assets', ['copy-bower_fonts'])

gulp.task('copy-bower_fonts', function() {
	return gulp.src(paths.bower_fonts)
	.pipe(rename({dirname: '/fonts'}))
	.pipe(gulp.dest('dist/lib'))
})

/**
* Handle custom files
*/
gulp.task('build-custom', [
	//'custom-images',
	'custom-js',
	'custom-less',
	'custom-templates'
])

gulp.task('custom-js', function() {
	return gulp.src(paths.scripts)
	//.pipe(minifyJs())
	.pipe(ngAnnotate({add: true, single_quotes: true}))
	.pipe(concat('app.js'))
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
gulp.task('build', [
	'usemin',
	'build-assets',
	'build-custom'
])

gulp.task('default', [
	'build',
	'webserver'
])






// gulp.task('analyze', function() {
// 	var jshint = analyzejshint([].concat(paths.js), './.jshintrc')
// 	return jshint
// })
//
// gulp.task('templatecache', function() {
// 	return gulp.src('./src/app/**/*.html')
// 	.pipe(ngHtml2Js({
// 		moduleName: 'app',
// 		prefix: ''
// 	}))
// 	.pipe($.concat('templates.js'))
// 	.pipe(gulp.dest('./src/app'))
// })
//
// gulp.task('scripts', ['templatecache', 'analyze'], function() {
// 	return gulp.src(paths.index)
	// .pipe($.usemin({
	// 	js: [$.ngAnnotate({
	// 		add: true,
	// 		single_quotes: true
	// 	})]
	// }))
// 	.pipe(gulp.dest(paths.dest))
// })
//
// gulp.task('vendor', function() {
// 	return gulp.src(paths.vendor)
// 	.pipe(gulp.dest(paths.dest))
// })
//
// gulp.task('favicon', function() {
// 	return gulp.src('./src/favicon.ico')
// 	.pipe(gulp.dest(paths.dest))
// })
//
// gulp.task('styles', function() {
// 	return gulp.src(paths.styles)
// 	.pipe($.less())
// 	.pipe(gulp.dest(paths.dest))
// })
//
// gulp.task('server', function() {
// 	$.connect.server({
// 		root: paths.dest,
// 		port: 3000,
// 		livereload: true
// 	})
// })
//
// function analyzejshint(sources, jshintrc) {
// 	return gulp.src(sources)
// 	.pipe($.jshint(jshintrc))
// 	.pipe($.jshint.reporter('jshint-stylish'))
// }
//
// gulp.task('copy', ['vendor', 'favicon'])
// gulp.task('build', ['styles', 'scripts', 'copy'])
// gulp.task('default', ['build'])
