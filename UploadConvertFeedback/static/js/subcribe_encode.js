socket = new Ws("ws://" + window.location.host + "/ws_endpoint");
var totalFrame;
var progressbar = $("#progressbar"),
    progressLabel = $(".progress-label");

progressbar.progressbar({
    value: 0,
    complete: function() {
        progressLabel.text( "Complete!" );
    }
});
socket.OnConnect(function () {
    console.log("Websocket connection established" + socket.id);

    socket.Emit("Encode", videoFileName);

    socket.On("TotalFrame",function (msg) {
        totalFrame = msg;
    })

    socket.On("CurrentFrame",function (currentFrame) {
        progressbar.progressbar({
            value: (currentFrame * 100)/totalFrame
        });

        progressLabel.text(currentFrame.toString() + " / " + totalFrame.toString() );
    })


});