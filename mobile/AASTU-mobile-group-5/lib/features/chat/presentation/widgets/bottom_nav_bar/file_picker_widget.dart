import 'package:file_picker/file_picker.dart';
import 'package:flutter/material.dart';

class FilePickerWidget extends StatelessWidget {
  const FilePickerWidget({super.key});

  Future<void> _pickFile() async {
    final result = await FilePicker.platform.pickFiles();
    if (result != null) {
      // Handle the file
      final file = result.files.single;
      // You can upload the file or do other operations here
    }
  }

  @override
  Widget build(BuildContext context) {
    return IconButton(
      icon: const Icon(Icons.attach_file),
      onPressed: _pickFile,
    );
  }
}
