
curl --request POST \
  --url https://api.github.com/repos/funapy-sandbox/actions-sandbox/statuses/${SHA} \
  -H "Authorization: token ${GITHUB_TOKEN}" \
  -H "Accept: application/vnd.github.v3+json" \
  --data '{
    "context": "lint",
    "state": "success",
    "description": "lint passed"
  }'
