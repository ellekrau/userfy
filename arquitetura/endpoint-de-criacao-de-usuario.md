---
description: Permite a adição de um usuário no banco de dados.
---

# Endpoint de criação de usuário

```
// Some codecurl --location --request POST 'localhost:5100/user' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.bnVsbA.TnGenwfJ_CqUtsjhiIO33r5i9eTj5pTFfGr3z2QAPBU' \
--header 'Content-Type: application/json' \
--data-raw '        {
            "name": "Adrielle Krauchuki",
            "cellphone": "5547923038078"
        }
'
```

### Fluxo

![](<../.gitbook/assets/image (2).png>)

### Validação de contrato

```
{
    "name": "Adrielle",
    "cellphone": "5541997060671"
}
```

#### Name

* Campo obrigatório

#### Cellphone

* Campo obrigatório
* Somente números
* Exatos 13 números

###

### Formatação do celular

O cliente pode possuir um padrão para a persistência do celular, e este padrão pode ser informado na variável de ambiente

```
CELLPHONE_PATTERN=
```

#### Regra para o padrão

* Se estiver vazio nenhuma formatação é aplicada
* Deve ser informado a letra X representando os números. Exemplo: +XX (XX) XXXXX-XXXX
* Deve possuir exatamente 13 letras X

#### Lógica

