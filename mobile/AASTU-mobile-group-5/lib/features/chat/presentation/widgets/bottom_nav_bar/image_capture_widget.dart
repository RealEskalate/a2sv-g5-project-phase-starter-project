
import 'package:flutter/material.dart';
import 'package:image_picker/image_picker.dart';

class ImageCaptureWidget extends StatefulWidget {
  const ImageCaptureWidget({super.key});

  @override
  _ImageCaptureWidgetState createState() => _ImageCaptureWidgetState();
}

class _ImageCaptureWidgetState extends State<ImageCaptureWidget> {
  XFile? _image;

  Future<void> _pickImage() async {
    final ImagePicker picker = ImagePicker();
    final XFile? image = await picker.pickImage(source: ImageSource.camera);

    if (image != null) {
      setState(() {
        _image = image;
      });
      // Handle the captured image
      print('Captured image: ${_image!.path}');
    }
  }

  @override
  Widget build(BuildContext context) {
    return IconButton(
      icon: const Icon(Icons.camera_alt),
      onPressed: _pickImage,
    );
  }
}
