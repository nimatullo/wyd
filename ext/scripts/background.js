chrome.runtime.onInstalled.addListener(async () => {
  await askForMasterPassword();
});

chrome.tabs.onActivated.addListener(updateCurrentAcitivty);

async function askForMasterPassword() {
  let url = chrome.runtime.getURL("../html/hello.html");

  await chrome.tabs.create({ url });
}

async function updateCurrentAcitivty() {
  let tab = await getCurrentTab();

  if (isTabBlacklisted(tab)) return;

  const activity = {
    name: tab.title,
    website: tab.url,
  };

  fetch("https://p-wyd.herokuapp.com/activity", {
    method: "POST",
    body: JSON.stringify(activity),
  })
    .then((response) => response.json())
    .then((data) => {
      chrome.storage.sync.set({ lastUpdate: tab.url });
    })
    .catch((error) => console.log(error));
}

function isTabBlacklisted(tab) {
  const blacklistedWebsites = [
    "chrome://",
    "chrome-extension://",
    "chrome-devtools://",
    "about:blank",
    "about:newtab",
    "https://nimatullo.com",
  ];

  return blacklistedWebsites.some((website) => tab.url.includes(website));
}

async function getCurrentTab() {
  let queryOptions = { active: true };

  let [tab] = await chrome.tabs.query(queryOptions);
  console.log(tab);

  return tab;
}

chrome.runtime.onMessage.addListener((msg, sender, response) => {
  if (msg.type === "add-link") {
    const params = new URLSearchParams({
      u: msg.url,
      p: msg.password,
    });

    fetch("https://nimatullo.com/api/a?" + params)
      .then((res) => {
        if (res.status !== 200) {
          res.json().then((data) => {
            response({ status: "error", message: data.message });
          });
        } else {
          res.json().then((data) => {
            response({ status: "success", message: data.message });
          });
        }
      })
      .catch((error) => response({ error: error }));
  }

  return true;
});
