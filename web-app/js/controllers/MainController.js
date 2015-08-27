app.controller('MainController', ['$scope','gameOfLife', function ($scope, gameOfLife){
   /*gameOfLife.getBoard().success(function(data) {
		$scope.board = data;
		console.log(data)
   })*/
	var gameId = 0;
	gameOfLife.getBoard().then(function(response){
		$scope.board = response.data;
		console.log($scope.board);
	});
	

	$scope.stepGame = function() {
		gameOfLife.updateGame($scope.board, gameId);
		gameOfLife.stepBoard(gameId).then(function(response){
			$scope.board = response.data;
		});
	}

	
}]);
