package main

import (
  "fmt"

  cohere "github.com/cohere-ai/cohere-go"
)

func main() {
  co, err := cohere.CreateClient("{apiKey}")
  if err != nil {
    fmt.Println(err)
    return
  }

  response, err := co.Generate(cohere.GenerateOptions{
    Model:             "large",
    Prompt:            `This is a spell check generator that fixes my personal common mispellings!\n\nSample: In a paralel plate capacitor, a voltage applied between two conductive plates creates a uniform electic feild between the plates. The electic feild strength in a capacitor is inversely proportional to the distance between the paralel plates and directly proportional to the voltage applied.\nCorrected Text: In a parallel-plate capacitor, a voltage applied between two conductive plates creates a uniform electric field between the plates. The electric field strength in a capacitor is inversely proportional to the distance between the plates and directly proportional to the voltage applied.\n--\nSample: The farming feild was coverd in soybeans. They were placed in paralel after it was decided the the feilds were not in the correct spot after a land survey.\nCorrected Text: The farming field was coverd in soybeans. They were placed in parallel after it was decided the the fields were not in the correct spot after a land survey.\n--\nSample: I always have issues spelling feilds as I flip the 2nd and 3rd characters around as they are in paralel with each other. This example should fix the issues in both paralel and feilds.\nCorrected Text: I always have issues spelling fields as I flip the 2nd and 3rd characters around as they are in parallel with each other. This example should fix the issues in both parallel and fields.\n--\nSample: The feild of text was created to help try and fix the issues with mispelling words with paralel characters.\nCorrected Text:`,
    MaxTokens:         50,
    Temperature:       0.3,
    NumGenerations:    {numGenerations},
    K:                 0,
    P:                 0.75,
    FrequencyPenalty:  0,
    PresencePenalty:   0,
    StopSequences:     []string{`--`},
    ReturnLikelihoods: "NONE",
  })
  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Println("Prediction:", response.Generations[0])
}
