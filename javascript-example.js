const cohere = require('cohere-ai');
cohere.init('{apiKey}');
(async () => {
  const response = await cohere.generate({
    model: 'large',
    prompt: 'This is a spell check generator that fixes my personal common mispellings!\n\nSample: In a paralel plate capacitor, a voltage applied between two conductive plates creates a uniform electic feild between the plates. The electic feild strength in a capacitor is inversely proportional to the distance between the paralel plates and directly proportional to the voltage applied.\nCorrected Text: In a parallel-plate capacitor, a voltage applied between two conductive plates creates a uniform electric field between the plates. The electric field strength in a capacitor is inversely proportional to the distance between the plates and directly proportional to the voltage applied.\n--\nSample: The farming feild was coverd in soybeans. They were placed in paralel after it was decided the the feilds were not in the correct spot after a land survey.\nCorrected Text: The farming field was coverd in soybeans. They were placed in parallel after it was decided the the fields were not in the correct spot after a land survey.\n--\nSample: I always have issues spelling feilds as I flip the 2nd and 3rd characters around as they are in paralel with each other. This example should fix the issues in both paralel and feilds.\nCorrected Text: I always have issues spelling fields as I flip the 2nd and 3rd characters around as they are in parallel with each other. This example should fix the issues in both parallel and fields.\n--\nSample: The feild of text was created to help try and fix the issues with mispelling words with paralel characters.\nCorrected Text:',
    max_tokens: 50,
    temperature: 0.3,
    k: 0,
    p: 0.75,
    frequency_penalty: 0,
    presence_penalty: 0,
    stop_sequences: ["--"],
    return_likelihoods: 'NONE'
  });
  console.log(`Prediction: ${response.body.generations[0].text}`);
})();
