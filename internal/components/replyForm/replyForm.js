{{ define "replyForm.js" }}
async function submitReply() {
    let formBody = document.getElementById("body").value;
    let parentPost = document.getElementById("postID").value;
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
                "parent": parentPost
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
{{end}}
