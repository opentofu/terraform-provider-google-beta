// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package dialogflowcx_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccDialogflowCXPage_dialogflowcxPageFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDialogflowCXPageDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDialogflowCXPage_dialogflowcxPageFullExample(context),
			},
			{
				ResourceName:            "google_dialogflow_cx_page.basic_page",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"knowledge_connector_settings.0.trigger_fulfillment.0.advanced_settings.0.logging_settings", "parent"},
			},
		},
	})
}

func testAccDialogflowCXPage_dialogflowcxPageFullExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_dialogflow_cx_agent" "agent" {
  display_name               = "tf-test-dialogflowcx-agent%{random_suffix}"
  location                   = "global"
  default_language_code      = "en"
  supported_language_codes   = ["fr", "de", "es"]
  time_zone                  = "America/New_York"
  description                = "Example description."
  avatar_uri                 = "https://cloud.google.com/_static/images/cloud/icons/favicons/onecloud/super_cloud.png"
  enable_stackdriver_logging = true
  enable_spell_correction    = true
  speech_to_text_settings {
    enable_speech_adaptation = true
  }
}


resource "google_dialogflow_cx_page" "basic_page" {
  parent       = google_dialogflow_cx_agent.agent.start_flow
  display_name = "MyPage"

  entry_fulfillment {
    messages {
      channel = "some-channel"
      text {
        text = ["Welcome to page"]
      }
    }
    messages {
      payload = <<EOF
        {"some-key": "some-value", "other-key": ["other-value"]}
      EOF
    }
    messages {
      conversation_success {
        metadata = <<EOF
          {"some-metadata-key": "some-value", "other-metadata-key": 1234}
        EOF
      }
    }
    messages {
      output_audio_text {
        text = "some output text"
      }
    }
    messages {
      output_audio_text {
        ssml = <<EOF
          <speak>Some example <say-as interpret-as="characters">SSML XML</say-as></speak>
        EOF
      }
    }
    messages {
      live_agent_handoff {
        metadata = <<EOF
          {"some-metadata-key": "some-value", "other-metadata-key": 1234}
        EOF
      }
    }
    messages {
      play_audio {
        audio_uri = "http://example.com/some-audio-file.mp3"
      }
    }
    messages {
      telephony_transfer_call {
        phone_number = "1-234-567-8901"
      }
    }

    set_parameter_actions {
      parameter = "some-param"
      value     = "123.45"
    }
    set_parameter_actions {
      parameter = "another-param"
      value     = jsonencode("abc")
    }
    set_parameter_actions {
      parameter = "other-param"
      value     = jsonencode(["foo"])
    }

    conditional_cases {
      cases = jsonencode([
        {
          condition = "$sys.func.RAND() < 0.5",
          caseContent = [
            {
              message = { text = { text = ["First case"] } }
            },
            {
              additionalCases = {
                cases = [
                  {
                    condition = "$sys.func.RAND() < 0.2"
                    caseContent = [
                      {
                        message = { text = { text = ["Nested case"] } }
                      }
                    ]
                  }
                ]
              }
            }
          ]
        },
        {
          caseContent = [
            {
              message = { text = { text = ["Final case"] } }
            }
          ]
        },
      ])
    }
  }

  event_handlers {
    event = "some-event"
    trigger_fulfillment {
      return_partial_responses = true
      messages {
        channel = "some-channel"
        text {
          text = ["Some text"]
        }
      }
      messages {
        payload = <<EOF
          {"some-key": "some-value", "other-key": ["other-value"]}
        EOF
      }
      messages {
        conversation_success {
          metadata = <<EOF
            {"some-metadata-key": "some-value", "other-metadata-key": 1234}
          EOF
        }
      }
      messages {
        output_audio_text {
          text = "some output text"
        }
      }
      messages {
        output_audio_text {
          ssml = <<EOF
            <speak>Some example <say-as interpret-as="characters">SSML XML</say-as></speak>
          EOF
        }
      }
      messages {
        live_agent_handoff {
          metadata = <<EOF
            {"some-metadata-key": "some-value", "other-metadata-key": 1234}
          EOF
        }
      }
      messages {
        play_audio {
          audio_uri = "http://example.com/some-audio-file.mp3"
        }
      }
      messages {
        telephony_transfer_call {
          phone_number = "1-234-567-8901"
        }
      }

      set_parameter_actions {
        parameter = "some-param"
        value     = "123.45"
      }
      set_parameter_actions {
        parameter = "another-param"
        value     = jsonencode("abc")
      }
      set_parameter_actions {
        parameter = "other-param"
        value     = jsonencode(["foo"])
      }

      conditional_cases {
        cases = jsonencode([
          {
            condition = "$sys.func.RAND() < 0.5",
            caseContent = [
              {
                message = { text = { text = ["First case"] } }
              },
              {
                additionalCases = {
                  cases = [
                    {
                      condition = "$sys.func.RAND() < 0.2"
                      caseContent = [
                        {
                          message = { text = { text = ["Nested case"] } }
                        }
                      ]
                    }
                  ]
                }
              }
            ]
          },
          {
            caseContent = [
              {
                message = { text = { text = ["Final case"] } }
              }
            ]
          },
        ])
      }
    }
  }

  form {
    parameters {
      display_name = "param1"
      entity_type  = "projects/-/locations/-/agents/-/entityTypes/sys.date"
      default_value = jsonencode("2000-01-01")
      fill_behavior {
        initial_prompt_fulfillment {
          messages {
            channel = "some-channel"
            text {
              text = ["Please provide param1"]
            }
          }
          messages {
            payload = <<EOF
              {"some-key": "some-value", "other-key": ["other-value"]}
            EOF
          }
          messages {
            conversation_success {
              metadata = <<EOF
                {"some-metadata-key": "some-value", "other-metadata-key": 1234}
              EOF
            }
          }
          messages {
            output_audio_text {
              text = "some output text"
            }
          }
          messages {
            output_audio_text {
              ssml = <<EOF
                <speak>Some example <say-as interpret-as="characters">SSML XML</say-as></speak>
              EOF
            }
          }
          messages {
            live_agent_handoff {
              metadata = <<EOF
                {"some-metadata-key": "some-value", "other-metadata-key": 1234}
              EOF
            }
          }
          messages {
            play_audio {
              audio_uri = "http://example.com/some-audio-file.mp3"
            }
          }
          messages {
            telephony_transfer_call {
              phone_number = "1-234-567-8901"
            }
          }

          set_parameter_actions {
            parameter = "some-param"
            value     = "123.45"
          }
          set_parameter_actions {
            parameter = "another-param"
            value     = jsonencode("abc")
          }
          set_parameter_actions {
            parameter = "other-param"
            value     = jsonencode(["foo"])
          }

          conditional_cases {
            cases = jsonencode([
              {
                condition = "$sys.func.RAND() < 0.5",
                caseContent = [
                  {
                    message = { text = { text = ["First case"] } }
                  },
                  {
                    additionalCases = {
                      cases = [
                        {
                          condition = "$sys.func.RAND() < 0.2"
                          caseContent = [
                            {
                              message = { text = { text = ["Nested case"] } }
                            }
                          ]
                        }
                      ]
                    }
                  }
                ]
              },
              {
                caseContent = [
                  {
                    message = { text = { text = ["Final case"] } }
                  }
                ]
              },
            ])
          }
        }
        reprompt_event_handlers {
          event = "sys.no-match-1"
          trigger_fulfillment {
            return_partial_responses = true
            webhook = google_dialogflow_cx_webhook.my_webhook.id
            tag = "some-tag"

            messages {
              channel = "some-channel"
              text {
                text = ["Please provide param1"]
              }
            }
            messages {
              payload = <<EOF
                {"some-key": "some-value", "other-key": ["other-value"]}
              EOF
            }
            messages {
              conversation_success {
                metadata = <<EOF
                  {"some-metadata-key": "some-value", "other-metadata-key": 1234}
                EOF
              }
            }
            messages {
              output_audio_text {
                text = "some output text"
              }
            }
            messages {
              output_audio_text {
                ssml = <<EOF
                  <speak>Some example <say-as interpret-as="characters">SSML XML</say-as></speak>
                EOF
              }
            }
            messages {
              live_agent_handoff {
                metadata = <<EOF
                  {"some-metadata-key": "some-value", "other-metadata-key": 1234}
                EOF
              }
            }
            messages {
              play_audio {
                audio_uri = "http://example.com/some-audio-file.mp3"
              }
            }
            messages {
              telephony_transfer_call {
                phone_number = "1-234-567-8901"
              }
            }

            set_parameter_actions {
              parameter = "some-param"
              value     = "123.45"
            }
            set_parameter_actions {
              parameter = "another-param"
              value     = jsonencode("abc")
            }
            set_parameter_actions {
              parameter = "other-param"
              value     = jsonencode(["foo"])
            }

            conditional_cases {
              cases = jsonencode([
                {
                  condition = "$sys.func.RAND() < 0.5",
                  caseContent = [
                    {
                      message = { text = { text = ["First case"] } }
                    },
                    {
                      additionalCases = {
                        cases = [
                          {
                            condition = "$sys.func.RAND() < 0.2"
                            caseContent = [
                              {
                                message = { text = { text = ["Nested case"] } }
                              }
                            ]
                          }
                        ]
                      }
                    }
                  ]
                },
                {
                  caseContent = [
                    {
                      message = { text = { text = ["Final case"] } }
                    }
                  ]
                },
              ])
            }
          }
        }
        reprompt_event_handlers {
          event = "sys.no-match-2"
          target_flow = google_dialogflow_cx_agent.agent.start_flow
        }
        reprompt_event_handlers {
          event = "sys.no-match-3"
          target_page = google_dialogflow_cx_page.my_page2.id
        }
      }
      required = "true"
      redact   = "true"
      advanced_settings {
        dtmf_settings {
          enabled      = true
          max_digits   = 1
          finish_digit = "#"
        }
      }
    }
  }

  transition_routes {
    condition = "$page.params.status = 'FINAL'"
    trigger_fulfillment {
      messages {
        channel = "some-channel"
        text {
          text = ["information completed, navigating to page 2"]
        }
      }
      messages {
        payload = <<EOF
          {"some-key": "some-value", "other-key": ["other-value"]}
        EOF
      }
      messages {
        conversation_success {
          metadata = <<EOF
            {"some-metadata-key": "some-value", "other-metadata-key": 1234}
          EOF
        }
      }
      messages {
        output_audio_text {
          text = "some output text"
        }
      }
      messages {
        output_audio_text {
          ssml = <<EOF
            <speak>Some example <say-as interpret-as="characters">SSML XML</say-as></speak>
          EOF
        }
      }
      messages {
        live_agent_handoff {
          metadata = <<EOF
            {"some-metadata-key": "some-value", "other-metadata-key": 1234}
          EOF
        }
      }
      messages {
        play_audio {
          audio_uri = "http://example.com/some-audio-file.mp3"
        }
      }
      messages {
        telephony_transfer_call {
          phone_number = "1-234-567-8901"
        }
      }

      set_parameter_actions {
        parameter = "some-param"
        value     = "123.45"
      }
      set_parameter_actions {
        parameter = "another-param"
        value     = jsonencode("abc")
      }
      set_parameter_actions {
        parameter = "other-param"
        value     = jsonencode(["foo"])
      }

      conditional_cases {
        cases = jsonencode([
          {
            condition = "$sys.func.RAND() < 0.5",
            caseContent = [
              {
                message = { text = { text = ["First case"] } }
              },
              {
                additionalCases = {
                  cases = [
                    {
                      condition = "$sys.func.RAND() < 0.2"
                      caseContent = [
                        {
                          message = { text = { text = ["Nested case"] } }
                        }
                      ]
                    }
                  ]
                }
              }
            ]
          },
          {
            caseContent = [
              {
                message = { text = { text = ["Final case"] } }
              }
            ]
          },
        ])
      }
    }
    target_page = google_dialogflow_cx_page.my_page2.id
  }

  advanced_settings {
    dtmf_settings {
      enabled      = true
      max_digits   = 1
      finish_digit = "#"
    }
  }
  knowledge_connector_settings {
    enabled = true
    trigger_fulfillment {
      messages {
        channel = "some-channel"
        knowledge_info_card {}
      }
      messages {
        channel = "some-channel"
        text {
          text = ["information completed, navigating to page 2"]
        }
      }
      messages {
        payload = <<EOF
          {"some-key": "some-value", "other-key": ["other-value"]}
        EOF
      }
      messages {
        conversation_success {
          metadata = <<EOF
            {"some-metadata-key": "some-value", "other-metadata-key": 1234}
          EOF
        }
      }
      messages {
        output_audio_text {
          text = "some output text"
        }
      }
      messages {
        output_audio_text {
          ssml = <<EOF
            <speak>Some example <say-as interpret-as="characters">SSML XML</say-as></speak>
          EOF
        }
      }
      messages {
        live_agent_handoff {
          metadata = <<EOF
            {"some-metadata-key": "some-value", "other-metadata-key": 1234}
          EOF
        }
      }
      messages {
        play_audio {
          audio_uri = "http://example.com/some-audio-file.mp3"
        }
      }
      messages {
        telephony_transfer_call {
          phone_number = "1-234-567-8902"
        }
      }
      webhook = google_dialogflow_cx_webhook.my_webhook.id
      return_partial_responses = true
      tag = "some-tag"
      set_parameter_actions {
        parameter = "some-param"
        value     = "123.45"
      }
      conditional_cases {
        cases = jsonencode([
          {
            condition = "$sys.func.RAND() < 0.5",
            caseContent = [
              {
                message = { text = { text = ["First case"] } }
              }
            ]
          },
          {
            caseContent = [
              {
                message = { text = { text = ["Final case"] } }
              }
            ]
          },
        ])
      }
      advanced_settings {
        speech_settings {
          endpointer_sensitivity        = 30
          no_speech_timeout             = "3.500s"
          use_timeout_based_endpointing = true
          models = {
            name : "wrench"
            mass : "1.3kg"
            count : "3"
          }
        }
        dtmf_settings {
          enabled      = true
          max_digits   = 1
          finish_digit = "#"
          interdigit_timeout_duration = "3.500s"
          endpointing_timeout_duration = "3.500s"
        }
        logging_settings {
          enable_stackdriver_logging     = true
          enable_interaction_logging     = true
          enable_consent_based_redaction = true
        }
      }
      enable_generative_fallback = true
    }
    data_store_connections {
      data_store_type = "PUBLIC_WEB"
      data_store = "projects/${data.google_project.project.number}/locations/${google_dialogflow_cx_agent.agent.location}/collections/default_collection/dataStores/datastore-page-full"
      document_processing_mode = "CHUNKS"
    }
    target_page = google_dialogflow_cx_page.my_page2.id
  }
}

resource "google_dialogflow_cx_page" "my_page2" {
  parent       = google_dialogflow_cx_agent.agent.start_flow
  display_name = "MyPage2"
}

resource "google_discovery_engine_data_store" "my_datastore" {
  location          = "global"
  data_store_id     = "tf-test-datastore-page-full%{random_suffix}"
  display_name      = "datastore-page-full"
  industry_vertical = "GENERIC"
  content_config    = "NO_CONTENT"
}

resource "google_dialogflow_cx_webhook" "my_webhook" {
  parent       = google_dialogflow_cx_agent.agent.id
  display_name = "MyWebhook"
  generic_web_service {
    uri = "https://example.com"
  }
}

data "google_project" "project" {
}
`, context)
}

func testAccCheckDialogflowCXPageDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_dialogflow_cx_page" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{DialogflowCXBasePath}}{{parent}}/pages/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("DialogflowCXPage still exists at %s", url)
			}
		}

		return nil
	}
}
