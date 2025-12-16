#!/bin/bash
cd /home/mmcnew/repos/forgexml-scap || exit 1
/home/mmcnew/go/bin/forgexml -output internal/generated -root-pkg github.com/aequo-labs/forgexml-scap/internal/generated -generate-tests schemas/arf/asset-reporting-format_1.1.0.xsd
