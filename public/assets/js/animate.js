function animatePage() {
    var params = [
        {
            type: dynamics.bounce,
            delay: 0,
            option: {
                translateX: $('.page-header').width(),
                backgroundColor: '#628f5b'
            }
        },
        {
            type: dynamics.gravity,
            delay: 2052,
            option: {
                translateX: $('.page-header').width(),
                backgroundColor: '#5bc0de'
            }
        },
        {
            type: dynamics.gravity,
            delay: 4104,
            option: {
                translateX: 0,
                backgroundColor: '#AB0000'
            }
        }
    ];
    for (i = 0; i < params.length; ++i) {
        dynamics.animate(document.querySelector('.dot-page-header'), params[i].option, {
            type: params[i].type,
            duration: 2000,
            frequency: 15,
            friction: 1,
            delay: params[i].delay
        });
    }
    dynamics.setTimeout(animatePage, 7210)
}

function animateDots() {
    var options = [
        {translateX: $('.container').width() - 6, translateY: 0, backgroundColor: '#337ab7'},
        {translateX: $('.container').width() - 6, translateY: $('.lweel').height(), backgroundColor: '#f0ad4e'},
        {translateY: $('.lweel').height() + 8, translateX: 0, backgroundColor: '#d9534f'},
        {translateX: 0, translateY: 0, backgroundColor: '#5bc0de'}
    ];
    for (i = 0; i < options.length; ++i) {
        dynamics.animate(document.querySelector('.dot'), options[i], {
            type: dynamics.spring,
            duration: 2050,
            frequency: 1,
            friction: 200,
            delay: i * 2100
        })
    }
    dynamics.setTimeout(animateDots, 8450)
}

$(function () {
    animatePage();
    animateDots();
});
