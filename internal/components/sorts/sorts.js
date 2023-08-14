{{ define "sorts.js" }}
function getStream(sortOrder) {
    window.location = window.location.origin + "/" + sortOrder;
}
// async function getStream(sortOrder) {
//     const response = await fetch("/api/getStream", {
//         method: "POST",
//         headers: { "Content-Type": "application/json" },
//         body: JSON.stringify({special:view}),
//     });

//     let res = await response.json();
//     if (res.success == "true") {
//         // do stuff
//     } else {
//         console.log("error");
//     }
// }
{{end}}
