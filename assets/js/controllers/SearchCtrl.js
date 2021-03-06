'use strict';

/* Controllers */

angular.module('emerald.controllers').
  controller('SearchCtrl', ['$scope', 'msg', '$http', '$routeParams', function($scope, msg, $http, $routeParams) {

    $scope.query = undefined;
    $scope.packetCount = '';
    $scope.searchResults = undefined;

    $scope.search = function () {
      $http.get('data/findPackets', {params : {query: $scope.query}}).success(function(response){
        if (response.success) {
          $scope.searchResults = response.data;
        } else {
          msg.error (response.message);
        }
      });
    };

    $scope.countPackets = function () {
      $http.get('data/countPackets').success(function(response){
        if (response.success) {
          $scope.packetCount = response.data;
        } else {
          msg.error (response.message);
        }
      });
    };
    $scope.countPackets();

    $scope.startDownload = function (item) {
      $http.post('downloads/downloadPacket', item).success(function(response){
        if (response.success) {
          msg.show ("Added '" + item.name + "' to Download-Queue.");
        } else {
          msg.error (response.message);
        }
      });
    };

    if ($routeParams.query) {
      $scope.query = $routeParams.query;
      $scope.search();
    }

  }]);
