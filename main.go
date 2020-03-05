package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()

	// Seed random
	rand.Seed(time.Now().Unix())
}

func main() {
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

var lowerReplace = []rune(reverse("zʎxʍʌnʇsɹbdouɯʅʞɾᴉɥƃⅎǝpɔqɐ"))
var upperReplace = []rune(reverse("Z⅄XϺɅՈꓕSꓤꝹԀONꟽ⅂ꓘᒋIH⅁ᖵƎᗡϽꓭ∀"))

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

var webhookCache = make(map[string]*discordgo.Webhook)

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if !strings.HasPrefix(m.Content, "!chaos ") {
		return
	}

	webhook, ok := webhookCache[m.ChannelID]
	if !ok {
		webhooks, err := s.ChannelWebhooks(m.ChannelID)
		if err != nil {
			_, _ = s.ChannelMessageSend(m.ChannelID, "oh no i can't get webhooks")
			return
		}
		if len(webhooks) > 0 {
			webhook = webhooks[0]
			webhookCache[m.ChannelID] = webhook
		} else {
			webhook, err = s.WebhookCreate(m.ChannelID, "ChaosBot", "")
			if err != nil {
				_, _ = s.ChannelMessageSend(m.ChannelID, "oh no i can't make a webhook")
				return
			}
			time.Sleep(300 * time.Millisecond)
			webhookCache[m.ChannelID] = webhook
		}
	}

	modifiedContent := reverse(strings.Map(func(r rune) rune {
		if r > 'a' && r <= 'z' {
			return lowerReplace[r-'a']
		}
		if r > 'A' && r <= 'Z' {
			return upperReplace[r-'A']
		}
		for i, v := range lowerReplace {
			if v == r {
				return rune('a' + i)
			}
		}
		for i, v := range upperReplace {
			if v == r {
				return rune('A' + i)
			}
		}
		return r
	}, strings.TrimPrefix(m.Content, "!chaos ")))

	params := discordgo.WebhookParams{
		Content:   modifiedContent,
		Username:  m.Author.Username,
		AvatarURL: m.Author.AvatarURL(""),
	}

	_, err := s.WebhookExecute(webhook.ID, webhook.Token, false, &params)
	if err != nil {
		_, _ = s.ChannelMessageSend(m.ChannelID, "oh no i can't send the webhook")
		return
	}

	_ = s.ChannelMessageDelete(m.ChannelID, m.ID)
}
