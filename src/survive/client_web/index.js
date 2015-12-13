
var decoder = new TextDecoder('utf-8'),
    msgSeed=10000;
function Message(typeName,msg){
    var self = this;//save the this ref

    self.typeName = typeName;
    self.msg = msg;
    self.desc = [typeName,":",msg].join('')
}

function ViewModel(){
    var self = this;//save the this ref

    self.toSend = ko.observable("");
    self.messages = ko.observableArray([]);
    self.ws=undefined;
    self.init();
}
ViewModel.prototype.init = function(){
    var self = this;//save the this ref

    self.ws  = new WebSocket("ws://127.0.0.1:13563");
    self.ws.binaryType = 'arraybuffer';
    self.ws.onopen = function (event) {
        self.addMessage("connect","connect success!");
    };
    self.ws.onmessage = function (event) {
        var data = decoder.decode(event.data);
        self.addMessage("receive",data);
    }
}

// ArrayBuffer转为字符串，参数为ArrayBuffer对象
function ab2str(buf) {
    return String.fromCharCode.apply(null, new Uint16Array(buf));
}

// 字符串转为ArrayBuffer对象，参数为字符串
function str2ab(str) {
    var buf = new ArrayBuffer(str.length+2); // 每个字符占用1个字节,2字节的长度
    //var bufView = new Uint16Array(buf);
    var bufDataView = new DataView(buf)
    bufDataView.setUint16(0,str.length);//使用大端序设置数据包长度(和leafserver约定的)
    for (var i = 1, strLen = str.length; i < strLen; i++) {
        bufDataView.setUint8(1+i,str.charCodeAt(i));//测试消息里都使用英文，则utf-16编码码制和utf-8一样，只需要1个字节
    }
    return buf;
}

//根据与后端约定的消息格式，产生2进制消息
function makeMessageBuffer(text){
    return str2ab(text);
}
ViewModel.prototype.send = function(text){
    var self = this;//save the this ref
    //var msg =makeMessageBuffer(text);
    var msg =text;
    self.ws.send(msg);
}
ViewModel.prototype.sendClick = function(){
    var self = this;//save the this ref

    var sendText = self.toSend();
    self.send(ko.toJSON({Login:{UserName:sendText,MsgId:(msgSeed++).toString()}}));
    self.addMessage("send",sendText);
}
ViewModel.prototype.addMessage = function(typeName,msg){
    var self = this;//save the this ref

    var item = new Message(typeName,msg);
    self.messages.push(item);
}

var vm = new ViewModel();
ko.applyBindings(vm,document.getElementById("root"));