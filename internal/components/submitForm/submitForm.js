{{ define "submitForm.js" }}
async function submitPost() {
    let formBody = document.getElementById("body").value;
    if (formBody.length < 5) {
        document.getElementById("errorField").innerHTML = "too short";
    } else if (formBody.length > 1000) {
        document.getElementById("errorField").innerHTML = "too long";
    } else {
        const response = await fetch("/api/submitForm", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({
                "title": "nil value",
                "bodytext": formBody,
                "parent": "root"
            }),
        });

        let res = await response.json();
        if (res.success == "true") {
            togglePostForm();
            window.location = window.location.origin + "/post/" + res.replyID;
        } else {
            console.log("error");
        }
    }
}

function testText() {
    let textfield = document.getElementById("body");
    if (textfield.value.length > 5) {

    }
}

function convertTime(unix_timestamp) {
    // Create a new JavaScript Date object based on the timestamp
    // multiplied by 1000 so that the argument is in milliseconds, not seconds.
    var date = new Date(unix_timestamp * 1000);
    // Hours part from the timestamp
    var hours = date.getHours();
    // Minutes part from the timestamp
    var minutes = "0" + date.getMinutes();
    // Seconds part from the timestamp
    var seconds = "0" + date.getSeconds();

    // Will display time in 10:30:23 format
    var formattedTime = hours + ':' + minutes.substr(-2) + ':' + seconds.substr(-2);

    console.log(formattedTime);
}
{{end}}
