# Multithreading

## Processo X Threads

Um processo é uma instância de um programa de computador em execução. Quando você executa um programa, o sistema operacional aloca espaço na memória para aquele programa e cria um processo que contém todas as informações necessárias para gerenciar a execução do programa. Isso inclui o código do programa, variáveis, contadores de instruções, informações de estado e uma área de memória própria, conhecida como espaço de endereçamento do processo.

Cada processo é executado de forma isolada dos outros processos, o que significa que os recursos (memória, entradas de sistema de arquivos, dispositivos, etc.) de um processo não são compartilhados diretamente com outros processos. Isso ajuda a manter a estabilidade e a segurança do sistema operacional, pois os processos mal comportados ou mal-intencionados têm mais dificuldade em afetar outros processos.

Uma thread, por outro lado, é uma entidade mais leve que existe dentro de um processo. Um processo pode ter várias threads, todas compartilhando o mesmo espaço de endereçamento do processo pai, mas cada thread executa uma sequência independente de instruções. As threads de um mesmo processo podem executar tarefas em paralelo e compartilhar recursos, como dados na memória e arquivos abertos, o que as torna extremamente úteis para realizar várias tarefas ao mesmo tempo dentro do mesmo programa.

Threads são unidades básicas de utilização de CPU e cada thread tem seu próprio contador do programa, registro e pilha de chamadas, mas elas compartilham outros recursos do processo. O uso de múltiplas threads pode aumentar a eficiência do programa ao aproveitar os núcleos de processadores multithread ou multi-core disponíveis em muitos computadores modernos, permitindo que as tarefas sejam executadas de forma mais rápida e eficiente.

Em resumo, um processo é como um programa em execução com seu próprio espaço de memória isolado, enquanto uma thread é uma linha de execução dentro de um processo que pode compartilhar recursos e ser usada para realizar tarefas paralelas.

Abaixo, temos a ilutração de dois processos. O processo A que possui duas threads, a thread A e a thread B e que acessam a mesma váriavel A.

![Processos e Threads](./imgs/Processos_Threads.png)

## Mutex e o controle de acesso um recurso compartilhado

Mutex, que é uma abreviação de "Mutual Exclusion object" (objeto de exclusão mútua), é um mecanismo de sincronização usado para controlar o acesso a um recurso compartilhado em ambientes de programação concorrentes. O conceito de mutex é essencial em sistemas operacionais e linguagens de programação que suportam multithreading para garantir que apenas uma thread (ou processo, em alguns casos) possa acessar um recurso ou seção de código crítico por vez.

Um mutex funciona da seguinte maneira:

1. **Bloqueio (lock):** Quando uma thread deseja entrar em uma seção crítica, ela tenta bloquear o mutex associado a essa seção. Se o mutex já estiver bloqueado por outra thread, a thread que está tentando adquirir o bloqueio será suspensa até que o mutex seja liberado.

2. **Seção Crítica:** Uma vez que a thread tem o mutex bloqueado, ela pode entrar e executar o código na seção crítica, que geralmente envolve acessar ou modificar recursos compartilhados.

3. **Liberação (unlock):** Quando a thread conclui sua tarefa na seção crítica, ela libera o mutex para que outras threads possam adquiri-lo e acessar a seção crítica.

Mutex são fundamentais para evitar race conditions, que ocorrem quando várias threads acessam e manipulam dados compartilhados simultaneamente, levando a resultados inesperados e potencialmente perigosos. Ao garantir a exclusão mútua, os mutex ajudam a manter a consistência dos dados e a integridade do programa.

Além disso, é importante gerenciar mutex com cuidado para evitar problemas comuns, como deadlocks, onde duas ou mais threads ficam bloqueadas indefinidamente, aguardando a liberação de mutexes que nunca ocorrerá.

Abaixo, segue uma breve ilustração de um mutex:

![Processos e Threads](./imgs/Processos_Threads2.png)

## Diferenças entre Concorrência e Paralelismo

Concorrência e paralelismo são conceitos relacionados à execução de múltiplas tarefas em sistemas de computador, mas têm diferenças importantes em como as tarefas são realizadas.

**Concorrência:**
Concorrência refere-se à capacidade de um sistema lidar com várias tarefas ao mesmo tempo. Em um sistema concorrente, várias tarefas podem estar em progresso em um dado momento, mas não necessariamente estão sendo processadas ao mesmo tempo. O processamento pode ser intercalado de forma que as tarefas avancem sem estar executando simultaneamente. Isso é especialmente útil em sistemas com um único processador ou núcleo, onde as tarefas são alternadas rapidamente (multitarefa), dando a impressão de que estão sendo realizadas simultaneamente.

**Paralelismo:**
Paralelismo, por outro lado, é quando tarefas múltiplas estão sendo executadas ao mesmo tempo em unidades de processamento separadas. Um sistema paralelo tira vantagem de múltiplos processadores ou núcleos para dividir uma tarefa em partes que podem ser processadas simultaneamente. O paralelismo pode ser alcançado em diferentes níveis, como múltiplos threads rodando em diferentes núcleos de um processador multicore, ou por distribuição de tarefas através de vários computadores em um cluster.

