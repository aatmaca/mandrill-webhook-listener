package main

type MandrillWebHook []MandrillWebHookContent

type MandrillWebHookContent struct {
	Ts    uint64                 `json:"ts"`
	Event string                 `json:"event"`
	Msg   MandrillWebHookMessage `json:"msg"`
}

type MandrillWebHookMessage struct {
	RawMsg      string                      `json:"raw_msg"`
	Headers     MandrillWebhookHeaders      `json:"headers"`
	Text        string                      `json:"text"`
	Html        string                      `json:"html"`
	FromEmail   string                      `json:"from_email"`
	FromName    string                      `json:"from_name"`
	To          []string                    `json:"to"`
	Email       string                      `json:"email"`
	Subject     string                      `json:"subject"`
	Tags        []string                    `json:"tags"`
	Sender      string                      `json:"sender"`
	Attachments []MandrillWebhookAttachment `json:"attachments"`
	Images      []MandrillWebhookAttachment `json:"images"`
	SpamReport  MandrillWebhookSpamReport   `json:"spam_report"`
	Dkim        MandrillWebhookDkim         `json:"dkim"`
	Spf         MandrillWebhookSpf          `json:"spf"`
}

type MandrillWebhookHeaders struct {
	ContentType        string   `json:"Content-Type"`
	Date               string   `json:"Date"`
	DkimSignature      []string `json:"Dkim-Signature"`
	DomainKeySignature string   `json:"Domainkey-Signature"`
	From               string   `json:"From"`
	ListUnsubscribe    string   `json:"List-Unsubscribe"`
	MessageId          string   `json:"Message-Id"`
	MimeVersion        string   `json:"Mime-Version"`
	Received           []string `json:"Received"`
	Sender             string   `json:"Sender"`
	Subject            string   `json:"Subject"`
	To                 string   `json:"To"`
	XReportAbuse       string   `json:"X-Report-Abuse"`
}

type MandrillWebhookAttachment struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Content string `json:"content"`
	Base64  bool   `json:"base64"`
}

type MandrillWebhookSpamReport struct {
	Score        int `json:"score"`
	MatchedRules []struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Score       int    `json:"score"`
		Spf         struct {
			Result string `json:"result"`
			Detail string `json:"detail"`
		} `json:"spf,omitempty"`
	} `json:"matched_rules"`
}

type MandrillWebhookDkim struct {
	Signed string `json:"signed"`
	Valid  string `json:"valid"`
}

type MandrillWebhookSpf struct {
	Detail string `json:"detail"`
	Result string `json:"result"`
}
