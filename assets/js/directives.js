'use strict';

/* Directives */


angular.module('xtv.directives', []).
  directive('downloadProgress', [function() {
    return {
      restrict: 'E',
      template: '<div class="progress xtv-progress"><div class="progress-bar progress-bar-{{status}}" role="progressbar" aria-valuenow="{{progress}}" aria-valuemin="0" aria-valuemax="100" style="width: {{progress}}%;">{{progress}}%</div></div>',
      scope: {
        item: '=ngModel'
      },
      link: function (scope) {

        scope.$watch ('item', function (newValue) {
          scope.progress = parseInt( (scope.item.bytesReceived / scope.item.size) * 100);

          scope.status = 'default';
          if (scope.item.status == 'FAILED') {
            scope.status = 'danger';
          } else  if (scope.item.status == 'COMPLETE') {
            scope.status = 'success';
          }
        }, true);

      }
    };
  }]).
  directive('confirmDialog', [function() {
    return {
      restrict: 'E',
      replace: true,
      transclude: true,
      template: '<div class="modal fade xtv-confirm-dialog" tabindex="-1" role="dialog" aria-labelledby="confirmDialogLabel" aria-hidden="true"><div class="modal-dialog"><div class="modal-content"> <div class="modal-header"><button type="button" class="close" data-dismiss="modal" aria-hidden="true">&times;</button><h4 class="modal-title" id="confirmDialogLabel">Confirm</h4></div><div class="modal-body" ng-transclude></div><div class="modal-footer"><button type="button" class="btn btn-default" data-dismiss="modal">Close</button><button type="button" class="btn btn-primary" ng-click="confirm()" data-dismiss="modal">Ok</button></div></div></div></div>',
      scope: {
        confirm: '&'
      }
    };
  }]);
