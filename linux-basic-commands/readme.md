## Comandos de Listagem e Navegação

1. **Listagem:**
    - `ls -l`: Lista pastas e arquivos com exibição detalhada.
    - `ls -lt`: Lista pastas e arquivos ordenados por timestamp com exibição detalhada.
    - `ls -l -S`: Lista pastas e arquivos ordenados por tamanho com exibição detalhada.

2. **Navegação entre Diretórios:**
    - `pwd`: Mostra o diretório atual.
    - `cd [nome do diretório]`: Muda para outro diretório.
    - `cd ..`: Volta para o diretório anterior.
    - `cd ~`: Vai para o diretório home.

## Comandos de Manipulação de Arquivos

3. **Manipulação de Arquivos:**
    - `touch [nome do arquivo]`: Cria um novo arquivo.
    - `mkdir [nome do diretório]`: Cria um novo diretório.
    - `rm [nome do arquivo]`: Exclui um arquivo.
    - `rm -r [nome do diretório]`: Exclui um diretório e seu conteúdo.
    - `mv [arquivo1] [arquivo2]`: Move ou renomeia um arquivo.

## Outros Comandos Úteis

4. **Medição de Tempo:**
    - `time [processo]`: Mede o tempo de execução de um processo.

5. **Download de Arquivos:**
    - `wget [URL]` ou `curl [URL]`: Baixa arquivos da web.

6. **Manipulação de Texto:**
    - `sed 's/antigo/novo/' arquivo`: Realiza substituições e transformações de texto.
    - `awk '{print $1}' arquivo`: Processamento de texto para extrair e manipular dados.

7. **Busca de Arquivos:**
    - `find diretório -name "*.txt"`: Encontra arquivos e diretórios por critérios.

8. **Filtragem de Texto:**
    - `grep "padrão" arquivo`: Pesquisa padrões em arquivos ou saída de comandos.

9. **Ordenação e Contagem de Linhas:**
    - `sort arquivo`: Classifica linhas de texto em arquivos.
    - `wc arquivo.txt`: Conta linhas, palavras e caracteres em um arquivo.

10. **Visualização de Arquivos:**
    - `tail arquivo.log`: Mostra as últimas linhas de um arquivo.
    - `tail -f arquivo.log`: Acompanha as novas linhas adicionadas em tempo real.
    - `head arquivo.txt`: Mostra as primeiras linhas de um arquivo.

11. **Comparação de Arquivos:**
    - `diff arquivo1.txt arquivo2.txt`: Compara dois arquivos linha por linha.

12. **Remoção de Duplicatas:**
    - `uniq arquivo.txt`: Filtra ou reporta linhas duplicadas em um arquivo.

13. **Execução Recorrente de Comandos:**
    - `watch -n 1 comando`: Executa repetidamente outro comando em intervalos regulares.

14. **Gerenciamento de Sessões de Terminal:**
    - `tmux` (ou `screen`): Gerencia sessões de terminal e múltiplos processos.

15. **Histórico de Comandos:**
    - `history`: Mostra os últimos comandos executados.

16. **Uso de Recursos do Sistema:**
    - `du -h diretório/`: Mostra uso de disco por diretório.
    - `free -h`: Mostra uso de memória.
    - `ps aux`: Lista processos em execução.

17. **Encerramento de Processos:**
    - `kill PID`: Mata um processo pelo seu ID.

18. **Listagem de Portas em Uso:**
    - `lsof -i`: Lista portas em uso.
