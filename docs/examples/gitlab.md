# How to configure GitLab Runner for GitLab CE integration tests

We will register the Runner using a confined Docker executor.

The registration token can be found at `https://gitlab.com/project_namespace/project_name/runners`.
You can export it as a variable and run the command below as is:

```shell
gitlab-runner register \
--non-interactive \
--url "https://gitlab.com" \
--registration-token "$REGISTRATION_TOKEN" \
--description "gitlab-ce-ruby-2.6" \
--executor "docker" \
--docker-image ruby:2.6 --docker-mysql latest \
--docker-services latest --docker-redis latest
```

You now have a GitLab CE integration testing instance with bundle caching.
Push some commits to test it.

For [advanced configuration](../configuration/advanced-configuration.md), look into
`/etc/gitlab-runner/config.toml` and tune it.
