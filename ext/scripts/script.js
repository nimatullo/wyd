async function onButtonClick() {
  showLoader();

  const tab = await getCurrentTab();

  // Get master password from storage and then tell background script to make a request to nimatullo.com
  chrome.storage.sync.get("password", function (data) {
    chrome.runtime.sendMessage(
      { type: "add-link", url: tab.url, password: data.password },
      (response) => setStatusMessage(response)
    );
  });
}

function showLoader() {
  const hoarderButton = document.getElementById("hoarder-button");

  hoarderButton.classList.add("loading");

  setTimeout(() => {
    hoarderButton.classList.remove("loading");
  }, 3000);
}

function hideLoader() {
  const hoarderButton = document.getElementById("hoarder-button");

  hoarderButton.innerHTML = "Add to Hoarder";
  hoarderButton.disabled = false;
}

function setStatusMessage(response) {
  hideLoader();
  const status = document.getElementById("status");

  status.innerText = response.message;
  status.classList.add(response.status);

  setTimeout(() => {
    status.innerText = "";
    status.classList.remove(response.status);
  }, 5000);
}

async function getCurrentTab() {
  let queryOptions = { active: true };

  let [tab] = await chrome.tabs.query(queryOptions);

  return tab;
}

document
  .getElementById("hoarder-button")
  .addEventListener("click", onButtonClick);
