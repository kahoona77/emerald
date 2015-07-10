'use strict';

/* Controllers */

angular.module('xtv.controllers').
  controller('LogFileCtrl', ['$scope', '$http', function($scope, $http) {

    $scope.loadLogFile = function () {
      $http.get('data/loadLogFile').success(function(response){
          $scope.logFile = response.logFile;
      });
    };
    $scope.loadLogFile();

    $scope.clearLogFile = function () {
      $http.get('data/clearLogFile').success(function(response){
          $scope.loadLogFile ();
      });
    };
}]);