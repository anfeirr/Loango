$(document).ready(function () {
    var myScroll = new IScroll('#wrapper', {
        probeType: 3
        , scrollbars: true
        , click: true
        , shrinkScrollbars: "clip"
    });

    /*----------------页面渲染-------------------*/
    /*分页加载*/
    var bool = true; // 拉动标记

    var pageSize = 1;
    var pageCode = 2;
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
});

var pageSize = 1;
var pageCode = 2;
function jiazai(){
    if (pageSize <= pageCode) {
        pageSize++;
        var sc=$('#sc').val();
        var sh=$('#screen').val();
        $.ajax({
            url:  LOCAL_HOME+"xjcsAjax/prods/",
            data: {ajaxdata: "list", t: 2,sc:sc,sh:sh, p: pageSize, ajax: Math.random()},
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
function seachProds(obj){
    var sc=$('#sc').val();
    var sh=$('#screen').val();
    var shmoney=obj.replace("元", "");
    $.ajax({
        url:  LOCAL_HOME+"xjcsAjax/prods/",
        data: {ajaxdata: "list", t: 2,sc:sc,sh:sh,shmoney:shmoney, p: pageSize, ajax: Math.random()},
        async: false,
        beforeSend: function () {
        },
        dataType: 'json',
        success: function (o) {
            if (o.status == 200) {
                $("#prolist").html(o.message);
                pageCode=o.data.totalPage;
            }
            else{
                cmCommon.toast(o.message,3000,2);
            }
            $("#PullUp").html('已加载全部产品');
        },
        complete: function () {
        },
        error: function () {
        }
    });

}

function seachProdsStyle(keywords){
    var sc=$('#sc').val();
    var sh=$('#screen').val();
    $.ajax({
        url:  LOCAL_HOME+"xjcsAjax/prods/",
        data: {ajaxdata: "list", t: 2,sc:sc,sh:sh,keywords:keywords, p: pageSize, ajax: Math.random()},
        async: false,
        beforeSend: function () {
        },
        dataType: 'json',
        success: function (o) {
            if (o.status == 200) {
                $("#prolist").html(o.message);
                pageCode=o.data.totalPage;
            }
            else{
                cmCommon.toast(o.message,3000,2);
            }
            $("#PullUp").html('已加载全部产品');
        },
        complete: function () {
        },
        error: function () {
        }
    });

}
function agree(obj,id){
    $(obj).removeAttr('onclick');
    $.ajax({
        url: LOCAL_HOME+"xjcsIndex/articleDigg/",
        data: {id:id,ajax: Math.random()},
        async: false,
        beforeSend: function () {
        },
        dataType: 'json',
        success: function (o) {
            if (o.status == 4000) {
                cmCommon.toast(o.message,3000,2);
                setTimeout("window.location.href='"+o.return_url+"/?sc="+ $("#m-suid").val()+"'", 1000);
            }else if (o.status == 200) {
                var val = parseInt($("#tagnum_"+id).text());
                if(!$("#icon_act_"+id).hasClass('active')){
                    $("#icon_act_"+id).addClass('active');
                    $("#tagnum_"+id).text(val+1).css({'color':'#ff6c6c'});
                }
                else
                {
                    $("#icon_act_"+id).removeClass('active');
                    $("#tagnum_"+id).text(val-1).css({'color':''});
                }
            } else {
                cmCommon.toast(o.message,3000,2);
            }
            $(obj).attr('onclick',"agree(this,"+id+")");
        }
    });
}