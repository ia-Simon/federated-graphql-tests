use apollo_router::plugin as apollo_plugin;

#[derive(serde::Deserialize, schemars::JsonSchema)]
struct ForwardHeadersConfig {
  headers: Vec<String>
}

struct ForwardHeaders {
  headers: Vec<String>
}

#[async_trait::async_trait]
impl apollo_plugin::Plugin for ForwardHeaders {
  type Config = ForwardHeadersConfig;

  async fn new(init: apollo_plugin::PluginInit<Self::Config>) -> Result<Self, tower::BoxError> {
    let ForwardHeadersConfig { headers } = init.config;
    Ok(Self { headers })
  }
}

apollo_router::register_plugin!(
  "local",
  "forward_headers",
  ForwardHeaders
);

#[cfg(test)]
mod tests {
  #[tokio::test]
  async fn plugin_registered() {
    let config = serde_json::json!({
      "plugins": {
        "local.forward_headers": {
          "headers": [
            "x-test-header"
          ]
        }
      }
    });

    apollo_router::TestHarness::builder()
      .configuration_json(config)
      .unwrap()
      .build_router()
      .await
      .unwrap();
  }
}