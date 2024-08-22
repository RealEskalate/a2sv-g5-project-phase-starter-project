import 'package:flutter/material.dart';

class VoiceRecordingWidget extends StatelessWidget {
  const VoiceRecordingWidget({super.key});

  @override
  Widget build(BuildContext context) {
    return IconButton(
      icon: const Icon(Icons.mic, color: Colors.grey),
      onPressed: () {
        // Add functionality to start voice recording here in the future
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(content: Text('Voice recording not implemented yet.')),
        );
      },
    );
  }
}
