'use strict';

angular.module('xtv.controllers', []);
angular.module('xtv.services', []);

// Declare app level module which depends on filters, and services
angular.module('xtv', [
  'ngRoute',
  'ngAnimate',
  'xtv.filters',
  'xtv.services',
  'xtv.directives',
  'xtv.controllers',
]).config(['$routeProvider', function($routeProvider) {

  $routeProvider.when('/home', {templateUrl: 'assets/partials/home.html', controller: 'HomeCtrl'});
  $routeProvider.when('/shows', {templateUrl: 'assets/partials/shows.html', controller: 'ShowsCtrl'});
  $routeProvider.when('/shows_recent', {templateUrl: 'assets/partials/shows_recent.html', controller: 'ShowsRecentCtrl'});
  $routeProvider.when('/search/:query?', {templateUrl: 'assets/partials/search.html', controller: 'SearchCtrl'});
  $routeProvider.when('/downloads', {templateUrl: 'assets/partials/downloads.html', controller: 'DownloadsCtrl'});
  $routeProvider.when('/logFile', {templateUrl: 'assets/partials/logFile.html', controller: 'LogFileCtrl'});
  $routeProvider.otherwise({redirectTo: '/home'});
}]).run(['$rootScope', function ($rootScope) {
        $rootScope.showSettingsDialog = function () {
            $rootScope.$broadcast ('xtv:showSettingsDialog');
        }
    }]);
