package main

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v3"
)

//RateLimitConfig ...
type RateLimitConfig struct {
	Descriptors []ServiceDescriptors `json:"descriptors" yaml:"descriptors"`
}

//ServiceDescriptors ...
type ServiceDescriptors struct {
	Key         string        `json:"key" yaml:"key"`
	Value       string        `json:"value" yaml:"value"`
	Descriptors []Descriptors `json:"descriptors" yaml:"descriptors"`
}

//
type Descriptors struct {
	Key       string    `json:"key" yaml:"key"`
	Value     string    `json:"value" yaml:"value"`
	RateLimit RateLimit `json:"rate_limit" yaml:"rate_limit"`
}

type RateLimit struct {
	Unit           string `json:"unit" yaml:"unit"`
	RequestPerUnit string `json:"requests_per_unit" yaml:"requests_per_unit"`
}

func main() {
	objRateLimit := RateLimit{Unit: "sec", RequestPerUnit: "2"}
	objDescriptors := Descriptors{Key: "maps-ratelimit-by-clients", Value: "PRD1714_CTA", RateLimit: objRateLimit}
	objDescriptors2 := Descriptors{Key: "maps-ratelimit-by-clients", Value: "PRD1713_PLTS", RateLimit: objRateLimit}
	objServiceDescriptor := ServiceDescriptors{}
	objServiceDescriptor.Key = "generic_key"
	objServiceDescriptor.Value = "maps-ratelimit"
	objServiceDescriptor.Descriptors = append(objServiceDescriptor.Descriptors, objDescriptors)
	objServiceDescriptor.Descriptors = append(objServiceDescriptor.Descriptors, objDescriptors2)
	objRateLimitConfig := RateLimitConfig{}
	objRateLimitConfig.Descriptors = append(objRateLimitConfig.Descriptors, objServiceDescriptor)
	yfile2, err := ioutil.ReadFile("ratelimit.yaml")
	if err != nil {
		fmt.Println(err)
	}
	objRateLimitConfig2 := RateLimitConfig{}
	err2 := yaml.Unmarshal(yfile2, &objRateLimitConfig2)
	if err2 != nil {
		fmt.Println(err2)
	}
	processIndexerlist := -1
	for index, descriptor := range objRateLimitConfig2.Descriptors {
		if descriptor.Value == objRateLimitConfig.Descriptors[0].Value {
			processIndexerlist = index
			break
		}
	}
	if processIndexerlist > -1 {
		objRateLimitConfig2.Descriptors[processIndexerlist] = objRateLimitConfig.Descriptors[0]
	} else {
		objRateLimitConfig2.Descriptors = append(objRateLimitConfig2.Descriptors, objRateLimitConfig.Descriptors[0])
	}
	data, err := yaml.Marshal(&objRateLimitConfig2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
	err2 = ioutil.WriteFile("ratelimitnew.yaml", data, 0644)
	if err2 != nil {
		fmt.Println(err2)
	}
}
