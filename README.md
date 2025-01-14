
# CommitIA

CommitIA é uma ferramenta em Go que utiliza llm para analisar mudanças no código e gerar mensagens de commit claras e descritivas. Simplifique o processo de criação de commits no Git com mensagens automáticas ou ajustadas ao contexto fornecido.

## Instalação

### Instalação via Código

1. Clone este repositório:
   ```bash
   git clone https://github.com/seu-usuario/commitia.git
   cd commitia
   ```

2. Dê permissão de execução ao instalador (se necessario):
   ```bash
   chmod +x ./install
   ```

3. Compile o projeto:
   ```bash
   ./build
   ```

4. Instale o binário:
   ```bash
   ./install
   ```


### Instalação via Release

1. Baixe a versão mais recente do [CommitIA Releases](https://github.com/seu-usuario/commitia/releases).
2. Extraia o arquivo:
   ```bash
   tar -xvf commitia-{versão}.tar.gz
   cd commitia
   ```

3. Dê permissão de execução ao instalador:
   ```bash
   chmod +x ./install
   ```

4. Instale o binário:
   ```bash
   ./install
   ```

O binário será movido para `/usr/local/bin`.

## Uso

1. No diretório do projeto onde deseja fazer o commit, execute:
   ```bash
   commitia
   ```

   A llm gerará automaticamente uma mensagem de commit baseada nas mudanças do código.

2. Caso queira fornecer mais contexto ou explicações adicionais sobre as alterações realizadas, utilize o parâmetro `-d`:
   ```bash
   commitia -d "Mensagem explicativa sobre as alterações feitas"
   ```

## Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues e enviar pull requests.

## Licença

Este projeto está licenciado sob a [Licença MIT](LICENSE).

---
