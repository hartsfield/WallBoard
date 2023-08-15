{{ define "stream.js" }}
async function submitReply(postID) {
    let formBody = document.getElementById("body_"+postID).value;
    if (formBody.length < 5) {
        document.getElementById("errorField_"+postID).innerHTML = "too short";
    } else if (formBody.length > 1000) {
        document.getElementById("errorField_"+postID).innerHTML = "too long";
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
            console.log(postID);
            let submitForm = document.getElementById("reply-form_"+postID);
            submitForm.remove();
            window.location = window.location.origin + "/post/" + postID;
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
function isElementInViewport (el) {
    // Special bonus for those using jQuery
    if (typeof jQuery === "function" && el instanceof jQuery) {
        el = el[0];
    }

    var rect = el.getBoundingClientRect();

    return (
        rect.top >= 0 &&
        rect.left >= 0 &&
        rect.bottom <= (window.innerHeight || document.documentElement.clientHeight) && /* or $(window).height() */
        rect.right <= (window.innerWidth || document.documentElement.clientWidth) /* or $(window).width() */
    );
}

{{end}}
