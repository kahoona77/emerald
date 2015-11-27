'use strict';

/* Controllers */

angular.module('emerald.controllers').
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

    $scope.showServerDialog = function (server) {
      if (!server) {
        server = {
           name: "",
           port: "",
           status: 'Not Connected',
           channels: []
       };
      }

      $scope.editServer = server;
      $('#serverDialog').modal('show');
    };

    $scope.saveServer = function () {
      $http.post ('data/saveServer', $scope.editServer).success (function (response) {
        if (response.status = 'ok') {
          $('#serverDialog').modal('hide');
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
      if (server.status) {
         return 'fa-globe';
      }
      return 'fa-ban';
    };

    $scope.toggleConnection = function (server) {
      $http.post ('irc/toggleConnection', angular.copy (server)).success (function (response) {
        if (response.success) {
          $scope.getServerStatus (server);
        } else {
          msg.error (response.message);
        }
      });
    };

    $scope.getServerStatus = function (server) {
      $http.post ('irc/getServerStatus', angular.copy (server)).success (function (response) {
        if (response.success) {
          server.status = response.data.connected;
        } else {
          msg.error (response.message);
        }
      });
    };

    $scope.loadConsole = function (server) {
      $scope.selectedServerConsole = undefined;
      $http.post ('irc/getServerConsole', angular.copy (server)).success (function (response) {
        if (response.success) {
          $scope.selectedServerConsole = response.data;
        } else {
          msg.error (response.message);
        }
      });
    };


    $scope.showAddDownloadDialog = function (server) {
      $scope.newDownload = {
        message: '',
        server: server.name
      };
      $('#addDownloadDialog').modal ('show');
    };

    $scope.addDownload = function (newDownload) {
      $http.post ('downloads/startDirectDownload', newDownload).success (function (response) {
        if (response.success) {
          $scope.selectedServerConsole = response.data;
        } else {
          msg.error (response.message);
        }
      });
      $('#addDownloadDialog').modal ('hide');
    }

  }]);
