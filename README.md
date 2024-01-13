## Web UI for MailHedgehog

SPA with reference to MailHedgehog API

## Usage

```shell
cp .env.example .env
yarn dev
```

## Build

```shell
yarn && yarn eslint
yarn build
```

## Configuration

All configuration parameters are optional

File [mh-configuration.json](public%2Fmh-configuration.json) (or you can set other name in .env) allow next configuration:

```json
{
  "http": {
    "auth": "internal",
    "baseUrl": "http://0.0.0.0:8026"
  },
  "ui": {
    "appName": "LocalApp",
    "copyrightsText": "<a href=\"https://test.com/\" class=\"underline transition-colors duration-500 hover:text-primary-600\" target=\"_blank\">Example Communications Ltd</a>"
    "appLogoUrl": "https://me/image.svg",
    "appLogoStyle": [
      "height: 80px"
    ]
  },
  "sharing": {
    "enabled": true
  }
}
```

### http

`http.auth`: Enum *"internal"* or *"basic"* to flag what type of login form display
`http.baseUrl`: API base url

### ui

`ui.appName`: String described application name (your company name)
`ui.copyrightsText`: String/HTML described application name (your company name)
`ui.appLogoUrl`: String, link to custom url
`ui.appLogoStyle`: Array/Object of css syles of logo element (you can set custom width ot height)

### sharing

`sharing.enabled`: Enable sharing UI

## Credits

- [![Think Studio](https://yaroslawww.github.io/images/sponsors/packages/logo-think-studio.png)](https://think.studio/)
