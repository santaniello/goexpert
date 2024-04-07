# Eventos

***Definição:*** Em engenharia de software, eventos são ações ou ocorrências que são detectadas pelo software que podem ser geradas por fontes diversas, como interação do usuário, mensagens de sistema, atualizações de estado, entre outros. Os eventos podem ser síncronos ou assíncronos, e são uma forma de comunicação na qual um componente de software transmite informação para outro(s) que estejam interessados em saber da ocorrência do evento. Aqueles que estão interessados em um evento em particular "ouvem" ou "observam" o evento.

- Algo que ocorreu no passado;

Exemplo: 

| Ação                     | Evento                        |
|--------------------------|-------------------------------|
| Inseri o registro        | Registro inserido             |
| Inseri um novo cliente   | Publicar uma mensagem na fila |
| Inseri um novo cliente   | Notificar um usuário no Slack |


- Evento (Carregar dados)
- Operações que serão executadas quando um evento e chamado.
- Gerenciador dos nossos eventos/operações
  - Registrar os eventos e suas operações;
  - Despachar / Fire no evento para suas operações serem executadas;





