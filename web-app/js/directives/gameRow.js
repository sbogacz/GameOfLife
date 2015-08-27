app.directive('gameRow', function() { 
  return { 
    restrict: 'E', 
    scope: { 
      info: '=' 
    }, 
    templateUrl: 'js/directives/gameRow.html' 
  }; 
});
