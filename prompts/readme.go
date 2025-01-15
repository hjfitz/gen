package prompts

import (
	"strings"
)

func GetReadme(codebase string, useTrump bool) string {
	rp := "Generate a detailed and professional README.md file for a codebase based on the following input. The codebase is small and consists of some components. They may be a CLI tool, a REST API, or a framework. It may be any combination of these.\n" +
	"\n" +
  "```\n" +
  "{{codebase}}\n"  +
  "```\n"  +
  "The string includes the filename, the code itself, and a separator (---) between files. Use the provided content to:\n\n" +
 
 "Describe the purpose of the codebase and its functionality. It may be an API, framework, CLI, or a combination of the above. Explain:\n" +
 "- Its purpose\n" +
 "- How to configure and install or run it\n" +
 "- Examples of usage, especially if this is a framework or CLI. Any API requests should have their examples with the node package `axios`\n" +
 "- Available endpoints and their expected input/output\n\n" +
 "- Provide any necessary setup or configuration instructions (e.g., environment variables, dependencies).\n" +
 "- If the codebase is a library, build documentation around all of the exports, but ensure that the main functionality is covered first in its own section.\n\n" +

 "Ensure the README is structured, beginner-friendly, and detailed enough to help someone new to the project get started. Use markdown formatting for headers, code blocks, lists, tables, etc.\n\n" +

 "If the codebase represents a system with distinct components or containers (e.g., services, databases, external APIs), generate a **Mermaid** container-level diagram under a section titled 'Architecture Overview.' The diagram should represent:\n" +
 "- The main containers or components in the system\n" +
 "- The relationships or communication pathways between them (e.g., database, APIs, CLIs)\n" +
 "- Any external integrations (e.g., third-party APIs, messaging systems, etc.)\n" +
 "- If adding styles, ensure that they are human readable with dark colours.`\n" +
 "- Use **double quotes** around all text inside square brackets (`[ ]`) or round brackets (`( )`) to ensure the diagram is valid for GitHub's markdown rendering and complies with Mermaid's syntax requirements.\n\n" +

 "When documenting configuration via environment variables, use the following table structure:\n" +
 "- Variable Name \n" +
 "- Description \n" +
 "- Required \n" +
 "- Example\n\n" +

 "Below the table, include a sample code block for `.env.local` configuration.\n" +
 "Do not add redundant information or sections such as dependency lists or \"Tools Used.\"\n" +
 "Ensure the `README.md` is written with certainty, maintaining a professional tone.\n\n" +

"If the codebase is a library, describe exports relative to the repository structure, with clear explanations of their main functionality.\n\n" +

"This is internal, so no license section is needed.\n\n" +

"Ensure that the README.md always follows this order:\n" +
"0. Table of contents (via GitHub's ToC macro)\n" +
"1. Purpose\n" +
"2. Architecture Overview\n" +
"3. Getting Started\n" +
"4. Configuration\n" +
"5. Any kind of usage (CLI, API interations, etc.)\n" +
"6. Exports documentation\n" +
"7. Anything else that you feel is necessary\n" +
"{{trump}}"

	if useTrump {
		rp = strings.ReplaceAll(rp, "{{trump}}",  `Ensure to write this as though you're Donald Trump, with all of his mannerisms`)
	} else {
		rp = strings.ReplaceAll(rp, "{{trump}}", "")
	}

	rp = strings.ReplaceAll(rp, "{{codebase}}", codebase)

	return rp
}
