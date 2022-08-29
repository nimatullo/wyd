// Chrome API Functions

chrome.runtime.onInstalled.addListener(async () => {
  await askForMasterPassword();
  chrome.storage.sync.set({ logging: false }); // default to false
});

chrome.webNavigation.onCompleted.addListener(updateCurrentAcitivty);

chrome.tabs.onActivated.addListener(updateCurrentAcitivty);

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

// Helper Functions
async function askForMasterPassword() {
  let url = chrome.runtime.getURL("../html/onboarding.html");

  await chrome.tabs.create({ url });
}

async function updateCurrentAcitivty() {
  await chrome.storage.sync.get("logging", async function (data) {
    if (!data.logging) {
      return;
    } else {
      const tab = await getCurrentTab();

      if (isTabBlacklisted(tab)) return;

      const activity = {
        name: tab.title,
        website: tab.url,
      };

      fetch("https://wyd.nimatullo.com/activity", {
        method: "POST",
        body: JSON.stringify(activity),
      })
        .then((response) => response.json())
        .then((data) => console.log(data))
        .catch((error) => console.log(error));
    }
  });
}

async function getCurrentTab() {
  let queryOptions = { active: true };

  let [tab] = await chrome.tabs.query(queryOptions);
  return tab;
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
