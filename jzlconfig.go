package jzlconfig

import (
    "fmt"
    "github.com/ewangplay/config"
)

// JZLConfig implementaion
type JZLConfig struct {
    data map[string]string
}

func (this *JZLConfig) Read(filename string) error {
    this.data = make(map[string]string)
    if cfg, err := config.ReadDefault(filename); err == nil {
        sections := cfg.Sections()
        for _, section := range sections {
            if cfg.HasSection(section) {
                if options, err := cfg.SectionOptions(section); err == nil {
                    for _, key := range options {
                        if value, err := cfg.String(section, key); err == nil {
                            this.data[key] = value
                        }
                    }
                }
            }
        }
    } else {
        return err
    }

    return nil
}

func (this *JZLConfig) Get(key string) (string, bool) {
    if value, ok := this.data[key]; ok {
        return value, true
    }
    return "", false
}

func (this JZLConfig) String() string {
    var fmt_str string
    for k, v := range this.data {
        fmt_str += fmt.Sprintf("%v\t%v\n", k, v)
    }
    return fmt_str
}

