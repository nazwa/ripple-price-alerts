# ripple-price-alerts
Get notified via SMS when ripple prices hit your targets. Saves you having to watch ripple charts all day long.

# Usage
## Config file
Copy and paste the sample config file (**jobs.json.sample** -> **jobs.json**)

  ```json
  {
  	"watcher": {
  		"url": "https://api.ripplecharts.com/api/exchange_rates",
  		"notifications": [
  			{
  				"service": "Twilio",
  				"settings": {
  					"sid": "",
  					"token": "",
  					"from": "+",
  					"to": "+"
  				}
  			}
  		]
  	},
  	"jobs": [
  		{
  			"base": {
  				"currency": "XRP"
  			},
  			"counter": {
  				"currency": "USD",
  				"issuer": "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B"
  			},
  			"live": true,
  			"value": 0.0083,
  			"mode": "above"			
  		},
  		
  		{
  			"base": {
  				"currency": "XRP"
  			},
  			"counter": {
  				"currency": "USD",
  				"issuer": "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B"
  			},
  			"live": true,
  			"value": 0.008,
  			"mode": "below"			
  		}
  	]
  		
  }
```

### Notifications
Currently only twilio is supported, but email would be good as well. Notifications is an array, so you can forward your alerts to multiple recipients.

```json

{
	"service": "Twilio",
	"settings": {
		"sid": "",
		"token": "",
		"from": "+",
		"to": "+"
	}
}
```

### Jobs
Set **mode** to "above" or "below" to trigger the notification if price drops below your value or raises above it. Jobs field is an array, so you can have as many jobs associated with this group as you want.
```json

{
	"base": {
		"currency": "XRP"
	},
	"counter": {
		"currency": "USD",
		"issuer": "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B"
	},
	"live": true,
	"value": 0.0083,
	"mode": "above"			
}
```

## Run

The tool uses https://github.com/kardianos/service to install itself as a service, so you have the option to **run it directly in terminal**, or set it up 'forever' as a service:

```
ripple-price-alerts -service=install
ripple-price-alerts -service=start
ripple-price-alerts -service=restart
ripple-price-alerts -service=stop
ripple-price-alerts -service=uninstall
```


# Tell me 


# todo
Proper readme
