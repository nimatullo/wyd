function savePassword() {
  const password = document.getElementById("password").value;

  chrome.storage.sync.set(
    {
      password: password,
    },
    function () {
      chrome.tabs.query({ active: true, currentWindow: true }, function (tabs) {
        chrome.tabs.remove(tabs[0].id);
      });
    }
  );
}

document
  .getElementById("password-button")
  .addEventListener("click", savePassword);
