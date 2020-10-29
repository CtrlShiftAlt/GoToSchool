window.addEventListener("load", function (evt) {
    var token;
    $("#login").submit(function (e) {
        e.preventDefault();
        username = $("#login #username").val()
        password = $("#login #password").val()
        $.ajax({
            "url": "/home/signin",
            "type": "POST",
            "async": false,
            "data": {
                "username": username,
                "password": password,
            },
            success: function (data) {
                if (data.message == "OK") {
                    token = data.data.token
                    return
                } else {
                    alert(data.message)
                }
            },
        })
    });
    $("#register").submit(function (e) {
        e.preventDefault();
        username = $("#register #username").val()
        password = $("#register #password").val()
        $.ajax({
            "url": "/home/signup",
            "type": "POST",
            "async": false,
            "data": {
                "username": username,
                "password": password,
            },
            success: function (data) {
                if (data.message == "OK") {
                    return
                } else {
                    alert(data.message)
                }
            },
        })
    });
    // HTML .
    var input = document.getElementById("input");
    var output = document.getElementById("output");
    var print = function (message) {
        var d = document.createElement("div");
        d.textContent = message;
        output.appendChild(d);
    };
    // websocket .
    var lockReconnect = false; // 避免重复连接
    var ws = null;
    var run = function () {
        var ws_url = "ws://localhost:8080/ws?token=" + token;
        try {
            if ('WebSocket' in window) {
                ws = new WebSocket(ws_url);
            } else if ('MozWebSocket' in window) {
                ws = new MozWebSocket(ws_url);
            } else {
                ws = new SocketJS(ws_url);
            }
            initEventHandle();
        } catch (e) {
            reconnect();
        }
        return false;
    };
    var reconnect = function () {
        if (lockReconnect) return;
        lockReconnect = true;
        setTimeout(function () {
            run();
            lockReconnect = false;
        }, 2000);
    };
    // 心跳机制
    var heartCheck = {
        timeout: 10000,
        timeoutObj: null,
        serverTimeoutObj: null,
        reset: function () {
            clearTimeout(this.timeoutObj);
            clearTimeout(this.serverTimeoutObj);
            return this;
        },
        start: function () {
            var self = this;
            this.timeoutObj = setTimeout(function () {
                ws.send('{"event":"ping"}');
                print('{"event":"ping"}')
                self.serverTimeoutObj = setTimeout(function () {
                    print("js CLOSE");
                    ws.close();
                }, self.timeout);
            }, this.timeout);
        }
    }

    var initEventHandle = function () {
        ws.onopen = function () {
            print("OPEN");
            heartCheck.reset().start();
        };
        ws.onmessage = function (evt) {
            print("RESPONSE: " + evt.data);
            heartCheck.reset().start();
        };
        ws.onclose = function (evt) {
            print("CLOSE");
            reconnect();
        };
        ws.onerror = function (evt) {
            print("ERROR: " + evt.data);
            reconnect();
        };
    };

    document.getElementById("open").onclick = run;

    document.getElementById("send").onclick = function (evt) {
        if (!ws) return false;
        print("SEND: " + input.value);
        ws.send(input.value);
        return false;
    };

    document.getElementById("close").onclick = function (evt) {
        if (!ws) return false;
        ws.close();
        return false;
    };
});