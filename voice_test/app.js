'use strict';

const message = setupVoice();
const recognition = new webkitSpeechRecognition();
const num = getArbitraryNumber();
console.log(num);
const jokeArray = [{ part1: "canoe",
                     part2: "Canoe help me with my homework?" },
                   { part1: "merry",
                     part2: "Merry Christmas" },
                   { part1: "orange",
                     part2: "orange you going to let me in?" },
                   { part1: "anee",
                     part2: "Anee one you like" },
                   { part1: "iva",
                     part2: "I've a sore hand from knocking!" },
                   { part1: "needle",
                     part2: "Needle little money for the movies" },
                   { part1: "henrietta",
                     part2: "Henrietta worm that was in his apple" },
                   { part1: "avenue",
                     part2: "Avenue knocked on this door before?" },
                   { part1: "harry",
                     part2: "Harry up, it’s cold out here!" },
                   { part1: "a herd",
                     part2: "A herd you were home, so I came over!" },
                   { part1: "adore",
                     part2: "Adore is between us. Open up!" },
                   { part1: "otto",
                     part2: "Otto know. I’ve got amnesia." },
                   { part1: "king Tutt",
                     part2: "King Tutt-key fried chicken" },
                   { part1: "lettuce",
                     part2: "Lettuce in it’s cold out here" },
                   { part1: "noah",
                     part2: "Noah good place we can get something to eat?" },
                   { part1: "robin",
                     part2: "Robin the piggy bank again." },
                   { part1: "dwayne",
                     part2: "Dwayne the bathtub, It’s overflowing!" }]

window.onload = () => {
  sayKnockKnock()
    .then(listenForSpeech)
    .then(respondWithJoke)
    .then(listenForSpeech)
    .then(respondWithPunchline);
}

function getArbitraryNumber() {
  return Math.floor(Math.random() * (16 - 0) + 0);
}

function setupVoice() {
  const message = new SpeechSynthesisUtterance();
  const voices = window.speechSynthesis.getVoices();
  message.voice = voices[5];
  message.volume = 1;  // 0 to 1
  message.rate = 1;  // 0.1 to 10
  message.pitch = 1;   //0 to 2
  message.lang = 'en-US';
  return message;
}

function sayKnockKnock() {
  const promise = new Promise(resolve => message.onend = resolve);
  message.text = "Knock knock...";
  speechSynthesis.speak(message);
  return promise;
}

function listenForSpeech() {
  const promise = new Promise(resolve => recognition.onresult = resolve);
  recognition.start();
  return promise;
}

function respondWithJoke(event) {
  const promise = new Promise(resolve => message.onend = resolve);
  const spoken = event.results[0][0].transcript;
  const confidence = event.results[0][0].confidence;

  if (spoken.toLowerCase() == "who's there") {
    message.text = jokeArray[num].part1;
  } else if (confidence < 0.5) {
    message.text = "I couldn't really understand what you said, so I'll say it for you. Who's there? Then I'll say," + jokeArray[num].part1;
  } else {
    message.text = "Silly human, don't you understand knock knock jokes? You're supposed to say who's there. Then I say," + jokeArray[num].part1;
  }

  speechSynthesis.speak(message);
  console.log(spoken, confidence);
  return promise;
}

function respondWithPunchline(event) {
  const promise = new Promise(resolve => message.onend = resolve);
  const spoken = event.results[0][0].transcript;
  const confidence = event.results[0][0].confidence;

  if (spoken.toLowerCase() == jokeArray[num].part1 + "who") {
    message.text = jokeArray[num].part2;
  } else if (confidence < 0.5) {
    message.text = "I couldn't really understand what you said, so I'll say it for you."  + jokeArray[num].part1 + "who? Then I'll say," + jokeArray[num].part2;
  } else {
    message.text = "You're really messing this joke up. You're supposed to say" + jokeArray[num].part1 + "who? Then I'll say," + jokeArray[num].part2;
  }

  speechSynthesis.speak(message);
  console.log(spoken, confidence);
  return promise;
}
