$("#SubmitTask").click(
    function() {

        task = {"Task" : $("#txtnum").val()};
        $.ajax({
            url: '/task',
            type: 'post',
            dataType: 'json',
            success: successCallBack,
            data: JSON.stringify(task)
        });
    }


);
function successCallBack() {
    console.log("Submit done");
}