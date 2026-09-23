# m07l04-02: Encoding health results

You can encode health results as stable JSON for pipelines, dashboards, and later automation.

The result fields carry tags for the external report names. Marshal encodes the complete slice as one JSON document, preserving the target name and health decision. The checker can now feed this output to a pipeline or a dashboard without asking the consumer to parse log prose.



Run: `sh run.sh`.

Verification: noVerify (Go toolchain is unavailable on this machine)
