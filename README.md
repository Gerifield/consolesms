# consolesms
SMS from console via Pushbullet and go

#Build

Use glide!

`glide install`

Then buld it.

`go buld`

#Usage:

Copy the config.samle.yaml and add a Pushbullet token to it and set some numbers.
Place this file next to the binary.

Then use the console:

`echo "SMS text" | gopbsms me`

(where me is a name from the `config.yaml`.)

Or

`echo "SMS text" | gopbsms +1234567890`
