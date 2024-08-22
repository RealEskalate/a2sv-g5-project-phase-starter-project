import 'package:flutter/material.dart';
import 'file_picker_widget.dart';
import 'image_capture_widget.dart';
import 'voice_recording_widget.dart';

class CustomBottomNavigationBar extends StatelessWidget {
  const CustomBottomNavigationBar({super.key});

  @override
  Widget build(BuildContext context) {
    return Container(
      color: Colors.white,
      padding: const EdgeInsets.all(8),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: [
          const FilePickerWidget(),
          Expanded(
            child: TextField(
              decoration: InputDecoration(
                hintText: 'Write a message...',
                border: OutlineInputBorder(
                  borderRadius: BorderRadius.circular(10),
                ),
                contentPadding: const EdgeInsets.symmetric(horizontal: 10),
              ),
              // Add any text editing controllers or functionality here
            ),
          ),
          const ImageCaptureWidget(),
          const VoiceRecordingWidget(),
          IconButton(
            icon: const Icon(Icons.send),
            onPressed: () {
              // Handle send action
            },
          ),
        ],
      ),
    );
  }
}
