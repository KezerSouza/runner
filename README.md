# Sistema Runner

## 1. Visão Geral

O projeto tem como objetivo simular um cenário real de integração entre sistemas, permitindo a execução simplificada de aplicações Java por meio de uma interface de linha de comando.

O sistema foi projetado para abstrair a complexidade envolvida na execução manual de aplicações Java, oferecendo uma experiência mais simples e automatizada ao usuário, sem a necessidade de conhecimento aprofundado sobre configuração de ambiente ou execução de comandos específicos.

---

## 2. Componentes do Sistema

* **Assinatura (CLI):** Interface de linha de comando desenvolvida em Go, responsável por receber comandos do usuário e coordenar a execução das operações do sistema de forma simples e multiplataforma (Windows, Linux e macOS).

* **Assinador (Java):** Aplicação `assinador.jar` responsável por simular a criação e validação de assinaturas digitais, incluindo validação de parâmetros e retorno estruturado em formato JSON.

* **Simulador:** Componente responsável pelo gerenciamento do ciclo de vida de um sistema externo (`simulador.jar`), permitindo iniciar, parar e consultar o status da aplicação.

---

## 3. Principais Funcionalidades

* **Execução Simplificada:** Permite executar operações de assinatura e validação por meio de comandos diretos no terminal, eliminando a necessidade de interação manual com o Java.

* **Simulação de Assinaturas Digitais:** O sistema simula a criação e validação de assinaturas digitais, garantindo a verificação de parâmetros e comportamento consistente para testes.

* **Integração Go + Java:** O CLI desenvolvido em Go invoca automaticamente o `assinador.jar`, realizando a comunicação entre diferentes tecnologias de forma transparente ao usuário.

* **Gerenciamento do Simulador:** Permite iniciar, encerrar e verificar o status do simulador externo, incluindo controle de processo via PID.

* **Portabilidade:** O sistema pode ser executado em diferentes sistemas operacionais, garantindo flexibilidade e facilidade de uso.

* **Automação de Execução:** O sistema localiza automaticamente o Java e o arquivo `.jar`, simplificando a configuração do ambiente.

---