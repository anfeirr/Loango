(function(win,doc){
    var common = {
        toastOn : false,//标识当前toast是否被移除，默认值false：被移除
        /**
         * 创建一个toast提示层，当前文档中的toast未被移除前不会创建
         * @param text 提示文本
         * @param type 成功或失败提示，1表示成功，2表示失败，不传默认为成功
         * @param ms toast层显示时间，单位：ms，默认3000
         */
        toast : function(text,ms,type){
            var _that = this;
            if(!text || _that.toastOn){
                return;
            }
            _that.toastOn = true;
            (function(){
                ms = ms || 3000;
                type = (type || 1)+'';
                var _toast_box_ele = document.createElement('DIV');
                _toast_box_ele.setAttribute('class','global-toast-box');
                if(type == '2'){
                    _toast_box_ele.innerHTML = '<p class="icon-error"></p><p>'+text+'</p>';
                }else{
                    _toast_box_ele.innerHTML = '<p class="icon-success"></p><p>'+text+'</p>';
                }
                doc.body.appendChild(_toast_box_ele);
                setTimeout(function(){
                    try{
                        doc.body.removeChild(_toast_box_ele);
                    }catch(e){
                        console.warn('Warn: the toast element is not exists.');
                    }
                    _that.toastOn = false;
                },ms);
            })();

        }
    };
    win.cmCommon = common;
})(window,document);
