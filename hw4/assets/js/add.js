$(function () {
    $("input[name=OK]").click(function () {
        $.ajax({
            cache: true,
            url: "/api/add",
            data: $("input[name='username'], input[name='password'], input[name='item']").serialize(),
            type: "POST",
            processData: false,
            contentType: "application/x-www-form-urlencoded",
            error: function (error) {
                alert("Connection error" + error);
            },
            success: function (data) {
                $('#result').text("@" + data.Username + ": " + data.Message);
            }
        });
    });
});