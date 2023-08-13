{{ define "submitForm.js" }}
async function submitPost() {
    let formBody = document.getElementById("body").value;
    console.log(formBody.length)
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
            console.log("true");
        } else {
            console.log("error");
        }
    }
}
{{end}}
