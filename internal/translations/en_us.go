package translations

func initPortugueseTranslation() Translation {
	translation := createTranslation()

	translation.put("requires-js", "Este site requer JavaScript para funcionar corretamente.")

	translation.put("start-the-game", "Preparar!")
	translation.put("force-start", "Iniciar à força")
	translation.put("force-restart", "Reiniciar à força")
	translation.put("game-not-started-title", "O jogo não começou")
	translation.put("waiting-for-host-to-start", "Aguarde o host da sala iniciar o jogo.")

	translation.put("now-spectating-title", "Você está assistindo")
	translation.put("now-spectating-text", "Você pode sair do modo espectador pressionando o botão do olho no topo.")
	translation.put("now-participating-title", "Você está participando")
	translation.put("now-participating-text", "Você pode entrar no modo espectador pressionando o botão do olho no topo.")

	translation.put("spectation-requested-title", "Modo espectador solicitado")
	translation.put("spectation-requested-text", "Você será um espectador após esta rodada.")
	translation.put("participation-requested-title", "Participação solicitada")
	translation.put("participation-requested-text", "Você participará após esta rodada.")

	translation.put("spectation-request-cancelled-title", "Solicitação de espectador cancelada")
	translation.put("spectation-request-cancelled-text", "Sua solicitação de espectador foi cancelada, você continuará participando.")
	translation.put("participation-request-cancelled-title", "Solicitação de participação cancelada")
	translation.put("participation-request-cancelled-text", "Sua solicitação de participação foi cancelada, você continuará assistindo.")

	translation.put("round", "Rodada")
	translation.put("toggle-soundeffects", "Ativar/desativar efeitos sonoros")
	translation.put("toggle-pen-pressure", "Ativar/desativar pressão da caneta")
	translation.put("change-your-name", "Apelido")
	translation.put("randomize", "Aleatório")
	translation.put("apply", "Aplicar")
	translation.put("save", "Salvar")
	translation.put("toggle-fullscreen", "Ativar/desativar tela cheia")
	translation.put("toggle-spectate", "Ativar/desativar modo espectador")
	translation.put("show-help", "Mostrar ajuda")
	translation.put("votekick-a-player", "Votar para expulsar um jogador")

	translation.put("last-turn", "(Última rodada: %s)")

	translation.put("drawer-kicked", "Como o jogador expulso estava desenhando, ninguém ganhará pontos nesta rodada.")
	translation.put("self-kicked", "Você foi expulso")
	translation.put("kick-vote", "(%s/%s) jogadores votaram para expulsar %s.")
	translation.put("player-kicked", "O jogador foi expulso.")
	translation.put("owner-change", "%s é o novo dono da sala.")

	translation.put("change-lobby-settings-tooltip", "Alterar configurações da sala")
	translation.put("change-lobby-settings-title", "Configurações da sala")
	translation.put("lobby-settings-changed", "Configurações da sala alteradas")
	translation.put("advanced-settings", "Configurações avançadas")
	translation.put("chill", "Relax")
	translation.put("competitive", "Competitivo")
	translation.put("chill-alt", "Embora ser rápido seja recompensado, não há problema se você for um pouco mais lento.\nA pontuação base é alta, concentre-se em se divertir!")
	translation.put("competitive-alt", "Quanto mais rápido você for, mais pontos ganhará.\nA pontuação base é mais baixa e a queda é mais rápida.")
	translation.put("score-calculation", "Pontuação")
	translation.put("word-language", "Idioma")
	translation.put("drawing-time-setting", "Tempo de desenho")
	translation.put("rounds-setting", "Rodadas")
	translation.put("max-players-setting", "Máximo de jogadores")
	translation.put("public-lobby-setting", "Sala pública")
	translation.put("custom-words", "Palavras personalizadas")
	translation.put("custom-words-info", "Digite palavras adicionais separadas por vírgulas")
	translation.put("custom-words-per-turn-setting", "Palavras personalizadas por rodada")
	translation.put("players-per-ip-limit-setting", "Limite de jogadores por IP")
	translation.put("save-settings", "Salvar configurações")
	translation.put("input-contains-invalid-data", "Seu input contém dados inválidos:")
	translation.put("please-fix-invalid-input", "Corrija os dados inválidos e tente novamente.")
	translation.put("create-lobby", "Criar Sala")
	translation.put("create-public-lobby", "Criar Sala Pública")
	translation.put("create-private-lobby", "Criar Sala Privada")

	translation.put("players", "Jogadores")
	translation.put("refresh", "Atualizar")
	translation.put("join-lobby", "Entrar na Sala")

	translation.put("message-input-placeholder", "Digite suas respostas e mensagens aqui")

	translation.put("word-choice-warning", "Palavra caso não escolha a tempo")
	translation.put("choose-a-word", "Escolha uma palavra")
	translation.put("waiting-for-word-selection", "Aguardando seleção da palavra")
	translation.put("is-choosing-word", "está escolhendo uma palavra.")

	translation.put("close-guess", "'%s' está muito perto.")
	translation.put("correct-guess", "Você acertou a palavra.")
	translation.put("correct-guess-other-player", "'%s' acertou a palavra.")
	translation.put("round-over", "Rodada encerrada, nenhuma palavra foi escolhida.")
	translation.put("round-over-no-word", "Rodada encerrada, a palavra era '%s'.")
	translation.put("game-over-win", "Parabéns, você venceu!")
	translation.put("game-over-tie", "Empate!")
	translation.put("game-over", "Você ficou em %s. com %s pontos")

	translation.put("change-active-color", "Mudar cor ativa")
	translation.put("use-pencil", "Usar lápis")
	translation.put("use-eraser", "Usar borracha")
	translation.put("use-fill-bucket", "Usar balde de tinta (Preenche a área selecionada com a cor escolhida)")
	translation.put("change-pencil-size-to", "Mudar tamanho do lápis/borracha para %s")
	translation.put("clear-canvas", "Limpar tela")
	translation.put("undo", "Reverter última alteração (Não funciona após \""+translation.Get("clear-canvas")+"\")")

	translation.put("connection-lost", "Conexão perdida!")
	translation.put("connection-lost-text", "Tentando reconectar... \n\nVerifique sua conexão com a internet.\nSe o problema persistir, contate o administrador.")
	translation.put("error-connecting", "Erro ao conectar ao servidor")
	translation.put("error-connecting-text", "O Scribble.rs não conseguiu estabelecer conexão.\n\nEmbora sua internet funcione, o servidor ou firewall não está configurado corretamente.\n\nPara tentar novamente, recarregue a página.")

	translation.put("message-too-long", "Sua mensagem é muito longa.")

	// Help dialog
	translation.put("controls", "Controles")
	translation.put("pencil", "Lápis")
	translation.put("eraser", "Borracha")
	translation.put("fill-bucket", "Balde de tinta")
	translation.put("switch-tools-intro", "Você pode alternar entre ferramentas usando atalhos")
	translation.put("switch-pencil-sizes", "Você também pode alternar os tamanhos do lápis usando as teclas %s a %s.")

	// Generic words
	translation.put("close", "Fechar")
	translation.put("no", "Não")
	translation.put("yes", "Sim")
	translation.put("system", "Sistema")

	translation.put("source-code", "Código Fonte")
	translation.put("help", "Ajuda")
	translation.put("submit-feedback", "Enviar Feedback")
	translation.put("stats", "Status")

	RegisterTranslation("pt", translation)

	return translation
}
