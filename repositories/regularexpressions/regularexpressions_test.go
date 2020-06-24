package regularexpressions_test

import (
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/espal-digital-development/espal-core/repositories/regularexpressions"
)

func TestValidRegexEmail(t *testing.T) {
	re, err := regularexpressions.New()
	if err != nil {
		t.Fatal(err)
	}

	emails := [...]string{
		"prettyandsimple@example.com",
		"very.common@example.com",
		"disposable.style.email.with+symbol@example.com",
		"other.email-with-dash@example.com",
		"fully-qualified-domain@example.com.",
		"x@example.com",
		// "\"much.more unusual\"@example.com",
		// "\"very.unusual.@.unusual.com\"@example.com",
		// "\"very.(),:;<>[]\\\".VERY.\\\"very@\\\\ \\\"very\\\".unusual\"@strange.example.com",
		"example-indeed@strange-example.com",
		// "admin@mailserver1",
		// "#!$%&'*+-/=?^_`{}|~@example.org",
		// "\"()<>[]:,;@\\\\\\\"!#$%&'-/=?^_`{}| ~.a\"@example.org",
		// "\" \"@example.org",
		// "example@localhost",
		"example@s.solutions",
		// "user@localserver",
		// "user@tt",
		// "user@[IPv6:2001:DB8::1]",
	}

	for k := range emails {
		if !re.GetEmail().MatchString(emails[k]) {
			t.Errorf("regular Expression on email %s failed", emails[k])
		}
	}
}

func TestInvalidRegexEmail(t *testing.T) {
	re, err := regularexpressions.New()
	if err != nil {
		t.Fatal(err)
	}

	emails := [...]string{
		"Abc.example.com",
		"A@b@c@example.com",
		"a\"b(c)d,e:f;g<h>i[j\\k]l@example.com",
		"just\"not\"right@example.com",
		"this is\"not\\allowed@example.com",
		"this\\ still\\\"not\\\\allowed@example.com",
		// "1234567890123456789012345678901234567890123456789012345678901234+x@example.com",
		// "john..doe@example.com",
		// "john.doe@example..com",
	}

	for k := range emails {
		if re.GetEmail().MatchString(emails[k]) {
			t.Errorf("invalid Regular Expression on email %s failed", emails[k])
		}
	}
}

