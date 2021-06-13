const { Octokit } = require("@octokit/action");

(async () => {
  const octokit = new Octokit();
  const sha = process.env.SHA;
 
  console.log({
    owner: 'funapy-sandbox',
    repo: 'docs',
    sha: sha,
    state: "success"
  });
  
  // See https://developer.github.com/v3/issues/#create-an-issue
  const { data } = await octokit.request('POST https://api.github.com/repos/{owner}/{repo}/statuses/{sha}', {
    owner: 'funapy-sandbox',
    repo: 'docs',
    sha: sha,
    state: "success"
  });
  console.log(JSON.stringify(data));
})();
