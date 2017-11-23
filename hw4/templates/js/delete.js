$(function () {
    $("input[name=pp]").click(function () {
        $.ajax({
            cache: true,
            url: "/api/delete",
            data: $("input[name='username'], input[name='password'], input[name='itemIndex']").serialize(),
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