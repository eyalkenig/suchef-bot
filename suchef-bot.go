package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/eyalkenig/suchef-bot/server/providers"
	"github.com/gorilla/mux"
	"gopkg.in/maciekmm/messenger-platform-go-sdk.v4"

	"github.com/eyalkenig/suchef-bot/server"
)

func main() {
	messenger := &messenger.Messenger{
		VerifyToken: os.Getenv("VERIFY_TOKEN"),
		AppSecret:   os.Getenv("APP_SECRET"),
		AccessToken: os.Getenv("PAGE_ACCESS_KEY"),
	}

	dbConnectionParams := providers.DBConnectionParams{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Address:  os.Getenv("DB_ADDRESS"),
		DBName:   os.Getenv("DB_NAME"),
	}

	dataProvider, err := providers.NewBotDataProvider(dbConnectionParams)
	if err != nil {
		fmt.Println("could not create suchef server. error: " + err.Error())
		return
	}

	accountID := int64(1)

	suchefServer := server.NewSuchefServer(accountID, messenger, dataProvider, dataProvider, dataProvider)
	fmt.Println("server started successfully")

	messenger.MessageReceived = suchefServer.BindMessageReceived()
	messenger.Postback = suchefServer.BindPostbackReceived()

	r := mux.NewRouter()
	r.HandleFunc("/webhook", messenger.Handler)
	r.HandleFunc("/privacy", privacy)
	r.HandleFunc("/accounts/{account_id}/courses", suchefServer.AddCourse).Methods("POST")
	http.Handle("/", r)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func privacy(w http.ResponseWriter, r *http.Request) {
	privacyText := `Suchef Privacy Policy
	This privacy policy has been compiled to better serve those who are concerned with how their 'Personally Identifiable Information' (PII) is being used online. PII, as described in US privacy law and information security, is information that can be used on its own or with other information to identify, contact, or locate a single person, or to identify an individual in context. Please read our privacy policy carefully to get a clear understanding of how we collect, use, protect or otherwise handle your Personally Identifiable Information in accordance with our website.

	What personal information do we collect from the people that visit our blog, website or app?

	When ordering or registering on our site, as appropriate, you may be asked to enter your name or other details to help you with your experience.

	When do we collect information?

	We collect information from you when you respond to a survey or enter information on our site.

	Provide us with feedback on our products or services
	How do we use your information?

	We may use the information we collect from you when you register, make a purchase, sign up for our newsletter, respond to a survey or marketing communication, surf the website, or use certain other site features in the following ways:

	• To personalize your experience and to allow us to deliver the type of content and product offerings in which you are most interested.
	• To allow us to better service you in responding to your customer service requests.
	• To ask for ratings and reviews of services or products
	• To follow up with them after correspondence (live chat, email or phone inquiries)

	How do we protect your information?

	We do not use vulnerability scanning and/or scanning to PCI standards.
	We only provide articles and information. We never ask for credit card numbers.
	We do not use Malware Scanning.

	We do not use an SSL certificate
	• We only provide articles and information. We never ask for personal or private information like names, email addresses, or credit card numbers.

	Do we use 'cookies'?

	We do not use cookies for tracking purposes

	You can choose to have your computer warn you each time a cookie is being sent, or you can choose to turn off all cookies. You do this through your browser settings. Since browser is a little different, look at your browser's Help Menu to learn the correct way to modify your cookies.

	If you turn cookies off, some features will be disabled. that make your site experience more efficient and may not function properly.

	However, you will still be able to place orders .


	Third-party disclosure

	We do not sell, trade, or otherwise transfer to outside parties your Personally Identifiable Information unless we provide users with advance notice. This does not include website hosting partners and other parties who assist us in operating our website, conducting our business, or serving our users, so long as those parties agree to keep this information confidential. We may also release information when it's release is appropriate to comply with the law, enforce our site policies, or protect ours or others' rights, property or safety.

	However, non-personally identifiable visitor information may be provided to other parties for marketing, advertising, or other uses.

	Third-party links

	Occasionally, at our discretion, we may include or offer third-party products or services on our website. These third-party sites have separate and independent privacy policies. We therefore have no responsibility or liability for the content and activities of these linked sites. Nonetheless, we seek to protect the integrity of our site and welcome any feedback about these sites.

	Google

	Google's advertising requirements can be summed up by Google's Advertising Principles. They are put in place to provide a positive experience for users. https://support.google.com/adwordspolicy/answer/1316548?hl=en

	We have not enabled Google AdSense on our site but we may do so in the future.

	California Online Privacy Protection Act

	CalOPPA is the first state law in the nation to require commercial websites and online services to post a privacy policy. The law's reach stretches well beyond California to require any person or company in the United States (and conceivably the world) that operates websites collecting Personally Identifiable Information from California consumers to post a conspicuous privacy policy on its website stating exactly the information being collected and those individuals or companies with whom it is being shared. - See more at: http://consumercal.org/california-online-privacy-protection-act-caloppa/#sthash.0FdRbT51.dpuf

	According to CalOPPA, we agree to the following:
	Users can visit our site anonymously.
	Once this privacy policy is created, we will add a link to it on our home page or as a minimum, on the first significant page after entering our website.
	Our Privacy Policy link includes the word 'Privacy' and can easily be found on the page specified above.

	You will be notified of any Privacy Policy changes:
	• On our Privacy Policy Page
	Can change your personal information:
	• By chatting with us or by sending us a support ticket

	How does our site handle Do Not Track signals?
	We honor Do Not Track signals and Do Not Track, plant cookies, or use advertising when a Do Not Track (DNT) browser mechanism is in place.

	Does our site allow third-party behavioral tracking?
	It's also important to note that we do not allow third-party behavioral tracking

	COPPA (Children Online Privacy Protection Act)

	When it comes to the collection of personal information from children under the age of 13 years old, the Children's Online Privacy Protection Act (COPPA) puts parents in control. The Federal Trade Commission, United States' consumer protection agency, enforces the COPPA Rule, which spells out what operators of websites and online services must do to protect children's privacy and safety online.

	We do not specifically market to children under the age of 13 years old.

	Fair Information Practices

	The Fair Information Practices Principles form the backbone of privacy law in the United States and the concepts they include have played a significant role in the development of data protection laws around the globe. Understanding the Fair Information Practice Principles and how they should be implemented is critical to comply with the various privacy laws that protect personal information.

	In order to be in line with Fair Information Practices we will take the following responsive action, should a data breach occur:
	We will notify the users via in-site notification
	• Within 7 business days

	We also agree to the Individual Redress Principle which requires that individuals have the right to legally pursue enforceable rights against data collectors and processors who fail to adhere to the law. This principle requires not only that individuals have enforceable rights against data users, but also that individuals have recourse to courts or government agencies to investigate and/or prosecute non-compliance by data processors.

	CAN SPAM Act

	The CAN-SPAM Act is a law that sets the rules for commercial email, establishes requirements for commercial messages, gives recipients the right to have emails stopped from being sent to them, and spells out tough penalties for violations.

	We collect your email address in order to:

	To be in accordance with CANSPAM, we agree to the following:

	If at any time you would like to unsubscribe from receiving future emails, you can email us at
	and we will promptly remove you from ALL correspondence.


	Contacting Us

	If there are any questions regarding this privacy policy, you may contact us using the information below.

	Suchef

	Last Edited on 2017-01-13`
	fmt.Fprintf(w, privacyText)
}
