'use strict';

/* Controllers */

angular.module('xtv.controllers').
  controller('SettingsCtrl', ['$scope', '$http', function($scope, $http) {

    $scope.loadSettings = function () {
      $http.get('data/loadSettings').success(function (response) {
        if (response.success) {
          $scope.settings = response.data;

        } else {
          msg.error(response.message);
        }
      });
    };
    $scope.loadSettings();

    $scope.saveSettings = function () {
      $http.post('data/saveSettings', $scope.settings).success(function (response) {
        if (response.success) {
          $scope.hideSettingsDialog();
          $scope.loadSettings();
        } else {
          msg.error(response.message);
        }
      });
    };

    $scope.$on('xtv:showSettingsDialog', function () {
      $scope.showSettingsDialog();
    });

    $scope.showSettingsDialog = function () {
      $('#settingsDialog').modal('show');
    };

    $scope.hideSettingsDialog = function () {
      $('#settingsDialog').modal('hide');
    };

  }]);
