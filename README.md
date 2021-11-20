目的

goのHelloWorldをHerokuでデプロイする

やり方

1.go htmlをいつも通り作成する
注：main.goのポートはHeroku側が勝手に降ってくるのでそれに対応する形に書き換える必要がある。
2.go mod init

3.heroku.ymlを作成
このファイルがないとgit push heroku mainでエラーになる
Dockerを使用しない場合は公式では採用されていない。
ただこのdeploypractice_phase1の書き方ならエラーにならない。
書き方は参照すること

4.git push heroku main

5.heroku open

これでできた。

Dockerを使用する場合は、heroku.ymlを修正すれば良さそう。
ただ一緒にDockerfileも編集が必要かと。
なので模写することから始めていこう。
