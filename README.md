# Netlify Glip Integration

Receive Netlify deploy notifications in Glip chat with a small service running on
[Google App Engine](https://cloud.google.com/appengine/)

## Installation

1. Set `GLIP_HOOK_URL` in `app.yaml`
2. `goapp deploy -application %YOUR_APP% -version %SOME_VERSION%`
3. Use *https://%YOUR_APP%.appspot.com* as **URL to notify** in **Outgoing webhook**.
