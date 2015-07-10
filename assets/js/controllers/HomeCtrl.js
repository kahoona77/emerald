'use strict';

/* Controllers */

angular.module('xtv.controllers').
  controller('HomeCtrl', ['$scope', '$http', 'msg', function($scope, $http, msg) {

    $scope.loadServers = function () {
      $http.get('data/loadServers').success(function(response){
        if (response.success) {
          $scope.servers = response.data;

          //reselect server
          angular.forEach ($scope.servers, function (server) {
             $scope.getServerStatus (server);
             if ($scope.selectedServer) {
               if (server.id == $scope.selectedServer.id) {
                 $scope.selectedServer = server;
                 $scope.loadConsole (server);
               }
             }
          });
        } else {
          msg.error (response.message);
        }
      });
    };
    $scope.loadServers();

    $scope.selectServer = function (server) {
      $scope.selectedServer = server;
      $scope.loadConsole (server);
    };

    $scope.showAddServerDialog = function () {
      $('#addServerDialog').modal('show');
    };

    $scope.addServer = function () {
      var newServer = {
          name: $scope.newServer.uri,
          port: $scope.newServer.port,
          status: 'Not Connected',
          channels: []
      };

      $http.post ('data/saveServer', newServer).success (function (response) {
        if (response.status = 'ok') {
          $('#addServerDialog').modal('hide');
          $scope.newServer = undefined;
          $scope.loadServers();
        } else {
          msg.error (response.message);
        }
      });
    };

    $scope.showAddChannelDialog = function () {
      $('#addChannelDialog').modal('show');
    };

    $scope.addChannel = function () {
      var channel = {
        name: $scope.newChannel.name
      };
      $scope.selectedServer.channels.push (channel);

      $http.post ('data/saveServer', $scope.selectedServer).success (function (response) {
        if (response.status = 'ok') {
          $('#addChannelDialog').modal('hide');
          $scope.newChannel = undefined;
          $scope.loadServers();
        } else {
          msg.error (response.message);
        }
      });
    };

    $scope.showDeleteServerConfirm = function (server) {
      $scope.serverToDelete = server;
      $('#deleteServerConfirmDialog').modal ('show');
    };

    $scope.deleteServer = function () {
      $http.post ('data/deleteServer', $scope.serverToDelete).success (function (response) {
        if (response.status = 'ok') {
          $scope.selectedServer = undefined;
          $scope.serverToDelete = undefined;
          $scope.selectedServerConsole = undefined;
          $scope.loadServers();
        } else {
          msg.error (response.message);
        }
      });
    };

    $scope.showDeleteChannelConfirm = function (channel) {
      $scope.channelToDelete = channel;
      $('#deleteChannelConfirmDialog').modal ('show');
    };

    $scope.deleteChannel = function () {
      // remove channel from Server
      $scope.selectedServer.channels = _.without($scope.selectedServer.channels, _.findWhere($scope.selectedServer.channels, {name: $scope.channelToDelete.name}));

      $http.post ('data/saveServer', $scope.selectedServer).success (function (response) {
        if (response.status = 'ok') {
          $scope.channelToDelete = undefined;
          $scope.loadServers();
        } else {
          msg.error (response.message);
        }
      });
    };

    $scope.getStatusClass = function (server) {
      if (server.status == 'Connected') {
         return 'mdi-social-public';
      }
      return 'mdi-notification-do-not-disturb';
    };

    $scope.toggleConnection = function (server) {
      $http.post ('irc/toggleConnection', angular.copy (server)).success (function (response) {
        if (response.status = 'ok') {
          server.status = response.result.status;
        } else {
          msg.error (response.message);
        }
      });
    };

    $scope.getServerStatus = function (server) {
      $http.post ('irc/getServerStatus', angular.copy (server)).success (function (response) {
        server.status = response.status;
      });
    };

    $scope.loadConsole = function (server) {
      $scope.selectedServerConsole = undefined;
      $http.post ('irc/getServerConsole', angular.copy (server)).success (function (response) {
        $scope.selectedServerConsole = response.console;
      });
    };
  }]);