func TestRegexPasswordRecoveryhash(t *testing.T) {
	re, err := regularexpressions.New()
	if err != nil {
		t.Fatal(err)
	}

	hashes := [...]string{
		"aeGK6La3iteIfHMcXmKanDhwUgSNCWto3NfaIkBn9BsU5ulhsXinqhmLlW13e5DATyS4iZwr",
		"P0Ih4VSXc9Str2I2jbbFNtdqPuYDWxbtw5sUQqItGgbVO5DgLcUHmBpFynYYx6ktVprUKg49",
		"Qx4HRo6FTHuhXY0RvGgnb7LI5QQy4hJ9EqGQXvKFMekRBV7oYeFSRzGDMZim8TTsvlGSxjfw",
		"lfuj7jczBUq5UnhqQ08h2AdBiEsP5Lwd5kTcmgZJCRtN8nuGDquMxeP4U3GcHAHqJIuKx1Yu",
		"9z0ELfhxFGarQtA6ftih0S6dUDZEU858SYyGcypAFDcKrYokZ6XoYDx35kmodkt3ftRUfdkc",
		"GRvnPHuB1SRIaCnMEhPGKHEDN6iGyxLtNWJK3TsYs9z6BLXGX6j7KLuIUXwBoHvKO9fAv12J",
		"RFcNmi6lhFu2eJytWnOGTZRv34bvAhe1UUQVaQqpKeS8lqi5za4TIeruB4fGGopYGBoHHJzU",
		"AIkHjGdXcWO1wqgbNAlFTS7NapNNoVw0RMTHzO4mjP2UkHre4F5KtWwfkLkrOd49mBBJpKtz",
		"faxi9L4fXsvPmLegx0BIYtt7UIWJC0K3Mhgnm0fGrmtjkUkkFWAMz5AMzUZJ666PT61405ap",
		"SXra6ILSH6IDrvUozyLKAoRAdGOZIQn4Arrx5NJnrHZWYoTmOhRK0vUrTtVKCQ4tRO465ac6",
		"8lV2HP77R0bFJ6fbvDNmZcbw3Tv69cqfIiutbcsHOTCunx2jvNzzlxLSGr0iCBMQXLLL4i5n",
		"nhy3hRQKc7man8d5MDHlIGiAtVcR1u1tqXbpMEgop14jKdFU31dD94cFvHRMZyDhcUOLCsrA",
		"WuhVxUVJAEFNeY7Q1U1Ru4dps9j74zqNXCxGUHQdjjNoSXdvlmO4vFhs7VbRF5HzlaupS7Kj",
		"mUCIrHZvgfLvWdiMtmpu7FEiffnv62nKha7wzelZpuLMsxIKkH5KfzyDv7vbm2I2yYbSwRL1",
		"yY10vjcU81wyLj5rwH54w7Nxe0eJSjOHnKLMu3mjSwMEKXEPLn1bMR5RCCFInqyz97rZTu9I",
		"vh0yzPu8QUKpVf3PUBUcOp2b8oToirDWcqV3ISgJ7uuCiKCROUg2o5uE3nre8QbVaV2xJJBY",
		"euPKz900VsoNJshC0wBBm9mV9VcSKXnBSUGdRUtBrjsMxgIoDqNt96DvHRYeP3Z1P9XKQ76D",
		"Y6Mh90HhJChZCeCQxY0pp3hcH52NWZ7v7m0aLkyQJyBBhV41NhOLuoiLaSqU91U5K7TNXHc2",
		"AQGWqDGMPMWwERpppHCPoGonKUdDciRTILqCYYFzAjRF3q0KHJ04SwtKJQngilQauwJIHPMB",
		"oyTUqCNtPQTPvttz6WmBuGiBTEQNN7qi3kQBrwv9MWv0EpuheRQSjpc6rO9ImisrGcpAHQH8",
	}

	for k := range hashes {
		if !re.GetPasswordRecoveryhash().MatchString(hashes[k]) {
			t.Errorf("regular Expression on password recovery hash %s failed", hashes[k])
		}
	}
}

func TestInvalidRegexPasswordRecoveryhash(t *testing.T) {
	re, err := regularexpressions.New()
	if err != nil {
		t.Fatal(err)
	}

	hashes := [...]string{
		"aeGK6La3iteIfHMcXmKanDhwUgSNCWto3NfaIkBn9BsU5ulhsXinqhmLlW13e5DATyS4i9ahd9asdasdZwr",
		"P0Ih4VSXc9Str2I2jbbFNtdqPuYD5sUQqItGgbVO5DgLcUHmBpFynYYg49",
		"92u(*D?!+)",
		"________________________________________________________________________",
	}

	for k := range hashes {
		if re.GetPasswordRecoveryhash().MatchString(hashes[k]) {
			t.Errorf("invalid Regular Expression on password recovery hash %s failed", hashes[k])
		}
	}
}

