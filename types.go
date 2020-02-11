package main

type Configuration struct {
  Database struct {
    Host      string `yaml:"host"`
    Database  string `yaml:"database"`
    Port      string `yaml:"port"`
    User      string `yaml:"user"`
    Pass      string `yaml:"pass"`
  } `yaml:"database"`
}
