#!/bin/bash

echo "VERCEL_GIT_COMMIT_REF: $VERCEL_GIT_COMMIT_REF"

if [[ "$VERCEL_GIT_COMMIT_REF" == "main" ]] ; then
  echo "✅ - Build can proceed"
  # 公式: https://vercel.com/docs/concepts/projects/overview#ignored-build-step によると
  # You can customize this behavior with a command that exits with code 1 (new Build needed) or code 0.
  # のため、新たにビルドが必要な場合はexit code 1を返す
  exit 1;

else
  echo "🛑 - Build canceled"
  # ビルド不要の場合はexit code 0を返す
  exit 0;
fi
