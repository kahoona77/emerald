'use strict';

/* Controllers */

angular.module('emerald.controllers').
  controller('LogFileCtrl', ['$scope', '$http', function($scope, $http) {

    $scope.loadLogFile = function () {
      $http.get('data/loadLogFile').success(function(response){
          $scope.logFile = response.data;
      });
    };
    $scope.loadLogFile();

    $scope.clearLogFile = function () {
      $http.get('data/clearLogFile').success(function(response){
          $scope.loadLogFile ();
      });
    };
}]);
