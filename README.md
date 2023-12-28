<h1 align="center">
  <img src="logo.png" alt="MediaMTX / rtsp-simple-server">

# MEDIAMTX DROID

This repository is help to build MEDIAMTX on Android,see offical page for more information.

### Android

Init the gomobile or follow the [offical guide](https://pkg.go.dev/golang.org/x/mobile/cmd/gomobile).

```
go install golang.org/x/mobile/cmd/gomobile@latest // install gomobile
goomobile init // init the enviroment
go get -u golang.org/x/mobile/... // if the version of gomobile is to old
```

Use this commmd in project folder.

```
gomobile bind -o mediamtxdroid.aar -target android -javapkg cn.wanghan2 -androidapi 21
```

Now you will see the AAR and JAR in the root if no any other issues.



By the way,if your add AAR in your project,don't forget to add permissions :-)

### Usage

Import the AAR and sources(optional),type codes: `Mediamtxdroid.` and see the completion.

Generally use `Mediamtxdroid.startServer(mediamtxYaml.path,logPath)` is enough.

Note:all this api are design for single proccess,so you may should call them from another process,or same time it will cause unexpected exit the whole app.

For example:

```
<service
	android:name=".MediaMTXService"
        android:enabled="true"
        android:exported="true"
        android:process=".MediaMTXService" />
```
