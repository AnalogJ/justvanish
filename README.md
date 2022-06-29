<p align="center">
  <a href="https://github.com/AnalogJ/justvanish">
  <img width="300" alt="vanish_view" src="docs/noun-ghost-1358159.png">
  </a>
</p>


# justvanish
Tell databrokers to F#@% Off. Your data is your data, they shouldn't be monetizing your personal information without your knowledge.


# Why?

- https://www.bloomberg.com/news/articles/2022-06-27/anti-abortion-centers-find-pregnant-teens-online-then-save-their-data?srnd=technology-vp
- https://slate.com/technology/2022/06/health-data-brokers-privacy.html
- https://www.vice.com/en/article/m7vzjb/location-data-abortion-clinics-safegraph-planned-parenthood
- https://www.wired.com/story/data-brokers-tracking-abortion-clinics-security-news/
- https://www.eff.org/deeplinks/2021/08/illinois-bought-invasive-phone-location-data-banned-broker-safegraph
- https://www.washingtonpost.com/news/the-switch/wp/2017/03/23/congress-is-poised-to-undo-landmark-rules-covering-your-internet-privacy/
- https://news.harvard.edu/gazette/story/2017/08/when-it-comes-to-internet-privacy-be-very-afraid-analyst-suggests/
- https://www.wsj.com/articles/amazon-starts-selling-software-to-mine-patient-health-records-1543352136
- https://www.fastcompany.com/90310803/here-are-the-data-brokers-quietly-buying-and-selling-your-personal-information
- https://www.kaspersky.com/resource-center/preemptive-safety/how-to-stop-data-brokers-from-selling-your-personal-information
- https://www.governing.com/now/ice-circumvents-colorado-sanctuary-laws-with-data-brokers
- 

# How?

There are a handful of regulations we can leverage to limit the ability of data brokers & other organizations to legally sell our data.
[General Data Protection Regulation (GDPR)](https://gdpr-info.eu/) and [California Consumer Privacy Act (CCPA)](https://oag.ca.gov/privacy/ccpa) are the two
most relevant, however there are other (sometimes less powerful) laws we can leverage as well (Colorado, Virginia, Utah, Connecticut, Indiana, and Ohio).

In general these laws provide privacy rights to consumers (GDPR is similar):

- The right to know about the personal information a business collects about them and how it is used and shared;
- The right to delete personal information collected from them (with some exceptions);
- The right to opt-out of the sale of their personal information; and 
- The right to non-discrimination for exercising their CCPA rights.

Unfortunately these laws only apply to California & European residents. Still, it's a good foundation to build ontop
of these laws.

# Goals

Data brokers exist because it's incredibly easy to collect information about you as you go about your day --
the websites you visit, the apps you use and even your phone are all sharing information with data brokers 24x7. 
However, while it's easy for data brokers to collect & sell your information, it's frustratingly hard to tell them to stop. 
You probably never directly interacted with the broker, you just purchased some groceries online. 

JustVanish hopes to fix this imbalance by providing users with a simple (automated) mechanism to:
- request a copy of your personal information stored by data brokers, government agencies and other organizations
- request that organizations restrict the collection & sale of your personal information
- request the deletion of your personal information from data brokers & other organizations

My intent is to have JustVanish become a community maintained global registry of databrokers
This databroker & organization information is stored following a [structured schema](organization-schema.json) allowing
other tools to leverage it (similar to how AdBlock filter lists are maintained https://github.com/topics/adblock-list).

Ideally there would be a national registry where consumers could "opt-out" themselves (similar to the [Do Not Call list](https://www.donotcall.gov/)), 
but until regulation supporting that passes, the only alternative is for consumers to exercise their rights themselves. 

# Getting Started


# What about ...?

### What about paid options?
- [Association of National Advertisers](https://www.ana.net/content/show?id=thedmaorg-redirect)
- [Brand Yourself](https://brandyourself.com/)
- [DeleteMeNow](https://deletemenow.com/)
- [DeleteMe](https://joindeleteme.com/)
- [EasyOptOuts](https://easyoptouts.com/)
- [IDX Privacy](https://www.idx.us/idx-privacy)
- [Kanary](https://www.thekanary.com/)
- [OneRep](https://onerep.com/)
- [Optery](https://www.optery.com)
- [Privacy Bee](https://privacybee.com)
- [Privacy Pros](https://privacypros.io/)
- [Removaly](https://removaly.com/)
- [Reputation Defender](https://www.reputationdefender.com/)
- [Reputation.com](Reputation.com)
- [Spartacus](https://spartacus.com)

# Violations?

- https://www.ftccomplaintassistant.gov/?utm_source=takeaction#crnt&panel1-1
- https://oag.ca.gov/contact/consumer-complaint-against-business-or-company
- https://ec.europa.eu/info/law/law-topic/data-protection/reform/rights-citizens/redress/what-should-i-do-if-i-think-my-personal-data-protection-rights-havent-been-respected_en

# References
- Logo: [Ghost by Royyan Wijaya](https://thenounproject.com/icon/ghost-1358159/)
