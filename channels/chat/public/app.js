window.addEventListener("DOMContentLoaded", (_) => {
  let websocket = new WebSocket("ws://" + window.location.host + "/websocket");
  websocket.onopen = () => {
    console.log("connected");
  };
  websocket.onerror = (err) => {
    console.error("failed to connect", error);
  };
  let room = document.getElementById("chat-text");

  websocket.onmessage = (e) => {
    // let data = JSON.parse(e.data);
    console.log("received", e.data);
    let li = document.createElement("li");

    li.innerText = e.data;
    document.querySelector("#chat-text").append(li);
    //room.scrollTop = room.scrollHeight; // Auto scroll to the bottom
  };

  let form = document.getElementById("input-form");
  form.addEventListener("submit", event => {
    event.preventDefault();
    let username = document.querySelector("#input-username").value;
    let text = document.querySelector("#input-text").value;
    websocket.send(username, text);
    document.querySelector("#input-username").value = "";
    document.querySelector("#input-text").value = "";
  });
});
