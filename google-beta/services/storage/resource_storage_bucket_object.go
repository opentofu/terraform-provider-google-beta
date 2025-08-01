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
//	Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/services/storage/resource_storage_bucket_object.go
//
//	DO NOT EDIT this file directly. Any changes made to this file will be
//	overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------
// SPDX-License-Identifier: MPL-2.0
package storage

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"net/http"

	"google.golang.org/api/googleapi"
	"google.golang.org/api/storage/v1"
)

func ResourceStorageBucketObject() *schema.Resource {
	return &schema.Resource{
		Create:        resourceStorageBucketObjectCreate,
		Read:          resourceStorageBucketObjectRead,
		Update:        resourceStorageBucketObjectUpdate,
		Delete:        resourceStorageBucketObjectDelete,
		CustomizeDiff: resourceStorageBucketObjectCustomizeDiff,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"bucket": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The name of the containing bucket.`,
			},

			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The name of the object. If you're interpolating the name of this object, see output_name instead.`,
			},

			"cache_control": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Optional:    true,
				Description: `Cache-Control directive to specify caching behavior of object data. If omitted and object is accessible to all anonymous users, the default will be public, max-age=3600`,
			},

			"content_disposition": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Optional:    true,
				Description: `Content-Disposition of the object data.`,
			},

			"content_encoding": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Optional:    true,
				Description: `Content-Encoding of the object data.`,
			},

			"content_language": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Optional:    true,
				Description: `Content-Language of the object data.`,
			},

			"content_type": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				Computed:      true,
				ConflictsWith: []string{"force_empty_content_type"},
				Description:   `Content-Type of the object data. Defaults to "application/octet-stream" or "text/plain; charset=utf-8".`,
			},

			"force_empty_content_type": {
				Type:          schema.TypeBool,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"content_type"},
				Description:   `Flag to set empty Content-Type.`,
			},

			"content": {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"source"},
				Sensitive:    true,
				Computed:     true,
				Description:  `Data as string to be uploaded. Must be defined if source is not. Note: The content field is marked as sensitive. To view the raw contents of the object, please define an output.`,
			},

			"generation": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `The content generation of this object. Used for object versioning and soft delete.`,
			},

			"crc32c": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Base 64 CRC32 hash of the uploaded data.`,
			},

			"md5hash": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Base 64 MD5 hash of the uploaded data.`,
			},

			"md5hexhash": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    false,
				Required:    false,
				Description: `Hex value of md5hash`,
			},

			"source": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ExactlyOneOf: []string{"content"},
				Description:  `A path to the data you want to upload. Must be defined if content is not.`,
			},

			"source_md5hash": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `User-provided md5hash, Base 64 MD5 hash of the object data.`,
			},

			// Detect changes to local file or changes made outside of Terraform to the file stored on the server.
			"detect_md5hash": {
				Type:       schema.TypeString,
				Deprecated: "`detect_md5hash` is deprecated and will be removed in future release. Start using `source_md5hash` instead",
				// This field is not Computed because it needs to trigger a diff.
				Optional: true,
				// Makes the diff message nicer:
				// detect_md5hash:       "1XcnP/iFw/hNrbhXi7QTmQ==" => "different hash" (forces new resource)
				// Instead of the more confusing:
				// detect_md5hash:       "1XcnP/iFw/hNrbhXi7QTmQ==" => "" (forces new resource)
				Default: "different hash",
				// 1. Compute the md5 hash of the local file
				// 2. Compare the computed md5 hash with the hash stored in Cloud Storage
				// 3. Don't suppress the diff iff they don't match
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					localMd5Hash := ""
					if d.GetRawConfig().GetAttr("source_md5hash") == cty.UnknownVal(cty.String) {
						return true
					}
					if v, ok := d.GetOk("source_md5hash"); ok && v != "" {
						return true
					}
					if source, ok := d.GetOkExists("source"); ok {
						localMd5Hash = tpgresource.GetFileMd5Hash(source.(string))
					}

					if content, ok := d.GetOkExists("content"); ok {
						localMd5Hash = tpgresource.GetContentMd5Hash([]byte(content.(string)))
					}

					// If `source` or `content` is dynamically set, both field will be empty.
					// We should not suppress the diff to avoid the following error:
					// 'Mismatch reason: extra attributes: detect_md5hash'
					if localMd5Hash == "" {
						return false
					}

					// `old` is the md5 hash we retrieved from the server in the ReadFunc
					if old != localMd5Hash {
						return false
					}

					return true
				},
			},

			"storage_class": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: `The StorageClass of the new bucket object. Supported values include: MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE. If not provided, this defaults to the bucket's default storage class or to a standard class.`,
			},

			"kms_key_name": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				Computed:         true,
				ConflictsWith:    []string{"customer_encryption"},
				DiffSuppressFunc: tpgresource.CompareCryptoKeyVersions,
				Description:      `Resource name of the Cloud KMS key that will be used to encrypt the object. Overrides the object metadata's kmsKeyName value, if any.`,
			},

			"customer_encryption": {
				Type:          schema.TypeList,
				MaxItems:      1,
				Optional:      true,
				Sensitive:     true,
				ConflictsWith: []string{"kms_key_name"},
				Description:   `Encryption key; encoded using base64.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"encryption_algorithm": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "AES256",
							ForceNew:    true,
							Description: `The encryption algorithm. Default: AES256`,
						},
						"encryption_key": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Sensitive:   true,
							Description: `Base64 encoded customer supplied encryption key.`,
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								_, err := base64.StdEncoding.DecodeString(val.(string))
								if err != nil {
									errs = append(errs, fmt.Errorf("Failed to decode (base64) customer_encryption, expecting valid base64 encoded key"))
								}
								return
							},
						},
					},
				},
			},

			"retention": {
				Type:          schema.TypeList,
				MaxItems:      1,
				Optional:      true,
				ConflictsWith: []string{"event_based_hold"},
				Description:   `Object level retention configuration.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"retain_until_time": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `Time in RFC 3339 (e.g. 2030-01-01T02:03:04Z) until which object retention protects this object.`,
						},
						"mode": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `The object retention mode. Supported values include: "Unlocked", "Locked".`,
						},
					},
				},
			},

			"event_based_hold": {
				Type:          schema.TypeBool,
				Optional:      true,
				ConflictsWith: []string{"retention"},
				Description:   `Whether an object is under event-based hold. Event-based hold is a way to retain objects until an event occurs, which is signified by the hold's release (i.e. this value is set to false). After being released (set to false), such objects will be subject to bucket-level retention (if any).`,
			},

			"temporary_hold": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `Whether an object is under temporary hold. While this flag is set to true, the object is protected against deletion and overwrites.`,
			},

			"metadata": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: `User-provided metadata, in key/value pairs.`,
			},

			"self_link": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `A url reference to this object.`,
			},

			// https://github.com/hashicorp/terraform/issues/19052
			"output_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The name of the object. Use this field in interpolations with google_storage_object_acl to recreate google_storage_object_acl resources when your google_storage_bucket_object is recreated.`,
			},

			"media_link": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `A url reference to download this object.`,
			},

			"deletion_policy": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  `The deletion policy for the object. Setting ABANDON allows the resource to be abandoned rather than deleted when removed from your Terraform configuration.`,
				ValidateFunc: validation.StringInSlice([]string{"ABANDON"}, false),
			},
		},
		UseJSONNumber: true,
	}
}

