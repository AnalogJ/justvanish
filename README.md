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

# How?

There are a handful of regulations we can leverage to limit the ability of data brokers & other organizations to legally sell our data.
[General Data Protection Regulation (GDPR)](https://gdpr-info.eu/) and [California Consumer Privacy Act (CCPA)](https://oag.ca.gov/privacy/ccpa) are the two
most relevant, however there are other (sometimes less powerful) laws we can leverage as well (Colorado, Virginia, Utah, Connecticut, Indiana, and Ohio).


In general these laws provide privacy rights to consumers (GDPR is similar):

- The right to know about the personal information a business collects about them and how it is used and shared;
- The right to delete personal information collected from them (with some exceptions);
- The right to opt-out of the sale of their personal information; and 
- The right to non-discrimination for exercising their CCPA rights.

Unfortunately these laws only apply to California & European residents. Still, its a good foundation to build our application ontop
of these laws.

# Goals

The goal for this app is to provide users with a simple mechanism to:
- request a copy of your personal information stored by data brokers, government agencies and other organizations
- request that organizations restrict the collection & sale of your personal information
- request the deletion of your personal information from data brokers & other organizations

My intent is to have justVanish become a community maintained global registry of databrokers
This databroker & organization information is stored following a [structured schema](organization-schema.json) allowing
other tools to leverage it (similar to how AdBlock filter lists are maintained https://github.com/topics/adblock-list).

Ideally there would be a national registry where consumers could "opt-out" themselves, but until regulation supporting that passes,
the only alternative is for consumers to exercise their rights themselves. 

# Getting Started