**Diferenças-chave:**
- **Execução:** Concorrência pode ocorrer em um única unidade de processamento através de alternância entre tarefas, enquanto paralelismo requer múltiplas unidades de processamento para executar tarefas ao mesmo tempo.
- **Escalabilidade:** O paralelismo se beneficia diretamente da adição de mais processadores ou núcleos, enquanto a concorrência se beneficia da organização eficiente e compartilhamento de recursos entre tarefas concorrentes.
- **Complexidade:** A gestão de tarefas em paralelo pode ser mais complexa, uma vez que requer coordenação e sincronização entre as unidades de processamento. A concorrência também tem seus desafios, especialmente quando se trata de compartilhamento de recursos e prevenção de condições de corrida, mas geralmente é mais simples de gerenciar do que o paralelismo puro.

Ambos os conceitos são aplicados na construção de sistemas de software modernos e sistemas operacionais para melhorar o desempenho e a eficiência, mas requerem técnicas cuidadosas de projeto e sincronização para prevenir problemas como condições de corrida, deadlocks e outros tipos de inconsistências.

**Paralelismo em um ambiente multicore:**
Um sistema com múltiplos núcleos (multicore) permite que paralelismo verdadeiro ocorra, já que diferentes threads ou processos podem ser executados em núcleos separados ao mesmo tempo. Isto é, cada núcleo pode processar uma tarefa diferente simultaneamente, o que melhora significativamente a eficiência e a velocidade do processamento comparado com a execução em um único núcleo.

**Concorrência em um ambiente multicore:**
A concorrência se refere ao gerenciamento de múltiplas tarefas que podem ou não estar rodando simultaneamente. Em sistemas multicore, a concorrência ainda é relevante porque geralmente temos mais threads ou processos concorrentes do que núcleos disponíveis. O sistema operacional é responsável por agendar as tarefas (e.g., threads) entre os núcleos disponíveis, alternando entre eles conforme necessário para lidar com a carga de trabalho. Isso significa que enquanto algumas tarefas podem ser executadas em paralelo, outras podem estar aguardando para serem executadas ou serem intercaladas nos núcleos, demonstrando concorrência.

**Concorrência e Paralelismo Simultâneos:**
Quando um sistema multicore está executando um aplicativo multithreaded, pode haver múltiplas threads rodando em paralelo nos diferentes núcleos (paralelismo), enquanto outras threads - talvez do mesmo aplicativo ou de outros aplicativos - podem estar aguardando sua vez de executar ou sendo alternadas de forma concorrente. A presença de múltiplas CPUs ou núcleos permite que as threads sejam executadas verdadeiramente em paralelo, mas o sistema operacional ainda precisa gerenciar a concorrência, que ocorre devido à quantidade limitada de recursos computacionais comparada ao número de tarefas.

Assim, concorrência e paralelismo são conceitos que se complementam em sistemas multicore. A concorrência otimiza o uso de recursos em um sistema enquanto o paralelismo acelera a execução através do processamento simultâneo. A complexidade do desenvolvimento de software em tais sistemas aumenta devido à necessidade de gerenciar corretamente os recursos compartilhados em um ambiente onde várias tarefas estão sendo executadas, tanto concorrentemente quanto em paralelo.

**Paralelismo em um ambiente multicore:**
Um sistema com múltiplos núcleos (multicore) permite que paralelismo verdadeiro ocorra, já que diferentes threads ou processos podem ser executados em núcleos separados ao mesmo tempo. Isto é, cada núcleo pode processar uma tarefa diferente simultaneamente, o que melhora significativamente a eficiência e a velocidade do processamento comparado com a execução em um único núcleo.

**Concorrência em um ambiente multicore:**
A concorrência se refere ao gerenciamento de múltiplas tarefas que podem ou não estar rodando simultaneamente. Em sistemas multicore, a concorrência ainda é relevante porque geralmente temos mais threads ou processos concorrentes do que núcleos disponíveis. O sistema operacional é responsável por agendar as tarefas (e.g., threads) entre os núcleos disponíveis, alternando entre eles conforme necessário para lidar com a carga de trabalho. Isso significa que enquanto algumas tarefas podem ser executadas em paralelo, outras podem estar aguardando para serem executadas ou serem intercaladas nos núcleos, demonstrando concorrência.

**Concorrência e Paralelismo Simultâneos:**
Quando um sistema multicore está executando um aplicativo multithreaded, pode haver múltiplas threads rodando em paralelo nos diferentes núcleos (paralelismo), enquanto outras threads - talvez do mesmo aplicativo ou de outros aplicativos - podem estar aguardando sua vez de executar ou sendo alternadas de forma concorrente. A presença de múltiplas CPUs ou núcleos permite que as threads sejam executadas verdadeiramente em paralelo, mas o sistema operacional ainda precisa gerenciar a concorrência, que ocorre devido à quantidade limitada de recursos computacionais comparada ao número de tarefas.

Assim, concorrência e paralelismo são conceitos que se complementam em sistemas multicore. A concorrência otimiza o uso de recursos em um sistema enquanto o paralelismo acelera a execução através do processamento simultâneo. A complexidade do desenvolvimento de software em tais sistemas aumenta devido à necessidade de gerenciar corretamente os recursos compartilhados em um ambiente onde várias tarefas estão sendo executadas, tanto concorrentemente quanto em paralelo.

Abaixo, segue uma breve ilustração de concorrência e paralelismo:

![Processos e Threads](./imgs/Processos_Threads3.png)