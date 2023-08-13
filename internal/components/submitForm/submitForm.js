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
            let submitForm = document.getElementById("section-submitForm");
            submitForm.remove();
            let prev = document.getElementById("section-stream").innerHTML;
            // let insert = document.createElement(
            let newelm = 
`<div class="section-outer postElement-outer section-postElement" id="section-postElement">
    <div class="post">
        <div class="body-outer">
            <div class="bodytext">` + formBody + `</div>
        </div>
    </div>
</div>`
            document.getElementById("section-stream").innerHTML = newelm + prev;
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
{{end}}
