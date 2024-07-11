package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/rs/cors"
)

// Post 構造体を定義
type Post struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Likes     int       `json:"likes"`
	CreatedAt time.Time `json:"created_at"`
}

// posts スライスを宣言
var posts []Post

var db *gorm.DB

func main() {
	// データベースに接続
	dsn := "root:_MySQLPassword1@tcp(localhost:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("データベースへの接続に失敗しました:", err)
	}

	// マイグレーションを実行
	db.AutoMigrate(&Post{})

	// ルートハンドラの登録
	//パスの設定
	http.HandleFunc("/posts", handlePosts)
	http.HandleFunc("/posts/", handlePost)

	// CORSの設定
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://127.0.0.1:5500"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
	})

	// ハンドラにCORSミドルウェアを適用
	handler := c.Handler(http.DefaultServeMux)

	// サーバーの起動
	log.Fatal(http.ListenAndServe(":8080", handler))
}

// handlePosts は /posts エンドポイントのリクエストを処理
func handlePosts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// 全ての投稿を取得
		getPosts(w, r)
	case http.MethodPost:
		// 新しい投稿を作成
		createPost(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// handlePost は /posts/:id エンドポイントのリクエストを処理
func handlePost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/posts/"):])
	if err != nil {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		// 特定の投稿を取得
		getPost(w, r, id)
	case http.MethodPut:
		// 特定の投稿を更新
		updatePost(w, r, id)
	case http.MethodDelete:
		// 特定の投稿を削除
		deletePost(w, r, id)
	case http.MethodPost:
		// 特定の投稿に対していいねを追加
		likesPost(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// getPosts は全ての投稿を取得
func getPosts(w http.ResponseWriter, r *http.Request) {
	result := db.Find(&posts)
	if result.Error != nil {
		http.Error(w, "クエリの実行に失敗しました", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(posts)
}

// createPost は新しい投稿を作成
func createPost(w http.ResponseWriter, r *http.Request) {
	var post Post

	//リクエストボディから新しい投稿の情報をデコード
	json.NewDecoder(r.Body).Decode(&post)

	//postIDの最大値を探す処理
	db.Find("id >?", post.ID).Last(&post)

	// postIDの最大値に1を加える
	post.ID++

	//現在の時刻を設定
	post.CreatedAt = time.Now()

	result := db.Create(&post)
	if result.Error != nil {
		http.NotFound(w, r)
		return
	}

	// 新たな投稿をpostsスライスに追加
	posts = append(posts, post)

	//作成された投稿の情報をJSONレスポンスとしてクライアントに送信
	json.NewEncoder(w).Encode(post)
}

// getPost は特定の投稿を取得
func getPost(w http.ResponseWriter, r *http.Request, id int) {
	for _, post := range posts {
		//指定されたIDに一致する投稿を取得
		if post.ID == id {
			//取得した投稿の情報を、JSONレスポンスとしてエンコードし、クライアントに送信
			json.NewEncoder(w).Encode(post)
			return
		}
	}
	//投稿が見つからない場合は、404エラーレスポンスを返す
	http.NotFound(w, r)
}

// updatePost は特定の投稿を更新
func updatePost(w http.ResponseWriter, r *http.Request, id int) {
	var post Post
	result := db.First(&post, id)
	if result.Error != nil {
		http.NotFound(w, r)
		return
	}

	var updatedPost Post
	json.NewDecoder(r.Body).Decode(&updatedPost)
	post.Title = updatedPost.Title
	post.Content = updatedPost.Content
	post.Likes = updatedPost.Likes
	result = db.Save(&post)
	if result.Error != nil {
		http.Error(w, "投稿の更新に失敗しました", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(post)
}

// deletePost は特定の投稿を削除
func deletePost(w http.ResponseWriter, r *http.Request, id int) {

	for _, post := range posts {
		//指定されたIDに一致する投稿をデータベースから取得
		if post.ID == id {
			//投稿が見つかった場合は、取得した投稿をデータベースから削除する
			db.Delete(&post)
			//投稿の削除が成功した場合は、w.WriteHeader(http.StatusNoContent) を使用して204 No Contentステータスコードを返す
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	//投稿が見つからない場合は、404エラーレスポンスを返す
	http.NotFound(w, r)
}

// likesPost は特定の投稿にいいねを付ける
func likesPost(w http.ResponseWriter, r *http.Request, id int) {
	var post Post
	result := db.First(&post, id)
	if result.Error != nil {
		http.NotFound(w, r)
		return
	}

	var likesPost Post
	json.NewDecoder(r.Body).Decode(&likesPost)
	//いいね項目の値をインクリメントする
	post.Likes++
	result = db.Save(&post)
	if result.Error != nil {
		http.Error(w, "投稿のいいねに失敗しました", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(post)
}
