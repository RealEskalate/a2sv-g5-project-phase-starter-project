import 'dart:async';
import 'dart:io';
import 'package:flutter/material.dart';
import 'package:flutter_sound/flutter_sound.dart';
import 'package:path_provider/path_provider.dart';
import 'package:permission_handler/permission_handler.dart';

class RealTimeAudioRecorder extends StatefulWidget {
  @override
  _RealTimeAudioRecorderState createState() => _RealTimeAudioRecorderState();
}

class _RealTimeAudioRecorderState extends State<RealTimeAudioRecorder> {
  FlutterSoundRecorder? _recorder;
  FlutterSoundPlayer? _player;
  bool _isRecording = false;
  bool _isPlaying = false;
  String? _filePath;
  Timer? _timer;
  int _recordDuration = 0;

  @override
  void initState() {
    super.initState();
    _recorder = FlutterSoundRecorder();
    _player = FlutterSoundPlayer();
    _initRecorder();
  }

  Future<void> _initRecorder() async {
    await Permission.microphone.request();
    Directory tempDir = await getTemporaryDirectory();
    _filePath = '${tempDir.path}/voice_message.aac';
    await _recorder!.openRecorder();
  }

  void _startRecording() async {
    setState(() {
      _isRecording = true;
      _recordDuration = 0;
    });
    _recorder!.startRecorder(toFile: _filePath).then((value) {
      _timer = Timer.periodic(Duration(seconds: 1), (Timer t) {
        setState(() {
          _recordDuration++;
        });
      });
    });
  }

  void _stopRecording() async {
    await _recorder!.stopRecorder();
    _timer?.cancel();
    setState(() {
      _isRecording = false;
    });
  }

  void _playAudio() async {
    if (_isPlaying) {
      await _player!.stopPlayer();
      setState(() {
        _isPlaying = false;
      });
    } else {
      await _player!.startPlayer(fromURI: _filePath, whenFinished: () {
        setState(() {
          _isPlaying = false;
        });
      });
      setState(() {
        _isPlaying = true;
      });
    }
  }

  @override
  void dispose() {
    _recorder!.closeRecorder();
    _player!.closePlayer();
    _timer?.cancel();
    super.dispose();
  }

  String _formatDuration(int seconds) {
    String twoDigits(int n) => n.toString().padLeft(2, "0");
    final minutes = twoDigits(seconds ~/ 60);
    final sec = twoDigits(seconds % 60);
    return "$minutes:$sec";
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            GestureDetector(
              onLongPress: _startRecording,
              onLongPressUp: _stopRecording,
              child: CircleAvatar(
                radius: 30,
                backgroundColor: Colors.purple,
                child: Icon(
                  _isRecording ? Icons.mic : Icons.mic_none,
                  color: Colors.white,
                ),
              ),
            ),
            SizedBox(height: 20),
            Text(
              _formatDuration(_recordDuration),
              style: TextStyle(fontSize: 24),
            ),
            SizedBox(height: 20),
            if (_filePath != null)
              GestureDetector(
                onTap: _playAudio,
                child: Row(
                  mainAxisSize: MainAxisSize.min,
                  children: [
                    CircleAvatar(
                      radius: 24,
                      backgroundColor: Colors.purple,
                      child: Icon(
                        _isPlaying ? Icons.pause : Icons.play_arrow,
                        color: Colors.white,
                      ),
                    ),
                    SizedBox(width: 16),
                    Text(_formatDuration(_recordDuration)),
                  ],
                ),
              ),
          ],
        ),
      ),
    );
  }
}
