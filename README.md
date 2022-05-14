## discord-bot
discord-bot

## 事前準備

discordのデベロッパーモードをONにする

参考：https://support.discord.com/hc/ja/articles/206346498-%E3%83%A6%E3%83%BC%E3%82%B6%E3%83%BC-%E3%82%B5%E3%83%BC%E3%83%90%E3%83%BC-%E3%83%A1%E3%83%83%E3%82%BB%E3%83%BC%E3%82%B8ID%E3%81%AF%E3%81%A9%E3%81%93%E3%81%A7%E8%A6%8B%E3%81%A4%E3%81%91%E3%82%89%E3%82%8C%E3%82%8B-

## Bot作成手順

1. 通知先のテキストチャンネルのIDを取得し、`.env.template`に記述
2. Bot用のアプリを作成（ https://discord.com/developers/applications ）
3. 左メニュー「bot」からBotユーザーを作成
4. Botユーザーのtokenを取得し、`.env.template`に記述
5. ブラウザで（ https://discordapp.com/oauth2/authorize?client_id=YOUR_APPLICATION_ID&scope=bot&permissions=0 ）にアクセスし、Botユーザーをチャンネルに招待
6. `.env.template` を `.env` にrenameする
7. docker-composeでアプリを起動
```shell
$ docker-compose up bot
```
※ APPLICATION_IDは、左メニュー「General Information」より取得可

