package discordlogger

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jonas747/yagpdb/bot"
	"github.com/jonas747/yagpdb/bot/eventsystem"
	"github.com/jonas747/yagpdb/common"
)

var (
	// Send bot leaves joins to this disocrd channel
	BotLeavesJoins int64

	logger = common.GetPluginLogger(&Plugin{})
)

func init() {
	BotLeavesJoins, _ = strconv.ParseInt(os.Getenv("YAGPDB_BOTLEAVESJOINS"), 10, 64)
}

func Register() {
	if BotLeavesJoins != 0 {
		logger.Info("Listening for bot leaves and join")
		common.RegisterPlugin(&Plugin{})
	}
}

var _ bot.BotInitHandler = (*Plugin)(nil)

type Plugin struct{}

func (p *Plugin) PluginInfo() *common.PluginInfo {
	return &common.PluginInfo{
		Name:     "Discord Logger",
		SysName:  "discord_logger",
		Category: common.PluginCategoryCore,
	}
}

func (p *Plugin) BotInit() {
	eventsystem.AddHandler(EventHandler, eventsystem.EventNewGuild, eventsystem.EventGuildDelete)
}

func EventHandler(evt *eventsystem.EventData) {
	count, err := common.GetJoinedServerCount()
	if err != nil {
		logger.WithError(err).Error("failed checking server count")
	}

	msg := ""
	switch evt.Type {
	case eventsystem.EventGuildDelete:
		if evt.GuildDelete().Unavailable {
			// Just a guild outage
			return
		}
		msg = fmt.Sprintf(":x: Left guild **%s** :(", evt.GuildDelete().Guild.Name)
	case eventsystem.EventNewGuild:
		msg = fmt.Sprintf(":white_check_mark: Joined guild **%s** :D", evt.GuildCreate().Guild.Name)
	}

	msg += fmt.Sprintf(" (now connected to %d servers)", count)
	common.BotSession.ChannelMessageSend(BotLeavesJoins, common.EscapeSpecialMentions(msg))
}
