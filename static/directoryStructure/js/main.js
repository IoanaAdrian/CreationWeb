function showOverlay() {
    let overlay = document.getElementById("mkDirOverlay")

    if (overlay.style.display === "none") {
        overlay.style.display = "flex";

    } else {
        overlay.style.display = "none";
    }

    document.getElementById("folderName").select()
}