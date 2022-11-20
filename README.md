# Hero Web Scanner

## Overview
*Hero Web Scanner* is a client/server application for controlling the audio received from an sdr-based scanning application and playing it in a **web browser**.

This project is designed to work with [Trunk Recorder](https://github.com/robotastic/trunk-recorder). You must have Trunk Recorder set up and running before using this tool. *Hero Web Scanner* uses the OpenMHz protocol to receive audio files from Trunk Recorder.

### Known Limitations
- Calls may only be played once (no instant recall)
- UI responsiveness is weak
- There is no concept of systems, therefore systems with conflicting talkgroup IDs cannot be supported
- There is no decoding or display of unit IDs

## Installation
Go to [releases](releases)

## Configuring Trunk Recorder
Some changes to the Trunk Recorder configuration are required to make *Hero Web Scanner* work, it is not a complete Trunk Recorder configuration. The configuration is the same as you would do for OpenMHZ (note: you cannot run *Hero Web Scanner* and OpenMHZ at the same time currently). The below configuration snippet contains only the fields required for *Hero Web Scanner* and is not a complete Trunk Recorder configuration.

**statusServer**: Websocket URI for the *Hero Web Scanner* service. If *Hero Web Scanner* and Trunk Recorder are running on the same server, you can use local host as shown below. Be sure to include "ws://" at the beginning of the URI and "/ws/recorder" at the end. Secure web sockets are hot currently supported

**uploadServer**: HTTP endpoint for uploading audio files to *Hero Web Scanner*. The URI must begin with "http://" and end with "/upload". HTTPS us not currently supported.

**apiKey**: While *Hero Web Scanner* does not currently support API keys, this field tells Trunk Recorder that you wish to upload the audio to an OpenMHZ server. You can specify any value.

**audioArchive**: we recommend setting this to false so your server does not fill up with audio files. *Hero Web Scanner* keeps its own cache of audio files and will never cull the files that Trunk Recorder saves. If you are using multiple uploaders, such as Broadcastify, the files will not be deleted until all uploads succeed. You may set this to true if you wish, it does not affect the operation of *Hero Web Scanner*.

config.json:
``` json
{
    "statusServer": "ws://localhost:8080/ws/recorder",
    "uploadServer": "http://localhost:8080/upload",
   
    "systems": [{
	"shortName":"MSP",
	    "apiKey": "1234",
	    "audioArchive": false
    }]
}
```

## Configuring Hero Web Scanner
*Hero Web Scanner* ships with a default **settings.toml** file. This is a complete configuration that will allow *Hero Web Scanner* to work out of the box if it is only accessed from the same machine where the service runs.

The comments in the configuration file describe each option currently available.

The settings **show_active_calls** and **play_unknown_talkgroups** are not yet implemented.

settings.toml:
``` toml 
# Settings for the server
[server]
    # The port that the server will listen on
    service_http_port=8080
    # The file the server will use to lookup talkgroup info
    # it will be checked each time a talkgroup is looked up
    # there is no need to restart the server when this file changes
    talkgroups_file="talkgroups.csv"
    console_logs=true

# Settings that will be sent to the client as defaults
# Once the client connects, it will cache these settings
# and they must we changed in the client UI
[client]
    # Talkgroups the client will avoid (can be overridden in the UI)
    disabled_talkgroups = []
    # Show the active calls in the UI
    show_active_calls= true
    # Play talkgroups even if they are not in the talkgroups file
    play_unknown_talkgroups= true

[client.server]
    # This is where the client will attempt to connect
    uri="localhost:8080"
```

### Talkgroups.csv