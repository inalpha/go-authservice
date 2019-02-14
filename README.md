`env $(cat .env | xargs) go run *.go`


https://webauthn.guide
https://github.com/markbates/goth

OAuth for Authorization to access user data on the provider server 
OAuth for Authentication to user the provider to make sure the user is real (login) 


CREATE TABLE IF NOT EXISTS `AuthenticationProvider` (
`ProviderKey` varchar(128) NOT NULL,
`userId` int(10) unsigned NOT NULL,
`ProviderType` enum('facebook','twitter', 'google') NOT NULL, 
PRIMARY KEY  (`ProviderKey`) )  
ENGINE=MyISAM  DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `users` (
  `userId` int(10) unsigned NOT NULL auto_increment,
  `username` varchar(128) NOT NULL default '',
  `password` varchar(32) NOT NULL default '',
  `email` text NOT NULL,
  `newsletter` tinyint(1) NOT NULL default '0',
  `banned` enum('yes','no') NOT NULL default 'no',
  `admin` enum('yes','no') NOT NULL default 'no',
  `signup_ip` varchar(20) NOT NULL default '',
  `activation_key` varchar(60) NOT NULL default '',
  `resetpassword_key` varchar(60) NOT NULL default '',
  `createdon` datetime NOT NULL default '0000-00-00 00:00:00',
  PRIMARY KEY  (`userId`)
) ENGINE=MyISAM  DEFAULT CHARSET=latin1 AUTO_INCREMENT=27 ;