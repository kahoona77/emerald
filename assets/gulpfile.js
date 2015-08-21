'use strict';

var gulp = require('gulp');
var sass = require('gulp-sass');

gulp.task('sass', function () {
  gulp.src('./css/bootstrap/**/*.scss')
    .pipe(sass().on('error', sass.logError))
    .pipe(gulp.dest('./css/bootstrap'));
});

gulp.task('sass:watch', function () {
  gulp.watch('./css/bootstrap/**/*.scss', ['sass']);
});
