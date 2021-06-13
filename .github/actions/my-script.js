const { Octokit } = require("@octokit/action");

const octokit = new Octokit();
const sha = process.env.SHA;

// See https://developer.github.com/v3/issues/#create-an-issue
const { data } = octokit.request('POST /repos/{owner}/{repo}/statuses/{sha}', {
  owner: 'funapy-sandbox',
  repo: 'docs',
  sha: sha,
  state: "success"
});

console.log(JSON.stringify(data));
