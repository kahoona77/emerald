'use strict';

/* Controllers */

angular.module('xtv.controllers').
  controller('SearchCtrl', ['$scope', 'msg', '$http', '$routeParams', function($scope, msg, $http, $routeParams) {

    $scope.query = undefined;
    $scope.packetCount = '';
    $scope.searchResults = undefined;

    $scope.search = function () {
      $http.get('packets/findPackets', {params : {query: $scope.query}}).success(function(response){
        if (response.status == 'ok') {
          $scope.searchResults = response.packets;
        } else {
          msg.error (response.message);
        }
      });
    };

    $scope.countPackets = function () {
      $http.get('packets/countPackets').success(function(response){
        if (response.status == 'ok') {
          $scope.packetCount = response.count;
        } else {
          msg.error (response.message);
        }
      });
    };
    $scope.countPackets();

    $scope.startDownload = function (item) {
      $http.post('downloads/downloadPacket', {data: item}).success(function(response){
        if (response.status == 'ok') {
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
