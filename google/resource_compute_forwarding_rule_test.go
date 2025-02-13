package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccComputeForwardingRule_update(t *testing.T) {
	t.Parallel()

	poolName := fmt.Sprintf("tf-%s", RandString(t, 10))
	ruleName := fmt.Sprintf("tf-%s", RandString(t, 10))

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeForwardingRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeForwardingRule_basic(poolName, ruleName),
			},
			{
				ResourceName:      "google_compute_forwarding_rule.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeForwardingRule_update(poolName, ruleName),
			},
			{
				ResourceName:      "google_compute_forwarding_rule.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeForwardingRule_ip(t *testing.T) {
	t.Parallel()

	addrName := fmt.Sprintf("tf-%s", RandString(t, 10))
	poolName := fmt.Sprintf("tf-%s", RandString(t, 10))
	ruleName := fmt.Sprintf("tf-%s", RandString(t, 10))
	addressRefFieldRaw := "address"
	addressRefFieldID := "id"

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeForwardingRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeForwardingRule_ip(addrName, poolName, ruleName, addressRefFieldID),
			},
			{
				ResourceName:            "google_compute_forwarding_rule.foobar",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"ip_address"}, // ignore ip_address because we've specified it by ID
			},
			{
				Config: testAccComputeForwardingRule_ip(addrName, poolName, ruleName, addressRefFieldRaw),
			},
			{
				ResourceName:      "google_compute_forwarding_rule.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeForwardingRule_networkTier(t *testing.T) {
	t.Parallel()

	poolName := fmt.Sprintf("tf-%s", RandString(t, 10))
	ruleName := fmt.Sprintf("tf-%s", RandString(t, 10))

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeForwardingRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeForwardingRule_networkTier(poolName, ruleName),
			},

			{
				ResourceName:      "google_compute_forwarding_rule.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeForwardingRule_basic(poolName, ruleName string) string {
	return fmt.Sprintf(`
resource "google_compute_target_pool" "foo-tp" {
  description = "Resource created for Terraform acceptance testing"
  instances   = ["us-central1-a/foo", "us-central1-b/bar"]
  name        = "foo-%s"
}

resource "google_compute_forwarding_rule" "foobar" {
  description = "Resource created for Terraform acceptance testing"
  ip_protocol = "UDP"
  name        = "%s"
  port_range  = "80-81"
  target      = google_compute_target_pool.foo-tp.self_link
}
`, poolName, ruleName)
}

func testAccComputeForwardingRule_update(poolName, ruleName string) string {
	return fmt.Sprintf(`
resource "google_compute_target_pool" "foo-tp" {
  description = "Resource created for Terraform acceptance testing"
  instances   = ["us-central1-a/foo", "us-central1-b/bar"]
  name        = "foo-%s"
}

resource "google_compute_target_pool" "bar-tp" {
  description = "Resource created for Terraform acceptance testing"
  instances   = ["us-central1-a/foo", "us-central1-b/bar"]
  name        = "bar-%s"
}

resource "google_compute_forwarding_rule" "foobar" {
  description = "Resource created for Terraform acceptance testing"
  ip_protocol = "UDP"
  name        = "%s"
  port_range  = "80-81"
  target      = google_compute_target_pool.bar-tp.self_link
}
`, poolName, poolName, ruleName)
}

func testAccComputeForwardingRule_ip(addrName, poolName, ruleName, addressRefFieldValue string) string {
	return fmt.Sprintf(`
resource "google_compute_address" "foo" {
  name = "%s"
}

resource "google_compute_target_pool" "foobar-tp" {
  description = "Resource created for Terraform acceptance testing"
  instances   = ["us-central1-a/foo", "us-central1-b/bar"]
  name        = "%s"
}

resource "google_compute_forwarding_rule" "foobar" {
  description = "Resource created for Terraform acceptance testing"
  ip_address  = google_compute_address.foo.%s
  ip_protocol = "TCP"
  name        = "%s"
  port_range  = "80-81"
  target      = google_compute_target_pool.foobar-tp.self_link
}
`, addrName, poolName, addressRefFieldValue, ruleName)
}

func testAccComputeForwardingRule_networkTier(poolName, ruleName string) string {
	return fmt.Sprintf(`
resource "google_compute_target_pool" "foobar-tp" {
  description = "Resource created for Terraform acceptance testing"
  instances   = ["us-central1-a/foo", "us-central1-b/bar"]
  name        = "%s"
}

resource "google_compute_forwarding_rule" "foobar" {
  description = "Resource created for Terraform acceptance testing"
  ip_protocol = "UDP"
  name        = "%s"
  port_range  = "80-81"
  target      = google_compute_target_pool.foobar-tp.self_link

  network_tier = "STANDARD"
}
`, poolName, ruleName)
}
