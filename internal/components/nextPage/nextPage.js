{{ define "nextPage.js" }}
let nextpagerButt = document.getElementById("nextPage");

let requestMade = false;
document.addEventListener("scroll", (event) => {
    if (isElementInViewport(nextpagerButt) && !requestMade) {
        requestMade = true;
        setTimeout(() => {
            console.log("inview");
            submitNext();
        }, 1000);
    }
});
let count = 20;
let lastPage = false;
async function submitNext() {
    if (!lastPage) {
        let postsWrapper = document.getElementById("postsWrapper")
        const response = await fetch("/chron?count=" + count, {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: "",
        });

        console.log("test:");
        let res = await response.json();
        if (res.success == "true") {
            postsWrapper.insertAdjacentHTML("beforeend", res.template);
            requestMade = false;
            if (res.count != "None") {
                count = parseInt(res.count);
            } else {
                lastPage = true;
                document.getElementById("nextPage").innerHTML = "no more posts";
                document.getElementById("nextPage").style.fontSize = "1em";
                document.getElementById("nextPage").style.animationIterationCount = "1";
            }
        } else {
            console.log("error");
        }
    }
}


{{end}}
