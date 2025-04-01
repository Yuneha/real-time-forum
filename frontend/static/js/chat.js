import { Throttle } from "./functions.js";

export function Chat(logUser, ws) {
  const content = document.querySelector(".content");
  content.id = "dm";
  content.innerHTML = `<h1>DM</h1>
  <div class="chat_content">
      <div class="chat_users">
      <ul class="users_list_header">
        <li>Users</li><br>
      </ul>
      <input type="text" placeholder="Search User" id="search_user"></input>
        <ul class="users_list"></ul>
      </div>
      <div class="chat_window">
      <span class="window_alert">Click on a user to start chatting !</span>
        <div class="chat_window_content">
          <div class="chat_window_header">
            <span class="username">Username</span>
          </div>
          <div class="chat_dialogue"></div>

          <div id="typingIndicator"></div>

          <div class="chat_action">
            <input type="text" placeholder="Enter a message" id="message" rows="1"></input>
            <input type="button" value="SEND" id="send_message"></input>
          </div>
        </div> 
      </div>
    </div>`;

  const usersList = document.querySelector(".users_list");
  const chatWindowContent = document.querySelector(".chat_window_content");
  const windowAlert = document.querySelector(".window_alert");
  const chatDialogue = document.querySelector(".chat_dialogue");
  const messageInput = document.getElementById("message");
  const sendMessage = document.getElementById("send_message");
  const notifications = document.getElementById("notification");
  let recipient;

  ws.onmessage = function (event) {
    const response = JSON.parse(event.data);

    if (response.Type && response.Type === "notif") {
      renderNotifications(response);
      notifications.innerHTML = `Notifs(${response.Count})`;
    }

    if (response.Type === "UserList") {
      renderUsers(response.UserList);
    }

    if (response.recipient === recipient || response.sender === recipient) {
      renderMessage(response, logUser);
    }

    if (response.Type === "typing") {
      updateTypingIndicator(response.Status, response.Sender);
    }

    chatDialogue.scrollTop = chatDialogue.scrollHeight;
  };

  async function fetchUsers() {
    try {
      const response = await fetch("/getUsers", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(logUser),
      });
      const users = await response.json();
      return users;
    } catch (error) {
      console.log("Error fetching users:", error);
    }
  }
  fetchUsers().then((data) => {
    renderUsers(data);
  });

  async function fetchNotifications() {
    try {
      const response = await fetch("/notification", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(logUser),
      });
      const unreadMessageCount = response.json();
      return unreadMessageCount;
    } catch (error) {
      console.log("Error fetching users:", error);
    }
  }

  let offset = 0;
  const limit = 10;

  async function fetchMessages(sender, recipient, offset, limit) {
    try {
      const response = await fetch(
        `/getMessages?sender=${sender}&recipient=${recipient}&offset=${offset}&limit=${limit}`
      );
      const messages = await response.json();
      return messages;
    } catch (error) {
      console.log("Error fetching messages:", error);
    }
  }

  function renderUsers(users) {
    usersList.innerHTML = "";
    users.forEach((user) => {
      if (user.username !== logUser.username) {
        const userElement = document.createElement("li");
        userElement.classList.add("chat_user");
        userElement.id = user.username;
        userElement.innerHTML = user.username;
        if (user.is_connected === 0) {
          userElement.style.color = "gray";
        }
        const notifspan = document.createElement("span");
        notifspan.id = user.username + "Notif";
        usersList.appendChild(userElement);
        userElement.appendChild(notifspan);
        userElement.addEventListener("click", () => {
          openChatWindow(user);
        });
      }
    });
    fetchNotifications().then((data) => {
      renderNotifications(data);
    });
  }

  function renderMessage(response, logUser, append) {
    if (
      response.read_status === false &&
      logUser.username === response.recipient
    ) {
      ws.send(
        JSON.stringify({ type: "mark_read", sender: `${response.sender}` })
      );
      ws.send(
        JSON.stringify({
          type: "notification",
          recipient: `${response.recipient}`,
        })
      );
    }
    const dateObject = new Date(response.timestamp);
    const date = dateObject.toLocaleDateString();
    const time = dateObject.toLocaleTimeString();
    const item = document.createElement("div");
    const itemNested = document.createElement("div");
    const dialogueContent = document.createElement("div");
    dialogueContent.style.position = "relative";
    dialogueContent.style.width = "100%";
    dialogueContent.style.height = "100%";
    const span = document.createElement("span");
    const dateAndTimeSpan = document.createElement("span");
    if (response.sender === logUser.username) {
      item.classList.add("bubble_right");
      itemNested.classList.add("bubble_dialogue_right");
    } else {
      item.classList.add("bubble_left");
      itemNested.classList.add("bubble_dialogue_left");
    }
    span.innerHTML = `${response.message}`;
    dateAndTimeSpan.innerHTML = `${date}, ${time}`;
    dialogueContent.appendChild(span);
    const line = document.createElement("br");
    dialogueContent.prepend(line);
    dialogueContent.prepend(dateAndTimeSpan);
    itemNested.appendChild(dialogueContent);
    item.appendChild(itemNested);
    chatDialogue.prepend(item);

    if (append) {
      chatDialogue.prepend(item); // Add older messages at the top
    } else if (offset >= 10) {
      chatDialogue.appendChild(item); // Add recent messages at the bottom
    }
  }

  function renderMessages(messages, logUser, append = false) {
    if (!append) {
      chatDialogue.innerHTML = "";
    }
    let preHeight = chatDialogue.scrollHeight;

    if (messages != null) {
      messages.forEach((response) => {
        renderMessage(response, logUser, append);
      });
    }

    if (!append) {
      chatDialogue.scrollTop = chatDialogue.scrollHeight; // Scroll to bottom for initial load
    } else {
      chatDialogue.scrollTop = chatDialogue.scrollHeight - preHeight - 400; // Scroll to loadded messages
    }
  }

  function renderNotifications(notif) {
    for (const key in notif.Notification) {
      const element = notif.Notification[key];
      const userNotif = document.getElementById(key + "Notif");
      userNotif.innerHTML = element;
    }
  }

  function openChatWindow(user) {
    chatWindowContent.style.display = "block";
    windowAlert.style.display = "none";
    document.querySelector(".username").innerHTML = user.username;
    recipient = user.username;
    const recipientNotif = document.getElementById(recipient + "Notif");
    recipientNotif.innerHTML = "";

    // Reset offset when opening a new chat
    offset = 0;

    fetchMessages(logUser.username, recipient, offset, limit).then(
      (messages) => {
        renderMessages(messages, logUser);
        offset += limit;
      }
    );

    let typingTimer;

    messageInput.addEventListener("input", () => {
      clearTimeout(typingTimer);
      sendTypingStatus("started");

      typingTimer = setTimeout(() => {
        sendTypingStatus("stopped");
      }, 3000);
    });

    function sendTypingStatus(status) {
      ws.send(
        JSON.stringify({
          sender: logUser.username,
          recipient: recipient,
          type: "typing",
          status: status,
        })
      );
    }

    sendMessage.addEventListener("click", () => {
      writeMessage(logUser, messageInput, recipient, ws);
    });

    messageInput.addEventListener("keypress", function (event) {
      if (event.key === "Enter" && !event.shiftKey) {
        event.preventDefault();
        writeMessage(logUser, messageInput, recipient, ws);
      }
    });
  }

  setupScrollListener();

  function writeMessage(logUser, messageInput, recipient, ws) {
    const message = {
      type: "message",
      sender: logUser.username,
      message: messageInput.value,
      recipient: recipient,
      timestamp: new Date().toISOString(),
    };
    if (message.message != "") {
      ws.send(JSON.stringify(message));
      ws.send(
        JSON.stringify({
          type: "updateUserList",
          sender: `${message.sender}`,
          recipient: `${message.recipient}`,
        })
      );
      ws.send(
        JSON.stringify({
          type: "notification",
          recipient: `${message.recipient}`,
        })
      );
      messageInput.value = "";
    }
  }

  function updateTypingIndicator(status, sender) {
    if (sender === recipient) {
      // Update UI to show/hide typing indicator
      const typingIndicator = document.getElementById("typingIndicator");
      if (status === "started") {
        typingIndicator.textContent = `${sender} is typing ...`;
      } else {
        typingIndicator.textContent = "";
      }
    }
  }

  function setupScrollListener() {
    chatDialogue.addEventListener(
      "scroll",
      Throttle(function () {
        if (chatDialogue.scrollTop === 0) {
          fetchMessages(logUser.username, recipient, offset, limit).then(
            (messages) => {
              if (messages && messages.length > 0) {
                renderMessages(messages, logUser, true); // Append older messages at the top
                offset += limit;
              }
            }
          );
        }
      }, 2000)
    );
  }
}
