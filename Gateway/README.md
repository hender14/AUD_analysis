# User manual

## Output in Markdown format

引数の -i はSwagger Editorで保存したYAMLファイル、 -f はコンバート後のファイル名（拡張子は自動的に設定されます）、 -c は先ほど作成したプロパティファイルです。
docker run の -v オプションでカレントディレクトリをDockerコンテナ上の /opt にマウントしているので、各引数のパスは /opt から始まるようにしています。

```bash
cd opt/
docker pull swagger2markup/swagger2markup
docker run --rm -v $(pwd):/opt swagger2markup/swagger2markup convert -i /opt/swagger.yaml -f /opt/swagger -c /opt/config.properties

```


## Check API specifications

```bash
docker compose up -d
localhost:8001 # サンプルのAPI仕様書の確認
localhost:8002 # ディレクトリ直下にある.yamlファイルのAPI仕様書の確認
localhost:8003/{path or parameter} # API動作の確認
```