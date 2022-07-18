---
description: Formata o nome do usuário conforme as especificações do cliente.
---

# Formatador de nome

### Formatação do nome

O cliente pode possuir um padrão para a persistência do nome dos usuários, e este padrão pode ser informado na variável de ambiente

```
NAME_PATTERN=
```

#### Tipos aceitos

* _uppercase :_ Todas as letras maiúsculas
* _lowercase :_ Todas as letras minúsculas
* Vazio : Nenhuma formação aplicada

Caso exista a necessidade, é possível estender o comportamento do formatador para que aceite outros tipos.

#### Lógica

Aplica a formatação informada em todos os caracteres do nome do cliente.

#### Exemplo

```
Padrão
NAME_PATTERN=lowercase

Entrada
Adrielle Krauchuki

Saída
adrielle krauchuki
```
