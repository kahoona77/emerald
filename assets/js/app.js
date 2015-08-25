'use strict';

angular.module('emerald.controllers', []);
angular.module('emerald.services', []);

// Declare app level module which depends on filters, and services
angular.module('emerald', [
  'ngRoute',
  'ngAnimate',
  'emerald.filters',
  'emerald.services',
  'emerald.directives',
  'emerald.controllers',
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
            $rootScope.$broadcast ('emerald:showSettingsDialog');
        }
    }]);
