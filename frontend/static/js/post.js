import { FormatDate } from "./functions.js";

export function Post(post, user) {
  const content = document.querySelector(".content");
  content.id = "post_page";
  content.innerHTML = `<div class="post_page_content">
    <div class="post_info_container">
      <h1 class="post_title"></h1>
      <div class="post_info">
        <span class="post_username"></span>
        <span class="post_publication_date"></span>
      </div>
      <textarea name="post_message" class="post_message" id="post_message" cols="60" rows="20" disabled></textarea>
      <div class="post_action">
        <span class="like_post">Like</span> 0
        <span class="dislike_post">Dislike</span> 0
      </div>
    </div>

    <div class="comment_section">
      <h1>Comment(s)</h1>
        <div class="create_comment">
          <textarea name="create_comment_message" class="create_comment_message" id="create_comment_message" cols="80" rows="9"></textarea>
          <input type="button" class="send_comment" value="SEND"></input>
        </div>

        <div class="all_comments">
          <template class="comment_template" id="comment_template">
            <div class="comment_card">
              <div class="comment_info">
                <span class="comment_username"></span>
                <span class="comment_publication_date"></span>
              </div>
              <textarea name="comment_message" class="comment_message" id="comment_message" cols="70" rows="1" disabled></textarea>
            </div>  
          </template>
        </div>     
    </div>
  </div>`;

  const sendComment = document.querySelector(".send_comment");
  const postTitle = document.querySelector(".post_title");
  const postUsername = document.querySelector(".post_username");
  const postPublicationDate = document.querySelector(".post_publication_date");
  const postMessage = document.querySelector(".post_message");
  const createCommentMessage = document.querySelector(
    ".create_comment_message"
  );
  const commentsSection = document.querySelector(".all_comments");
  const commentTemplate = document.getElementById("comment_template").content;

  postTitle.textContent = post.title;
  postUsername.textContent = user.username;
  postPublicationDate.textContent = FormatDate(post.publication_date).split(
    " "
  )[0];
  postMessage.textContent = post.message;

  async function fetchComments() {
    try {
      const response = await fetch("/comments", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(post),
      });
      const comments = await response.json();
      return comments;
    } catch (error) {
      console.log("Error fetching comments:", error);
    }
  }

  fetchComments().then((data) => {
    const allComments = data.comments;
    const users = data.users;
    if (allComments && allComments.length > 0) {
      for (let index = allComments.length - 1; index >= 0; index--) {
        const comment = allComments[index];
        const commentUser = users[index];
        createComment(comment, commentUser);
      }
    }
  });

  function createComment(comment, user) {
    const commentElement = document.importNode(commentTemplate, true);
    commentElement.querySelector(".comment_username").textContent =
      user.username;
    commentElement.querySelector(".comment_publication_date").textContent =
      comment.publication_date.split(" ")[0];
    commentElement.querySelector(".comment_message").textContent =
      comment.message;
    commentsSection.appendChild(commentElement);
  }

  sendComment.addEventListener("click", () => {
    const commentObject = {
      message: createCommentMessage.value,
    };
    fetch("http://localhost:5656/createComment", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ user: user, post: post, comment: commentObject }),
    })
      .then((response) => {
        return response.json();
      })
      .then(() => {
        Post(post, user);
      })
      .catch((error) => {
        console.error("fetch issue:", error);
      });
  });
}
