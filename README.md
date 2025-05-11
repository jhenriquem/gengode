# 🧠 gengode — Ferramenta CLI para interações com IA

`gengode` é uma ferramenta de linha de comando desenvolvida em Go que permite interações com uma IA diretamente do terminal. Um projeto que fiz como forma de estudo para o uso de IA's junto ao Go. 

_🚀 O foco da ferramenta é permite consultas com IA sem precisar sai do terminal, para um desenvolvedor que usar neovim, como eu, isso traz uma facilidade._

## ✨ Funcionalidades

- 💬 Chat interativo com IA diretamente no terminal
- 🧱 Geração de snippets de código com base em prompts
<!--- 🔁 Histórico de conversas salvo localmente-->
- ⚙️ Suporte a múltiplas linguagens (ex: Go, JavaScript, Python, etc.)
<!--- 🔐 Suporte à configuração de chave de API para uso com provedores como OpenAI-->

## 📦 Pacotes utilizados 
- [go-term-markdown](https://github.com/MichaelMure/go-term-markdown) Um pacote de renderização de markdown para o terminal 
- [github.com/fatih/color](https://github.com/fatih/color) Pacote de cores para Go (golang) 

## ⚙️ LLM utilizado (Microsoft: MAI DS R1)
Como LLM (Large Language Model) usei um diponivel pela [Open Router](https://openrouter.ai/), uma interface unificada de LLMs. 
- [Microsoft: MAI DS R1](https://openrouter.ai/microsoft/mai-ds-r1:free) : MAI-DS-R1 é uma variante pós-treinada do DeepSeek-R1 desenvolvida pela equipe de IA da Microsoft para melhorar a capacidade de resposta do modelo em tópicos bloqueados anteriormente, ao mesmo tempo em que aprimora seu perfil de segurança.


## 📦 Instalação

#### 📌 Dependências 
- Go instalado na sua máquina.
- [Conta na *Open Router*](https://openrouter.ai/) : A interface unificada para LLMs  *( ⚠️ Após criar uma conta, gere uma chave de API )*


Clone o repositório:
```bash
git clone https://github.com/jhenriquem/gengode.git
cd gengode
```

Gere o executável:
```bash
go build -o gengode .

gengode
```
Ou execute o arquivo main.go
```bash 
go run main.go
``` 
