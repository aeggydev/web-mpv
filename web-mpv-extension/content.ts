/*
chrome.runtime.onMessage.addListener((message, sender, sendResponse) => {
    if (message.type === "highlight") {
        let allImages = document.querySelectorAll("ytd-thumbnail") as HTMLDivElement[]
        const enter = (x: HTMLElement) => () => {
            x.style.border = "red solid 5px"
        }
        const leave = (x: HTMLElement) => () => {
            x.style.border = "yellow solid 5px"
        }
        allImages.forEach(x => x.addEventListener("mouseover", enter(x)))
        allImages.forEach(x => x.addEventListener("mouseleave", leave(x)))
    }
})*/
