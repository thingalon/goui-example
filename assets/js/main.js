$(document).ready(function() {
	goui.SetMessageHandler('showName', function(data) {
		$('#name').html(data.name);
	});
	goui.SetMessageHandler('threadOutput', function(data) {
		$('#thread-output')
			.append($('<div class="line">').html(data.text))
			.animate({scrollTop: $('#thread-output')[0].scrollHeight}, 10);
	});
	goui.Init();
	
	$( "#duration-slider" ).slider({
    	orientation: "horizontal",
    	min: 10,
    	max: 120,
    	value: 60,
    	slide: function(event, ui) {
			$( "#duration-description" ).html(ui.value + ' seconds');
		}
	});
	
	$( '#speed-slider' ).slider({
		min: 0.1,
		max: 5,
		value: 1.5,
		slide: function(event, ui) {
			$( "#speed-description" ).html(ui.value + ' seconds');
		}
	});
	
	$('#example-popup').click(function() {
		goui.SendMessage('examples.showPopup', {data:$('#modal-input').val()});
	});
	
	$('#spawn-thread').click(function() {
		var duration = $('#duration-slider').slider('option', 'value');
		var speed = $('#speed-slider').slider('option', 'value');
		goui.SendMessage('examples.spawnThread', {duration:duration, speed:speed});
	});
});