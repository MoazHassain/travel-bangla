/* changing sidebar active according to page */

console.log('test');

var navigations = document.querySelectorAll("#navbarCollapse a.nav-item");

navigations.forEach(navigation => {
    // var currentPageLocation = location.href;
    var locationUrlSplit = location.pathname.split("/");
    var rootPage = locationUrlSplit[1];
    
    var navUrlSplit = navigation.pathname.split("/");
    var navUrl = navUrlSplit[1];
    console.log(navUrl + " " + rootPage)
    if(navUrl === rootPage){
        navigation.classList.add("active")
    }
});

(function ($) {
    "use strict";
    
    // Dropdown on mouse hover
    $(document).ready(function () {
        function toggleNavbarMethod() {
            if ($(window).width() > 992) {
                $('.navbar .dropdown').on('mouseover', function () {
                    $('.dropdown-toggle', this).trigger('click');
                }).on('mouseout', function () {
                    $('.dropdown-toggle', this).trigger('click').blur();
                });
            } else {
                $('.navbar .dropdown').off('mouseover').off('mouseout');
            }
        }
        toggleNavbarMethod();
        $(window).resize(toggleNavbarMethod);
    });
    
    
    // Back to top button
    $(window).scroll(function () {
        if ($(this).scrollTop() > 100) {
            $('.back-to-top').fadeIn('slow');
        } else {
            $('.back-to-top').fadeOut('slow');
        }
    });
    $('.back-to-top').click(function () {
        $('html, body').animate({scrollTop: 0}, 1500, 'easeInOutExpo');
        return false;
    });


    // // Date and time picker
    // $('.date').datetimepicker({
    //     format: 'L'
    // });
    // $('.time').datetimepicker({
    //     format: 'LT'
    // });


    // Testimonials carousel
    $(".testimonial-carousel").owlCarousel({
        autoplay: true,
        smartSpeed: 1500,
        margin: 30,
        dots: true,
        loop: true,
        center: true,
        responsive: {
            0:{
                items:1
            },
            576:{
                items:1
            },
            768:{
                items:2
            },
            992:{
                items:3
            }
        }
    });
    
})(jQuery);



// "use strict";



// // Dropdown on mouse hover
// document.addEventListener('DOMContentLoaded', function () {
//     function toggleNavbarMethod() {
//         if (window.innerWidth > 992) {
//             document.querySelectorAll('.navbar .dropdown').forEach(function (dropdown) {
//                 dropdown.addEventListener('mouseover', function () {
//                     dropdown.querySelector('.dropdown-toggle').click();
//                 });
//                 dropdown.addEventListener('mouseout', function () {
//                     dropdown.querySelector('.dropdown-toggle').click();
//                     dropdown.querySelector('.dropdown-toggle').blur();
//                 });
//             });
//         } else {
//             document.querySelectorAll('.navbar .dropdown').forEach(function (dropdown) {
//                 dropdown.removeEventListener('mouseover', function () {});
//                 dropdown.removeEventListener('mouseout', function () {});
//             });
//         }
//     }
//     toggleNavbarMethod();
//     window.addEventListener('resize', toggleNavbarMethod);
// });

// // Back to top button
// window.addEventListener('scroll', function () {
//     if (window.scrollY > 100) {
//         document.querySelector('.back-to-top').style.display = 'block';
//     } else {
//         document.querySelector('.back-to-top').style.display = 'none';
//     }
// });
// document.querySelector('.back-to-top').addEventListener('click', function () {
//     window.scrollTo({top: 0, behavior: 'smooth'});
//     return false;
// });

// // Date and time picker (Using a library like Flatpickr for vanilla JS)
// // flatpickr('.date', {
// //     dateFormat: 'Y-m-d'
// // });
// // flatpickr('.time', {
// //     enableTime: true,
// //     noCalendar: true,
// //     dateFormat: 'H:i'
// // });

// // Testimonials carousel (Using a library like Glide.js for vanilla JS)
// document.addEventListener('DOMContentLoaded', function () {
//     const carousel = document.querySelector('.testimonial-carousel');
//     if (carousel) {
//         new Glide(carousel, {
//             type: 'carousel',
//             autoplay: 3000,
//             hoverpause: true,
//             animationDuration: 1500,
//             gap: 30,
//             perView: 3,
//             breakpoints: {
//                 992: {
//                     perView: 3
//                 },
//                 768: {
//                     perView: 2
//                 },
//                 576: {
//                     perView: 1
//                 },
//                 0: {
//                     perView: 1
//                 }
//             }
//         }).mount();
//     }
// });