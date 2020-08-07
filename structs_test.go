package typeform

import (
	"bytes"
)

var (
	correctSubmission = bytes.NewBufferString(`{
  "event_id": "01EDZHFQ34Q7XDJ5BB0C17P9E4",
  "event_type": "form_response",
  "form_response": {
    "form_id": "VheyFb",
    "token": "e2rkl48t53kb8denkcsp1s8e2rkl48ty",
    "landed_at": "2020-07-24T04:30:17Z",
    "submitted_at": "2020-07-24T04:32:16Z",
    "hidden": {
      "t": "0",
      "user": "auth0|5e44a49f67b2aa0e992d2d8e"
    },
    "definition": {
      "id": "VheyFb",
      "title": "ToDo-CR RCT",
      "fields": [
        {
          "id": "DBM7YzrNm0HU",
          "title": "Is this the* first time* you're doing this questionnaire?",
          "type": "yes_no",
          "ref": "2ea90a9d-da76-49be-b334-0f71852d7a64",
          "properties": {}
        },
        {
          "id": "olFYPnoUHGIv",
          "title": "Think about how a friend, a partner or work colleague might describe you. Don't take too long. Just select the ones that feel right.",
          "type": "multiple_choice",
          "allow_multiple_selections": true,
          "ref": "0f074ed0-dabe-4ac8-841c-4e74d4ab4947",
          "properties": {},
          "choices": [
            {
              "id": "qFLRbq1L1X7H",
              "label": "Firm"
            },
            {
              "id": "RYIaUrU0u1it",
              "label": "Unpredictable"
            },
            {
              "id": "A0YeC1YudXCQ",
              "label": "Individually-centred"
            },
            {
              "id": "urzDgnqjise9",
              "label": "Behave as others want you to"
            },
            {
              "id": "Ll3XeudgM6fm",
              "label": "Behave as you wish"
            },
            {
              "id": "R1fSibpVlWq6",
              "label": "Reactive"
            },
            {
              "id": "q8w6J3m7j3BL",
              "label": "Lively"
            },
            {
              "id": "MnFPwA9yJsAc",
              "label": "Definite"
            },
            {
              "id": "VlGsO8CVBt8m",
              "label": "Calm/Relaxed"
            },
            {
              "id": "ONuFDaVNVbBO",
              "label": "Gentle"
            },
            {
              "id": "HQamCmVxMwyO",
              "label": "Play it safe"
            },
            {
              "id": "ctoF2gLvwXLA",
              "label": "Proactive"
            },
            {
              "id": "jXG3rGUFxwVW",
              "label": "Not lively/Laid back"
            },
            {
              "id": "nIj88CDCRXOy",
              "label": "Open-minded"
            },
            {
              "id": "WE5oDlCBeFwE",
              "label": "Assertive"
            },
            {
              "id": "g4jVwdXGfaBa",
              "label": "Introverted"
            },
            {
              "id": "r2IT5KzfOODv",
              "label": "Systematic"
            },
            {
              "id": "N4e8JUuXBDU0",
              "label": "Extroverted"
            },
            {
              "id": "ol0qPyZ6X1l8",
              "label": "Predictable"
            },
            {
              "id": "SpeJP9SujmmF",
              "label": "Conventional"
            },
            {
              "id": "mTIH8MMcfzv6",
              "label": "Flexible"
            },
            {
              "id": "nlPxCBOdRWjj",
              "label": "Trusting"
            },
            {
              "id": "YgizW5tyYZSY",
              "label": "Group-centred"
            },
            {
              "id": "VXEtMkpOZUdi",
              "label": "Spontaneous"
            },
            {
              "id": "tUOm1WwBe6bc",
              "label": "Risk-taker"
            },
            {
              "id": "e3DDABpis27O",
              "label": "Wary of others"
            },
            {
              "id": "HkbTyJzkYpsw",
              "label": "Unconventional"
            },
            {
              "id": "Q45YmJuW9Iik",
              "label": "Single-minded"
            },
            {
              "id": "aogPaCw2ivJd",
              "label": "Unassertive"
            },
            {
              "id": "wvTqSOY6Setr",
              "label": "Energetic/Driven"
            }
          ]
        },
        {
          "id": "C5d64ajwbqyN",
          "title": "read something about your medical condition?",
          "type": "opinion_scale",
          "ref": "e996510d-55ca-41aa-bd98-8d8d1aa63072",
          "properties": {}
        },
        {
          "id": "lrYlqHaPLrOM",
          "title": "spend most evenings watching TV or in front of a screen?",
          "type": "opinion_scale",
          "ref": "4b8e2508-1949-4c58-a027-a119d14f9fc2",
          "properties": {}
        },
        {
          "id": "ygVnMNN8mDhr",
          "title": "feel stressed or worry a lot?",
          "type": "opinion_scale",
          "ref": "237d73cc-2c1e-4c71-b4a3-5b4a3a6abf64",
          "properties": {}
        },
        {
          "id": "TGygGGCH2etq",
          "title": "spend at least half an hour a day being active? (e.g. sport, walking, cycling)",
          "type": "opinion_scale",
          "ref": "682789a8-ec3b-4eff-b535-02dff5f8f564",
          "properties": {}
        },
        {
          "id": "RVcEzUGU5HaZ",
          "title": "take time to connect with new people?",
          "type": "opinion_scale",
          "ref": "44a0c328-6795-4063-9ef5-01467ccf2097",
          "properties": {}
        },
        {
          "id": "EqRfVuXn9Ts6",
          "title": "take time to notice the good things in your life?",
          "type": "opinion_scale",
          "ref": "7a4449b1-cab6-43df-94c1-ac1a493a0742",
          "properties": {}
        },
        {
          "id": "fks9DnlvUlI8",
          "title": "feel too tired or too weak to exercise?",
          "type": "opinion_scale",
          "ref": "aa6ec172-de3f-4f6f-adf7-0bd63eb187f2",
          "properties": {}
        },
        {
          "id": "RZNbJhESQFib",
          "title": "take long naps during the day? (more than 1 hour)",
          "type": "opinion_scale",
          "ref": "3d8ff842-45b6-4a7c-8369-4364f3276f88",
          "properties": {}
        },
        {
          "id": "CE8J85Z8UcoU",
          "title": "engage in light or heavy housework? (cooking, cleaning, gardening)",
          "type": "opinion_scale",
          "ref": "31bdd262-e885-4aa0-8a01-da6ca110963f",
          "properties": {}
        },
        {
          "id": "ktbMTzhWfD9I",
          "title": "sit for long periods?",
          "type": "opinion_scale",
          "ref": "4a681824-7438-4233-be15-a2b1543f923c",
          "properties": {}
        },
        {
          "id": "AEjfeIG1Qzhb",
          "title": "get at least 7 hours sleep a night?",
          "type": "opinion_scale",
          "ref": "db639f0b-a361-4dd1-906e-81a8e5ea4bb5",
          "properties": {}
        }
      ]
    },
    "answers": [
      {
        "type": "boolean",
        "boolean": false,
        "field": {
          "id": "DBM7YzrNm0HU",
          "type": "yes_no",
          "ref": "2ea90a9d-da76-49be-b334-0f71852d7a64"
        }
      },
      {
        "type": "choices",
        "choices": {
          "labels": [
            "Calm/Relaxed"
          ]
        },
        "field": {
          "id": "olFYPnoUHGIv",
          "type": "multiple_choice",
          "ref": "0f074ed0-dabe-4ac8-841c-4e74d4ab4947"
        }
      },
      {
        "type": "number",
        "number": 4,
        "field": {
          "id": "C5d64ajwbqyN",
          "type": "opinion_scale",
          "ref": "e996510d-55ca-41aa-bd98-8d8d1aa63072"
        }
      },
      {
        "type": "number",
        "number": 6,
        "field": {
          "id": "lrYlqHaPLrOM",
          "type": "opinion_scale",
          "ref": "4b8e2508-1949-4c58-a027-a119d14f9fc2"
        }
      },
      {
        "type": "number",
        "number": 3,
        "field": {
          "id": "ygVnMNN8mDhr",
          "type": "opinion_scale",
          "ref": "237d73cc-2c1e-4c71-b4a3-5b4a3a6abf64"
        }
      },
      {
        "type": "number",
        "number": 4,
        "field": {
          "id": "TGygGGCH2etq",
          "type": "opinion_scale",
          "ref": "682789a8-ec3b-4eff-b535-02dff5f8f564"
        }
      },
      {
        "type": "number",
        "number": 2,
        "field": {
          "id": "RVcEzUGU5HaZ",
          "type": "opinion_scale",
          "ref": "44a0c328-6795-4063-9ef5-01467ccf2097"
        }
      },
      {
        "type": "number",
        "number": 4,
        "field": {
          "id": "EqRfVuXn9Ts6",
          "type": "opinion_scale",
          "ref": "7a4449b1-cab6-43df-94c1-ac1a493a0742"
        }
      },
      {
        "type": "number",
        "number": 2,
        "field": {
          "id": "fks9DnlvUlI8",
          "type": "opinion_scale",
          "ref": "aa6ec172-de3f-4f6f-adf7-0bd63eb187f2"
        }
      },
      {
        "type": "number",
        "number": 1,
        "field": {
          "id": "RZNbJhESQFib",
          "type": "opinion_scale",
          "ref": "3d8ff842-45b6-4a7c-8369-4364f3276f88"
        }
      },
      {
        "type": "number",
        "number": 5,
        "field": {
          "id": "CE8J85Z8UcoU",
          "type": "opinion_scale",
          "ref": "31bdd262-e885-4aa0-8a01-da6ca110963f"
        }
      },
      {
        "type": "number",
        "number": 5,
        "field": {
          "id": "ktbMTzhWfD9I",
          "type": "opinion_scale",
          "ref": "4a681824-7438-4233-be15-a2b1543f923c"
        }
      },
      {
        "type": "number",
        "number": 1,
        "field": {
          "id": "AEjfeIG1Qzhb",
          "type": "opinion_scale",
          "ref": "db639f0b-a361-4dd1-906e-81a8e5ea4bb5"
        }
      }
    ]
  }
}`).Bytes()

	emptySubmission = bytes.NewBufferString("").Bytes()

	faultySubmission = bytes.NewBufferString(`{
  "event_id": "01EDZHFQ34Q7XDJ5BB0C17P9E4",
  "event_type": "form_response",
  "form_response": {
    "form_id": "VheyFb",
    "token": "e2rkl48t53kb8denkcsp1s8e2rkl48ty",
    "landed_at": "2020-07-24T04:30:17Z",
    "submitted_at": "2020-07-24T04:32:16Z",
    "hidden": {
      "t": "0",
      "user": "auth0|5e44a49f67b2aa0e992d2d8e"
    },
    "definition": {
      "id": "VheyFb",
      "title": "ToDo-CR RCT",
      "fields": []
    },
    "answers": []
  }
}`).Bytes()

	missingUserSubmission = bytes.NewBufferString(`{
  "event_id": "01EDZHFQ34Q7XDJ5BB0C17P9E4",
  "event_type": "form_response",
  "form_response": {
    "form_id": "VheyFb",
    "token": "e2rkl48t53kb8denkcsp1s8e2rkl48ty",
    "landed_at": "2020-07-24T04:30:17Z",
    "submitted_at": "2020-07-24T04:32:16Z",
    "hidden": {
      "t": "0"
    },
    "definition": {
      "id": "VheyFb",
      "title": "ToDo-CR RCT",
      "fields": []
    },
    "answers": []
  }
}`).Bytes()

	missingTSubmission = bytes.NewBufferString(`{
  "event_id": "01EDZHFQ34Q7XDJ5BB0C17P9E4",
  "event_type": "form_response",
  "form_response": {
    "form_id": "VheyFb",
    "token": "e2rkl48t53kb8denkcsp1s8e2rkl48ty",
    "landed_at": "2020-07-24T04:30:17Z",
    "submitted_at": "2020-07-24T04:32:16Z",
    "hidden": {
      "user": "auth0|5e44a49f67b2aa0e992d2d8e"
    },
    "definition": {
      "id": "VheyFb",
      "title": "ToDo-CR RCT",
      "fields": []
    },
    "answers": []
	}
}`).Bytes()
)
