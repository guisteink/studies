#!/bin/bash

# Função para gerar um slug a partir de uma string
slugify() {
    echo "$1" | tr -d '[:punct:]' | tr '[:upper:]' '[:lower:]' | tr ' ' '-'
}

# Obtém o diretório do script
script_dir=$(dirname "$(readlink -f "$0")")

# Navega para o diretório do script
cd "$script_dir" || exit 1

# Loop para renomear todos os arquivos no diretório do script
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

