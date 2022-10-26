# Feature Flags Service

In order to facilitate feature flagging at a regional level, we need to have an API a frontend can query to check whether certain features
or pages can be enabled or disabled for different users.

To deploy:

```
gcloud run deploy feature-flags-service --source . --region=<YOUR REGION(S)> --allow-unauthenticated
```
