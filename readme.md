## Go-todo-app

### 使用技術
- Go v1.22.4
- sqlite3

### 機能
- ログイン
- ログアウト
- セッション
- アクセス制限
  - ログインしているとlogin画面　siginup画面が表示されない　
  - ログアウトではtodo一覧ページが表示されない　

- TodoのCRUD機能


#### 実行手順

```
% git clone https://github.com/yuta82644/go-todo_app.git
% cd go-todo_app
% go mod init github.com/yuta82644/go-todo_app
% go mod tidy
% go run main.go

ブラウザ　http://localhost:8080/

```
