# platform

<a href="https://codecov.io/gh/morning-night-guild/platform" >
<img src="https://codecov.io/gh/morning-night-guild/platform/branch/main/graph/badge.svg?token=OZM65W4G5Q"/>
</a>

## document

- [api](https://github.com/morning-night-guild/platform/tree/gh-pages/api)
- [database](https://github.com/morning-night-guild/platform/tree/gh-pages/database)

## directory

```bash
├── README.md
├── .docker     // ローカル環境構築
├── .github     // CICD
├── api         // API定義
├── backend     // バックエンドアプリ
├── frontend    // フロントエンドアプリ
└── infra       // インフラストラクチャー
```

## commit message prefix

実装をする際にはissueを作成し番号をコミットメッセージに含めること。

| PREFIX           | 意味                                                          |
| ---------------- | ------------------------------------------------------------ |
| **feat(#x)**     | 新機能                                                        |
| **fix(#x)**      | バグ修正                                                      |
| **docs(#x)**     | ドキュメントの修正                                              |
| **style(#x)**    | コードの意味に影響を与えない変更（空白、書式設定、セミコロンの欠落など） |
| **refactor(#x)** | バグの修正でも機能の追加でもないコードの変更                         |
| **perf(#x)**     | パフォーマンスを向上させるコード変更                               |
| **test(#x)**     | 不足しているテストの追加や既存のテストの修正                         |
| **chore(#x)**    | ビルドプロセスやドキュメント生成などの補助ツールやライブラリの変更      |
