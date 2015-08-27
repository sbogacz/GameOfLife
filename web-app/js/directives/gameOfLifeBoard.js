app.directive('gameOfLifeBoard', function() {
  return {
    restrict: 'E',
    scope: {
      board: '='
    },
    templateUrl: 'js/directives/gameOfLifeBoard.html'
  };
});
