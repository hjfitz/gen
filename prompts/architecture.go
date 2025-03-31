package prompts

import (
	"fmt"
	"strings"
)

func GetArchitectureGenPrompt(terraform string) string {
	prompt := `
	You are an expert in GCP infrastructure and architectural visualization. I will provide a multi-line string of Terraform code. Your job is to generate a high-level Mermaid diagram that shows the core GCP architecture, ignoring low-level implementation details.

Instructions:

    Include major components only — such as:

        Compute (e.g., google_cloudfunctions_function, google_compute_instance, google_cloud_run_service)

        Storage (e.g., google_storage_bucket)

        Networking/API (e.g., google_api_gateway_api, google_compute_forwarding_rule)

        Databases (e.g., google_sql_database_instance, google_firestore_database)

        Queues and messaging (e.g., google_pubsub_topic)

    Exclude resources like IAM, service accounts, firewall rules, VPCs, etc.

    Use graph TD or graph LR, whichever fits better.

    Label nodes like: Cloud Function: my_func or Storage Bucket: my_bucket


    Show directional relationships between services (e.g., a Pub/Sub topic triggering a Cloud Function).

    Color-code nodes by category:

        Compute: light-blue

        Storage: orange

        Networking/API: purple

        Database: green

        Messaging/queues: red

    Use Mermaid’s classDef and class features to apply color classes.

    Output just the mermaid code. Do **not** include the markdown block or any commentary. Treat this as plain text output.

	Your response should begin with graph TD (or graph LR) and contain no other formatting.

Here is the Terraform code:
{{terraform}}
`

	formatted_prompt := strings.ReplaceAll(prompt, "{{terraform}}", fmt.Sprintf("```\n%s\n```", terraform))

	return formatted_prompt

}
