var navItems = document.querySelectorAll(".mobile-bottom-nav__item");
navItems.forEach(function (e, i) {
	e.addEventListener("click", function (e) {
		navItems.forEach(function (e2, i2) {
			e2.classList.remove("mobile-bottom-nav__item--active");
		});
		this.classList.add("mobile-bottom-nav__item--active");
	});
});

$('#home').on('click', function(){
	$('#main_body').load((window.location.href).split('/mobile/')[0] + '/mobile/dashboard' + ' #main_body');
})

$('#enforcers').on('click', function(){
	$('#main_body').load((window.location.href).split('/mobile/')[0] +'/mobile/enforcers'+ ' #main_body');
})

$('#help').on('click', function () {
	$('#main_body').load((window.location.href).split('/mobile/')[0] + '/mobile/help' + ' #main_body');
})

$('#about').on('click', function () {
	$('#main_body').load((window.location.href).split('/mobile/')[0] + '/mobile/about' + ' #main_body');
})