func objectGetID(object *storage.Object) string {
	return object.Bucket + "-" + object.Name
}

func resourceStorageBucketObjectCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	bucket := d.Get("bucket").(string)
	name := d.Get("name").(string)
	var media io.Reader

	if v, ok := d.GetOk("source"); ok {
		var err error
		media, err = os.Open(v.(string))
		if err != nil {
			return err
		}
	} else if v, ok := d.GetOk("content"); ok {
		media = bytes.NewReader([]byte(v.(string)))
	} else {
		return fmt.Errorf("Error, either \"content\" or \"source\" must be specified")
	}

	objectsService := storage.NewObjectsService(config.NewStorageClientWithTimeoutOverride(userAgent, d.Timeout(schema.TimeoutCreate)))
	object := &storage.Object{Bucket: bucket}

	if v, ok := d.GetOk("cache_control"); ok {
		object.CacheControl = v.(string)
	}

	if v, ok := d.GetOk("content_disposition"); ok {
		object.ContentDisposition = v.(string)
	}

	if v, ok := d.GetOk("content_encoding"); ok {
		object.ContentEncoding = v.(string)
	}

	if v, ok := d.GetOk("content_language"); ok {
		object.ContentLanguage = v.(string)
	}

	if v, ok := d.GetOk("content_type"); ok {
		object.ContentType = v.(string)
	}

	if v, ok := d.GetOk("metadata"); ok {
		object.Metadata = tpgresource.ConvertStringMap(v.(map[string]interface{}))
	}

	if v, ok := d.GetOk("storage_class"); ok {
		object.StorageClass = v.(string)
	}

	if v, ok := d.GetOk("kms_key_name"); ok {
		object.KmsKeyName = v.(string)
	}

	if v, ok := d.GetOk("retention"); ok {
		object.Retention = expandObjectRetention(v)
	}

	if v, ok := d.GetOk("event_based_hold"); ok {
		object.EventBasedHold = v.(bool)
	}

	if v, ok := d.GetOk("temporary_hold"); ok {
		object.TemporaryHold = v.(bool)
	}

	insertCall := objectsService.Insert(bucket, object)
	insertCall.Name(name)
	if v, ok := d.GetOk("force_empty_content_type"); ok && v.(bool) {
		insertCall.Media(media, googleapi.ContentType(""))
	} else {
		insertCall.Media(media)
	}

	// This is done late as we need to add headers to enable customer encryption
	if v, ok := d.GetOk("customer_encryption"); ok {
		customerEncryption := expandCustomerEncryption(v.([]interface{}))
		setEncryptionHeaders(customerEncryption, insertCall.Header())
	}

	_, err = insertCall.Do()

	if err != nil {
		return fmt.Errorf("Error uploading object %s: %s", name, err)
	}

	return resourceStorageBucketObjectRead(d, meta)
}

func resourceStorageBucketObjectUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	bucket := d.Get("bucket").(string)
	name := d.Get("name").(string)

	if d.HasChange("content") || d.HasChange("source_md5hash") || d.HasChange("detect_md5hash") {
		// The KMS key name are not able to be set on create :
		// or you get error: Error uploading object test-maarc: googleapi: Error 400: Malformed Cloud KMS crypto key: projects/myproject/locations/myregion/keyRings/mykeyring/cryptoKeys/mykeyname/cryptoKeyVersions/1, invalid
		d.Set("kms_key_name", nil)
		return resourceStorageBucketObjectCreate(d, meta)
	} else {

		objectsService := storage.NewObjectsService(config.NewStorageClientWithTimeoutOverride(userAgent, d.Timeout(schema.TimeoutUpdate)))
		getCall := objectsService.Get(bucket, name)

		res, err := getCall.Do()
		if err != nil {
			return fmt.Errorf("Error retrieving object during update %s: %s", name, err)
		}

		hasRetentionChanges := d.HasChange("retention")
		if hasRetentionChanges {
			if v, ok := d.GetOk("retention"); ok {
				res.Retention = expandObjectRetention(v)
			} else {
				res.Retention = nil
				res.NullFields = append(res.NullFields, "Retention")
			}
		}

		if d.HasChange("event_based_hold") {
			v := d.Get("event_based_hold")
			res.EventBasedHold = v.(bool)
		}

		if d.HasChange("temporary_hold") {
			v := d.Get("temporary_hold")
			res.TemporaryHold = v.(bool)
		}

		updateCall := objectsService.Update(bucket, name, res)
		if hasRetentionChanges {
			updateCall.OverrideUnlockedRetention(true)
		}
		_, err = updateCall.Do()

		if err != nil {
			return fmt.Errorf("Error updating object %s: %s", name, err)
		}

		return nil
	}
}

func resourceStorageBucketObjectRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	bucket := d.Get("bucket").(string)
	name := d.Get("name").(string)

	objectsService := storage.NewObjectsService(config.NewStorageClientWithTimeoutOverride(userAgent, d.Timeout(schema.TimeoutRead)))
	getCall := objectsService.Get(bucket, name)

	if v, ok := d.GetOk("customer_encryption"); ok {
		customerEncryption := expandCustomerEncryption(v.([]interface{}))
		setEncryptionHeaders(customerEncryption, getCall.Header())
	}

	res, err := getCall.Do()

	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("Storage Bucket Object %q", d.Get("name").(string)))
	}

	if err := d.Set("md5hash", res.Md5Hash); err != nil {
		return fmt.Errorf("Error setting md5hash: %s", err)
	}
	hash, err := base64.StdEncoding.DecodeString(res.Md5Hash)
	if err != nil {
		return fmt.Errorf("Error decoding md5hash: %s", err)
	}
	// encode
	md5HexHash := hex.EncodeToString(hash)
	if err := d.Set("md5hexhash", md5HexHash); err != nil {
		return fmt.Errorf("Error setting md5hexhash: %s", err)
	}
	if err := d.Set("detect_md5hash", res.Md5Hash); err != nil {
		return fmt.Errorf("Error setting detect_md5hash: %s", err)
	}
	if err := d.Set("source_md5hash", d.Get("source_md5hash")); err != nil {
		return fmt.Errorf("Error setting source_md5hash: %s", err)
	}
	if err := d.Set("generation", res.Generation); err != nil {
		return fmt.Errorf("Error setting generation: %s", err)
	}
	if err := d.Set("crc32c", res.Crc32c); err != nil {
		return fmt.Errorf("Error setting crc32c: %s", err)
	}
	if err := d.Set("cache_control", res.CacheControl); err != nil {
		return fmt.Errorf("Error setting cache_control: %s", err)
	}
	if err := d.Set("content_disposition", res.ContentDisposition); err != nil {
		return fmt.Errorf("Error setting content_disposition: %s", err)
	}
	if err := d.Set("content_encoding", res.ContentEncoding); err != nil {
		return fmt.Errorf("Error setting content_encoding: %s", err)
	}
	if err := d.Set("content_language", res.ContentLanguage); err != nil {
		return fmt.Errorf("Error setting content_language: %s", err)
	}
	if err := d.Set("content_type", res.ContentType); err != nil {
		return fmt.Errorf("Error setting content_type: %s", err)
	}
	if err := d.Set("storage_class", res.StorageClass); err != nil {
		return fmt.Errorf("Error setting storage_class: %s", err)
	}
	if err := d.Set("kms_key_name", res.KmsKeyName); err != nil {
		return fmt.Errorf("Error setting kms_key_name: %s", err)
	}
	if err := d.Set("self_link", res.SelfLink); err != nil {
		return fmt.Errorf("Error setting self_link: %s", err)
	}
	if err := d.Set("output_name", res.Name); err != nil {
		return fmt.Errorf("Error setting output_name: %s", err)
	}
	if err := d.Set("metadata", res.Metadata); err != nil {
		return fmt.Errorf("Error setting metadata: %s", err)
	}
	if err := d.Set("media_link", res.MediaLink); err != nil {
		return fmt.Errorf("Error setting media_link: %s", err)
	}
	if err := d.Set("retention", flattenObjectRetention(res.Retention)); err != nil {
		return fmt.Errorf("Error setting retention: %s", err)
	}
	if err := d.Set("event_based_hold", res.EventBasedHold); err != nil {
		return fmt.Errorf("Error setting event_based_hold: %s", err)
	}
	if err := d.Set("temporary_hold", res.TemporaryHold); err != nil {
		return fmt.Errorf("Error setting temporary_hold: %s", err)
	}

	d.SetId(objectGetID(res))

	return nil
}

func resourceStorageBucketObjectDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	if deletionPolicy := d.Get("deletion_policy"); deletionPolicy == "ABANDON" {
		log.Printf("[WARN] Object %q deletion_policy is set to 'ABANDON', object deletion has been abandoned", d.Id())
		d.SetId("")
		return nil
	}

	bucket := d.Get("bucket").(string)
	name := d.Get("name").(string)

	objectsService := storage.NewObjectsService(config.NewStorageClientWithTimeoutOverride(userAgent, d.Timeout(schema.TimeoutDelete)))

	DeleteCall := objectsService.Delete(bucket, name)
	err = DeleteCall.Do()

	if err != nil {
		if gerr, ok := err.(*googleapi.Error); ok && gerr.Code == 404 {
			log.Printf("[WARN] Removing Bucket Object %q because it's gone", name)
			// The resource doesn't exist anymore
			d.SetId("")

			return nil
		}

		return fmt.Errorf("Error deleting contents of object %s: %s", name, err)
	}

	return nil
}

func setEncryptionHeaders(customerEncryption map[string]string, headers http.Header) {
	decodedKey, _ := base64.StdEncoding.DecodeString(customerEncryption["encryption_key"])
	keyHash := sha256.Sum256(decodedKey)
	headers.Set("x-goog-encryption-algorithm", customerEncryption["encryption_algorithm"])
	headers.Set("x-goog-encryption-key", customerEncryption["encryption_key"])
	headers.Set("x-goog-encryption-key-sha256", base64.StdEncoding.EncodeToString(keyHash[:]))
}

func getFileMd5Hash(filename string) string {
	return tpgresource.GetFileMd5Hash(filename)
}

func getContentMd5Hash(content []byte) string {
	return tpgresource.GetContentMd5Hash(content)
}

func expandCustomerEncryption(input []interface{}) map[string]string {
	expanded := make(map[string]string)
	if input == nil {
		return expanded
	}
	for _, v := range input {
		original := v.(map[string]interface{})
		expanded["encryption_key"] = original["encryption_key"].(string)
		expanded["encryption_algorithm"] = original["encryption_algorithm"].(string)
	}
	return expanded
}

func expandObjectRetention(configured interface{}) *storage.ObjectRetention {
	retentions := configured.([]interface{})
	if len(retentions) == 0 {
		return nil
	}
	retention := retentions[0].(map[string]interface{})

	objectRetention := &storage.ObjectRetention{
		RetainUntilTime: retention["retain_until_time"].(string),
		Mode:            retention["mode"].(string),
	}

	return objectRetention
}

func flattenObjectRetention(objectRetention *storage.ObjectRetention) []map[string]interface{} {
	retentions := make([]map[string]interface{}, 0, 1)

	if objectRetention == nil {
		return retentions
	}

	retention := map[string]interface{}{
		"mode":              objectRetention.Mode,
		"retain_until_time": objectRetention.RetainUntilTime,
	}

	retentions = append(retentions, retention)
	return retentions
}

func resourceStorageBucketObjectCustomizeDiff(ctx context.Context, d *schema.ResourceDiff, meta interface{}) error {
	localMd5Hash := ""

	if (d.GetRawConfig().GetAttr("source_md5hash") == cty.UnknownVal(cty.String)) || d.HasChange("source_md5hash") {
		return showDiff(d)
	}

	if source, ok := d.GetOkExists("source"); ok {
		localMd5Hash = tpgresource.GetFileMd5Hash(source.(string))
	}
	if content, ok := d.GetOkExists("content"); ok {
		localMd5Hash = tpgresource.GetContentMd5Hash([]byte(content.(string)))
	}
	if localMd5Hash == "" {
		return nil
	}

	oldMd5Hash, ok := d.GetOkExists("md5hash")
	if ok && oldMd5Hash == localMd5Hash {
		return nil
	}
	return showDiff(d)
}

func showDiff(d *schema.ResourceDiff) error {
	err := d.SetNewComputed("md5hash")
	if err != nil {
		return fmt.Errorf("Error re-setting md5hash: %s", err)
	}
	err = d.SetNewComputed("crc32c")
	if err != nil {
		return fmt.Errorf("Error re-setting crc32c: %s", err)
	}
	err = d.SetNewComputed("generation")
	if err != nil {
		return fmt.Errorf("Error re-setting generation: %s", err)
	}

	return nil
}
