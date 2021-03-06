const root_url = "http://localhost:6969"

chrome.runtime.onInstalled.addListener(() => {
    chrome.contextMenus.create({
        "id": "aeggy.web-mpv",
        "title": "Play with MPV",
        "contexts": ["page", "link", "video", "audio"]
    })
})

chrome.browserAction.onClicked.addListener(() => {
    chrome.tabs.query({active: true, currentWindow: true}, tabs => {
        // @ts-ignore
        chrome.tabs.sendMessage(tabs[0].id, {type: "highlight"}, function (response) {
        })
    })
})

async function playUrl(url: string) {
    await fetch(`${root_url}/play?url=${url}`)
}

chrome.contextMenus.onClicked.addListener(function(info, tab) {
    playUrl(info["linkUrl"] || info["srcUrl"]|| info["pageUrl"]) // TODO: Support for pausing video
})