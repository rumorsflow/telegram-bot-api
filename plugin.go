package tgbotapi

import (
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/roadrunner-server/errors"
	"github.com/rumorsflow/contracts/config"
	"go.uber.org/zap"
)

const PluginName = "tgbotapi"

type Plugin struct {
	cfg *Config
	bot *tgbotapi.BotAPI
}

func (p *Plugin) Init(cfg config.Configurer, log *zap.Logger) error {
	const op = errors.Op("tgbotapi plugin init")

	if !cfg.Has(PluginName) {
		return errors.E(op, errors.Disabled)
	}

	var err error

	if err = cfg.UnmarshalKey(PluginName, &p.cfg); err != nil {
		return errors.E(op, errors.Init, err)
	}

	_ = tgbotapi.SetLogger(&logger{l: log})

	p.bot, err = tgbotapi.NewBotAPI(p.cfg.Token)
	if err != nil {
		return errors.E(op, errors.Init, err)
	}

	return nil
}

// Name returns user-friendly plugin name
func (p *Plugin) Name() string {
	return PluginName
}

// Provides declares factory methods.
func (p *Plugin) Provides() []any {
	return []any{
		p.ServiceBotAPI,
	}
}

func (p *Plugin) ServiceBotAPI() *tgbotapi.BotAPI {
	return p.bot
}
