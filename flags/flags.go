package main

import "flag"

// Config will be the holder for our flags
type Config struct {
  subject      string
  isAwesome    bool
  howAwesome   int
  countTheWays CountTheWays
}

// Setup initializes a config from flags that are passed in
func (c *Config) Setup() {
  // you can set a flag directly like so:
  // var someVar = flag.String("flag_name", "default_val", "description")
  // but in practice putting it in a struct is generally better longhand
  flag.StringVar(&c.subject, "subject", "", "subject is a string, it defaults to empty")
  // shorthand
  flag.StringVar(&c.subject, "s", "", "subject is a string, it defaults to empty (shorthand)")
  
  flag.BoolVar(&c.isAwesome, "is_awesome", false, "is it awesomme or what?")
  flag.IntVar(&c.howAwesome, "how_awesome", 0 , "how awesome is?")

  // Custom variable type 
  flag.Var(&c.countTheWays, "c", "comma separated list of integers")
}

// GetMessage uses all of the internal config vars and returns a sentence 
func (c *Config) GetMessage() string {
  msg := c.subject
  if c.isAwesome {
    msg += " is awesome"
  } else {
    msg += " is not awesome :C"
  }
  msg = fmt.Sprintf("%s, with a certainty of %d out of 10. Let me count the ways %s", msg, c.howAwesome, c.countTheWays).String())
  return msg
} 

