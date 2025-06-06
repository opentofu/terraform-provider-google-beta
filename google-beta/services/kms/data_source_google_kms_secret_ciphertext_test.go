// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
// ----------------------------------------------------------------------------
//
//	***     AUTO GENERATED CODE    ***    Type: Handwritten     ***
//
// ----------------------------------------------------------------------------
//
//	This code is generated by Magic Modules using the following:
//
//	Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/services/kms/data_source_google_kms_secret_ciphertext_test.go
//
//	DO NOT EDIT this file directly. Any changes made to this file will be
//	overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------
package kms_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccDataKmsSecretCiphertext_basic(t *testing.T) {
	t.Parallel()

	kms := acctest.BootstrapKMSKey(t)

	plaintext := fmt.Sprintf("secret-%s", acctest.RandString(t, 10))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testGoogleKmsSecretCiphertext_datasource(kms.CryptoKey.Name, plaintext),
				Check: func(s *terraform.State) error {
					plaintext, err := testAccDecryptSecretDataWithCryptoKey(t, s, kms.CryptoKey.Name, "data.google_kms_secret_ciphertext.acceptance", "")

					if err != nil {
						return err
					}

					return resource.TestCheckResourceAttr("data.google_kms_secret_ciphertext.acceptance", "plaintext", plaintext)(s)
				},
			},
		},
	})
}

func testGoogleKmsSecretCiphertext_datasource(cryptoKeyTerraformId, plaintext string) string {
	return fmt.Sprintf(`
data "google_kms_secret_ciphertext" "acceptance" {
  crypto_key = "%s"
  plaintext  = "%s"
}
`, cryptoKeyTerraformId, plaintext)
}
