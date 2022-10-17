# Ad Library Middleware

This repository contains middleware that extracts data from the Facebook [Ad Library](https://www.facebook.com/ads/library/api/?source=archive-landing-page) and inserts it into BigQuery.

This is a contracted project I worked on for [Saguaro Strategies](https://www.saguarostrategies.com/).

### Google Credentials

A valid Google Cloud service account key file is needed. Set the following environment variable to the path of the key file before running:

```sh
export GOOGLE_APPLICATION_CREDENTIALS=[path to JSON key file]
```

See more at [Authenticating as a Service Account](https://cloud.google.com/docs/authentication/production#passing_variable).

### Facebook Credentials

The credentials for the Facebook project must also be provided. Create a JSON file at the root of the repository called `fb-credentials.json` and store the following information:

```json
{
    "app_id": "[the app id]",
    "app_secret": "[the app secret]",
    "access_token": "[a valid access token]"
}
```

The access token in this file will periodically be exchanged for a long-lived (60 day) access token.
