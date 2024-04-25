### Estratégias de Controle de Concorrência
P garantir a consistência dos dados e evitar condições de corrida:
1. **Bloqueio (Locking)**: Garante q apenas uma parte do código tenha acesso a um recurso compartilhado por vez. P isso podemos usar mecanismos como mutexes, semáforos ou locks exclusivos. Qnd uma parte do código adquire o bloqueio, outras partes devem esperar até que o bloqueio seja liberado antes de acessar o recurso compartilhado.
 `lock()` `unlock()`

2. **Transações Isoladas em Bancos de Dados**: Em sistemas q utilizam bancos de dados, transações isoladas podem ser usadas p garantir que várias operações ocorram de forma consistente, como se fossem executadas em sequência, mesmo que ocorram simultaneamente. Isso ajuda a evitar problemas de concorrência e inconsistências nos dados.

### Técnicas de Sincronização e Fila de Transações
P garantir acesso exclusivo a recursos compartilhados e minimizar problemas de concorrência:
1. **Semáforos**: Abstração de sincronização que permitem o controle de acesso a um recurso compartilhado através do uso de contadores. Eles podem ser usados para limitar o número de threads que podem acessar um recurso simultaneamente.

2. **Filas de Transações**: São estruturas de dados que permitem que operações sejam executadas de forma assíncrona e atômica. Em sistemas de processamento de transações, as operações são colocadas em uma fila e processadas em ordem, garantindo consistência e atomicidade. `kafka` `aws sqs` `rabbitmq` `redis`

3. **Sync**: Conceito fundamental na programação concorrente. Refere-se à coordenação de múltiplas threads ou processos para garantir a execução ordenada das operações. Uma variedade de mecanismos de sincronização está disponível, incluindo bloqueios, semáforos e variáveis de condição. `async` `await`

4. **Mutexes**: Mutex (abreviação de "mutual exclusion") é um mecanismo de sincronização usado para garantir que apenas uma thread por vez possa acessar um recurso compartilhado. Uma thread que deseja acessar o recurso adquire o mutex antes de acessá-lo e o libera quando termina. `Mutex.Lock()` `Mutex.Unlock()`

### Mecanismos de Retry com Backoff
Lidando com falhas em sistemas distribuidos de forma eficaz e resiliente

1. **Retry**:  Mecanismo de retry consiste em tentar novamente uma operação que falhou após um intervalo de tempo específico. Quando uma falha temporária ocorre, o sistema tenta a operação novamente após um curto período de espera, na esperança de que a falha tenha sido resolvida e a operação possa ser concluída com sucesso.

2. **Backoff**: Tecnica que ajusta dinamicamente o intervalo de tempo entre as tentativas de retry com base em certos critérios. Em vez de tentar novamente imediatamente após uma falha, o sistema espera um período de tempo crescente antes de tentar novamente. Isso ajuda a reduzir a carga no sistema e evita sobrecarregar o serviço que está sendo acessado.

    a. `Linear: Intervalo de tempo de espera aumenta por uma quantidade fixa a cada tentativa.`

    b. `Exponencial: Intervalo de tempo entre as tentativas de retry aumenta exponencialmente a cada tentativa.`

    c. `Tentativa Limitada: Tenta um número máximo predefinido de vezes antes de desistir e considerar a operação como falha permanente. Isso impede que o sistema fique preso em loops de retry infinitos que podem prejudicar o desempenho e a disponibilidade do sistema.`


### Técnicas de Cache
Armazenamento temporario de dados frequentemente acessados ​​em uma área de memória de acesso rápido.

`Reduz carga do sistema pq diminui numero de operacoes I/O.`

`Aumenta eficiencia e permite que o sistema manipule maior numero de solicitacoes p/ seg.`

Ps: Importante nesse caso envolver estrategias de invalidacao e expiracao para garantir dados removidos ou atualizados apos determinado tempo. Pq a memoria em cache tem limite menor;;;

###  Tipos de Escalonamento
Abordagem para lidar com concorrencia e desempenho

- Vertical: Aumento de CPU, RAM ou memoria do sistema.

  `Bom pq eh simples de fazer e tem menor complexidade de comunicacao no sistema`

  `Ruim pq ponto unico de falha e tem limitacao fisica de escalabilidade`

- Horizontal: Aumento de maquinas, ou seja, de instancias no sistema e consequente balanceamento de carga

  `Bom escala ao infinito e alem, e tolera melhor falhas`

  `Ruim pq eh caro e piora comunicacao entre componentes`

### Tolerância a Falhas
Técnicas como circuit-breakers e failover para garantir a disponibilidade do sistema mesmo diante de falhas inesperadas
  - Circuit-breaker: Funciona como um disjuntor eletrico, interrompe o acesso a um servico q esta com falhas persistentes. Qndo eh ativado, permite que o servico se recupere e corrija as falhas.

    `Bom pq previne sobrecarga e melhora resiliencia`

  - Failover: Tecnica para comutar automaticamente de um sistema para outro. Por ex, um cluster de servidores, se 1 falhar, o failover redireciona o trafego para outro servidor funcional.

    `Bom pq garante disponibilidade continua e minimiza inatividade [+++ Service Level Indicators +++]`

    `Ruim pq eh complexo e pode gerar inconsistencia nos bancos`

### Estrategias c/ DLQs (Dead Letter Queues)
`Primeiro, o que eh uma DLQ?`  é uma fila especial em sistemas de mensageria que armazena mensagens que não puderam ser processadas com sucesso após um número específico de tentativas. Essas mensagens são consideradas "mortas" ou "não entregáveis" devido a erros durante o processamento ou outras circunstâncias, e são encaminhadas para a DLQ para posterior análise e resolução.

Dito isso, praticas para identificar padroes de erro e ajustar processamento em sistemas de mensageria

1. Estabelecer metricas e alertas relevantes

    `num de msg processadas` `tempo medio de processamento` `tam da fila a ser processada` `etc`

2. Identificar padroes de erro

3. Melhorar processos de processamento

    `ou seja, otimizar processos com base no que se sabe sobre os padroes de erro`

4. Automatizar recuperacao de mensagens

    `retries automaticos` `estrategia de backoff`

5. Monitoramento do desempenho continuo
