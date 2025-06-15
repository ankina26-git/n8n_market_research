// Package searchconsole は、SEOパフォーマンス分析に関する
// ドメインロジックおよびデータ構造を管理するパッケージです。
//
// このパッケージは、バージョンごとにサブディレクトリ（v1, v2など）で構成されており、
// 各バージョンごとに独自のルーター・ハンドラー・サービスを実装しています。
//
// モデル（models.go）はこのパッケージ直下に定義され、基本的にはバージョン間で共有されますが、
// 変更が必要な場合は明示的に別バージョン用モデルを分けて管理してください。
//
// 利用方法の例：
//
//	import searchconsole_v1 "n8n_project_go/app/searchconsole/v1"
//	searchconsole_v1.RegisterRoutes(app.Group("/api/v1/searchconsole"))
//
// 他ドメイン（例：UserやAuth）から直接このパッケージを参照することは避け、
// バージョン付きのAPIルート経由で利用してください。
package searchconsole
