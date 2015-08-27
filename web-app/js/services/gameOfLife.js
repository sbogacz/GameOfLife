app.factory('gameOfLife', function($http) {
	var gameOfLife = {};

	gameOfLife.data = {};
    
	var gameId = 0;
/*	gameOfLife.getBoard =  function() {
		$http.get("http://199.233.247.129:8080/getGameJSON/1")
			.success(function(data) {
				 gameOfLife.data.board = data;
			})
			.error(function(error) {
				return error;
			});
			return gameOfLife.data;
	}*/
	gameOfLife.getBoard = function() { 
		return $http.get("http://199.233.247.129:8080/getGameJSON/" + gameId);
	};
	
	gameOfLife.stepBoard = function(id) {
		return $http.put("http://199.233.247.129:8080/stepGame/" + id);
	}

	gameOfLife.newBoard = function() {
		return $http.put("http://199.233.247.129:8080/initGame/classic");
	}

	gameOfLife.updateGame = function(board, id) {
		console.log(angular.toJson(board));
		return $http.put("http://199.233.247.129:8080/updateGame/" + id, angular.toJson(board));
	}

	return gameOfLife;
});
