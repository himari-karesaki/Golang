// 全ての記事を取得して表示する関数 
function getPosts() {
    fetch('http://localhost:8080/posts')
    .then(response => response.json())
    .then(posts => {
      const tbody = document.querySelector('#postTable tbody');
      tbody.innerHTML = '';
      posts.forEach(post => {
        const row = document.createElement('tr');
        row.innerHTML = `
                  <td>${post.id}</td>
                  <td>${post.title}</td>
                  <td>${post.content}</td>
                  <td>${post.likes}</td>
                  <td>${new Date(post.created_at).toLocaleString()}</td>
                  <td><button class="like-btn" id="${post.id}">Like</button></td>
                  <td><button class="delete-btn" id="${post.id}">Delete</button></td>`;
        tbody.appendChild(row);
});

// いいねボタンのクリックイベントを追加
const likeButtons = document.querySelectorAll('.like-btn'); likeButtons.forEach(button => {
        button.addEventListener('click', () => {
          const postId = button.id;
          console.log(button.id);
          likePost(postId);
}); });
// }); }

// 削除ボタンのクリックイベントを追加
const deleteButtons = document.querySelectorAll('.delete-btn'); deleteButtons.forEach(button => {
  button.addEventListener('click', () => {
    const postId = button.id;
    console.log(button.id);
    deletePost(postId);
    location.reload()
}); });
}); }

// 新しい記事を追加する関数
function addPost(title, content) {
  const post = {
    title,
  content };
  
  fetch('http://localhost:8080/posts', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(post)
})
    .then(() => {
getPosts();
      document.querySelector('#postForm').reset();
    });
}

//特定の投稿を更新する関数
function updatePost(id,title, content) {
  const postID = id;
  const post = {
    id,
    title,
content };
  fetch(`http://localhost:8080/posts/${postID}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(post)
})
    .then(() => {
getPosts();
      document.querySelector('#updateForm').reset();
    });
}

// 記事にいいねを追加する関数 
function likePost(postId) {
  fetch(`http://localhost:8080/posts/${postId}`, {
    method: 'POST',
  })
    .then(() => {
      getPosts();
    });
}
// 記事を削除する関数 
function deletePost(postId) {
  fetch(`http://localhost:8080/posts/${postId}`, {
    method: 'DELETE',
  }
  )
}
// ページ読み込み時に全ての記事を取得して表示 
document.addEventListener('DOMContentLoaded', getPosts);

// 新しい記事を追加するフォームの送信イベント 
document.querySelector('#postForm').addEventListener('submit', event => {
  event.preventDefault();
  const title = document.querySelector('#title').value;
  const content = document.querySelector('#content').value;
  addPost(title, content);
});

//特定の投稿を更新するフォームの送信イベント
document.querySelector('#updateForm').addEventListener('submit', event => {
  event.preventDefault();
  const id = document.querySelector('#updateId').value;
  const title = document.querySelector('#updateTitle').value;
  const content = document.querySelector('#updateContent').value;
  updatePost(id,title, content);
})

// 特定の記事を取得して表示する関数
function getPost(id) {
  
  const getID = id;

  fetch(`http://localhost:8080/posts/${getID}`)
  .then(response => response.json())
  .then(post => {
    const tbody = document.querySelector('#getTable tbody');
    // tbody.innerHTML = '';
      const row = document.createElement('tr');
      row.innerHTML = `
                <td>${post.id}</td>
                <td>${post.title}</td>
                <td>${post.content}</td>
                <td>${post.likes}</td>
                <td>${new Date(post.created_at).toLocaleString()}</td>`;
                tbody.appendChild(row);
});
}; 
// 特定の記事を取得するフォームの送信イベント
document.querySelector('#getForm').addEventListener('submit', event => {
  event.preventDefault();
  const id = document.querySelector('#getId').value;
  console.log(id);
  getPost(id);
})
