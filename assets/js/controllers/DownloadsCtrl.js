'use strict';

/* Controllers */

angular.module('emerald.controllers').
  controller('DownloadsCtrl', ['$scope', '$timeout', '$http', 'msg', function($scope, $timeout, $http, msg) {

    $scope.downloads = [];

    $scope.loadDownloads = function () {
        $http.get('downloads/listDownloads',{params: { 'nocache': new Date().getTime() }}).success(function(response){
            if (response.success) {
                $scope.downloads = response.data;
                //reslect the download
                if ($scope.selectedDownload) {
                  angular.forEach ($scope.downloads, function (item) {
                     if (item.file == $scope.selectedDownload.file) {
                       $scope.selectedDownload = item;
                     }
                  });
                }
            } else {
                msg.error (response.message);
            }
        });
    };
    $scope.loadDownloads();

    $scope.startReloadTimer = function () {
      $scope.reloadTimer = $timeout (function () {
          $scope.loadDownloads();
          $scope.startReloadTimer();
      }, 1000)
    };
    $scope.startReloadTimer();

    $scope.$on ('$locationChangeStart', function () {
      if ($scope.reloadTimer) {
        $timeout.cancel ($scope.reloadTimer);
      }
    });

    $scope.selectDownload = function (item) {
      $scope.selectedDownload = item;
    };

    $scope.stopDownload = function () {
        $http.post('downloads/stopDownload', $scope.selectedDownload).success(function(response){
            if (response.success) {
                $scope.selectedDownload = undefined;
                $scope.loadDownloads();
            } else {
                msg.error (response.message);
            }
        });
    };

    $scope.resumeDownload = function () {
        $http.post('downloads/resumeDownload', $scope.selectedDownload).success(function(response){
            if (response.success) {
                $scope.selectedDownload = undefined;
                $scope.loadDownloads();
            } else {
                msg.error (response.message);
            }
        });
    };

    $scope.showCancelConfirm = function () {
      $('#downloadDeleteConfirmDialog').modal ('show');
    };

    $scope.cancelDownload = function () {
        $http.post('downloads/cancelDownload', $scope.selectedDownload).success(function(response){
            if (response.success) {
                $scope.selectedDownload = undefined;
                $scope.loadDownloads();
            } else {
                msg.error (response.message);
            }
        });
    };

    $scope.clearDownloads = function () {
      var completed = [];
      angular.forEach ($scope.downloads, function (item) {
         if (item.status == 'COMPLETE') {
           completed.push (item);
         }
      });

      angular.forEach (completed, function (item) {
         $http.post('downloads/cancelDownload', item).success(function(response){
              if (response.success) {
                  $scope.selectedDownload = undefined;
                  $scope.loadDownloads();
              } else {
                  msg.error (response.message);
              }
          });
      });
    };

    $scope.calcTimeRemaining = function (item) {
      var remainingKBytes = (item.size - item.bytesReceived) / 1024;
      var remainingSeconds = remainingKBytes / item.speed;

      var min = Math.floor(remainingSeconds / 60);
      var sec = Math.round(remainingSeconds % 60);

      if (min < 10) {
        min = '0' + min
      }

      if (sec < 10) {
        sec = '0' + sec
      }

      return min + ":" + sec + " Minutes";
    };


    //Files
    $scope.files = [];

    $scope.loadFiles = function () {
        $http.get('downloads/loadFiles').success(function(response){
            if (response.success) {
                $scope.files = response.data;
            } else {
                msg.error (response.message);
            }
        });
    };
    $scope.loadFiles();

    $scope.selectedFiles = [];
    $scope.selectFile = function (file) {
      if ($scope.isSelected(file)) {
        $scope.selectedFiles = _.without ($scope.selectedFiles, file);
      } else {
        $scope.selectedFiles.push(file);
      }
    };

    $scope.isSelected = function (file) {
      return _.contains ($scope.selectedFiles, file);
    };

    $scope.updateEpisodes = function () {
     $http.get('/shows/updateEpisodes').success(function(response){
        if (response.success) {
          msg.show ("Updating episodes started...");
        } else {
          msg.error (response.message);
        }
      });
    };


    $scope.showFileDelteConfirm = function () {
      $('#fileDeteConfirmDialog').modal ('show');
    };

    $scope.deleteSelectedFiles = function () {
        $http.post('downloads/deleteFiles', $scope.selectedFiles).success(function(response){
            if (response.success) {
                msg.show ("All files deleted!");
                $scope.selectedFiles = [];
                $scope.loadFiles();
            } else {
                msg.error (response.status);
            }
        });
    };

    $scope.showMoveFilesConfirm = function () {
      $('#moveFilesConfirmDialog').modal ('show');
    };

    $scope.moveFilesToMovies = function () {
        $http.post('downloads/moveFilesToMovies', $scope.selectedFiles).success(function(response){
            if (response.success) {
                msg.show ("All files moved!");
                $scope.selectedFiles = [];
                $scope.loadFiles();
            } else {
                msg.error (response.status);
            }
        });
    };
  }]);
