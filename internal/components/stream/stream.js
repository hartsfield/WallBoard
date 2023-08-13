{{ define "stream.js" }}
async function submitReply(postID) {
    console.log(postID);
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
                "parent": postID
            }),
        });

        let res = await response.json();
        if (res.success == "true") {
            console.log("true");
        } else {
            console.log("error");
        }
    }
}
function toggleReplyForm(postID) {
    let toggleStatus = document.getElementById("reply-form_"+postID).style.display;
    console.log(toggleStatus);
    if (toggleStatus != "block") {
        document.getElementById("reply-form_"+postID).style.display = "block";
    } else {
        document.getElementById("reply-form_"+postID).style.display = "none";
    }
}
{{end}}
