// Extension event listeners are a little different from the patterns you may have seen in DOM or
// Node.js APIs. The below event listener registration can be broken in to 4 distinct parts:
//
// * chrome      - the global namespace for Chrome's extension APIs
// * runtime     – the namespace of the specific API we want to use
// * onInstalled - the event we want to subscribe to
// * addListener - what we want to do with this event
//
// See https://developer.chrome.com/docs/extensions/reference/events/ for additional details.

chrome.tabs.onActivated.addListener(updateCurrentAcitivty);

async function updateCurrentAcitivty() {
  let tab = await getCurrentTab();
  const activity = {
    name: tab.title,
    website: tab.url,
  };

  fetch("http://localhost:8080/activity", {
    method: "POST",
    body: JSON.stringify(activity),
  })
    .then((response) => response.json())
    .then((data) => console.log(data))
    .catch((error) => console.log(error));
}

async function getCurrentTab() {
  let queryOptions = { active: true, lastFocusedWindow: true };

  let [tab] = await chrome.tabs.query(queryOptions);
  console.log(tab);

  return tab;
}
