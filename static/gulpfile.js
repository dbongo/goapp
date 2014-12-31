var gulp = require('gulp')
var ngHtml2Js = require('gulp-ng-html2js')
var $ = require('gulp-load-plugins')()

var paths = {
	styles: './src/less/app.less',
	index: './src/index.html',
	vendor: [
		'!./src/vendor/**/*.{js,css,less,scss,json,md,gzip,txt}',
		'!./src/vendor/angular-scenario',
		'!./src/vendor/angular-mocks',
		'./src/vendor/**/*.*'
	],
	js: [
		'./src/app/**/*.module.js',
		'./src/app/**/*.js',
		'./src/common/**/*.js'
	],
	dest: '/Users/dbongo/www'
}

gulp.task('analyze', function() {
	var jshint = analyzejshint([].concat(paths.js), './.jshintrc')
	return jshint
})

gulp.task('templatecache', function() {
	return gulp.src('./src/app/**/*.html')
	.pipe(ngHtml2Js({
		moduleName: 'app',
		prefix: ''
	}))
	.pipe($.concat('templates.js'))
	.pipe(gulp.dest('./src/app'))
})

gulp.task('scripts', ['templatecache', 'analyze'], function() {
	return gulp.src(paths.index)
	.pipe($.usemin({
		js: [$.ngAnnotate({
			add: true,
			single_quotes: true
		})]
	}))
	.pipe(gulp.dest(paths.dest))
})

gulp.task('vendor', function() {
	return gulp.src(paths.vendor)
	.pipe(gulp.dest(paths.dest))
})

gulp.task('favicon', function() {
	return gulp.src('./src/favicon.ico')
	.pipe(gulp.dest(paths.dest))
})

gulp.task('styles', function() {
	return gulp.src(paths.styles)
	.pipe($.less())
	.pipe(gulp.dest(paths.dest))
})

gulp.task('server', function() {
	$.connect.server({
		root: paths.dest,
		port: 3000,
		livereload: true
	})
})

function analyzejshint(sources, jshintrc) {
	return gulp.src(sources)
	.pipe($.jshint(jshintrc))
	.pipe($.jshint.reporter('jshint-stylish'))
}

gulp.task('copy', ['vendor', 'favicon'])
gulp.task('build', ['styles', 'scripts', 'copy'])
gulp.task('default', ['build'])
