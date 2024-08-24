import 'package:flutter/material.dart';
import 'package:flutter_sound/flutter_sound.dart';
import 'package:permission_handler/permission_handler.dart';

class AudioRecorder extends StatefulWidget {
  @override
  _AudioRecorderState createState() => _AudioRecorderState();
}

class _AudioRecorderState extends State<AudioRecorder> {
  FlutterSoundRecorder? _recorder;
  bool _isRecording = false;

  @override
  void initState() {
    super.initState();
    _recorder = FlutterSoundRecorder();
    _initializeRecorder();
  }

  Future<void> _initializeRecorder() async {
    await Permission.microphone.request();
    await _recorder!.openRecorder();
  }

  @override
  void dispose() {
    _recorder!.closeRecorder();
    _recorder = null;
    super.dispose();
  }

  Future<void> _startRecording() async {
    await _recorder!.startRecorder(toFile: 'audio.aac');
    setState(() {
      _isRecording = true;
    });
  }

  Future<void> _stopRecording() async {
    await _recorder!.stopRecorder();
    setState(() {
      _isRecording = false;
    });
  }

  @override
  Widget build(BuildContext context) {
    // return IconButton(
    //   icon: const Icon(Icons.mic, color: Colors.grey),
    //   onPressed: () {
    //     // Add functionality to start voice recording here in the future
    //     ScaffoldMessenger.of(context).showSnackBar(
    //       const SnackBar(content: Text('Voice recording not implemented yet.')),
    //     );
    //   },
    // );
    return IconButton(
      icon: Icon(_isRecording ? Icons.stop : Icons.mic),
      onPressed: _isRecording ? _stopRecording : _startRecording,
      iconSize: 30.0,
      color: _isRecording ? Colors.red : Colors.blue,
    );

    
  }
}

