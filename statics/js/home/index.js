;
(function($){
	/*刷新加载*/
    var myScroll = new IScroll('#wrapper', {
        probeType: 3
        , scrollbars: false
        , click: true
    });
    /*单行广告轮播*/
    function poster() {
        var liHeight = $(".font_inner li").height(); //一个li的高度
        var clickEndFlag = true; //设置每张走完才能再点击
        setInterval(function () {
            $(".font_inner").animate({
                top: -liHeight
            }, 500, function () {
                $(".font_inner li").eq(0).appendTo($(".font_inner")); //克隆第一个放到最后(实现无缝滚动)
                $(".font_inner").css({
                    "top": 0
                });
            })
        }, 3000)
    };

	var mySwiper = new Swiper('#lunbo', {
        direction: 'horizontal'
        , loop: true
        , autoplay: 5000
        , pagination: '.swiper-pagination' // 如果需要分页器
    });

    var swiper = new Swiper('#newSelected', {
        slidesPerView: 3.3
        , spaceBetween: 12
        , freeMode: true
    });

	poster();
    indexAjax();
	function indexAjax() {
		topGradient();
		//头部渐变背景
		function topGradient() {
		    var finished = true;
		    var refresh = false;
		    myScroll.on('scrollStart', function () {
		        startY = myScroll.y
		    });
		    myScroll.on('scroll', function () {
		        var scrollY = -myScroll.y
		        var wrapperHeight = $("#thelist").height() - $("#wrapper").height();
		        //下拉刷新
		        if (finished && scrollY < -10 && startY == 0) {
		            $('.head_top').hide();
		        }
		        if (finished && scrollY < -40 && startY == 0) {
		            finished = false;
		            $('#pullDown-msg').text('松开刷新');
		            $('#PullDown').css({
		                "position": "static"
		            });
		            myScroll.scrollToElement(PullDown, 0);
		            $('#wrapper').on('touchend', function () {
		                $('#pullDown-msg').text('刷新中...');
		                refresh = true;
		                setTimeout(function () {
		                    window.location.reload();
		                }, 1000)
		            });
		        }
		        myScroll.on('scrollEnd', function () {
		            if (scrollY >= -10 && !refresh) {
		                $('.head_top').show();
		            }
		        });

		        //导航栏背景
		        if (scrollY > 100) {
		            $(".head_top").addClass('gradient');
		        }
		        else {
		            $(".head_top").removeClass('gradient');
		        }
		    });
		    myScroll.on('scrollEnd', function () {
		        myScroll.refresh();
		    });
		};
	}

})(jQuery);

var pageSize = 1;
var pageCode = 2;
function jiazai(){
if (pageSize <= pageCode) {
		pageSize++;
         var sc=$('#sc').val();
         $.ajax({
            url:  LOCAL_HOME+"xjcsAjax/prods/",
            data: {ajaxdata: "list", t: 0,sc:sc, p: pageSize, ajax: Math.random()},
            async: false,
            beforeSend: function () {
            },
            dataType: 'json',
            success: function (o) {
                 if (o.status == 200) {
                    $("#prolist").append(o.message);
                    pageCode=o.data.totalPage;
                }
                else{
                    $("#PullUp").html('已加载全部产品');
                }
            },
            complete: function () {
            },
            error: function () {
            }
        });
    }
    else {
        $("#PullUp").html('已加载全部产品');
    }
}