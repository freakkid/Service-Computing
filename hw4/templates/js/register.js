$(function () {
    $("input[name=pp]").click(function () {
        $.ajax({
            cache: true,
            url: "/api/register",
            data: $("input[name='username'], input[name='password']").serialize(),
            type: "POST",
            processData: false,
            contentType: "application/x-www-form-urlencoded",
            error: function (error) {
                alert("Connection error" + error);
            },
            success: function (data) {
                $('#result').text("@" + data.username + ": " + data.message);
            }
        });
    });
});