func TestRegexActivateAccounthash(t *testing.T) {
	re, err := regularexpressions.New()
	if err != nil {
		t.Fatal(err)
	}

	hashes := [...]string{
		"aeGK6La3iteIfHMcXmKanDhwUgSNCWto3NfaIkBn9BsU5ulhsXinqhmLlW13e5DATyS4iZwr",
		"P0Ih4VSXc9Str2I2jbbFNtdqPuYDWxbtw5sUQqItGgbVO5DgLcUHmBpFynYYx6ktVprUKg49",
		"Qx4HRo6FTHuhXY0RvGgnb7LI5QQy4hJ9EqGQXvKFMekRBV7oYeFSRzGDMZim8TTsvlGSxjfw",
		"lfuj7jczBUq5UnhqQ08h2AdBiEsP5Lwd5kTcmgZJCRtN8nuGDquMxeP4U3GcHAHqJIuKx1Yu",
		"9z0ELfhxFGarQtA6ftih0S6dUDZEU858SYyGcypAFDcKrYokZ6XoYDx35kmodkt3ftRUfdkc",
		"GRvnPHuB1SRIaCnMEhPGKHEDN6iGyxLtNWJK3TsYs9z6BLXGX6j7KLuIUXwBoHvKO9fAv12J",
		"RFcNmi6lhFu2eJytWnOGTZRv34bvAhe1UUQVaQqpKeS8lqi5za4TIeruB4fGGopYGBoHHJzU",
		"AIkHjGdXcWO1wqgbNAlFTS7NapNNoVw0RMTHzO4mjP2UkHre4F5KtWwfkLkrOd49mBBJpKtz",
		"faxi9L4fXsvPmLegx0BIYtt7UIWJC0K3Mhgnm0fGrmtjkUkkFWAMz5AMzUZJ666PT61405ap",
		"SXra6ILSH6IDrvUozyLKAoRAdGOZIQn4Arrx5NJnrHZWYoTmOhRK0vUrTtVKCQ4tRO465ac6",
		"8lV2HP77R0bFJ6fbvDNmZcbw3Tv69cqfIiutbcsHOTCunx2jvNzzlxLSGr0iCBMQXLLL4i5n",
		"nhy3hRQKc7man8d5MDHlIGiAtVcR1u1tqXbpMEgop14jKdFU31dD94cFvHRMZyDhcUOLCsrA",
		"WuhVxUVJAEFNeY7Q1U1Ru4dps9j74zqNXCxGUHQdjjNoSXdvlmO4vFhs7VbRF5HzlaupS7Kj",
		"mUCIrHZvgfLvWdiMtmpu7FEiffnv62nKha7wzelZpuLMsxIKkH5KfzyDv7vbm2I2yYbSwRL1",
		"yY10vjcU81wyLj5rwH54w7Nxe0eJSjOHnKLMu3mjSwMEKXEPLn1bMR5RCCFInqyz97rZTu9I",
		"vh0yzPu8QUKpVf3PUBUcOp2b8oToirDWcqV3ISgJ7uuCiKCROUg2o5uE3nre8QbVaV2xJJBY",
		"euPKz900VsoNJshC0wBBm9mV9VcSKXnBSUGdRUtBrjsMxgIoDqNt96DvHRYeP3Z1P9XKQ76D",
		"Y6Mh90HhJChZCeCQxY0pp3hcH52NWZ7v7m0aLkyQJyBBhV41NhOLuoiLaSqU91U5K7TNXHc2",
		"AQGWqDGMPMWwERpppHCPoGonKUdDciRTILqCYYFzAjRF3q0KHJ04SwtKJQngilQauwJIHPMB",
		"oyTUqCNtPQTPvttz6WmBuGiBTEQNN7qi3kQBrwv9MWv0EpuheRQSjpc6rO9ImisrGcpAHQH8",
	}

	for k := range hashes {
		if !re.GetActivateAccounthash().MatchString(hashes[k]) {
			t.Errorf("hash %s failed", hashes[k])
		}
	}
}

func TestInvalidRegexActivateAccounthash(t *testing.T) {
	re, err := regularexpressions.New()
	if err != nil {
		t.Fatal(err)
	}

	hashes := [...]string{
		"aeGK6La3iteIfHMcXmKanDhwUgSNCWto3NfaIkBn9BsU5ulhsXinqhmLlW13e5DATyS4i9ahd9asdasdZwr",
		"P0Ih4VSXc9Str2I2jbbFNtdqPuYD5sUQqItGgbVO5DgLcUHmBpFynYYg49",
		"92u(*D?!+)",
		"________________________________________________________________________",
	}

	for k := range hashes {
		if re.GetActivateAccounthash().MatchString(hashes[k]) {
			t.Errorf("hash %s failed", hashes[k])
		}
	}
}

func TestRegexRouteIds(t *testing.T) {
	re, err := regularexpressions.New()
	if err != nil {
		t.Fatal(err)
	}
	ids := [...]string{
		gofakeit.UUID(),
		gofakeit.UUID() + "," + gofakeit.UUID() + "," + gofakeit.UUID(),
	}

	for k := range ids {
		if !re.GetRouteIDs().MatchString(ids[k]) {
			t.Errorf("route ids %s failed", ids[k])
		}
	}
}

func TestRegexInvalidRouteIds(t *testing.T) {
	re, err := regularexpressions.New()
	if err != nil {
		t.Fatal(err)
	}

	ids := [...]string{
		"",
		",",
		"12,",
		"12,13,14,",
		"12.01,12.3",
	}

	for k := range ids {
		if re.GetRouteIDs().MatchString(ids[k]) {
			t.Errorf("route ids %s failed", ids[k])
		}
	}
}
