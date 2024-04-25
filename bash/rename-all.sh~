#!/bin/bash

# Função para gerar um slug a partir de uma string
slugify() {
    echo "$1" | tr -d '[:punct:]' | tr '[:upper:]' '[:lower:]' | tr ' ' '-'
}

# Verifica se um diretório foi fornecido como argumento
if [ $# -eq 0 ]; then
    echo "Por favor, forneça o caminho do diretório como argumento."
    exit 1
fi

# Navega para o diretório fornecido
cd "$1" || exit 1

# Loop para renomear todos os arquivos no diretório
for file in *; do
    # Ignora se não for um arquivo regular
    if [ -f "$file" ]; then
        # Gera um slug a partir do nome do arquivo
        slug=$(slugify "$file")

        # Renomeia o arquivo
        mv "$file" "$slug"
        echo "Renomeado: $file para $slug"
    fi
done
